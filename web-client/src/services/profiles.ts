import { callApi } from "./api";
import { type Profile } from "../shared/types";

export async function getProfile(userId: string): Promise<Profile> {
	const profile = await callApi("GET", `profile/${userId}`);
	return profile;
}

export async function createProfile(
	p: Profile,
	bearerToken: string
): Promise<Profile> {
	return await callApi("POST", `profile`, p, bearerToken);
}

export async function updateProfile(
	p: Profile,
	userId: string,
	bearerToken: string
): Promise<Profile> {
	return await callApi("PUT", `profile/${userId}`, p, bearerToken);
}

export async function deleteProfile(userId: string, bearerToken: string) {
	await callApi("DELETE", `profile/${userId}`, null, bearerToken);
}

export async function searchProfiles(query: string): Promise<Profile[]> {
	return await callApi("GET", `profile/search?q=${query}`);
}
