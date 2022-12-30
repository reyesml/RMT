<script lang="ts">
	import type { ActionData } from './$types';
	export let form: ActionData;
</script>

<svelte:head>
	<title>RMT: New Journal</title>
	<meta name="description" content="RMT new journal" />
</svelte:head>

<section>
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl space-y-8">
			{#if !form || !form.success}
				<div>
					<h1>New Journal</h1>
				</div>
				<form action="?/create" method="POST">
					<div>
						<label for="title" class="sr-only">Title</label>
						<input
							id="title"
							name="title"
							type="text"
							required
							placeholder="title"
							class="relative block w-full appearance-none rounded-md border border-gray-300 px-3 py-2 text-white focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
						/>
					</div>
					<div class="mt-3">
						<label for="body" class="sr-only">Body</label>
						<textarea
							id="body"
							name="body"
							placeholder="Dear diary..."
							required
							class="relative block w-full h-52 appearance-none rounded-md border border-gray-300 px-3 py-2 text-white focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
						/>
					</div>
					<div class="mt-3 flex">
						<div>
							<label for="mood" class="sr-only">Mood</label>
							<select
								id="mood"
								name="mood"
								class="relative block w-44 appearance-none rounded-md border border-gray-300 px-3 py-2 text-white focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
							>
								<option value="" disabled selected>Mood</option>
								<option value="happy">happy</option>
								<option value="excited">excited</option>
								<option value="hopeful">hopeful</option>
								<option value="calm">calm</option>
								<option value="bleh">bleh</option>
								<option value="anxious">anxious</option>
								<option value="sad">sad</option>
								<option value="angry">angry</option>
							</select>
						</div>
						<div class="ml-auto">
							<button type="submit" class="bg-green-700 w-44 p-2 rounded-md focus:z-10 focus:border-indigo-500 focus:ring-indigo-500">submit</button>
						</div>
					</div>
				</form>
			{/if}
		</div>
	</div>
</section>
