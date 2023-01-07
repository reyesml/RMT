<script lang="ts">
	import type { Note } from '$lib/models/note';
	export let notes: Note[] = [];
</script>

{#if notes && notes.length > 0}
	{#each notes as { uuid, title, body, createdAt }}
		<div class="mt-2 p-2 bg-black bg-opacity-20 rounded-xl">
			<h3 class="text-lg font-bold break-words">{title}</h3>
			<div class="mt-1">
				<div class="whitespace-pre-wrap">
					{body}
				</div>
				<div class="flex text-gray-400">
					<div class="ml-auto">{new Date(createdAt).toLocaleString()}</div>
				</div>
			</div>
		</div>
	{/each}
{:else}
	<div>no notes</div>
{/if}
