<script lang="ts">
  import { goto } from "$app/navigation";
  import CtaButton from "@rilldata/web-common/components/calls-to-action/CTAButton.svelte";
  import CtaContentContainer from "@rilldata/web-common/components/calls-to-action/CTAContentContainer.svelte";
  import CtaLayoutContainer from "@rilldata/web-common/components/calls-to-action/CTALayoutContainer.svelte";
  import CtaMessage from "@rilldata/web-common/components/calls-to-action/CTAMessage.svelte";
  import type { ExploreState } from "@rilldata/web-common/features/dashboards/stores/explore-state";
  import Spinner from "@rilldata/web-common/features/entity-management/Spinner.svelte";
  import { EntityStatus } from "@rilldata/web-common/features/entity-management/types";
  import { mapQueryToDashboard } from "@rilldata/web-common/features/explore-mappers/map-to-explore";
  import { getExplorePageUrlSearchParams } from "@rilldata/web-common/features/explore-mappers/utils";
  import type { PageData } from "./$types";

  export let data: PageData;

  $: ({
    report: reportResource,
    organization,
    project,
    reportId,
    executionTime,
    token,
    exploreName,
  } = data);

  let dashboardStateForReport: ReturnType<typeof mapQueryToDashboard>;
  $: dashboardStateForReport = mapQueryToDashboard(
    exploreName,
    reportResource?.report?.spec?.queryName,
    reportResource?.report?.spec?.queryArgsJson,
    executionTime,
    reportResource?.report?.spec?.annotations ?? {},
  );

  $: if ($dashboardStateForReport?.data) {
    void gotoExplorePage(
      $dashboardStateForReport.data.exploreName,
      $dashboardStateForReport.data.exploreState,
    );
  }

  async function gotoExplorePage(
    exploreName: string,
    exploreState: ExploreState,
  ) {
    const url = new URL(window.location.origin);
    if (token) {
      url.pathname = `/${organization}/${project}/-/share/${token}/explore/${exploreName}`;
    } else {
      url.pathname = `/${organization}/${project}/explore/${exploreName}`;
    }

    const exploreStateParams = await getExplorePageUrlSearchParams(
      exploreName,
      exploreState,
    );

    url.search = exploreStateParams.toString();

    return goto(url.toString());
  }
</script>

<CtaLayoutContainer>
  <CtaContentContainer>
    {#if $dashboardStateForReport.isLoading}
      <div class="h-36 mt-10">
        <Spinner status={EntityStatus.Running} size="7rem" duration={725} />
      </div>
    {:else if $dashboardStateForReport.error}
      <div class="flex flex-col gap-y-2">
        <h2 class="text-lg font-semibold">Unable to open report</h2>
        <CtaMessage>
          {$dashboardStateForReport.error}
        </CtaMessage>
      </div>
      <CtaButton
        href={`/${organization}/${project}/-/reports/${reportId}`}
        variant="secondary"
      >
        Go to report page
      </CtaButton>
    {/if}
  </CtaContentContainer>
</CtaLayoutContainer>
