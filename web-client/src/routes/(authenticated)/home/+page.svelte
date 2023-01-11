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
		keys: ['firstName', 'lastName', 'qualities.name'],
	};
	console.log(data.people?.length);
	const fuse = new Fuse(data.people || [], options);
	$: if(data.people) fuse.setCollection(data.people);
	$: if(query || data.people) {
		results = fuse.search(query).map((val) => val.item);
	}
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
		<div class="text-lg">The impossible is unknown.</div>
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
		{#if results}
		<div class="flex flex-col divide-y mt-6 w-full">
			{#each results as result}
				<SearchResult item={result} />
			{/each}
		</div>
		{/if}
	</div>
</section>
