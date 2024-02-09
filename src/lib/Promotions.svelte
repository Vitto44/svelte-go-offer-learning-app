<script context="module" lang="ts">
	import { load } from "../routes/+page";
	import type { Promotion } from "../types";

	// let serverPromos: Promotion[];

	// (async () => {
	// 	const result = await load();
	// 	serverPromos = result.data;
	// })();
</script>

<script>
	import Card from "./Card.svelte";
	// import { getPromotions } from "../api/promotions";

	// export let promotions = serverPromos;
	// console.log(promotions);

	// export let promotions;
</script>

<section
	class="max-w-5xl mx-auto gap-6 p-8 my-8 rounded-2xl flex flex-row flex-wrap justify-center shadow-[rgba(0,_0,_0,_0.24)_0px_3px_8px]"
>
	<!-- {#each promotions as promotion}
		<Card />
	{/each} -->
</section>
