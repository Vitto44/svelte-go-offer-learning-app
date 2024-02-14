import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	try {
		const res = await fetch("http://localhost:3000/esports");
		const promotions = await res.json();

		return { promotions };
	} catch (error) {
		console.error("Error fetching promotions:", error);
		return { error };
	}
};
