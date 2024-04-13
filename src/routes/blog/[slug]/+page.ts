import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ params }) => {
	try {
		const res = await fetch(`http://localhost:3000/blog/${params.slug}`);
		const blog = await res.json();

		return {
			blog,
		};
	} catch (error) {
		console.error("Error fetching reviews", error);
		return { error };
	}
};
