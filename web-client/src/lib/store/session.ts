// IMPORTANT: SvelteKit SHARES ALL GLOBAL VARIABLES ACROSS ALL CLIENTS
// when performing SSR. That means things like stores are LEAKED to
// to different clients. To avoid this, we wrap the stores in a "context"
// so that they are no longer globally scoped, and thus not leaked to
// different clients.
import { setContext, getContext } from 'svelte';
import { type Writable, writable } from 'svelte/store';
import type { User } from '$lib/models/user';

export interface Session {
	user: User;
	expiration: Date;
}

export function createSessionStore() {
	return setContext('session-store', writable<Session>());
}

export function getSessionStore() {
	return getContext<Writable<Session>>('session-store');
}
