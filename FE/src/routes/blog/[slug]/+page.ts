import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
	const { slug } = params;

	try {
		const res = await fetch(`http://localhost:3000/blog/${slug}`).catch(() => {
			throw new Error("Failed to fetch blog post");
		});

		if (res.status === 404) {
			throw new Error("404");
		}

		if (!res.ok) {
			throw new Error("Failed to fetch blog post");
		}

		const blog = await res.json();

		return {
			blog,
			loading: false,
			error: null,
		};
	} catch (error: unknown) {
		let errorMessage = "Unknown error";
		if (error instanceof Error) {
			errorMessage = error.message;
		}
		console.error("Error fetching blog post:", errorMessage);
		return {
			error: errorMessage,
			post: null,
		};
	}
};
