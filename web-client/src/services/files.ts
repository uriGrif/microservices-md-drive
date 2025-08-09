import type { File } from "../shared/types";
import { callApi } from "./api";

export async function getFiles(): Promise<File[]> {
	// mock
	return [
		{
			id: "1",
			name: "file1.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user1"
		},
		{
			id: "2",
			name: "file2.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user2"
		},
		{
			id: "3",
			name: "file1.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user1"
		},
		{
			id: "4",
			name: "file2.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user2"
		},
		{
			id: "5",
			name: "file1.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user1"
		},
		{
			id: "6",
			name: "file2.txt",
			createdAt: new Date().toISOString(),
			updatedAt: new Date().toISOString(),
			userId: "user2"
		}
	];
	return await callApi("GET", "files/by-user");
}
