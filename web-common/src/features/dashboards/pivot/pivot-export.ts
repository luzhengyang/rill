import { mergeDimensionAndMeasureFilters } from "@rilldata/web-common/features/dashboards/filters/measure-filters/measure-filter-utils";
import { sanitiseExpression } from "@rilldata/web-common/features/dashboards/stores/filter-utils";
import type { ExploreState } from "@rilldata/web-common/features/dashboards/stores/explore-state";
import { useTimeControlStore } from "@rilldata/web-common/features/dashboards/time-controls/time-control-store";
import { mapSelectedTimeRangeToV1TimeRange } from "@rilldata/web-common/features/dashboards/time-controls/time-range-mappers";
import { type TimeRangeString } from "@rilldata/web-common/lib/time/types";
import {
  V1TimeGrain,
  type V1MetricsViewAggregationRequest,
  type V1MetricsViewAggregationSort,
  type V1Query,
  type V1TimeRange,
} from "@rilldata/web-common/runtime-client";
import { get } from "svelte/store";
import { runtime } from "../../../runtime-client/runtime-store";
import type { StateManagers } from "../state-managers/state-managers";
import { getPivotConfig } from "./pivot-data-config";
import { prepareMeasureForComparison } from "./pivot-utils";
import {
  COMPARISON_DELTA,
  COMPARISON_PERCENT,
  PivotChipType,
  type PivotChipData,
  type PivotState,
} from "./types";

export function getPivotExportQuery(ctx: StateManagers, isScheduled: boolean) {
  const metricsViewName = get(ctx.metricsViewName);
  const validSpecStore = get(ctx.validSpecStore);
  const timeControlState = get(useTimeControlStore(ctx));
  const dashboardState = get(ctx.dashboardStore);
  const configStore = get(getPivotConfig(ctx));
  const rows = get(ctx.selectors.pivot.rows);
  const columns = get(ctx.selectors.pivot.columns);

  if (!validSpecStore.data?.explore || !timeControlState.ready)
    return undefined;

  const enableComparison = configStore.enableComparison;
  const isFlat = configStore.isFlat;
  const comparisonTime = configStore.comparisonTime;
  const pivotState = configStore.pivot;

  const metricsViewSpec = validSpecStore.data?.metricsView ?? {};
  const exploreSpec = validSpecStore.data?.explore ?? {};

  let timeRange: V1TimeRange | undefined;
  if (isScheduled) {
    timeRange = mapSelectedTimeRangeToV1TimeRange(
      timeControlState.selectedTimeRange,
      dashboardState.selectedTimezone,
      exploreSpec,
    );
  } else {
    timeRange = {
      start: timeControlState.timeStart,
      end: timeControlState.timeEnd,
    };
  }

  if (!timeRange) return undefined;

  const query: V1Query = {
    metricsViewAggregationRequest: getPivotAggregationRequest(
      metricsViewName,
      metricsViewSpec.timeDimension ?? "",
      dashboardState,
      timeRange,
      rows,
      columns,
      enableComparison,
      comparisonTime,
      isFlat,
      pivotState,
    ),
  };

  return query;
}

function getPivotAggregationRequest(
  metricsView: string,
  timeDimension: string,
  exploreState: ExploreState,
  timeRange: V1TimeRange,
  rows: PivotChipData[],
  columns: { dimension: PivotChipData[]; measure: PivotChipData[] },
  enableComparison: boolean,
  comparisonTime: TimeRangeString | undefined,
  isFlat: boolean,
  pivotState: PivotState,
): undefined | V1MetricsViewAggregationRequest {
  const measures = columns.measure.flatMap((m) => {
    const measureName = m.id;
    const group = [{ name: measureName }];

    if (enableComparison) {
      group.push(
        { name: `${measureName}${COMPARISON_DELTA}` },
        { name: `${measureName}${COMPARISON_PERCENT}` },
      );
    }

    return group;
  });

  const allDimensions = [...rows, ...columns.dimension].map((d) =>
    d.type === PivotChipType.Time
      ? {
          name: timeDimension,
          timeGrain: d.id as V1TimeGrain,
          timeZone: exploreState.selectedTimezone,
          alias: isFlat ? `${timeDimension}_rill_${d.id}` : `Time ${d.title}`,
        }
      : {
          name: d.id,
        },
  );

  const pivotOn = isFlat
    ? undefined
    : columns.dimension.map((d) =>
        d.type === PivotChipType.Time ? `Time ${d.title}` : d.id,
      );

  const rowDimensions = rows.map((d) =>
    d.type === PivotChipType.Time
      ? {
          name: timeDimension,
          timeGrain: d.id as V1TimeGrain,
          timeZone: exploreState.selectedTimezone,
          alias: `Time ${d.title}`,
        }
      : {
          name: d.id,
        },
  );

  let sort: V1MetricsViewAggregationSort[] = [];

  if (isFlat) {
    if (pivotState.sorting.length > 0) {
      sort = [
        {
          name: pivotState.sorting[0].id,
          desc: pivotState.sorting[0].desc,
        },
      ];
    } else {
      sort = [
        {
          desc: measures?.[0] ? true : false,
          name: measures?.[0]?.name || allDimensions?.[0]?.name,
        },
      ];
    }
  } else {
    // Sort by the dimensions in the pivot's rows
    sort = rowDimensions.map((d) => {
      return {
        name: d.alias ? d.alias : d.name,
        desc: pivotState.sorting.find((s) => s.id === d.name)?.desc ?? false,
      };
    });
  }

  return {
    instanceId: get(runtime).instanceId,
    metricsView,
    timeRange,
    comparisonTimeRange: comparisonTime,
    measures: enableComparison
      ? prepareMeasureForComparison(measures)
      : measures,
    dimensions: allDimensions,
    where: sanitiseExpression(
      mergeDimensionAndMeasureFilters(
        exploreState.whereFilter,
        exploreState.dimensionThresholdFilters,
      ),
      undefined,
    ),
    pivotOn,
    sort,
    offset: "0",
  };
}
