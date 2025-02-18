import { callApi } from "./api";
import { type Profile } from "../shared/types";

export async function getProfile(userId: string): Promise<Profile> {
	const profile = await callApi("GET", `profiles/${userId}`);
	return profile;
}

export async function createProfile(
	p: Profile,
	bearerToken: string
): Promise<Profile> {
	return await callApi("POST", `profiles`, p, bearerToken);
}

export async function updateProfile(
	p: Profile,
	userId: string,
	bearerToken: string
): Promise<Profile> {
	return await callApi("PUT", `profiles/${userId}`, p, bearerToken);
}

export async function deleteProfile(userId: string, bearerToken: string) {
	await callApi("DELETE", `profiles/${userId}`, null, bearerToken);
}

export async function searchProfiles(query: string): Promise<Profile[]> {
	return await callApi("GET", `profiles/search?q=${query}`);
}
