<script lang="ts">
	import { onMount } from 'svelte';
	import { createSessionStore } from '$lib/store/session';
	import { goto } from '$app/navigation';

	export let mandatory = true;

	let session = createSessionStore();

	let saveSession = false;
	$: if (saveSession) {
		if ($session) {
			window.sessionStorage.setItem('session', JSON.stringify($session));
		} else {
			window.sessionStorage.removeItem('session');
		}
	}

	onMount(async () => {
		//Fetch the session from storage
		let ses = window.sessionStorage.getItem('session');

		if (!ses && mandatory) {
			//redirect to login if session is missing and mandatory
			goto('/login');
			return;
		}

		//Flag the session for persistence (set during rerender)
		saveSession = true;

		if (!ses) {
			//exit early if session is missing
			return;
		}

		//Parse session and check expiration
		$session = JSON.parse(ses);

		let currTime = new Date().getTime();
		let duration = new Date($session!.expiration).getTime() - currTime;
		if (duration <= 0) {
			goto('/login');
			return;
		}

		//Set a timer to clear the session after expiration
		let timerId = setTimeout(() => {
			window.sessionStorage.removeItem('session');
			session.set(null);
			goto('/login');
		}, duration);

		//Destroy the timer if this component is unmounted.
		return () => {
			clearTimeout(timerId);
		};
	});
</script>

{#if !mandatory || $session}
	<slot />
{/if}
