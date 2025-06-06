<script lang="ts">
  import { createEventDispatcher, getContext } from "svelte";
  import ColumnHeader from "../core/ColumnHeader.svelte";
  import type {
    VirtualizedTableColumns,
    VirtualizedTableConfig,
  } from "../types";

  const dispatch = createEventDispatcher();

  export let columns: VirtualizedTableColumns[];
  export let pinnedColumns: VirtualizedTableColumns[] = [];
  export let virtualColumnItems;
  export let noPin = false;
  export let showDataIcon = false;
  export let sortByMeasure: string | null = null;

  const config: VirtualizedTableConfig = getContext("config");

  const getColumnHeaderProps = (header) => {
    const column = columns[header.index];
    const name = column.label || column.name;
    const isEnableResizeDefined = "enableResize" in column;
    const enableResize = isEnableResizeDefined ? column.enableResize : true;
    const enableSorting =
      "enableSorting" in column
        ? column.enableResize
        : config.table === "DimensionTable";
    return {
      name,
      enableResize,
      enableSorting,
      type: column.type,
      description: column.description || "",
      pinned: pinnedColumns.some((pinCol) => pinCol.name === column.name),
      isSelected: sortByMeasure === column.name,
      sorted: column.sorted,
    };
  };
</script>

<div class="w-full sticky top-0 z-10">
  {#each virtualColumnItems as header (header.key)}
    {@const props = getColumnHeaderProps(header)}
    <ColumnHeader
      on:resize-column
      on:reset-column-width
      {...props}
      {header}
      {noPin}
      {showDataIcon}
      on:pin={() => {
        dispatch("pin", columns[header.index]);
      }}
      on:click-column={() => {
        dispatch("click-column", columns[header.index]?.name);
      }}
    />
  {/each}
</div>
