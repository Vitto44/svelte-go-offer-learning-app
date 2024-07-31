import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
	const { slug } = params;

	try {
		const res = await fetch(`${process.env.VITE_API_URL}/companies/${slug}`).catch(() => {
			throw new Error("Failed to fetch Review");
		});

		if (res.status === 404) {
			throw new Error("404");
		}

		if (!res.ok) {
			throw new Error("Failed to fetch blog post");
		}

		const review = await res.json();

		return {
			error: null,
			review,
		};
	} catch (error: unknown) {
		let errorMessage = "Unknown error";
		if (error instanceof Error) {
			errorMessage = error.message;
		}
		console.error("Error fetching reviews", error);
		return { error: errorMessage, review: null };
	}
};
