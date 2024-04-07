<script lang="ts">
	import { goto } from "$app/navigation";
	import Stars from "./Stars.svelte";

	export let review: Review;
	const { title, image, rating, category, id, availableCountries } = review;

	function goToReview() {
		goto(`/reviews/${id}`);
	}
</script>

<div
	class="flex flex-col align-center w-64 rounded-xl pb-2 text-center border-2 border-gray-900 shadow-[5px_5px_rgba(0,_64,_128,_0.4),_10px_10px_rgba(0,_64,_128,_0.3),_15px_15px_rgba(0,_64,_128,_0.2),_20px_20px_rgba(0,_64,_128,_0.1),_25px_25px_rgba(0,_64,_128,_0.05)]
"
>
	<img src={image} alt={`${title} logo`} class="rounded-t-lg mx-auto h-44 w-64" />
	<h3 class="items-center mx-auto my-3 text-gray-100 bg-gray-800 py-1 px-2 rounded-md font-bold">
		{title}
	</h3>
	<Stars {rating} />

	<p class="mt-1 w-fit mx-auto px-1 py-1 font-bold">{category}</p>
	{#each availableCountries as country}
		<img
			src={`https://www.countryflags.io/${country}/flat/64.png`}
			alt={`${country} flag`}
			class="w-8 h-8 mx-auto"
		/>
	{/each}
	<button
		on:click={() => goToReview()}
		class="mt-1 bg-secondary-200 w-fit mx-auto px-12 py-2 mb-2 font-bold rounded-md text-xl hover:bg-secondary-100 hover:scale-105 transition-all duration-200 ease-in-out"
	>
		More Info
	</button>
</div>
