<script lang="ts">
  import Button from "@rilldata/web-common/components/button/Button.svelte";
  import * as DropdownMenu from "@rilldata/web-common/components/dropdown-menu/";
  import type { LineStatus } from "@rilldata/web-common/components/editor/line-status/state";
  import Input from "@rilldata/web-common/components/forms/Input.svelte";
  import InputLabel from "@rilldata/web-common/components/forms/InputLabel.svelte";
  import CancelCircle from "@rilldata/web-common/components/icons/CancelCircle.svelte";
  import CaretDownIcon from "@rilldata/web-common/components/icons/CaretDownIcon.svelte";
  import Close from "@rilldata/web-common/components/icons/Close.svelte";
  import Search from "@rilldata/web-common/components/icons/Search.svelte";
  import Trash from "@rilldata/web-common/components/icons/Trash.svelte";
  import { LIST_SLIDE_DURATION } from "@rilldata/web-common/layout/config";
  import { clamp } from "@rilldata/web-common/lib/clamp";
  import { TIMESTAMPS } from "@rilldata/web-common/lib/duckdb-data-types";
  import { eventBus } from "@rilldata/web-common/lib/event-bus/event-bus";
  import { queryClient } from "@rilldata/web-common/lib/svelte-query/globalQueryClient";
  import { TIME_GRAIN } from "@rilldata/web-common/lib/time/config";
  import {
    createConnectorServiceOLAPListTables,
    createQueryServiceTableColumns,
    createRuntimeServiceAnalyzeConnectors,
    createRuntimeServiceGetInstance,
    type MetricsViewSpecDimension,
    type V1Resource,
  } from "@rilldata/web-common/runtime-client";
  import { runtime } from "@rilldata/web-common/runtime-client/runtime-store";
  import { PlusIcon } from "lucide-svelte";
  import { tick } from "svelte";
  import { slide } from "svelte/transition";
  import { parseDocument, Scalar, YAMLMap, YAMLSeq } from "yaml";
  import ConnectorExplorer from "../connectors/ConnectorExplorer.svelte";
  import { connectorExplorerStore } from "../connectors/connector-explorer-store";
  import { OLAP_DRIVERS_WITHOUT_MODELING } from "../connectors/olap/olap-config";
  import { FileArtifact } from "../entity-management/file-artifact";
  import {
    ResourceKind,
    useResource,
  } from "../entity-management/resource-selectors";
  import { featureFlags } from "../feature-flags";
  import { useModels } from "../models/selectors";
  import { useSources } from "../sources/selectors";
  import AlertConfirmation from "../visual-metrics-editing/AlertConfirmation.svelte";
  import MetricsTable from "../visual-metrics-editing/MetricsTable.svelte";
  import VisualMetricsSidebar from "../visual-metrics-editing/VisualMetricsSidebar.svelte";
  import type { Confirmation, ItemType } from "../visual-metrics-editing/lib";
  import {
    editingItemData,
    types,
    YAMLDimension,
    YAMLMeasure,
  } from "../visual-metrics-editing/lib";

  const { clickhouseModeling } = featureFlags;
  const store = connectorExplorerStore.duplicateStore(
    (connector, database, schema, table) => {
      if (!table || !schema) return;

      confirmation = {
        action: "switch",
        connector,
        database,
        schema,
        model: table,
      };

      tableSelectionOpen = false;
    },
  );

  export let fileArtifact: FileArtifact;
  export let errors: LineStatus[];
  export let switchView: () => void;
  export let unsavedChanges = false;

  let searchValue = "";
  let confirmation: Confirmation | null = null;
  let tableSelectionOpen = false;
  let collapsed = {
    measures: false,
    dimensions: false,
  };
  let selected = {
    measures: new Set<number>(),
    dimensions: new Set<number>(),
  };
  let storedProperties: Record<string, unknown> = {};

  $: ({ instanceId } = $runtime);

  $: instance = createRuntimeServiceGetInstance(instanceId, {
    sensitive: true,
  });

  $: olapConnector = $instance.data?.instance?.olapConnector;

  $: totalSelected = selected.measures.size + selected.dimensions.size;

  $: ({ editorContent, updateEditorContent, getResource } = fileArtifact);

  // YAML Parsing
  $: parsedDocument = parseDocument($editorContent ?? "");

  $: raw = {
    measures: parsedDocument.get("measures"),
    dimensions: parsedDocument.get("dimensions"),
  };

  $: modelingDisabled =
    olapConnector &&
    OLAP_DRIVERS_WITHOUT_MODELING.filter(
      (driver) => driver === "clickhouse" && $clickhouseModeling,
    ).includes(olapConnector);

  $: rawSmallestTimeGrain = parsedDocument.get("smallest_time_grain");
  $: rawTimeDimension = parsedDocument.get("timeseries");
  $: rawDatabaseSchema = parsedDocument.get("database_schema");
  $: rawModel = parsedDocument.get("model");
  $: rawTable = parsedDocument.get("table");
  $: rawDatabase = parsedDocument.get("database");
  $: rawConnector = parsedDocument.get("connector");

  $: timeDimension = stringGuard(rawTimeDimension);
  $: databaseSchema = stringGuard(rawDatabaseSchema);
  $: database = stringGuard(rawDatabase);
  $: yamlConnector = stringGuard(rawConnector);
  $: modelOrSourceOrTableName = stringGuard(rawModel) || stringGuard(rawTable);
  $: smallestTimeGrain = stringGuard(rawSmallestTimeGrain);

  $: noTableProperties = !yamlConnector && !database && !databaseSchema;

  $: modelsQuery = useModels(instanceId);
  $: sourcesQuery = useSources(instanceId);
  $: metricsViewQuery = getResource(queryClient, instanceId);

  $: modelNames = $modelsQuery?.data?.map(resourceToOption) ?? [];
  $: sourceNames = $sourcesQuery?.data?.map(resourceToOption) ?? [];
  $: dimensions = $metricsViewQuery?.data?.metricsView?.spec?.dimensions ?? [];
  $: hasSourceSelected =
    noTableProperties &&
    sourceNames.some(({ value }) => value === modelOrSourceOrTableName);
  $: hasModelSelected =
    noTableProperties &&
    modelNames.some(({ value }) => value === modelOrSourceOrTableName);

  $: modelAndSourceOptions = [...modelNames, ...sourceNames];

  $: hasValidModelOrSourceSelection = hasSourceSelected || hasModelSelected;

  $: connectorsQuery = createRuntimeServiceAnalyzeConnectors(instanceId, {
    query: {
      select: (data) => {
        if (!data?.connectors) return [];

        const connectors = data.connectors.filter(
          (connector) =>
            connector.driver?.implementsOlap &&
            connector.driver.name !== "duckdb",
        );
        return connectors;
      },
    },
  });

  $: resourceKind = hasSourceSelected
    ? ResourceKind.Source
    : hasModelSelected
      ? ResourceKind.Model
      : undefined;

  $: hasNonDuckDBOLAPConnector = Boolean($connectorsQuery.data?.length);

  $: resourceQuery =
    resourceKind &&
    useResource(instanceId, modelOrSourceOrTableName, resourceKind);

  $: connector =
    yamlConnector ||
    (hasModelSelected
      ? $resourceQuery?.data?.model?.spec?.outputConnector
      : $resourceQuery?.data?.source?.spec?.sinkConnector) ||
    olapConnector;

  $: columnsQuery = createQueryServiceTableColumns(
    instanceId,
    modelOrSourceOrTableName,
    {
      connector,
      database,
      databaseSchema,
    },
    {
      query: {
        enabled: Boolean(modelOrSourceOrTableName && connector),
      },
    },
  );

  $: ({ data: columnsResponse } = $columnsQuery);

  $: columns = columnsResponse?.profileColumns ?? [];

  $: timeOptions = columns
    .filter(({ type }) => type && TIMESTAMPS.has(type))
    .map(({ name }) => ({ value: name ?? "", label: name ?? "" }));

  /** display the main error (the first in this array) at the bottom */
  $: mainError = errors?.at(0);

  $: itemGroups = {
    measures:
      raw.measures instanceof YAMLSeq
        ? raw.measures.items.map((item) => {
            return new YAMLMeasure(item instanceof YAMLMap ? item : undefined);
          })
        : [],
    dimensions:
      raw.dimensions instanceof YAMLSeq
        ? createDimensions(raw.dimensions, dimensions)
        : [],
  };

  $: dimensionNamesAndLabels = itemGroups.dimensions.reduce(
    (acc, { name, display_name, resourceName }) => {
      acc.name = Math.max(acc.name, name.length || resourceName?.length || 0);
      acc.label = Math.max(acc.label, display_name.length);
      return acc;
    },
    { name: 0, label: 0 },
  );

  $: measureNamesAndLabels = itemGroups.measures.reduce(
    (acc, { name, display_name }) => {
      acc.name = Math.max(acc.name, name.length);
      acc.label = Math.max(acc.label, display_name.length);
      return acc;
    },
    { name: 0, label: 0 },
  );

  $: longestName = Math.max(
    dimensionNamesAndLabels.name,
    measureNamesAndLabels.name,
  );
  $: longestLabel = Math.max(
    dimensionNamesAndLabels.label,
    measureNamesAndLabels.label,
  );

  $: tablesQuery = createConnectorServiceOLAPListTables(
    {
      instanceId,
      connector,
    },
    {
      query: {
        enabled: !!instanceId && !!connector && !hasValidModelOrSourceSelection,
      },
    },
  );

  $: tables = $tablesQuery.data?.tables ?? [];

  $: hasValidOLAPTableSelected =
    !hasValidModelOrSourceSelection &&
    modelOrSourceOrTableName &&
    tables.find(
      (table) =>
        table.name === modelOrSourceOrTableName &&
        table.database === database &&
        (table.databaseSchema === databaseSchema ||
          (!databaseSchema && table.databaseSchema === "default")),
    );

  $: tableMode = Boolean(hasValidOLAPTableSelected);

  function createDimensions(
    rawDimensions: YAMLSeq<YAMLMap<string, string>>,
    metricsViewDimensions: MetricsViewSpecDimension[],
  ) {
    return rawDimensions.items.map((item, i) => {
      return new YAMLDimension(
        item instanceof YAMLMap ? item : undefined,
        metricsViewDimensions[i],
      );
    });
  }

  function stringGuard(value: unknown | undefined): string {
    return value && typeof value === "string" ? value : "";
  }

  async function setEditing(index: number, type: ItemType, field?: string) {
    if (unsavedChanges) {
      confirmation = {
        action: "cancel",
        index,
        type,
        field,
      };
      return;
    }

    const item = itemGroups[type][index];

    if (item) {
      editingItemData.set({ index, type });
    }

    if (field) {
      await tick();
      document.getElementById(`vme-${field}`)?.focus();
    }
  }

  function resetEditing() {
    editingItemData.set(null);
    unsavedChanges = false;
  }

  function updateProperties(
    newRecord: Record<string, unknown>,
    removeProperties?: string[],
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
        parsedDocument.delete(prop);
      });
    }

    updateEditorContent(parsedDocument.toString(), false, true);
  }

  $: editingItem = $editingItemData
    ? itemGroups[$editingItemData.type][$editingItemData?.index]
    : null;

  function resourceToOption(resource: V1Resource) {
    const value = resource.meta?.name?.name ?? "";
    return {
      value,
      label: value,
    };
  }

  function reorderList(
    initIndexes: number[],
    newIndex: number,
    type: ItemType,
  ) {
    const sequence = raw[type];

    if (!(sequence instanceof YAMLSeq)) {
      return;
    }

    const items = sequence.items as Array<YAMLMap | null | Scalar>;

    const itemsCopy = [...items];
    const sortedIndices = [...initIndexes].sort((a, b) => a - b);
    const editingItemIndex = sortedIndices.indexOf(
      $editingItemData?.index ?? -1,
    );

    const clampedIndex = clamp(0, newIndex, itemsCopy.length);

    const movedItems: Array<YAMLMap | null | Scalar> = [];

    sortedIndices.forEach((index) => {
      movedItems.push(itemsCopy[index]);
      itemsCopy[index] = null;
    });

    itemsCopy.splice(clampedIndex, 0, ...movedItems);

    const countBeforeClamped = sortedIndices.filter(
      (i) => i < clampedIndex,
    ).length;

    const newIndexes = sortedIndices.map((_, dragPosition) => {
      return clampedIndex + dragPosition - countBeforeClamped;
    });

    if (editingItemIndex !== -1) {
      editingItemData.set({ index: newIndexes[editingItemIndex], type });
    }

    if (selected[type].size) {
      selected[type] = new Set(newIndexes);
    }

    // Remove nulls and scalars
    updateProperties({
      [type]: itemsCopy.filter((i) => i && i instanceof YAMLMap),
    });

    eventBus.emit("notification", { message: "Item moved", type: "success" });
  }

  function triggerDelete(index?: number, type?: ItemType) {
    if (totalSelected) {
      confirmation = {
        action: "delete",
      };
    } else {
      confirmation = {
        action: "delete",
        index,
        type,
      };
    }
  }

  function deleteItems(items: Partial<typeof selected>) {
    let deletedEditingItem = false;

    Object.entries(items).forEach(([type, indices]) => {
      const sequence = raw[type];

      if (!(sequence instanceof YAMLSeq)) {
        eventBus.emit("notification", {
          message: "Error deleting items",
          type: "error",
        });
        return;
      }
      const items = sequence.items as Array<YAMLMap>;

      if ($editingItemData !== null) {
        deletedEditingItem =
          indices.has($editingItemData.index) &&
          type === $editingItemData?.type;
      }

      const filtered = items.filter((_, i) => !indices.has(i));

      parsedDocument.set(type, filtered);

      indices.forEach((i) => {
        selected[type].delete(i);
      });
    });

    selected = selected;

    updateEditorContent(parsedDocument.toString(), false, true);

    if (deletedEditingItem) {
      resetEditing();
    }

    eventBus.emit("notification", { message: "Item deleted", type: "success" });
  }

  function duplicateItem(index: number, type: ItemType) {
    const sequence = raw[type];

    if (!(sequence instanceof YAMLSeq)) {
      return;
    }

    const items = [...sequence.items] as Array<YAMLMap>;

    const originalItem = items[index];
    const name = stringGuard(originalItem.get("name"));
    const split = name.split("_copy");
    const potentialNumber = split[1] ? split[1].split("_")[1] : undefined;
    const number: number | undefined = isNaN(Number(potentialNumber))
      ? undefined
      : Number(potentialNumber);

    const canUseOriginalName =
      split.length <= 2 && (!split[1] || number !== undefined);

    const originalName = canUseOriginalName ? split[0] : name;
    const newItem = originalItem.clone() as YAMLMap;

    const itemNames = items.map((i) => stringGuard(i.get("name")));
    let count = number ?? 0;
    let newName = `${originalName}_copy`;
    newItem.set("name", newName);

    while (itemNames.includes(newName)) {
      count++;
      newName = `${originalName}_copy_${count}`;
      newItem.set("name", newName);
    }

    items.splice(index + 1, 0, newItem);

    updateProperties({ [type]: items });

    eventBus.emit("notification", {
      message: "Item duplicated",
      type: "success",
    });
  }

  function switchTableMode() {
    const mode = tableMode;

    const currentProperties = {
      model: rawModel,
      database: rawDatabase,
      connector: rawConnector,
      database_schema: rawDatabaseSchema,
    };
    updateProperties(storedProperties);

    storedProperties = currentProperties;
    tableMode = !mode;
  }
</script>

<div class="wrapper">
  <div class="main-area">
    <div class="flex gap-x-4 border-b pb-4">
      {#if tableMode || modelingDisabled}
        <div class="flex flex-col gap-y-1 w-full">
          <InputLabel label="Table" id="table">
            <svelte:fragment slot="mode-switch">
              {#if !modelingDisabled}
                <button
                  on:click={switchTableMode}
                  class="ml-auto text-primary-600 font-medium text-xs"
                >
                  Select model
                </button>
              {/if}
            </svelte:fragment>
          </InputLabel>
          <DropdownMenu.Root bind:open={tableSelectionOpen}>
            <DropdownMenu.Trigger asChild let:builder>
              <button
                use:builder.action
                {...builder}
                class="flex px-3 gap-x-2 h-8 max-w-full items-center text-sm border-gray-300 border rounded-[2px]
                focus:ring-2 focus:ring-primary-100 focus:border-primary-600 break-all overflow-hidden
               "
              >
                {#if !hasValidOLAPTableSelected}
                  <span class="text-gray-400 truncate">Select table</span>
                {:else}
                  <span class="text-gray-700 truncate">
                    {modelOrSourceOrTableName}
                  </span>
                {/if}
                <CaretDownIcon
                  size="12px"
                  className="!fill-gray-600 ml-auto flex-none"
                />
              </button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content
              sameWidth
              align="start"
              class="!min-w-64  overflow-hidden p-1"
            >
              <div class="size-full overflow-y-auto max-h-72">
                <ConnectorExplorer {store} />
              </div>
            </DropdownMenu.Content>
          </DropdownMenu.Root>
        </div>
      {:else}
        {#key confirmation}
          <Input
            sameWidth
            full
            truncate
            value={noTableProperties ? modelOrSourceOrTableName : undefined}
            options={modelAndSourceOptions}
            placeholder="Select a model"
            label="Model"
            onChange={async (newModelOrSourceName) => {
              if (modelOrSourceOrTableName === newModelOrSourceName) return;
              if (!modelOrSourceOrTableName) {
                updateProperties({ model: newModelOrSourceName }, [
                  "table",
                  "database",
                  "connector",
                  "database_schema",
                ]);
              } else {
                confirmation = {
                  action: "switch",
                  model: newModelOrSourceName,
                };
              }
            }}
          >
            <svelte:fragment slot="mode-switch">
              {#if hasNonDuckDBOLAPConnector}
                <button
                  on:click={switchTableMode}
                  class="ml-auto text-primary-600 font-medium text-xs"
                >
                  Select table
                </button>
              {/if}
            </svelte:fragment>
          </Input>
        {/key}
      {/if}

      <Input
        sameWidth
        full
        enableSearch
        truncate
        value={timeDimension}
        options={timeOptions}
        placeholder="Select time column"
        label="Time column"
        disabledMessage={!hasValidModelOrSourceSelection
          ? "No model selected"
          : "No timestamp columns in model"}
        hint="Column from model that will be used as primary time dimension in dashboards"
        onChange={async (value) => {
          await updateProperties({ timeseries: value });
        }}
      />

      <Input
        sameWidth
        full
        truncate
        value={smallestTimeGrain}
        options={Object.entries(TIME_GRAIN).map(([_, { label }]) => ({
          value: label,
          label: label.charAt(0).toUpperCase() + label.slice(1),
        }))}
        placeholder="Select time grain"
        label="Smallest time grain"
        hint="The smallest time unit by which your charts and tables can be bucketed"
        onChange={async (value) => {
          await updateProperties({ smallest_time_grain: value });
        }}
      />
    </div>

    <div class="grid grid-cols-3 gap-4 relative">
      <div class="col-span-3 sm:col-span-2 lg:col-span-1">
        <Input
          full
          textClass="text-sm"
          placeholder="Search"
          bind:value={searchValue}
          onInput={(value) => {
            searchValue = value;
          }}
        >
          <Search slot="icon" size="16px" color="#374151" />
        </Input>
      </div>

      {#if totalSelected}
        <div
          class="bg-surface rounded-[2px] z-20 shadow-md flex gap-x-0 h-8 text-gray-700 border border-slate-100 absolute right-0"
        >
          <div class="px-2 flex items-center">
            {totalSelected}
            {totalSelected > 1 ? "items" : "item"} selected:
          </div>
          <button
            on:click={() => {
              triggerDelete();
            }}
            class="flex gap-x-2 text-inherit items-center px-2 border-l border-slate-100 hover:bg-gray-50 cursor-pointer"
          >
            <Trash size="16px" />
            Delete
          </button>

          <button
            on:click={() => {
              selected = {
                measures: new Set(),
                dimensions: new Set(),
              };
            }}
            class="flex gap-x-2 text-inherit items-center px-2 border-l border-slate-100 hover:bg-gray-50 cursor-pointer"
          >
            <Close size="14px" />
          </button>
        </div>
      {/if}
    </div>

    <div
      class="flex flex-col gap-y-4 h-fit w-full flex-shrink overflow-y-scroll"
    >
      {#each types as type (type)}
        {@const items = itemGroups[type]}
        <div class="section">
          <header class="flex gap-x-1 items-center flex-none">
            <Button
              type="ghost"
              square
              gray
              noStroke
              onClick={() => {
                collapsed[type] = !collapsed[type];
              }}
            >
              <span
                class="transition-transform"
                class:-rotate-90={collapsed[type]}
              >
                <CaretDownIcon size="16px" className="!fill-gray-700" />
              </span>
            </Button>

            <h1 class="capitalize font-medium select-none pointer-events-none">
              {type}
            </h1>

            <Button
              type="ghost"
              square
              gray
              noStroke
              label="Add new {type.slice(0, -1)}"
              onClick={() => {
                editingItemData.set({
                  type,
                  index: -1,
                });
              }}
            >
              <PlusIcon size="16px" />
            </Button>
          </header>

          {#if !collapsed[type]}
            <MetricsTable
              selected={selected[type]}
              {type}
              {items}
              {searchValue}
              longest={{ name: longestName, label: longestLabel }}
              editingIndex={$editingItemData?.type === type
                ? $editingItemData.index
                : null}
              {reorderList}
              onDuplicate={duplicateItem}
              onDelete={triggerDelete}
              onEdit={setEditing}
              onCheckedChange={(checked, index) => {
                if (index === undefined) {
                  if (checked) {
                    selected[type] = new Set(
                      Array.from({ length: items.length }, (_, i) => i),
                    );
                  } else {
                    selected[type] = new Set();
                  }
                } else {
                  selected[type][checked ? "add" : "delete"](index);
                  selected[type] = selected[type];
                }
              }}
            />
          {/if}
        </div>
      {/each}
    </div>

    {#if mainError}
      <div
        role="status"
        transition:slide={{ duration: LIST_SLIDE_DURATION }}
        class=" flex items-center gap-x-2 ui-editor-text-error ui-editor-bg-error border border-red-500 border-l-4 px-2 py-5 max-h-40 overflow-auto"
      >
        <CancelCircle />
        {mainError.message}
      </div>
    {/if}
  </div>

  {#if $editingItemData !== null}
    {@const { index, type } = $editingItemData}
    {#key editingItem}
      <VisualMetricsSidebar
        item={editingItem ??
          (type === "measures" ? new YAMLMeasure() : new YAMLDimension())}
        {type}
        {index}
        {columns}
        {fileArtifact}
        editing={index !== -1}
        bind:unsavedChanges
        {switchView}
        {resetEditing}
        onDelete={() => {
          triggerDelete(index, type);
        }}
        onCancel={(unsavedChanges) => {
          if (unsavedChanges) {
            confirmation = {
              action: "cancel",
              index,
              type,
            };
          } else {
            resetEditing();
          }
        }}
      />
    {/key}
  {/if}
</div>

{#if confirmation}
  <AlertConfirmation
    {confirmation}
    onCancel={() => (confirmation = null)}
    onConfirm={async () => {
      if (confirmation?.action === "delete") {
        await deleteItems(
          confirmation?.index !== undefined && confirmation.type
            ? {
                [confirmation.type]: new Set([confirmation.index]),
              }
            : selected,
        );
        resetEditing();
      } else if (confirmation?.action === "switch") {
        await updateProperties(
          {
            model: confirmation.model,
            database: confirmation.database,
            connector: confirmation.connector,
            database_schema: confirmation.schema,
          },
          ["table"],
        );
        resetEditing();
      } else if (confirmation?.action === "cancel") {
        if (
          confirmation?.field &&
          confirmation?.index !== undefined &&
          confirmation?.type
        ) {
          unsavedChanges = false;
          await setEditing(
            confirmation.index,
            confirmation.type,
            confirmation.field,
          );
        }
      }

      confirmation = null;
    }}
  />
{/if}

<style lang="postcss">
  .wrapper {
    @apply size-full max-w-full max-h-full flex-none;
    @apply overflow-hidden;
    @apply flex gap-x-2;
  }

  h1 {
    @apply text-[16px] font-medium;
  }

  .main-area {
    @apply flex flex-col gap-y-4 size-full p-4 bg-surface border;
    @apply flex-shrink overflow-hidden rounded-[2px] relative;
  }

  .section {
    @apply flex flex-none flex-col gap-y-2 justify-start w-full h-fit max-w-full;
  }
</style>
