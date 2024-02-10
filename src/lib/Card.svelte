<script lang="ts">
	import type { Promotion } from "../types";
	import Stars from "./Stars.svelte";

	export let promotion: Promotion;
	const { title, image, code, disclaimer, investment, rating, reward, rewardType, url, isMoney } =
		promotion;
</script>

<div
	class="flex flex-col align-center w-72 md:w-52 p-4 rounded-xl text-center border-2 border-gray-900 shadow-[5px_5px_rgba(0,_64,_128,_0.4),_10px_10px_rgba(0,_64,_128,_0.3),_15px_15px_rgba(0,_64,_128,_0.2),_20px_20px_rgba(0,_64,_128,_0.1),_25px_25px_rgba(0,_64,_128,_0.05)]
"
>
	<img
		src={image}
		alt={`${title} logo`}
		class="border-2 border-gray-900 rounded-full w-32 mx-auto -mt-10"
	/>
	<h3 class="font-bold mb-1 text-xl">{title}</h3>
	<Stars {rating} />
	<p class="font-bold text-lg">Bet ${investment}</p>
	<div class="relative flex items-center text-3xl">
		<div class="flex-grow border-2 border-secondary-200" />
		<span class="flex-shrink mx-2 text-gray-100 bg-gray-800 py-1 px-2 rounded-md font-bold"
			>Get {isMoney ? "$" : ""}{reward}</span
		>
		<div class="flex-grow border-2 border-secondary-200" />
	</div>
	<p class="font-bold text-lg">{isMoney ? "in " : ""}{rewardType}</p>

	<button class="mt-4 bg-primary text-text-primary -skew-x-12 w-fit mx-auto px-4 py-1"
		>More Info</button
	>
	<a href={url} target="_blank">
		<button
			class="mt-1 bg-secondary-200 w-fit mx-auto px-12 py-2 mb-2 font-bold rounded-md text-xl"
		>
			Claim
		</button>
	</a>

	{#if code}
		<p class="text-xs font-semibold text-gray-800">Or use code: {code}</p>
	{/if}

	<p class="text-xs text-gray-400 max-h-12 overflow-y-scroll p-1 my-4">{disclaimer}</p>
</div>
