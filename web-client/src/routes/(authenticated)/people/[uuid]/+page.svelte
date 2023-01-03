<script lang="ts">
	import type { Person } from '$lib/models/person';
	import type { PageData } from './$types';

	export let data: PageData;
	let person: Person;
	$: person = data.person;
</script>

<svelte:head>
	<title>RMT: Person</title>
	<meta name="description" content="RMT Person" />
</svelte:head>

<section class="relative">
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl">
			{#if data.error}
				<h1>{data.error}</h1>
			{:else if person}
				<h1 class="text-4xl break-words">{`${person.firstName} ${person.lastName}`.trim()}</h1>
				<div class="mt-6">Qualities list here</div>
				<div class="mt-6">Notes list here</div>
			{/if}
		</div>
	</div>
</section>
