<script lang="ts">
	import type { ActionData } from './$types';
	export let form: ActionData;
</script>

<svelte:head>
	<title>RMT: New Journal</title>
	<meta name="description" content="RMT new journal" />
</svelte:head>

<section>
	{#if !form || !form.success}
		<h1>New Journal</h1>
		<form action="?/create" method="POST">
			<div>
				<label for="title" class="sr-only">Title</label>
				<input id="title" name="title" type="text" required placeholder="title" />
			</div>
			<div>
				<label for="body" class="sr-only">Body</label>
				<textarea id="body" name="body" placeholder="Dear diary..." />
			</div>
			<div>
				<label for="mood" class="sr-only">Mood</label>
				<select id="mood" name="mood">
					<option value="">N/A</option>
					<option value="happy">happy</option>
					<option value="sad">sad</option>
					<option value="frisky">frisky 😈</option>
				</select>
			</div>
			<div>
                <button type="submit">submit</button>
            </div>
		</form>
	{/if}
</section>
