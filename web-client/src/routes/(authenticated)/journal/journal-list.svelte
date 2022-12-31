<script lang="ts">
	let journals = [
		{
			mood: 'foo',
			title: 'some title really REAAAAAALlllllllllllllllyyyyyyy loooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
    {
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
    {
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		},
		{
			mood: 'foo',
			title: 'some title',
			createdAt: new Date()
		}
	];
</script>

<div class="flex flex-col divide-y">
	{#each journals as { mood, title, createdAt, id }}
		<div class="py-2 w-full overflow-auto truncate">
			<a href='/journal/${id}' class="text-2xl text-green-400 ">{title}</a>
			<div class="flex ">
				<div>{mood}</div>
				<div class="ml-auto">{createdAt.toLocaleDateString()}</div>
			</div>
		</div>
	{/each}
</div>
