<script lang="ts">
  import { page } from "$app/stores";
  import type { V1OrganizationMemberUser } from "@rilldata/web-admin/client";
  import {
    createAdminServiceAddUsergroupMemberUser,
    createAdminServiceCreateUsergroup,
    getAdminServiceListOrganizationMemberUsergroupsQueryKey,
    getAdminServiceListOrganizationMemberUsersQueryKey,
    getAdminServiceListUsergroupMemberUsersQueryKey,
  } from "@rilldata/web-admin/client";
  import Avatar from "@rilldata/web-common/components/avatar/Avatar.svelte";
  import { Button } from "@rilldata/web-common/components/button/index.js";
  import Combobox from "@rilldata/web-common/components/combobox/Combobox.svelte";
  import {
    Dialog,
    DialogContent,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@rilldata/web-common/components/dialog";
  import Input from "@rilldata/web-common/components/forms/Input.svelte";
  import { eventBus } from "@rilldata/web-common/lib/event-bus/event-bus";
  import { useQueryClient } from "@tanstack/svelte-query";
  import { defaults, superForm } from "sveltekit-superforms";
  import { yup } from "sveltekit-superforms/adapters";
  import { object, string } from "yup";
  import { SLUG_REGEX } from "../constants";

  export let open = false;
  export let groupName: string;
  export let organizationUsers: V1OrganizationMemberUser[] = [];
  export let currentUserEmail: string = "";

  let searchText = "";
  let selectedUsers: V1OrganizationMemberUser[] = [];
  let pendingAdditions: string[] = [];
  let pendingRemovals: string[] = [];

  $: organization = $page.params.organization;

  const queryClient = useQueryClient();
  const createUserGroup = createAdminServiceCreateUsergroup();
  const addUsergroupMemberUser = createAdminServiceAddUsergroupMemberUser();

  async function handleCreate(newName: string) {
    try {
      await $createUserGroup.mutateAsync({
        organization: organization,
        data: {
          name: newName,
        },
      });

      // Apply pending user changes after group creation
      await applyPendingChanges(newName);

      await queryClient.invalidateQueries({
        queryKey: getAdminServiceListOrganizationMemberUsergroupsQueryKey(
          organization,
          {
            includeCounts: true,
          },
        ),
      });

      groupName = "";
      selectedUsers = [];
      pendingAdditions = [];
      pendingRemovals = [];
      open = false;

      eventBus.emit("notification", { message: "User group created" });
    } catch (error) {
      eventBus.emit("notification", {
        message: `Error: ${error.response.data.message}`,
        type: "error",
      });
    }
  }

  async function applyPendingChanges(usergroup: string) {
    try {
      // Add pending users to the group
      for (const email of pendingAdditions) {
        await $addUsergroupMemberUser.mutateAsync({
          organization: organization,
          usergroup: usergroup,
          email: email,
          data: {},
        });
      }

      await queryClient.invalidateQueries({
        queryKey:
          getAdminServiceListOrganizationMemberUsersQueryKey(organization),
      });

      await queryClient.invalidateQueries({
        queryKey: getAdminServiceListUsergroupMemberUsersQueryKey(
          organization,
          usergroup,
        ),
      });

      eventBus.emit("notification", {
        message: "User group changes saved successfully",
      });
    } catch (error) {
      eventBus.emit("notification", {
        message: `Error: ${error.response.data.message}`,
        type: "error",
      });
    }
  }

  async function handleRemove(email: string) {
    selectedUsers = selectedUsers.filter((user) => user.userEmail !== email);
    pendingRemovals = [...pendingRemovals, email];
    pendingAdditions = pendingAdditions.filter((e) => e !== email);
  }

  async function handleAdd(email: string) {
    const user = organizationUsers.find((u) => u.userEmail === email);

    // Don't add if already in selectedUsers
    if (
      user &&
      !selectedUsers.some((selected) => selected.userEmail === email)
    ) {
      selectedUsers = [...selectedUsers, user];
      pendingAdditions = [...pendingAdditions, email];
      pendingRemovals = pendingRemovals.filter((e) => e !== email);
    }
  }

  const formId = "create-user-group-form";

  const initialValues = {
    name: "",
  };

  const schema = yup(
    object({
      name: string()
        .required("Name is required")
        .min(1, "Name must be at least 1 character")
        .max(40, "Name must be at most 40 characters")
        .matches(
          SLUG_REGEX,
          "Name can only include letters, numbers, underscores, and hyphens — no spaces or special characters",
        ),
    }),
  );

  const { form, enhance, submit, errors, submitting } = superForm(
    defaults(initialValues, schema),
    {
      SPA: true,
      validators: schema,
      validationMethod: "oninput",
      async onUpdate({ form }) {
        if (!form.valid) return;
        const values = form.data;

        await handleCreate(values.name);
        open = false;
      },
    },
  );

  $: coercedUsersToOptions = organizationUsers.map((user) => ({
    value: user.userEmail,
    label: user.userName,
  }));

  function getMetadata(email: string) {
    const user = organizationUsers.find((user) => user.userEmail === email);
    return user
      ? { name: user.userName, photoUrl: user.userPhotoUrl }
      : undefined;
  }

  // Check if form has been modified
  $: hasFormChanges = $form.name !== initialValues.name;

  function handleClose() {
    open = false;
    searchText = "";
    selectedUsers = [];
    pendingAdditions = [];
    pendingRemovals = [];
    // Only reset the form if it has been modified
    if (hasFormChanges) {
      $form.name = initialValues.name;
    }
  }
</script>

<Dialog
  bind:open
  onOutsideClick={(e) => {
    e.preventDefault();
  }}
  onOpenChange={(open) => {
    if (!open) {
      handleClose();
    }
  }}
>
  <DialogTrigger asChild>
    <div class="hidden"></div>
  </DialogTrigger>
  <DialogContent class="translate-y-[-200px]">
    <DialogHeader>
      <DialogTitle>Create a group</DialogTitle>
    </DialogHeader>
    <form
      id={formId}
      class="w-full"
      on:submit|preventDefault={submit}
      use:enhance
    >
      <div class="flex flex-col gap-4 w-full">
        <Input
          bind:value={$form.name}
          id="create-user-group-name"
          label="Name"
          placeholder="Untitled"
          errors={$errors.name}
          alwaysShowError={true}
        />

        <div class="flex flex-col gap-y-1">
          <label
            for="user-group-users"
            class="line-clamp-1 text-sm font-medium text-gray-800"
          >
            Users
          </label>
          <Combobox
            searchValue={searchText}
            options={coercedUsersToOptions}
            placeholder="Search for users"
            {getMetadata}
            selectedValues={[
              ...new Set(
                [
                  ...selectedUsers.map((user) => user.userEmail),
                  ...pendingAdditions,
                ].filter((email) => !pendingRemovals.includes(email)),
              ),
            ]}
            onSelectedChange={(values) => {
              if (!values) return;

              const newEmails = values.map((v) => v.value);
              const currentEmails = selectedUsers.map((u) => u.userEmail);

              // Find emails to add (in new but not in current)
              newEmails
                .filter((email) => !currentEmails.includes(email))
                .forEach((email) => handleAdd(email));

              // Find emails to remove (in current but not in new)
              currentEmails
                .filter((email) => !newEmails.includes(email))
                .forEach((email) => handleRemove(email));
            }}
          />
        </div>
      </div>
    </form>

    <div class="flex flex-col gap-2 w-full">
      {#if selectedUsers.length > 0}
        <div class="flex flex-row items-center gap-x-1">
          <div class="text-xs font-semibold uppercase text-gray-500">
            {selectedUsers.length} User{selectedUsers.length === 1 ? "" : "s"}
          </div>
        </div>
      {/if}
      <div class="max-h-[208px] overflow-y-auto">
        <div class="flex flex-col gap-2">
          {#each selectedUsers as user (user.userEmail)}
            <div class="flex flex-row justify-between gap-2 items-center">
              <div class="flex items-center gap-2">
                <Avatar
                  avatarSize="h-7 w-7"
                  alt={user.userName}
                  src={user.userPhotoUrl}
                />
                <div class="flex flex-col text-left">
                  <span class="text-sm font-medium text-gray-900">
                    {user.userName}
                    <span class="text-gray-500 font-normal">
                      {user.userEmail === currentUserEmail ? "(You)" : ""}
                    </span>
                  </span>
                  <span class="text-xs text-gray-500">{user.userEmail}</span>
                </div>
              </div>
              <Button
                type="text"
                danger
                onClick={() => handleRemove(user.userEmail)}
              >
                Remove
              </Button>
            </div>
          {/each}
        </div>
      </div>
    </div>

    <DialogFooter>
      <Button type="plain" onClick={handleClose}>Cancel</Button>
      <Button
        type="primary"
        disabled={$submitting || $form.name.trim() === "" || !!$errors.name}
        form={formId}
        submitForm
      >
        Create
      </Button>
    </DialogFooter>
  </DialogContent>
</Dialog>
