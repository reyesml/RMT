const host = 'http://localhost:3000';

export function list(auth: string) {
	return fetch(`${host}/people`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function get(auth: string, uuid: string) {
	return fetch(`${host}/people/${uuid}`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function create(auth: string, firstName: string, lastName: string) {
  return fetch(`${host}/people`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		},
		body: JSON.stringify({ firstName: firstName, lastName: lastName })
	});
}

export function getQualities(auth: string, uuid: string) {
	return fetch(`${host}/people/${uuid}/qualities`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function getNotes(auth: string, uuid: string) {
	return fetch(`${host}/people/${uuid}/notes`, {
		method: 'GET',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		}
	});
}

export function createNote(auth: string, uuid: string, title: string, body: string) {
	return fetch(`${host}/people/${uuid}/notes`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		},
		body: JSON.stringify({ title: title, body: body })
	});
}

export function createQuality(auth: string, uuid: string, name: string, type: string) {
	return fetch(`${host}/people/${uuid}/qualities`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${auth}`
		},
		body: JSON.stringify({ name: name, type: type })
	});
}