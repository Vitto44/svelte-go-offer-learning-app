import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	try {
		const res = await fetch(`http://localhost:3000/blog`).catch(() => {
			throw new Error("Failed to fetch blog list");
		});
		if (!res.ok) {
			throw new Error("Failed to fetch blog list");
		}
		const blogList: BlogCard[] = await res.json();

		return {
			blogList,
			error: null,
			loading: false,
		};
	} catch (error: unknown) {
		let errorMessage = "Unknown error";
		if (error instanceof Error) {
			errorMessage = error.message;
		}
		console.error("Error fetching review list:", error);
		return {
			blogList: [],
			loading: false,
			error: errorMessage,
		};
	}
};
