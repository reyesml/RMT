<script lang="ts">
	import PlusIcon from '$lib/components/icons/plus-icon.svelte';
	import type { PageData } from './$types';
	import PeopleList from './people-list.svelte';

	export let data: PageData;
</script>

<svelte:head>
	<title>RMT: People</title>
	<meta name="description" content="RMT People" />
</svelte:head>

<section class="relative">
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl">
			<h1>People</h1>
			{#if data.error}
				Something went wrong.
			{:else if data.people}
				<PeopleList people={data.people} />
			{/if}
		</div>
	</div>
	<a
		href="/people/new"
		class="fixed bottom-20 left-64 text-white bg-purple-700 rounded-full w-14 h-14 flex items-center justify-center"
	>
		<PlusIcon class="h-8 w-8" />
	</a>
</section>
