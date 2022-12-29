<script lang="ts">
	import { getSessionStore } from '$lib/store/session';
	let session = getSessionStore();
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="RMT Welcome" />
</svelte:head>

<section class="text-purple-800">
	<h1>
		Welcome to RMT, {$session.user.username}
	</h1>

	<div class="flex flex-col items-center">
		<h2>
			You can do anything. The only limit is <span class="italic underline">yourself</span>.
		</h2>
	</div>
</section>
