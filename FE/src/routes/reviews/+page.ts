import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	try {
		const res = await fetch(`${process.env.VITE_API_URL}/companies`).catch(() => {
			throw new Error("Failed to fetch review list");
		});
		if (!res.ok) {
			throw new Error("Failed to fetch review list");
		}
		const reviewList: Review[] = await res.json();

		return {
			reviewList,
			loading: false,
			error: null,
		};
	} catch (error: unknown) {
		let errorMessage = "Unknown error";
		if (error instanceof Error) {
			errorMessage = error.message;
		}
		console.error("Error fetching review list:", errorMessage);
		return {
			reviewList: [],
			loading: false,
			error: errorMessage,
		};
	}
};
