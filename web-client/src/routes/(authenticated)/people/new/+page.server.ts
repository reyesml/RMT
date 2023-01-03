import { gateways } from '$lib/gateways';
import { redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ request, cookies }) => {
		const data = await request.formData();
		const firstName = data.get('firstName') ?? '';
		const lastName = data.get('lastName') ?? '';
		if (!firstName && !lastName) {
			return { success: false, error: 'at least one name is required' };
		}
		let auth = cookies.get('session') ?? '';
		const res = await gateways.people.create(auth, firstName.toString(), lastName.toString());

		if (!res.ok) {
			return { success: false, error: res.statusText };
		}
		throw redirect(307, '/people');
	}
};