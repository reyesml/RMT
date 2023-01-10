<script lang="ts">
	import Portal from '$lib/components/portal/portal.svelte';
	import Session from '$lib/components/session.svelte';
</script>

<Session>
	<div class="mt-4 pb-10">
		<main class="h-full w-full">
			<slot />
		</main>

		<footer class="fixed inset-x-0 bottom-0 pb-6 bg-gradient-to-t from-slate-900 to-transparent">
			<div class="flex flex-col items-center">
				<div class="flex flex-row items-center">
					<a href="/home" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>Home</a
					>
					<a href="/people" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>People</a
					>
					<a href="/journal" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>Journal</a
					>
					<a href="/initiatives" class="bg-purple-700 rounded-md text-white w-32 py-1 mx-1 text-center"
						>Initiatives</a
					>
				</div>
			</div>
		</footer>
	</div>
	<Portal id="modal" />
</Session>
