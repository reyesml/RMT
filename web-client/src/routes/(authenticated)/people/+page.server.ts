import { gateways } from '$lib/gateways';
import type { Person } from '$lib/models/person';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const auth = cookies.get('session') ?? '';
	const res = await gateways.people.list(auth);
	if (!res.ok) {
		return { error: res.statusText };
	}
	const body = await res.json();
	return {
		people: body.people as Person[]
	};
}) satisfies PageServerLoad;
