<script lang="ts">
	import { onMount } from 'svelte';
	import { createSessionStore } from '$lib/store/session';
	import { goto } from '$app/navigation';

	export let mandatory = true;

	let session = createSessionStore();

	let saveSession = false;
	$: if (saveSession && $session) {
		window.sessionStorage.setItem('session', JSON.stringify($session));
	}
	onMount(async () => {
		let ses = window.sessionStorage.getItem('session');
		if (ses) {
			$session = JSON.parse(ses);
		} else if (mandatory) {
			goto('/login');
		}
		saveSession = true;
	});
</script>

{#if !mandatory || $session}
	<slot />
{/if}
