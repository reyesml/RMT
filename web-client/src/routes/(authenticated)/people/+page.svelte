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

<section>
	<div class="flex min-h-full items-center justify-center py-12 px-4">
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
					<PeopleList people={data.people} />
				{/if}
			</div>
		</div>
	</div>
</section>
