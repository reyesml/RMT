<script lang="ts">
	import type { Note } from '$lib/models/note';
	export let notes: Note[] = [];
</script>

{#if notes}
	<section class="w-full">
		<h2 class="text-xl">Notes:</h2>
		<div class="p-2 bg-black bg-opacity-20 rounded-xl">
			{#each notes as { uuid, title, body, createdAt }}
				<h3 class="text-lg font-bold break-words">{title}</h3>
				<div class="mt-1">
					<div class="whitespace-pre-wrap">
						{body}
					</div>
					<div class="flex text-gray-400">
						<div class="ml-auto">{new Date(createdAt).toLocaleString()}</div>
					</div>
				</div>
			{/each}
		</div>
	</section>
{/if}
