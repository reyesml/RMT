const host = 'http://localhost:3000';

export function create(auth: string, mood: string, title: string, body: string) {
	return fetch(`${host}/journal`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		},
		body: JSON.stringify({ mood: mood, title: title, body: body })
	});
}

export function list(auth: string) {
	return fetch(`${host}/journal/`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function get(auth: string, uuid: string) {
	return fetch(`${host}/journal/${uuid}`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}
