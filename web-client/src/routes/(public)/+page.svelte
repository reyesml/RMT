<script lang="ts">
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { getSessionStore } from '$lib/store/session';

	let session = getSessionStore();

	if(browser){
		if ($session) {
			goto('/home');
		} else {
			goto('/login');
		}
	}
</script>
