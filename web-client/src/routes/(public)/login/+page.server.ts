import { gateways } from '$lib/gateways';
import type { Session } from '$lib/store/session';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ cookies }) => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	login: async ({ cookies, request }) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');

		const res = await gateways.auth.login(username?.toString() ?? '', password?.toString() ?? '');
		if (!res.ok) {
			return { success: false };
		}

		let body = await res.json();
		cookies.set('session', body.token, { expires: new Date(body.expiration) });
		return {
			success: true,
			session: {
				user: {
					UUID: body.user.UUID,
					username: body.user.username,
					admin: body.user.admin
				},
				expiration: new Date(body.expiration)
			} as Session
			
		};
	}
};
