<script lang="ts">
  import Add from "@rilldata/web-common/components/icons/Add.svelte";
  import RillLogoSquareNegative from "@rilldata/web-common/components/icons/RillLogoSquareNegative.svelte";
  import RadixH1 from "@rilldata/web-common/components/typography/RadixH1.svelte";
  import Subheading from "@rilldata/web-common/components/typography/Subheading.svelte";
  import { behaviourEvent } from "../../metrics/initMetrics";
  import {
    BehaviourEventAction,
    BehaviourEventMedium,
  } from "../../metrics/service/BehaviourEventTypes";
  import { MetricsEventSpace } from "../../metrics/service/MetricsTypes";
  import { addSourceModal } from "../sources/modal/add-source-visibility";

  async function openShowAddSourceModal() {
    addSourceModal.open();
    await behaviourEvent?.fireSplashEvent(
      BehaviourEventAction.SourceModal,
      BehaviourEventMedium.Button,
      MetricsEventSpace.Workspace,
    );
  }
</script>

<section class="flex flex-col gap-y-6 items-center text-center">
  <RillLogoSquareNegative size="84px" gradient />
  <RadixH1>
    <span
      class="bg-gradient-to-r from-primary-900 to-primary-800 text-transparent bg-clip-text opacity-75"
    >
      Welcome to Rill
    </span>
  </RadixH1>
  <div class="flex flex-col gap-y-2">
    <Subheading twColor="text-slate-600">
      Build fast operational dashboards that your team will actually use.
    </Subheading>
  </div>
  <button
    class="pl-2 pr-4 py-2 rounded-sm bg-gradient-to-b from-primary-400 to-primary-500 hover:from-primary-500 hover:to-primary-500"
    on:click={openShowAddSourceModal}
  >
    <div
      class="flex flex-row gap-x-1 items-center text-sm font-medium text-white"
    >
      <Add className="text-white" />
      Connect your data
    </div>
  </button>
</section>
