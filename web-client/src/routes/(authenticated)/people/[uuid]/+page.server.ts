import { gateways } from '$lib/gateways';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, params }) => {
	let auth = cookies.get('session') ?? '';
	const res = await gateways.people.get(auth, params.uuid);
	if (!res.ok) {
		return { error: res.statusText };
	}

	const body = await res.json();
	return {
		person: body.person
	};
}) satisfies PageServerLoad;
