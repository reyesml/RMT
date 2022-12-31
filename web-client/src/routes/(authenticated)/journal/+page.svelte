<script lang="ts">
	import PlusIcon from '$lib/components/icons/plus-icon.svelte';
	import JournalList from './journal-list.svelte';
</script>

<svelte:head>
	<title>RMT: Journal</title>
	<meta name="description" content="RMT Journal" />
</svelte:head>

<section class="relative">
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl space-y-8">
			<h1>Journal</h1>
			<JournalList />
		</div>
	</div>
	<a
		href="/journal/new"
		class="fixed bottom-20 left-64 text-white bg-purple-700 rounded-full w-14 h-14 flex items-center justify-center"
	>
		<PlusIcon class="h-8 w-8" />
	</a>
</section>
