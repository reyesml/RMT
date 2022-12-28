const host = 'http://localhost:3000';

export function login(username: string, password: string) {
  return fetch(`${host}/login`, {
    method: "POST",
    cache: "no-cache",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: 'include',
    body: JSON.stringify({ username: username, password: password, tokenInBody: true }),
  });
}