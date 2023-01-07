import { gateways } from '$lib/gateways';
import type { PageServerLoad } from './$types';
import type { Actions } from '@sveltejs/kit';

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

	res = await gateways.people.getNotes(auth, params.uuid);
	let notes;
	if (!res.ok) {
		notes = [{ title: 'Error Loading Notes', body: '', uuid: '' }];
	} else {
		notes = (await res.json()).notes;
	}

	return {
		person: person,
		qualities: qualities,
		notes: notes
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	createNote: async ({ request, cookies }) => {
		const data = await request.formData();
		const title = data.get('title');
		const body = data.get('body') ?? '';
		const uuid = data.get('uuid') ?? '';
		if (!title) {
			return { success: false, error: 'title is required' };
		}
		let auth = cookies.get('session') ?? '';
		const res = await gateways.people.createNote(
			auth,
			uuid.toString(),
			title.toString(),
			body.toString()
		);
		if (!res.ok) {
			return { success: false, error: res.statusText };
		}
		return { success: true };
	},
	createQuality: async ({ request, cookies }) => {
		const data = await request.formData();
		const name = data.get('name');
		const uuid = data.get('uuid');
		if (!name) {
			return { success: false, error: 'name is required' };
		}
		let auth = cookies.get('session') ?? '';
		const res = await gateways.people.createQuality(auth, uuid!.toString()!, name.toString());
		if (!res.ok) {
			return { success: false, error: res.statusText };
		}
		return { success: true };
	}
};
