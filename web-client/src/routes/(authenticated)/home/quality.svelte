<script lang="ts">
	import type { SearchableQuality } from "$lib/models/person";

  export let item: SearchableQuality;
</script>

<div class="text-md text-center text-white bg-green-600 py-1 px-2 mr-2 rounded-full">
  {item.name}
</div>