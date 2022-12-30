import { gateways } from '$lib/gateways';
import { redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ request, cookies }) => {
		const data = await request.formData();
		const mood = data.get('mood') ?? '';
		const title = data.get('title');
		const body = data.get('body');
		if (!title || !body) {
			return { success: false, error: 'title and body are required' };
		}
		let auth = cookies.get('session') ?? '';
		const res = await gateways.journal.create(auth, mood.toString(), title.toString(), body.toString());

		if (!res.ok) {
			return { success: false, error: res.statusText };
		}
		throw redirect(307, '/journal');
	}
};
