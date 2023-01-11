<script lang="ts">
	import ChevronDown from '$lib/components/icons/chevron-down.svelte';
	import ChevronUp from '$lib/components/icons/chevron-up.svelte';
	import type { Person } from '$lib/models/person';
	import Quality from './quality.svelte';

	export let item: Person;

	let expanded = false;
</script>

<div class="w-full overflow-auto p-2">
	<div class="flex items-center">
		<div class="flex-1">
			<a href="/people/{item.uuid}" class="text-2xl text-green-400"
				>{`${item.firstName} ${item.lastName}`.trim()}</a
			>
		</div>
		<button
			type="button"
			on:click={() => {
				expanded = !expanded;
			}}
		>
			{#if expanded}
				<ChevronUp class="w-6 h-6" />
			{:else}
				<ChevronDown class="w-6 h-6" />
			{/if}
		</button>
	</div>
	{#if expanded}
		<div class="flex flex-wrap gap-y-1 items-center mt-1">
			{#if !item.qualities || item.qualities.length === 0}
				<Quality item={{ name: 'n/a', type: 'default' }} />
			{:else}
				{#each item.qualities as quality}
					<Quality item={quality} />
				{/each}
			{/if}
		</div>
	{/if}
</div>
