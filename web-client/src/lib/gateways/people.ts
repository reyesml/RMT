const host = 'http://localhost:3000';

export function list(auth: string) {
	return fetch(`${host}/people/`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}
