import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ params }) => {
	try {
		const res = await fetch(`http://localhost:3000/companies/${params.slug}`);
		const review = await res.json();

		return {
			review,
		};
	} catch (error) {
		console.error("Error fetching reviews", error);
		return { error };
	}
};
