<script lang="ts">
  import OnboardingWorkspace from "@rilldata/web-common/features/onboarding/OnboardingWorkspace.svelte";
  import ProjectCards from "@rilldata/web-common/features/welcome/ProjectCards.svelte";
  import TitleContent from "@rilldata/web-common/features/welcome/TitleContent.svelte";
  import { fly } from "svelte/transition";
  import type { LayoutData } from "../$types";

  export let data: LayoutData;
</script>

<svelte:head>
  <title>Rill Developer</title>
</svelte:head>

{#if data.initialized}
  <OnboardingWorkspace />
{:else}
  <div class="scroll" in:fly={{ duration: 1600, delay: 400, y: 8 }}>
    <div class="wrapper column p-10 2xl:py-16">
      <TitleContent />
      <div class="column" in:fly={{ duration: 1600, delay: 1200, y: 4 }}>
        <ProjectCards />
      </div>
    </div>
  </div>
{/if}

<style lang="postcss">
  .scroll {
    @apply size-full overflow-x-hidden overflow-y-auto;
  }

  .wrapper {
    @apply w-full h-fit min-h-screen bg-no-repeat bg-cover;
    background-image: url("/img/welcome-bg-art.png");
  }

  .column {
    @apply flex flex-col items-center gap-y-6;
  }
</style>
