<script lang="ts">
  import SimpleDataGraphic from "@rilldata/web-common/components/data-graphic/elements/SimpleDataGraphic.svelte";
  import { HistogramPrimitive } from "@rilldata/web-common/components/data-graphic/marks";
  import { ScaleType } from "@rilldata/web-common/components/data-graphic/state";
  import Tooltip from "@rilldata/web-common/components/tooltip/Tooltip.svelte";
  import TooltipContent from "@rilldata/web-common/components/tooltip/TooltipContent.svelte";
  import { COLUMN_PROFILE_CONFIG } from "@rilldata/web-common/layout/config";
  import { INTEGERS } from "@rilldata/web-common/lib/duckdb-data-types";
  import type { NumericHistogramBinsBin } from "@rilldata/web-common/runtime-client";

  export let compact = false;
  export let data: NumericHistogramBinsBin[];
  export let type: string;

  $: summaryWidthSize =
    COLUMN_PROFILE_CONFIG.summaryVizWidth[compact ? "small" : "medium"];
</script>

{#if data}
  <Tooltip location="right" distance={8}>
    <SimpleDataGraphic
      xType={ScaleType.NUMBER}
      yType={ScaleType.NUMBER}
      yMin={0}
      width={summaryWidthSize}
      height={18}
      bodyBuffer={0}
      marginBuffer={0}
      left={1}
      right={1}
      top={4}
      bottom={1}
      let:config
    >
      <g class="text-red-200">
        <line
          x1={config.plotLeft}
          x2={config.plotRight}
          y1={config.plotBottom}
          y2={config.plotBottom}
          stroke="currentColor"
          stroke-width={0.5}
        />
      </g>
      <g class="text-red-300">
        <HistogramPrimitive
          {data}
          xLowAccessor="low"
          xHighAccessor="high"
          yAccessor="count"
          lineThickness={0.5}
          separator={data.length < 20 && INTEGERS.has(type) ? 0.25 : 0}
          color="currentColor"
          stopOpacity={0.5}
        />
      </g>
    </SimpleDataGraphic>
    <TooltipContent slot="tooltip-content">
      the distribution of the values of this column
    </TooltipContent>
  </Tooltip>
{/if}
