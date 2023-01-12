import { gateways } from '$lib/gateways';
import type { Journal } from '$lib/models/journal';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const auth = cookies.get('session') ?? '';
	const res = await gateways.journal.list(auth);
	if (!res.ok) {
		return { error: res.statusText };
	}
	const body = await res.json();
	return {
		journals: body.journals as Journal[]
	};
}) satisfies PageServerLoad;
