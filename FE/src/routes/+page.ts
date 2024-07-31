import { filters } from "$lib/stores/store";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ url }) => {
	try {
		const location = url.searchParams.get("location") || "Gb";
		const type = url.searchParams.get("type") || "all";
		const niche = url.searchParams.get("niche") || "sport";

		// const res = await fetch(`${process.env.VITE_API_URL}/${niche}`);
		// const promotions = await res.json();

		filters.update((data) => ({ ...data, location, type, niche }));
		// return { promotions };
	} catch (error) {
		console.error("Error fetching promotions:", error);
		return { error };
	}
};
