<script lang="ts">
  import { getRandomBgColor } from "@rilldata/web-common/features/themes/color-config";
  import { cn } from "@rilldata/web-common/lib/shadcn";

  export let name: string;
  export let usersCount: number;

  function getInitials(name: string) {
    return name.charAt(0).toUpperCase();
  }
</script>

<div class="flex items-center gap-2 py-2 pl-2">
  <div
    class={cn(
      "h-7 w-7 rounded-sm flex items-center justify-center",
      getRandomBgColor(name),
    )}
  >
    <span class="text-sm text-white font-semibold">{getInitials(name)}</span>
  </div>
  <div class="flex flex-col text-left">
    <span class="text-sm font-medium text-gray-900 flex flex-row gap-x-1">
      {name}
    </span>
    <div class="flex flex-row items-center gap-x-1">
      <span class="text-xs text-gray-500">
        {usersCount} user{usersCount > 1 ? "s" : ""}
      </span>
    </div>
  </div>
</div>
