<script lang="ts">
	import { onDestroy } from 'svelte';
	import { portal } from './portal/actions';

	let modal: Element;

	const handle_keydown = (e: KeyboardEvent) => {
		
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			e.preventDefault();
			const nodes = modal.querySelectorAll<HTMLElement>('*');
			const tabbable = Array.from(nodes).filter((n) => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement as HTMLElement);
			console.log('currtab', index);

			let nextEl: HTMLElement;
			if (e.shiftKey) {
				if (index < 0) index = 0; //no element was selected, so 'fake' that the first element was selected.
				//going backwards
				if (index <= 0) {
					//we're at the beginning, so wrap to the end.
					console.log('backwards', tabbable.length - 1);
					nextEl = tabbable[tabbable.length - 1];
				} else {
					console.log('backwards', index-1);
					nextEl = tabbable[index - 1];
				}
			} else {
				if (index < tabbable.length - 1) {
					//move foreward one step
					console.log('forewards', index+1);
					nextEl = tabbable[index + 1];
				} else {
					//we will step out of bounds, so wrap to start.
					console.log('fowewards-wrap', 0)
					nextEl = tabbable[0];
				}
			}
			console.log(nextEl);
			nextEl && nextEl.focus();
		}
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
	class="fixed top-0 w-full h-full bg-black bg-opacity-75 transition-opacity"
>
	<div class="flex h-full w-full flex-col items-center justify-center">
		<div class="flex items-center">
			<slot />
		</div>
	</div>
</div>
