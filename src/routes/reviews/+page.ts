import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	try {
		const res = await fetch(`http://localhost:3000/companies`);
		const reviewList: Review[] = await res.json();

		console.log(reviewList);

		return {
			reviewList,
		};
	} catch (error) {
		console.error("Error fetching review list:", error);
		return { error };
	}
};
