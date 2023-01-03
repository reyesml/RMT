<script lang="ts">
	import type { Person } from "$lib/models/person";
  export let people: Person[] = [];
</script>

{#if people}
<div class="flex flex-col divide-y">
  {#each people as { firstName, lastName, uuid }}
    <div class="py-2 px-2 w-full overflow-auto truncate">
      <a href="/person/{uuid}" class="text-2xl text-green-400 ">{`${firstName} ${lastName}`.trim()}</a>
    </div>
  {/each}
</div>
{/if}