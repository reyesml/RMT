<script lang="ts">
	import type { Journal } from '$lib/models/journal';
	export let journals: Journal[] = [];
</script>

{#if journals}
	<div class="flex flex-col divide-y">
		{#each journals as { mood, title, createdAt, uuid: UUID }}
			<div class="py-2 px-2 w-full overflow-auto truncate">
				<a href="/journal/{UUID}" class="text-2xl text-green-400 ">{title}</a>
				<div class="flex ">
					<div>{mood}</div>
					<div class="ml-auto">{new Date(createdAt).toLocaleDateString()}</div>
				</div>
			</div>
		{/each}
	</div>
{/if}
