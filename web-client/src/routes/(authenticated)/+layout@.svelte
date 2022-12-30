<script lang="ts">
	import Session from '$lib/components/session.svelte';
</script>

<Session>
	<div class="mt-4">
		<main>
			<slot />
		</main>

		<footer class="fixed inset-x-0 bottom-0 p-4">
			<div class="flex flex-col items-center">
				<div class="flex flex-row items-center">
					<a href="/journal" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>Journal</a
					>
					<a href="/people" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>People</a
					>
					<a href="/settings" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>Settings</a
					>
				</div>
			</div>
		</footer>
	</div>
</Session>
