export interface Profile {
	userId?: string;
	nickname: string | null;
	country: string | null;
	description: string | null;
}

export interface File {
	id: string;
	name: string;
	created_at: string;
	updated_at: string;
	userId: string;
}

export interface CreateFileDTO {
	name: string;
}

export interface EditFileDTO {
	name: string;
}