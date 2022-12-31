<script lang="ts">
	import JournalList from "./journal-list.svelte";

</script>

<svelte:head>
	<title>RMT: Journal</title>
	<meta name="description" content="RMT Journal" />
</svelte:head>

<section>
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl space-y-8">
			<h1>Journal</h1>
			<JournalList />
		</div>
	</div>
</section>
