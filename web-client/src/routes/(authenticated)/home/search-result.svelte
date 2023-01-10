<script lang="ts">
	import type { SearchablePerson } from '$lib/models/person';
	import Quality from './quality.svelte';

	export let item: SearchablePerson;

	let expanded = false;
</script>

<div class="w-full overflow-auto">
	<div class="w-full">
		<a href="/people/{item.uuid}" class="text-2xl text-green-400"
			>{`${item.firstName} ${item.lastName}`.trim()}</a
		>
		<button
			type="button"
			on:click={() => {
				expanded = !expanded;
			}}>V</button
		>
	</div>
	{#if expanded}
		<div class="flex flex-wrap gap-y-1 items-center mt-1">
			{#each item.qualities as quality}
				<Quality item={quality} />
			{/each}
		</div>
	{/if}
</div>
