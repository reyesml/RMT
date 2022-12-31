import { gateways } from '$lib/gateways';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, params }) => {
	let auth = cookies.get('session') ?? '';
	const res = await gateways.journal.get(auth, params.uuid);
	if (!res.ok) {
		return { error: res.statusText };
	}

	const body = await res.json();
	return {
		journal: body.journal
	};
}) satisfies PageServerLoad;
