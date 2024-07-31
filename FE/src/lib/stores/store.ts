import { derived, writable } from "svelte/store";

const initialData: StoreData = {
	location: "Gb",
	type: "all",
	niche: "sport",
};

export const filters = writable(initialData);

export const pomos = derived(
	filters,
	($filters, set) => {
		let cancelled = false;

		const fetchData = async () => {
			set({ loading: true, error: null, data: null });
			try {
				const res = await fetch(`${process.env.VITE_API_URL}/${$filters.niche}`).catch(() => {
					throw new Error("Oops! Looks like you're offline.");
				});
				if (!cancelled) {
					if (!res.ok) {
						throw new Error("Oops! Something went wrong.");
					}
					const data = await res.json();
					set({ loading: false, error: null, data });
				}
			} catch (error: Error) {
				if (!cancelled) {
					set({ loading: false, error: error.message, data: null });
				}
			}
		};

		fetchData();

		return () => {
			cancelled = true;
		};
	},
	{ loading: false, error: null, data: null },
);
