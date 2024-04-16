import { multiLayerBaseSpec } from "./utils";

export function buildStackedArea(
  timeField: string,
  quantitativeField: string,
  nominalField: string,
) {
  const baseSpec = multiLayerBaseSpec();

  baseSpec.encoding = {
    x: { field: timeField, type: "temporal" },
    y: {
      field: quantitativeField,
      type: "quantitative",
      stack: "zero",
    },
    color: { field: nominalField, type: "nominal", legend: null },
  };
  baseSpec.layer = [
    {
      mark: { type: "area", clip: true },
    },
    {
      mark: { type: "line", strokeWidth: 1, clip: true },
      encoding: {
        stroke: { field: nominalField, type: "nominal", legend: null },
      },
    },
    {
      mark: { type: "rule", color: "transparent", clip: true },
      encoding: {
        color: { value: "transparent" },
        tooltip: [
          { field: quantitativeField, type: "quantitative" },
          { field: "ts", type: "temporal", title: "Time" },
          { field: nominalField, type: "nominal" },
        ],
      },
      params: [
        {
          name: "x_hover",
          select: {
            type: "point",
            encodings: ["x"],
            nearest: true,
            on: "pointerover",
            clear: "pointerout",
          },
        },
      ],
    },
    {
      transform: [{ filter: { param: "x-hover", empty: false } }],
      mark: { type: "point", filled: true, opacity: 1, size: 40, clip: true },
    },
  ];

  return baseSpec;
}