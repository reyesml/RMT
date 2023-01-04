import { tick } from 'svelte';

const portal_map = new Map();

export function createPortal(node: Node, id = 'default') {
	if (portal_map.has(id)) throw `duplicate portal key "${id}"`;
	else portal_map.set(id, node);
	return { destroy: portal_map.delete.bind(portal_map, id) };
}

export function portal(node: Node, id = 'default') {
	let destroy: () => void;
	if (!portal_map.has(id))
		// TODO: better error handling when portal doesn't exist (yet)
		tick().then(() => {
			destroy = mount(node, id);
		});
	else destroy = mount(node, id);
	return { destroy: () => destroy?.() };
}

function mount(node: Node, id: string) {
	if (!portal_map.has(id)) throw `unknown portal ${id}`;
	const host = portal_map.get(id);
	host.insertBefore(node, null);
	return () => host.contains(node) && host.removeChild(node);
}
