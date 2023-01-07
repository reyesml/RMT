<script lang="ts">
	import { clickOutside } from '$lib/actions/click-outside';
	import { createEventDispatcher, onDestroy } from 'svelte';
	import { portal } from './portal/actions';

	const dispatch = createEventDispatcher();
	const close = () => dispatch('close');

	let modal: Element;

	const handle_keydown = (e: KeyboardEvent) => {
		if (e.key === 'Escape') {
			close();
			return;
		}
		if (e.key !== 'Tab') return;

		// trap focus
		e.preventDefault();
		const nodes = modal.querySelectorAll<HTMLElement>('*');
		const tabbable = Array.from(nodes).filter((n) => n.tabIndex >= 0);

		let index = tabbable.indexOf(document.activeElement as HTMLElement);

		let nextEl: HTMLElement;
		if (e.shiftKey) {
			if (index < 0) index = 0; //no element was selected, so 'fake' that the first element was selected.
			//going backwards
			if (index <= 0) {
				//we're at the beginning, so wrap to the end.
				nextEl = tabbable[tabbable.length - 1];
			} else {
				nextEl = tabbable[index - 1];
			}
		} else {
			if (index < tabbable.length - 1) {
				//move foreward one step
				nextEl = tabbable[index + 1];
			} else {
				//we will step out of bounds, so wrap to start.
				nextEl = tabbable[0];
			}
		}
		nextEl && nextEl.focus();
	};

	const previously_focused =
		typeof document !== 'undefined' && (document.activeElement as HTMLElement);

	if (previously_focused) {
		onDestroy(() => {
			previously_focused.focus();
		});
	}
</script>

<svelte:window on:keydown={handle_keydown} />

<div
	use:portal={'modal'}
	bind:this={modal}
	class="fixed top-0 w-full h-full flex flex-col items-center justify-center bg-black bg-opacity-75 transition-opacity"
>
	<div use:clickOutside on:outclick={close} class="flex items-center">
		<slot />
	</div>
</div>
