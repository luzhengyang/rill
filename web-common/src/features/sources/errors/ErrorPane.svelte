<script lang="ts">
  import CancelCircle from "../../../components/icons/CancelCircle.svelte";
  import { runtime } from "../../../runtime-client/runtime-store";
  import { humanReadableErrorMessage } from "./errors";
  import { useSourceFromYaml } from "../selectors";

  export let filePath: string;
  export let errorMessage: string;

  $: ({ instanceId } = $runtime);

  // Parse Source YAML client-side
  $: sourceFromYaml = useSourceFromYaml(instanceId, filePath);

  // Try to extract the connector type
  $: connectorType = $sourceFromYaml.data?.type;

  // Try to create an actionable error message
  $: prettyMessage = humanReadableErrorMessage(connectorType, 3, errorMessage);
</script>

<div class="w-full h-full bg-surface flex-col justify-center inline-flex p-3">
  <div class="flex-col justify-start items-center gap-1 flex text-red-500">
    <CancelCircle size="24px" />
    <div class="text-center text-sm font-medium">
      {prettyMessage ?? errorMessage}
    </div>
  </div>
</div>
