<script lang="ts">
  import Shortcut from "@rilldata/web-common/components/tooltip/Shortcut.svelte";
  import Tooltip from "@rilldata/web-common/components/tooltip/Tooltip.svelte";
  import TooltipContent from "@rilldata/web-common/components/tooltip/TooltipContent.svelte";
  import TooltipShortcutContainer from "@rilldata/web-common/components/tooltip/TooltipShortcutContainer.svelte";
  import { Inspect } from "lucide-svelte";
  import { createEventDispatcher } from "svelte";
  import { Button } from "../../../components/button";

  export let disabled = false;
  export let areAllTableRowsSelected: boolean;

  const dispatch = createEventDispatcher();
</script>

<Tooltip distance={4} location="top">
  <TooltipContent slot="tooltip-content">
    {#if areAllTableRowsSelected}
      <div>Deselect all selections</div>
    {:else}
      <TooltipShortcutContainer pad={false}>
        <div>Select all</div>
        <Shortcut
          ><span style="font-family: var(--system);"> ⌘ </span> + A</Shortcut
        >
      </TooltipShortcutContainer>
    {/if}
  </TooltipContent>
  <Button
    type="toolbar"
    on:click={() => dispatch("toggle-all-search-items")}
    {disabled}
  >
    <div class="ui-copy-icon">
      <Inspect size={16} />
    </div>
    {areAllTableRowsSelected && !disabled ? "Deselect all" : "Select all"}
  </Button>
</Tooltip>
