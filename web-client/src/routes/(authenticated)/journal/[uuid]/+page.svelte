<script lang="ts">
	import type { Journal } from '$lib/models/journal';
	import type { PageData } from './$types';

	export let data: PageData;
	let journal: Journal;
	$: journal = data.journal;
</script>

<svelte:head>
	<title>RMT: Journal</title>
	<meta name="description" content="RMT Journal" />
</svelte:head>

<section class="relative">
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl">
			{#if data.error}
				<h1>{data.error}</h1>
			{:else if journal}
				<h1 class="text-4xl break-words">{journal.title}</h1>
                <div class="bg-black bg-opacity-20 rounded-3xl p-4 mt-6">
                    <div class="whitespace-pre-wrap min-h-[300px]">
                        {journal.body}
                    </div>
                    <div class="flex text-gray-400">
                        <div>
                            {journal.mood}
                        </div>
                        <div class="ml-auto">{new Date(journal.createdAt).toLocaleString()}</div>
                    </div>
                </div>
			{/if}
		</div>
	</div>
</section>
