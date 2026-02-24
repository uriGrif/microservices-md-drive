import type { CreateFileDTO, EditFileDTO, File } from "../shared/types";
import { callApi } from "../shared/api";

export async function getFiles(): Promise<File[]> {
	return await callApi("GET", "files/by-user");
}

export async function createFile(file: CreateFileDTO): Promise<File> {
	return await callApi("POST", "files/", file);
}

export async function editFile(id: string, file: EditFileDTO): Promise<File> {
	return await callApi("PUT", `files/${id}`, file);
}

export async function deleteFile(id: string): Promise<File> {
	return await callApi("DELETE", `files/${id}`);
}