import { gateways } from '$lib/gateways';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from '../$types';

export const POST = (async ({ cookies }) => {
	const auth = cookies.get('session') ?? '';
	await gateways.auth.logout(auth);
	cookies.delete('session', { path: '/' });
	return json({});
}) satisfies RequestHandler;
