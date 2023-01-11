import type { Quality } from "./quality";

export interface Person {
	uuid: string;
	firstName: string;
	lastName: string;
	qualities: Quality[]; //The root Quality record, not PersonQuality join table
	createdAt: string;
	updatedAt: string;
}