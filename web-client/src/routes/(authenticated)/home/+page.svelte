<script lang="ts">
	import MagIcon from '$lib/components/icons/mag-icon.svelte';
	import { getSessionStore } from '$lib/store/session';
	import type { PageData } from './$types';
	import Fuse from 'fuse.js';
	import type { Person } from '$lib/models/person';
	import SearchResult from './search-result.svelte';
	let session = getSessionStore();

	export let data: PageData;

	let query: string = '';
	let results: Person[] = [];

	const options = {
		keys: [{ name: 'firstName', weight: 6 }, { name: 'lastName', weight: 6 }, 'qualities.name']
	};
	const fuse = new Fuse(data.people || [], options);
	$: if (data.people) fuse.setCollection(data.people);
	$: if (query || data.people) {
		results = fuse.search(query).map((val) => val.item);
	}

	const greetings = [
		'Welcome to RMT!',
		'This is RMT.',
		'The infinite is possible at RMT.',
		'You can do anything at RMT, anything at all.',
		'The only limit is yourself.',
		'This is RMT, welcome!',
		'The unobtainable is unknown at RMT.'
	];
	let greeting = greetings[Math.floor(Math.random() * greetings.length)];
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="RMT Home" />
</svelte:head>

<section class="mt-40 flex flex-col items-center">
	<h1>
		Welcome, {$session?.user.username ?? 'friend'} 👋
	</h1>

	<div class="w-full max-w-xl h-full flex flex-col justify-center items-center">
		<div class="text-lg">{greeting}</div>
		<div
			class="rounded-full w-full h-16 bg-white mt-12 flex items-center py-3 px-3 border-4 border-transparent focus-within:border-indigo-500 focus-within:outline-none focus-within:ring-sky-500"
		>
			<MagIcon class="w-8 h-8 text-gray-600" />
			<label for="search" class="sr-only">Search</label>
			<input
				bind:value={query}
				type="text"
				name="search"
				id="search"
				class="w-full h-full bg-inherit text-black text-2xl ml-4 focus:outline-none"
				placeholder="search..."
			/>
		</div>
		{#if results && results.length > 0}
			<div class="flex flex-col divide-y mt-6 w-full">
				{#each results as result}
					<SearchResult item={result} />
				{/each}
			</div>
		{:else if query.length > 0}
			<div class="mt-6">No results 😕</div>
		{/if}
	</div>
</section>
