export async function getPromotions(
	params: StoreData,
	fetch: (input: RequestInfo, init?: RequestInit) => Promise<Response>,
): Promise<void> {
	try {
		const res = await fetch(`${process.env.VITE_API_URL}/${params.niche}`);
		const data: Promotion[] = await res.json();

		console.log("data", data);
	} catch (error) {
		console.error("Error fetching pomos", error);
	}
}
