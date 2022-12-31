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
				<h1>New Journal</h1>
				<form action="?/create" method="POST">
					<div>
						<label for="title" class="sr-only">Title</label>
						<input
							id="title"
							name="title"
							type="text"
							required
							placeholder="title"
							class="relative block w-full appearance-none rounded-3xl bg-black bg-opacity-20 px-8 py-4 text-white focus:z-10 border border-transparent focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 text-xl font-bold"
						/>
					</div>
					<div class="mt-3 bg-black bg-opacity-20 rounded-3xl px-8 py-6 border border-transparent focus-within:z-10 focus-within:border-indigo-500 focus-within:outline-none focus-within:ring-indigo-500">
						<label for="body" class="sr-only">Body</label>
						<textarea
							id="body"
							name="body"
							placeholder="Dear diary..."
							required
							class="bg-transparent relative block w-full min-h-[300px] appearance-none rounded-md text-white inherit focus:outline-none"
						/>
					</div>
					<div class="mt-3 flex">
						<div>
							<label for="mood" class="sr-only">Mood</label>
							<select
								id="mood"
								name="mood"
								required
								class="relative bg-black bg-opacity-20 text-center block w-44 appearance-none rounded-3xl px-3 py-2 border border-transparent focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
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
							<button
								type="submit"
								class="bg-green-600 w-44 p-2 rounded-xl font-bold border border-transparent focus:z-10 focus:border-indigo-500 focus:ring-indigo-500"
								>submit</button
							>
						</div>
					</div>
				</form>
			{/if}
		</div>
	</div>
</section>
