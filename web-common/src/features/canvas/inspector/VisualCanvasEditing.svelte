<script lang="ts">
  import { replaceState } from "$app/navigation";
  import ComponentsEditor from "@rilldata/web-common/features/canvas/inspector/ComponentsEditor.svelte";
  import PageEditor from "@rilldata/web-common/features/canvas/inspector/PageEditor.svelte";
  import { getCanvasStore } from "@rilldata/web-common/features/canvas/state-managers/state-managers";
  import type { FileArtifact } from "@rilldata/web-common/features/entity-management/file-artifact";
  import { Inspector } from "@rilldata/web-common/layout/workspace";
  import { parseDocument } from "yaml";
  import { runtime } from "@rilldata/web-common/runtime-client/runtime-store";

  export let fileArtifact: FileArtifact;
  export let autoSave: boolean;
  export let canvasName: string;

  $: ({ instanceId } = $runtime);

  $: ({
    canvasEntity: { selectedComponent, components },
  } = getCanvasStore(canvasName, instanceId));

  $: ({ editorContent, updateEditorContent, saveLocalContent, path } =
    fileArtifact);

  $: parsedDocument = parseDocument($editorContent ?? "");
  $: component = components.get($selectedComponent ?? "");

  async function updateProperties(
    newRecord: Record<string, unknown>,
    removeProperties?: Array<string | string[]>,
  ) {
    Object.entries(newRecord).forEach(([property, value]) => {
      if (!value) {
        parsedDocument.delete(property);
      } else {
        parsedDocument.set(property, value);
      }
    });

    if (removeProperties) {
      removeProperties.forEach((prop) => {
        try {
          if (Array.isArray(prop)) {
            parsedDocument.deleteIn(prop);
          } else {
            parsedDocument.delete(prop);
          }
        } catch (e) {
          console.warn(e);
        }
      });
    }

    killState();
    updateEditorContent(parsedDocument.toString(), false, autoSave);
    await saveLocalContent();
  }

  function killState() {
    replaceState(window.location.origin + window.location.pathname, {});
  }
</script>

<Inspector minWidth={320} filePath={path}>
  {#if component}
    <ComponentsEditor {component} />
  {:else}
    <PageEditor {canvasName} {fileArtifact} {updateProperties} />
  {/if}
</Inspector>
