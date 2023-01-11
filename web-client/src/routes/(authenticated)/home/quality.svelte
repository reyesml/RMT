<script lang="ts">
	import type { Quality } from "$lib/models/quality";

  export let item: Quality;
</script>

<div class="text-md text-center text-white bg-green-600 py-1 px-2 mr-2 rounded-full">
  {item.name}
</div>