import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	try {
		const res = await fetch(`http://localhost:3000/blog`);
		const blogList: BlogCard[] = await res.json();

		return {
			blogList,
		};
	} catch (error) {
		console.error("Error fetching review list:", error);
		return { error };
	}
};
