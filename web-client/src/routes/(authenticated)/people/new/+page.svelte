<script lang="ts">
	import type { ActionData } from './$types';
	export let form: ActionData;
</script>

<svelte:head>
	<title>RMT: New Person</title>
	<meta name="description" content="RMT new person" />
</svelte:head>

<section>
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl space-y-8">
			{#if !form || !form.success}
				<h1>New Person</h1>
				<form action="?/create" method="POST">
                    <div class="flex">
                        <div>
                            <label for="firstName" class="sr-only">First Name</label>
                            <input
                                id="firstName"
                                name="firstName"
                                type="text"
                                placeholder="first name"
                                class="relative block w-56 appearance-none rounded-3xl bg-black bg-opacity-20 px-8 py-4 text-white focus:z-10 border border-transparent focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 text-xl font-bold"
                            />
                        </div>
                        <div class="ml-auto">
                            <label for="lastName" class="sr-only">Last Name</label>
                            <input
                                id="lastName"
                                name="lastName"
                                type="text"
                                placeholder="last name"
                                class="relative block w-56 appearance-none rounded-3xl bg-black bg-opacity-20 px-8 py-4 text-white focus:z-10 border border-transparent focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 text-xl font-bold"
                            />
                        </div>
                    </div>
					<div class="mt-3 flex">
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
