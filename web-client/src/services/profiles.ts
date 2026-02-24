import { callApi } from "../shared/api";
import { type Profile } from "../shared/types";

export async function getProfile(userId: string): Promise<Profile> {
	const profile = await callApi("GET", `profiles/${userId}`);
	return profile;
}

export async function createProfile(
	p: Profile,
): Promise<Profile> {
	return await callApi("POST", `profiles`, p);
}

export async function updateProfile(
	p: Profile,
	userId: string,
): Promise<Profile> {
	return await callApi("PUT", `profiles/${userId}`, p);
}

export async function deleteProfile(userId: string) {
	await callApi("DELETE", `profiles/${userId}`, null);
}

export async function searchProfiles(query: string): Promise<Profile[]> {
	return await callApi("GET", `profiles/search?q=${query}`);
}
