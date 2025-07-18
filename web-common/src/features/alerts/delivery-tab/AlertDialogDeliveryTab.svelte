<script lang="ts">
  import FormSection from "@rilldata/web-common/components/forms/FormSection.svelte";
  import Input from "@rilldata/web-common/components/forms/Input.svelte";
  import MultiInput from "@rilldata/web-common/components/forms/MultiInput.svelte";
  import Select from "@rilldata/web-common/components/forms/Select.svelte";
  import { getHasSlackConnection } from "@rilldata/web-common/features/alerts/delivery-tab/notifiers-utils";
  import { SnoozeOptions } from "@rilldata/web-common/features/alerts/delivery-tab/snooze";
  import type { AlertFormValues } from "@rilldata/web-common/features/alerts/form-utils";
  import { runtime } from "@rilldata/web-common/runtime-client/runtime-store";
  import type { SuperForm } from "sveltekit-superforms/client";

  export let superFormInstance: SuperForm<AlertFormValues>;

  $: ({ form, errors } = superFormInstance);

  $: ({ instanceId } = $runtime);

  $: hasSlackNotifier = getHasSlackConnection(instanceId);
</script>

<div class="flex flex-col gap-y-3">
  <FormSection title="Alert name">
    <Input
      alwaysShowError
      errors={$errors["name"]}
      id="name"
      title="Alert name"
      placeholder="My alert"
      bind:value={$form["name"]}
    />
  </FormSection>
  <FormSection
    description="We'll check for this alert whenever the data refreshes"
    title="Trigger"
  />
  <FormSection
    description="Set a snooze period to silence repeat notifications for the same alert."
    title="Snooze"
  >
    <Select
      bind:value={$form["snooze"]}
      id="snooze"
      label=""
      options={SnoozeOptions}
    />
  </FormSection>
  {#if $hasSlackNotifier.data}
    <FormSection
      bind:enabled={$form["enableSlackNotification"]}
      showSectionToggle
      title="Slack notifications"
    >
      <MultiInput
        id="slackChannels"
        placeholder="# Enter a Slack channel name"
        description="We’ll send alerts directly to these channels."
        contentClassName="relative"
        bind:values={$form.slackChannels}
        errors={$errors.slackChannels}
        singular="channel"
        plural="channels"
        preventFocus={true}
      />
      <MultiInput
        id="slackUsers"
        placeholder="Enter an email address"
        description="We’ll alert them with direct messages in Slack."
        contentClassName="relative"
        bind:values={$form.slackUsers}
        errors={$errors.slackUsers}
        singular="user"
        plural="users"
        preventFocus={true}
      />
    </FormSection>
  {:else}
    <FormSection title="Slack notifications">
      <svelte:fragment slot="description">
        <span class="text-sm text-slate-600">
          Slack has not been configured for this project. Read the <a
            href="https://docs.rilldata.com/explore/alerts/slack"
            target="_blank"
          >
            docs
          </a> to learn more.
        </span>
      </svelte:fragment>
    </FormSection>
  {/if}
  <FormSection
    bind:enabled={$form["enableEmailNotification"]}
    description="We’ll email alerts to these addresses. Make sure they have access to your project."
    showSectionToggle
    title="Email notifications"
  >
    <MultiInput
      id="slackUsers"
      placeholder="Enter an email address"
      contentClassName="relative"
      bind:values={$form.emailRecipients}
      errors={$errors.emailRecipients}
      singular="email"
      plural="emails"
      preventFocus={true}
    />
  </FormSection>
</div>
