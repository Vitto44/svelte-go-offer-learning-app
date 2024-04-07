import { derived, writable } from "svelte/store";

const initialData: StoreData = {
	location: "Gb",
	type: "all",
	niche: "sport",
};

export const filters = writable(initialData);

export const pomos = derived(filters, async ($filters) => {
	try {
		const res = await fetch(`http://localhost:3000/${$filters.niche}`);
		const data = await res.json();
		return data;
	} catch (error) {
		console.error("Error fetching pomos", error);
	}
});
