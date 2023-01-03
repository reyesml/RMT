<script lang="ts">
	import type { PersonQuality } from '$lib/models/person-quality';
	export let qualities: PersonQuality[] = [];
</script>

{#if qualities}
	<section>
		<h2 class="text-xl">Qualities:</h2>
		<div class="flex items-center">
			{#each qualities as { name, uuid }}
				<a href="/person-quality/{uuid}" class="text-md text-center text-white bg-green-600 py-1 px-2 ml-2 rounded-full"
					>{name}</a
				>
			{/each}
		</div>
	</section>
{/if}
