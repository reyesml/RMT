<script lang="ts">
	import '../../app.postcss';
	import '../styles.css';
	import Session from '$lib/components/session.svelte';
</script>

<Session>
	<div class="app">
		<main>
			<slot />
		</main>

		<footer />
	</div>
</Session>
