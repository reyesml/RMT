import { gateways } from '$lib/gateways';
import { getNotes } from '$lib/gateways/people';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies, params }) => {
	let auth = cookies.get('session') ?? '';
	let res = await gateways.people.get(auth, params.uuid);
	if (!res.ok) {
		return { error: res.statusText };
	}
	const person = (await res.json()).person;

	res = await gateways.people.getQualities(auth, params.uuid);
	let qualities;
	if (!res.ok) {
		qualities = [{ name: 'Error Loading qualities', uuid: 'not-found' }];
	} else {
		qualities = (await res.json()).personQualities;
	}
	console.log('quals', qualities);

	res = await gateways.people.getNotes(auth, params.uuid);
	let notes;
	if (!res.ok) {
		notes = [{ title: 'Error Loading Notes', body: '', uuid: '' }];
	} else {
		notes = (await res.json()).notes;
	}

	console.log('notes', notes);

	return {
		person: person,
		qualities: qualities,
		notes: notes
	};
}) satisfies PageServerLoad;
