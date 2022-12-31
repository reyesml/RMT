import { gateways } from '$lib/gateways';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	const auth = cookies.get('session') ?? '';
	const res = await gateways.journal.list(auth);
	if (!res.ok) {
		return { error: res.statusText };
	}
	const body = await res.json();
  console.log(body)
	return {
		journals: body.journals
	};
}) satisfies PageServerLoad;
