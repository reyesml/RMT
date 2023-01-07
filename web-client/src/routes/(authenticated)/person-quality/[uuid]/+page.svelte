<script lang="ts">
	import PlusIcon from '$lib/components/icons/plus-icon.svelte';
	import Modal from '$lib/components/modal.svelte';
	import type { Note } from '$lib/models/note';
	import type { Person } from '$lib/models/person';
	import type { PersonQuality } from '$lib/models/person-quality';
    import type { ActionData, PageData } from './$types';
    import NoteList from '$lib/components/note-list.svelte';
	import { page } from '$app/stores';

    
    export let data: PageData;
    console.log(data);
    let person: Person;
	$: person = data.person;
    
    let quality: PersonQuality;
    $: quality = data.quality;

    let notes: Note[];
	$: notes = data.notes;

	let isNoteModalOpen = false;
	export let form: ActionData;
	if(isNoteModalOpen && (form?.success)){
		isNoteModalOpen = false;
	}
</script>

<svelte:head>
	<title>RMT: Person-Quality</title>
	<meta name="description" content="RMT Person-Quality" />
</svelte:head>

<section class="relative">
	<div class="flex min-h-full items-center justify-center py-12 px-4">
		<div class="w-full max-w-xl">
			{#if data.error}
				<h1>{data.error}</h1>
			{:else if person}
				<h1 class="text-4xl break-words">
                    <a href="/people/{person.uuid}" class="text-green-500">{`${person.firstName} ${person.lastName}`.trim()}</a>
                    -- {quality.name}
                </h1>
				<section class="mt-6 w-full">
					<div class="flex items-center">
						<h2 class="text-xl font-bold">Notes</h2>
						<button
							class="ml-2 bg-purple-600 rounded-full p-1"
							on:click={() => {
								isNoteModalOpen = true;
							}}><PlusIcon class="h-4 w-4" /></button
						>
					</div>
					<NoteList {notes} />
				</section>
			{/if}
		</div>
	</div>
</section>

{#if isNoteModalOpen}
	<Modal on:close={() => {isNoteModalOpen = false}}>
		<section
			role="dialog"
			class="w-[700px] bg-gray-800 border-2 border-purple-600 rounded-xl p-5 focus:outline-none"
		>
			<h2 class="text-center text-xl font-bold">Add Note</h2>
			<form action="?/createNote" method="POST">
				<input tabindex="-1" type="hidden" name="uuid" value={$page.params.uuid} />
				<div class="mt-5">
					<label for="title" class="sr-only">Title</label>
					<input
						id="title"
						name="title"
						type="text"
						required
						placeholder="title"
						class="relative block w-full appearance-none rounded-3xl bg-black px-8 py-4 text-white focus:z-10 border border-transparent focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 text-xl font-bold"
					/>
				</div>
				<div
					class="mt-3 bg-black rounded-3xl border border-transparent focus-within:z-10 focus-within:border-indigo-500 focus-within:outline-none focus-within:ring-indigo-500"
				>
					<label for="body" class="sr-only">Body</label>
					<textarea
						id="body"
						name="body"
						placeholder="Say more..."
						class="bg-transparent relative block w-full px-8 mt-3 pb-3 min-h-[300px] appearance-none rounded-md text-white inherit focus:outline-none"
					/>
				</div>
				<div class="mt-3 flex">
					<div>
						<button
							type="button"
							on:click={() => {
								isNoteModalOpen = false;
							}}
							class="bg-red-600 w-44 p-2 rounded-xl font-bold border border-transparent focus:z-10 focus:border-indigo-500 focus:ring-indigo-500"
							>cancel</button
						>
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
		</section>
	</Modal>
{/if}
