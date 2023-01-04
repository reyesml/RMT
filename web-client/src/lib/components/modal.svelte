<script lang="ts">
	import { portal } from './portal/actions';
</script>

<div
	use:portal={'modal'}
	class="fixed top-0 w-full h-full bg-gray-500 bg-opacity-75 transition-opacity"
>
	<div class="flex h-full w-full flex-col items-center justify-center">
		<div class="flex items-center">
			<slot />
		</div>
	</div>
</div>
