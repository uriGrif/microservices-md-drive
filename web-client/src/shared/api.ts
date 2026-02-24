import auth0Client from "./auth0-client";

export class ApiError extends Error {
	statusCode: number;

	constructor(statusCode: number, message: string) {
		super(message); // Call the parent class constructor (Error)
		this.name = "ApiError"; // Set the error name
		this.statusCode = statusCode;

		// Set the prototype explicitly for compatibility with `instanceof`
		Object.setPrototypeOf(this, ApiError.prototype);
	}
}

export const callApi = async (
	method: string,
	path: string,
	body?: any,
): Promise<any> => {
	const { getAccessTokenSilently } = auth0Client;
	try {
		const response = await fetch(
			`${import.meta.env.VITE_API_SERVER_URL}/${path}`,
			{
				method: method,
				headers: {
					Authorization: "Bearer " + await getAccessTokenSilently(),
					"Content-Type": "application/json"
				},
				body: JSON.stringify(body)
			}
		);
		if (!response.ok) {
			throw new ApiError(
				response.status,
				`Error: ${(await response.json()).message}`
			);
		}
		if (response.status === 204) {
			return;
		}
		const data = await response.json();
		return data;
	} catch (error) {
		if (error instanceof ApiError) {
			throw error;
		}
		throw new ApiError(0, `Error: failed to fetch: ${error}`);
	}
};
