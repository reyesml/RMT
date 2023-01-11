<script lang="ts">
	import { goto } from '$app/navigation';
	import LogOutIcon from '$lib/components/icons/log-out-icon.svelte';
	import { getSessionStore } from '$lib/store/session';
	let session = getSessionStore();

  async function PromptLogout() {
    if(window.confirm("Log out?")){
      await Logout()
    }
  }

	async function Logout() {
		await fetch('/api/logout', {
			method: 'POST',
			headers: {
				'content-type': 'application/json'
			}
		});

		$session = null;
		goto('/login');
	}
</script>

<button on:click={PromptLogout} title="log out"><LogOutIcon class="w-6 h-6" /></button>
