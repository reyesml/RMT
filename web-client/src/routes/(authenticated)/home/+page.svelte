<script lang="ts">
	import MagIcon from '$lib/components/icons/mag-icon.svelte';
	import { getSessionStore } from '$lib/store/session';
	let session = getSessionStore();
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="RMT Welcome" />
</svelte:head>

<section class="mt-40">
	<h1>
		Welcome, {$session.user.username} 👋
	</h1>

	<div class="w-full h-full flex flex-col justify-center items-center">
		<div class="text-lg">The impossible is unknown.</div>
		<div class="rounded-full w-1/2 h-16 bg-white mt-12 flex items-center py-3 px-3 border-4 border-transparent focus-within:border-indigo-500 focus-within:outline-none focus-within:ring-sky-500">
			<MagIcon class="w-8 h-8 text-gray-600" />
			<form class="flex-1 h-full flex items-center">
				<label for="search" class="sr-only">Search</label>
				<input type="text" name="search" id="search" class="w-full h-full bg-inherit text-black text-2xl ml-4 focus:outline-none" placeholder="search..."/>
			</form>
		</div>
	</div>
</section>
