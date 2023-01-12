<script lang="ts">
	import MagIcon from '$lib/components/icons/mag-icon.svelte';
	import PlusIcon from '$lib/components/icons/plus-icon.svelte';
	import type { Person } from '$lib/models/person';
	import Fuse from 'fuse.js';
	import type { PageData } from './$types';
	import PeopleList from './people-list.svelte';

	export let data: PageData;
	let query: string = '';
	let results: Person[] = [];
	const options = {
		keys: ['firstName', 'lastName']
	};

	const fuse = new Fuse(data.people || [], options);
	$: if (data.people) {
		results = fuse.search(query).map((val) => val.item);
	}
</script>

<svelte:head>
	<title>RMT: People</title>
	<meta name="description" content="RMT People" />
</svelte:head>

<section class="flex min-h-full items-center justify-center py-12 px-4">
	<div class="w-full max-w-xl">
		<h1>People</h1>
		<div class="relative">
			<div class="absolute -left-20 h-full">
				<a
					href="/people/new"
					class="sticky top-5 mt-5 text-white bg-purple-700 rounded-full w-14 h-14 flex items-center justify-center"
				>
					<PlusIcon class="h-8 w-8" />
				</a>
			</div>
			{#if data.error}
				Something went wrong.
			{:else if data.people}
				<div
					class="rounded-full w-full h-12 bg-white mt-12 flex items-center py-3 px-3 border-4 border-transparent focus-within:border-indigo-500 focus-within:outline-none focus-within:ring-sky-500"
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
				<div class="mt-6">
					{#if query.length === 0}
						<PeopleList people={data.people} />
					{:else if results.length > 0}
						<PeopleList people={results} />
					{:else}
						<div class="flex justify-center">No results 😕</div>
					{/if}
				</div>
			{/if}
		</div>
	</div>
</section>
