<script lang="ts">
	import { getSessionStore } from '$lib/store/session';

	let session = getSessionStore();
</script>

<h1>Home Page (Placeholder)</h1>
<p>
	You probably want {#if $session}
		<a href="/home">Home</a>
	{:else}
		<a href="/login">Login</a>
	{/if} instead.
</p>
