export interface Person {
	uuid: string;
	firstName: string;
	lastName: string;
	createdAt: string;
	updatedAt: string;
}

export interface SearchablePerson {
	uuid: string;
	firstName: string;
	lastName: string;
	qualities: SearchableQuality[];
}

export interface SearchableQuality {
	name: string;
	type: string;
}