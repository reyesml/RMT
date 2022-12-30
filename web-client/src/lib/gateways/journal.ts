const host = 'http://localhost:3000';

export function create(auth: string, mood: string, title: string, body: string) {
  console.log('calling create journal...')
	return fetch(`${host}/journal`, {
		method: 'POST',
		cache: 'no-cache',
		headers: {
			'Content-Type': 'application/json',
      'Authorization': `Bearer ${auth}`
		},
		body: JSON.stringify({ mood: mood, title: title, body: body })
	});
}
