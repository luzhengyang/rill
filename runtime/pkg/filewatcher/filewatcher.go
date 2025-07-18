package filewatcher

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"
	runtimev1 "github.com/rilldata/rill/proto/gen/rill/runtime/v1"
	"github.com/rilldata/rill/runtime/drivers"
	"go.uber.org/zap"
	"golang.org/x/exp/maps"
	"golang.org/x/sys/unix"
)

const batchInterval = 250 * time.Millisecond

const maxBufferSize = 1000

// Watcher implements a recursive, batching file watcher on top of fsnotify.
type Watcher struct {
	logger           *zap.Logger
	root             string
	ignorePaths      []string
	watcher          *fsnotify.Watcher
	closed           atomic.Bool
	done             chan struct{}
	err              error
	mu               sync.Mutex
	subscribers      map[int]WatchCallback
	nextSubscriberID int
	buffer           map[string]WatchEvent
}

type WatchCallback func(event []WatchEvent)

type WatchEvent struct {
	Type     runtimev1.FileEvent
	FullPath string
	RelPath  string
	Dir      bool
	isCreate bool
}

// NewWatcher creates a new watcher for the given root directory.
// The root directory must be an absolute path.
func NewWatcher(root string, ignorePaths []string, logger *zap.Logger) (*Watcher, error) {
	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	w := &Watcher{
		logger:      logger,
		root:        root,
		ignorePaths: ignorePaths,
		watcher:     fsw,
		done:        make(chan struct{}),
		subscribers: make(map[int]WatchCallback),
		buffer:      make(map[string]WatchEvent),
	}

	err = w.addDir(root, false, true)
	if err != nil {
		w.watcher.Close()
		return nil, err
	}

	go w.run()

	return w, nil
}

func (w *Watcher) Close() {
	w.closeWithErr(nil)
}

func (w *Watcher) closeWithErr(err error) {
	// Support multiple calls, but only actually close once.
	// Not using w.mu here because someday someone will try to close the watcher from a callback.
	if w.closed.Swap(true) {
		return
	}

	closeErr := w.watcher.Close()
	w.err = errors.Join(err, closeErr)
	if w.err == nil {
		w.err = fmt.Errorf("file watcher closed")
	}

	close(w.done)
}

func (w *Watcher) Subscribe(ctx context.Context, fn WatchCallback) error {
	w.mu.Lock()
	if w.err != nil {
		w.mu.Unlock()
		return w.err
	}
	id := w.nextSubscriberID
	w.nextSubscriberID++
	w.subscribers[id] = fn
	w.mu.Unlock()

	defer func() {
		w.mu.Lock()
		delete(w.subscribers, id)
		w.mu.Unlock()
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-w.done:
		return w.err
	}
}

// flush emits buffered events to all subscribers.
// Note it is called in the event loop in runInner, so new events will not be appended to w.buffer while a flush is running.
// Calls to flush block until all subscribers have processed the events. This is an acceptable trade-off for now, but we may want to revisit it in the future.
func (w *Watcher) flush() {
	if len(w.buffer) == 0 {
		return
	}

	for p, event := range w.buffer {
		if !event.isCreate {
			continue
		}
		// check for directory for CREATE events
		info, err := os.Stat(event.FullPath)
		event.Dir = err == nil && info.IsDir()
		if event.Dir {
			// add directory to tracking paths
			err = w.addDir(event.FullPath, true, false)
			if err != nil {
				delete(w.buffer, p)
				continue
			}
		}
		w.buffer[p] = event
	}

	events := maps.Values(w.buffer)

	w.mu.Lock()
	defer w.mu.Unlock()

	for _, fn := range w.subscribers {
		fn(events)
	}

	w.buffer = make(map[string]WatchEvent)
}

func (w *Watcher) run() {
	err := w.runInner()
	w.closeWithErr(err)
}

func (w *Watcher) runInner() error {
	ticker := time.NewTicker(batchInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ticker.Stop()
			w.flush()
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return nil
			}
			if err == nil || isNotExists(err) {
				continue
			}
			return err
		case e, ok := <-w.watcher.Events:
			if !ok {
				return nil
			}

			we := WatchEvent{}
			if e.Has(fsnotify.Remove) || e.Has(fsnotify.Rename) {
				we.Type = runtimev1.FileEvent_FILE_EVENT_DELETE
			} else if e.Has(fsnotify.Create) || e.Has(fsnotify.Write) || e.Has(fsnotify.Chmod) {
				we.Type = runtimev1.FileEvent_FILE_EVENT_WRITE
			} else {
				continue
			}
			we.isCreate = e.Has(fsnotify.Create)

			p, err := filepath.Rel(w.root, e.Name)
			if err != nil {
				w.logger.Warn("ignoring watcher event: failed to get relative path", zap.String("root", w.root), zap.String("event_name", e.Name), zap.String("event_op", e.Op.String()))
				continue
			}

			p = path.Join("/", p)
			we.RelPath = p
			we.FullPath = e.Name

			// Do not send files for ignored paths
			if drivers.IsIgnored(p, w.ignorePaths) {
				continue
			}

			existing, ok := w.buffer[p]
			if ok && existing.isCreate && we.Type == runtimev1.FileEvent_FILE_EVENT_WRITE {
				// copy over `IsCreate` within the batch for a path
				we.isCreate = existing.isCreate
			}
			w.buffer[p] = we

			// Reset the timer so we only flush when no events have been observed for batchInterval.
			// (But to avoid the buffer growing infinitely in edge cases, we enforce a max buffer size.)
			if len(w.buffer) < maxBufferSize {
				ticker.Reset(batchInterval)
			} else {
				ticker.Stop()
				w.flush()
			}
		}
	}
}

func (w *Watcher) addDir(p string, replay, errIfNotExist bool) error {
	// Check if it is an ignored path
	// We do not watch files for ignored paths
	relPath, err := filepath.Rel(w.root, p)
	if err != nil {
		// should never happen unless some bug in the code
		return fmt.Errorf("failed to add path %q to watcher: failed to get relative path", p)
	}
	relPath = filepath.Join(string(filepath.Separator), relPath)
	if drivers.IsIgnored(relPath, w.ignorePaths) {
		w.logger.Debug("watcher: ignoring path", zap.String("path", p))
		return nil
	}

	err = w.watcher.Add(p)
	if err != nil {
		// Need to check unix.ENOENT (and probably others) since fsnotify doesn't always use cross-platform syscalls.
		if !errIfNotExist && isNotExists(err) {
			return nil
		}
		return err
	}

	entries, err := os.ReadDir(p)
	if err != nil {
		if !errIfNotExist && isNotExists(err) {
			return nil
		}
		return err
	}

	for _, e := range entries {
		fullPath := filepath.Join(p, e.Name())

		if replay {
			ep, err := filepath.Rel(w.root, fullPath)
			if err != nil {
				return err
			}
			ep = path.Join("/", ep)

			w.buffer[ep] = WatchEvent{
				Type:     runtimev1.FileEvent_FILE_EVENT_WRITE,
				FullPath: fullPath,
				RelPath:  ep,
				Dir:      e.IsDir(),
			}
		}

		if e.IsDir() {
			err := w.addDir(fullPath, replay, errIfNotExist)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isNotExists(err error) bool {
	return os.IsNotExist(err) || errors.Is(err, unix.ENOENT)
}
