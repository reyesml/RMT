<script lang="ts">
	import { getSessionStore } from '$lib/store/session';
	import type { ActionData } from './$types';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';

	export let form: ActionData;
	let session = getSessionStore();

	if (form?.success) {
		session.set(form.session!);
	}
</script>

<svelte:head>
	<title>Login</title>
	<meta name="description" content="RMT Login" />
</svelte:head>

<section>
	<div class="flex min-h-full items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
		<div class="w-full max-w-md space-y-8">
			{#if $session && browser}
				{goto('/home')}
			{:else}
				<div>
					<h2 class="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900">
						Sign in to your account
					</h2>
				</div>
				<form class="mt-8 space-y-6" action="?/login" method="POST">
					<div class="-space-y-px rounded-md shadow-sm">
						<div>
							<label for="username" class="sr-only">Username</label>
							<input
								id="username"
								name="username"
								type="text"
								required
								class="relative block w-full appearance-none rounded-none rounded-t-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
								placeholder="Username"
							/>
						</div>
						<div>
							<label for="password" class="sr-only">Password</label>
							<input
								id="password"
								name="password"
								type="password"
								autocomplete="current-password"
								required
								class="relative block w-full appearance-none rounded-none rounded-b-md border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
								placeholder="Password"
							/>
						</div>
					</div>

					<div>
						<button
							type="submit"
							class="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
						>
							<span class="absolute inset-y-0 left-0 flex items-center pl-3">
								<!-- Heroicon name: mini/lock-closed -->
								<svg
									class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400"
									xmlns="http://www.w3.org/2000/svg"
									viewBox="0 0 20 20"
									fill="currentColor"
									aria-hidden="true"
								>
									<path
										fill-rule="evenodd"
										d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z"
										clip-rule="evenodd"
									/>
								</svg>
							</span>
							Sign in
						</button>
					</div>
					{#if form && !form.success}
						<div class="w-full rounded-md bg-red-400 text-center p-2">Login failed.</div>
					{/if}
				</form>
			{/if}
		</div>
	</div>
</section>
