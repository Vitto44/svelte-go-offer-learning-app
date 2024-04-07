import { derived } from "svelte/store";

export default function (stores, initial_value) {
	return derived(
		stores,
		($stores, set) => {
			set(initial_value);
			// Promise.resolve(callback($stores)).then((value) => {
			// 	if (guard === inner) {
			// 		set(value);
			// 	}
			// });

			async function load() {
				try {
					const res = await fetch(`http://localhost:3000/${$stores.niche}`);
					const data = await res.json();
					set(data);
				} catch (error) {
					console.error("Error fetching pomos", error);
				}
			}

			load();
		},
		initial_value,
	);
}
