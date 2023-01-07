const host = 'http://localhost:3000';

export function get(auth: string, uuid: string) {
	return fetch(`${host}/person-quality/${uuid}`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function getNotes(auth: string, uuid: string) {
	return fetch(`${host}/person-quality/${uuid}/notes`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function createNote(auth: string, uuid: string, title: string, body: string) {
	return fetch(`${host}/person-quality/${uuid}/notes`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		},
		body: JSON.stringify({ title: title, body: body })
	});
}