// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rill/ui/v1/dashboard.proto

package uiv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	runtimev1 "github.com/rilldata/rill/proto/gen/rill/runtime/v1"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = runtimev1.TimeGrain(0)
)

// Validate checks the field values on DashboardState with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DashboardState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DashboardState with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DashboardStateMultiError,
// or nil if none found.
func (m *DashboardState) ValidateAll() error {
	return m.validate(true)
}

func (m *DashboardState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimeRange()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "TimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "TimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeRange()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DashboardStateValidationError{
				field:  "TimeRange",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilters()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "Filters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "Filters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DashboardStateValidationError{
				field:  "Filters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWhere()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "Where",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "Where",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWhere()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DashboardStateValidationError{
				field:  "Where",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHaving() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("Having[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("Having[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardStateValidationError{
					field:  fmt.Sprintf("Having[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TimeGrain

	if all {
		switch v := interface{}(m.GetCompareTimeRange()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "CompareTimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DashboardStateValidationError{
					field:  "CompareTimeRange",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCompareTimeRange()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DashboardStateValidationError{
				field:  "CompareTimeRange",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ActivePage

	// no validation rules for PivotExpanded

	for idx, item := range m.GetPivotSort() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotSort[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotSort[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardStateValidationError{
					field:  fmt.Sprintf("PivotSort[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPivotRowAllDimensions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotRowAllDimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotRowAllDimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardStateValidationError{
					field:  fmt.Sprintf("PivotRowAllDimensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPivotColumnAllDimensions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotColumnAllDimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  fmt.Sprintf("PivotColumnAllDimensions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardStateValidationError{
					field:  fmt.Sprintf("PivotColumnAllDimensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.LeaderboardMeasure != nil {
		// no validation rules for LeaderboardMeasure
	}

	if m.SelectedDimension != nil {
		// no validation rules for SelectedDimension
	}

	if m.ShowTimeComparison != nil {
		// no validation rules for ShowTimeComparison
	}

	if m.AllMeasuresVisible != nil {
		// no validation rules for AllMeasuresVisible
	}

	if m.AllDimensionsVisible != nil {
		// no validation rules for AllDimensionsVisible
	}

	if m.LeaderboardContextColumn != nil {
		// no validation rules for LeaderboardContextColumn
	}

	if m.SelectedTimezone != nil {
		// no validation rules for SelectedTimezone
	}

	if m.ScrubRange != nil {

		if all {
			switch v := interface{}(m.GetScrubRange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  "ScrubRange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardStateValidationError{
						field:  "ScrubRange",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetScrubRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardStateValidationError{
					field:  "ScrubRange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.LeaderboardSortDirection != nil {
		// no validation rules for LeaderboardSortDirection
	}

	if m.LeaderboardSortType != nil {
		// no validation rules for LeaderboardSortType
	}

	if m.LeaderboardMeasureCount != nil {
		// no validation rules for LeaderboardMeasureCount
	}

	if m.LeaderboardShowContextForAllMeasures != nil {
		// no validation rules for LeaderboardShowContextForAllMeasures
	}

	if m.ComparisonDimension != nil {
		// no validation rules for ComparisonDimension
	}

	if m.ExpandedMeasure != nil {
		// no validation rules for ExpandedMeasure
	}

	if m.PinIndex != nil {
		// no validation rules for PinIndex
	}

	if m.ChartType != nil {
		// no validation rules for ChartType
	}

	if m.PivotColumnPage != nil {
		// no validation rules for PivotColumnPage
	}

	if m.PivotTableMode != nil {
		// no validation rules for PivotTableMode
	}

	if m.PivotEnableComparison != nil {
		// no validation rules for PivotEnableComparison
	}

	if len(errors) > 0 {
		return DashboardStateMultiError(errors)
	}

	return nil
}

// DashboardStateMultiError is an error wrapping multiple validation errors
// returned by DashboardState.ValidateAll() if the designated constraints
// aren't met.
type DashboardStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DashboardStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DashboardStateMultiError) AllErrors() []error { return m }

// DashboardStateValidationError is the validation error returned by
// DashboardState.Validate if the designated constraints aren't met.
type DashboardStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DashboardStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DashboardStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DashboardStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DashboardStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DashboardStateValidationError) ErrorName() string { return "DashboardStateValidationError" }

// Error satisfies the builtin error interface
func (e DashboardStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDashboardState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DashboardStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DashboardStateValidationError{}

// Validate checks the field values on DashboardTimeRange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DashboardTimeRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DashboardTimeRange with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DashboardTimeRangeMultiError, or nil if none found.
func (m *DashboardTimeRange) ValidateAll() error {
	return m.validate(true)
}

func (m *DashboardTimeRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.Name != nil {
		// no validation rules for Name
	}

	if m.TimeStart != nil {

		if all {
			switch v := interface{}(m.GetTimeStart()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardTimeRangeValidationError{
						field:  "TimeStart",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardTimeRangeValidationError{
						field:  "TimeStart",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTimeStart()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardTimeRangeValidationError{
					field:  "TimeStart",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.TimeEnd != nil {

		if all {
			switch v := interface{}(m.GetTimeEnd()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DashboardTimeRangeValidationError{
						field:  "TimeEnd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DashboardTimeRangeValidationError{
						field:  "TimeEnd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTimeEnd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DashboardTimeRangeValidationError{
					field:  "TimeEnd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DashboardTimeRangeMultiError(errors)
	}

	return nil
}

// DashboardTimeRangeMultiError is an error wrapping multiple validation errors
// returned by DashboardTimeRange.ValidateAll() if the designated constraints
// aren't met.
type DashboardTimeRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DashboardTimeRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DashboardTimeRangeMultiError) AllErrors() []error { return m }

// DashboardTimeRangeValidationError is the validation error returned by
// DashboardTimeRange.Validate if the designated constraints aren't met.
type DashboardTimeRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DashboardTimeRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DashboardTimeRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DashboardTimeRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DashboardTimeRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DashboardTimeRangeValidationError) ErrorName() string {
	return "DashboardTimeRangeValidationError"
}

// Error satisfies the builtin error interface
func (e DashboardTimeRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDashboardTimeRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DashboardTimeRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DashboardTimeRangeValidationError{}

// Validate checks the field values on DashboardDimensionFilter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DashboardDimensionFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DashboardDimensionFilter with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DashboardDimensionFilterMultiError, or nil if none found.
func (m *DashboardDimensionFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *DashboardDimensionFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DashboardDimensionFilterValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DashboardDimensionFilterValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DashboardDimensionFilterValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DashboardDimensionFilterMultiError(errors)
	}

	return nil
}

// DashboardDimensionFilterMultiError is an error wrapping multiple validation
// errors returned by DashboardDimensionFilter.ValidateAll() if the designated
// constraints aren't met.
type DashboardDimensionFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DashboardDimensionFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DashboardDimensionFilterMultiError) AllErrors() []error { return m }

// DashboardDimensionFilterValidationError is the validation error returned by
// DashboardDimensionFilter.Validate if the designated constraints aren't met.
type DashboardDimensionFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DashboardDimensionFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DashboardDimensionFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DashboardDimensionFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DashboardDimensionFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DashboardDimensionFilterValidationError) ErrorName() string {
	return "DashboardDimensionFilterValidationError"
}

// Error satisfies the builtin error interface
func (e DashboardDimensionFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDashboardDimensionFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DashboardDimensionFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DashboardDimensionFilterValidationError{}

// Validate checks the field values on PivotColumnSort with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PivotColumnSort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PivotColumnSort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PivotColumnSortMultiError, or nil if none found.
func (m *PivotColumnSort) ValidateAll() error {
	return m.validate(true)
}

func (m *PivotColumnSort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Desc

	// no validation rules for Id

	if len(errors) > 0 {
		return PivotColumnSortMultiError(errors)
	}

	return nil
}

// PivotColumnSortMultiError is an error wrapping multiple validation errors
// returned by PivotColumnSort.ValidateAll() if the designated constraints
// aren't met.
type PivotColumnSortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PivotColumnSortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PivotColumnSortMultiError) AllErrors() []error { return m }

// PivotColumnSortValidationError is the validation error returned by
// PivotColumnSort.Validate if the designated constraints aren't met.
type PivotColumnSortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PivotColumnSortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PivotColumnSortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PivotColumnSortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PivotColumnSortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PivotColumnSortValidationError) ErrorName() string { return "PivotColumnSortValidationError" }

// Error satisfies the builtin error interface
func (e PivotColumnSortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPivotColumnSort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PivotColumnSortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PivotColumnSortValidationError{}

// Validate checks the field values on PivotElement with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PivotElement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PivotElement with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PivotElementMultiError, or
// nil if none found.
func (m *PivotElement) ValidateAll() error {
	return m.validate(true)
}

func (m *PivotElement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Element.(type) {
	case *PivotElement_PivotTimeDimension:
		if v == nil {
			err := PivotElementValidationError{
				field:  "Element",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for PivotTimeDimension
	case *PivotElement_PivotDimension:
		if v == nil {
			err := PivotElementValidationError{
				field:  "Element",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for PivotDimension
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return PivotElementMultiError(errors)
	}

	return nil
}

// PivotElementMultiError is an error wrapping multiple validation errors
// returned by PivotElement.ValidateAll() if the designated constraints aren't met.
type PivotElementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PivotElementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PivotElementMultiError) AllErrors() []error { return m }

// PivotElementValidationError is the validation error returned by
// PivotElement.Validate if the designated constraints aren't met.
type PivotElementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PivotElementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PivotElementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PivotElementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PivotElementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PivotElementValidationError) ErrorName() string { return "PivotElementValidationError" }

// Error satisfies the builtin error interface
func (e PivotElementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPivotElement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PivotElementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PivotElementValidationError{}
