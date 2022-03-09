interface GoAPIResponse {
	data: { success: true; message: string };
}

/** @type {import('./index').RequestHandler} */
export async function get(): Promise<{ body: GoAPIResponse } | { status: number }> {
	const url = 'http://127.0.0.1:8080/api';
	const response = await fetch(url);
	const data = await response.json();

	if (data) {
		return {
			body: { data }
		};
	}

	return {
		status: 404
	};
}
