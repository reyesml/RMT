<script>
	import Session from '$lib/components/session.svelte';
	import '../../app.postcss';
	import '../styles.css';
</script>

<Session mandatory={false}>
	<div class="app">
		<main>
			<slot />
		</main>

		<footer />
	</div>
</Session>
