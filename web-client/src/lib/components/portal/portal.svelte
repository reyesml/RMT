<script lang="ts">
	import { createPortal } from './actions';

	export let id: string = 'default';
</script>

<div use:createPortal={id} />
