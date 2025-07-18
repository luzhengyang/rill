<script lang="ts">
  import type { NumberParts } from "@rilldata/web-common/lib/number-formatting/humanizer-types";
  import Base from "./Base.svelte";
  import { PERC_DIFF, isPercDiff } from "./type-utils";

  export let isNull = false;
  export let inTable = false;
  export let showPosSign = false;
  export let color = "!text-gray-500";
  export let customStyle = "";
  export let value:
    | string
    | number
    | undefined
    | null
    | NumberParts
    | PERC_DIFF;
  export let tabularNumber = true;
  export let assembled = true;

  let diffIsNegative = false;
  let intValue: string;
  let negSign = "";
  let posSign = "";
  let approxSign = "";
  let suffix = "";

  $: isNoData = isPercDiff(value) || value === null || value === undefined;

  $: if (
    !isNoData &&
    // expanding this out in full provides type narrowing
    !isPercDiff(value) &&
    value !== null &&
    value !== undefined &&
    typeof value !== "number"
  ) {
    // in this case, we have a NumberParts object.
    // We have a couple cases to consider:
    // * If the NumberParts object has approxZero===true,
    // we want to show e.g. "~0%" WITHOUT a negative sign
    // * However, in this case we show the number in red to indicate a
    // small negative change.
    //
    // Otherwise, we format the number as usual.
    let intPart = +value.int;
    let fracPart = +value.frac / 10 ** value.frac.length;
    intValue = Math.round(intPart + fracPart).toString();

    diffIsNegative = value?.neg === "-";
    negSign = diffIsNegative && !value?.approxZero ? "-" : "";
    approxSign = value?.approxZero ? "~" : "";
    posSign = !diffIsNegative && !approxSign && showPosSign ? "+" : "";
    suffix = value?.suffix ?? "";
  } else if (typeof value === "number") {
    // FIXME: this seems to only come up in the tool tip,
    // for percentages in the dimension table,
    // but this whole thing is a mess and needs to be cleaned up.

    diffIsNegative = value < 0;
    intValue = Math.round(100 * value).toString();
    approxSign = Math.abs(value) < 0.005 ? "~" : "";
    posSign = !diffIsNegative && !approxSign && showPosSign ? "+" : "";
    negSign = "";
    suffix = "";
  }
</script>

<!-- FIXME: !color to override the .ui-copy class -->
<Base
  {isNull}
  {color}
  classes="{tabularNumber
    ? 'ui-copy-number'
    : ''} w-full {customStyle} {inTable &&
    'block text-right'} pointer-events-none"
>
  <slot name="value">
    {#if isNoData}
      <span class="text-gray-400">-</span>
    {:else if value !== null && assembled}
      <span class:text-red-500={diffIsNegative}>
        {approxSign}{negSign}{posSign}{intValue}{suffix}<span class="opacity-50"
          >%</span
        >
      </span>
    {/if}
  </slot>
</Base>
