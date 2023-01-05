<script lang="ts">
	import type { PersonQuality } from '$lib/models/person-quality';
	export let qualities: PersonQuality[] = [];
</script>

<div class="flex flex-wrap gap-y-1 items-center mt-1 w-full">
	<slot />
	{#each qualities as { name, uuid }}
		<a
			href="/person-quality/{uuid}"
			class="text-md text-center text-white bg-green-600 py-1 px-2 mr-2 rounded-full">{name}</a
		>
	{/each}
</div>
