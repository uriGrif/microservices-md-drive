export interface Profile {
	userId?: string;
	nickname: string | null;
	country: string | null;
	description: string | null;
}

export interface File {
	id: string;
	name: string;
	createdAt: string;
	updatedAt: string;
	userId: string;
}
