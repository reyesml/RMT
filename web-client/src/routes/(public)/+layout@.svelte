<script>
	import Session from '$lib/components/session.svelte';
</script>

<Session mandatory={false}>
	<div class="app">
		<main>
			<slot />
		</main>

		<footer />
	</div>
</Session>
