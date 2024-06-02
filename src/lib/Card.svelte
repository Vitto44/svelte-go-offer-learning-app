<script lang="ts">
	import Stars from "./Stars.svelte";

	export let useShadow = false;
	export let isCard = false;
	export let useMoreInfo = false;
	export let promotion: Promotion;
	export let setCardID = (id: number) => {};

	const {
		title,
		id,
		image,
		code,
		disclaimer,
		investment,
		rating,
		reward,
		rewardType,
		url,
		isMoney,
	} = promotion;
</script>

<div
	class={`flex flex-col  mx-auto align-center bg-gray-100 w-72 p-4 rounded-xl text-center border-2 border-gray-900 ${!isCard && "md:grid md:gap-6 md:grid-cols-12 md:h-44 md:w-full md:flex-row"} ${useShadow && "shadow-[5px_5px_rgba(0,_64,_128,_0.4),_10px_10px_rgba(0,_64,_128,_0.3),_15px_15px_rgba(0,_64,_128,_0.2),_20px_20px_rgba(0,_64,_128,_0.1),_25px_25px_rgba(0,_64,_128,_0.05)]"}`}
>
	<div class="mx-auto pb-3 col-span-3">
		<img
			src={image}
			alt={`${title} logo`}
			class="border-2 border-gray-900 rounded-3xl w-32 mx-auto -mt-10 shadow-primary"
		/>
		<h3 class="font-bold my-1 text-xl">{title}</h3>
		<Stars {rating} />
	</div>

	<div class="col-span-3 my-auto">
		<p class="font-bold text-lg">Bet ${investment}</p>
		<div class="relative flex items-center text-3xl">
			<div class="flex-grow border-2 border-secondary-200" />
			<span class="flex-shrink mx-2 text-gray-100 bg-gray-800 py-1 px-2 rounded-md font-bold"
				>Get {isMoney ? "$" : ""}{reward}</span
			>
			<div class="flex-grow border-2 border-secondary-200" />
		</div>
		<p class="font-bold text-lg">{isMoney ? "in " : ""}{rewardType}</p>
	</div>

	<p
		class={`text-xs text-gray-400 max-h-12 overflow-y-scroll p-1 my-2 ${!isCard && "col-span-3 md:max-h-full"}`}
	>
		{disclaimer}
	</p>

	<div class="flex flex-col col-span-3">
		{#if useMoreInfo}
			<button
				on:click={() => {
					setCardID(id);
				}}
				class="mt-4 bg-brand text-text-primary -skew-x-12 w-fit mx-auto px-4 py-1"
			>
				How To Claim
			</button>
		{/if}
		<a href={url} target="_blank">
			<button
				class="mt-1 bg-secondary-200 w-fit mx-auto px-12 py-2 mb-2 font-bold rounded-md text-xl hover:bg-secondary-100 hover:scale-105 transition-all duration-200 ease-in-out"
			>
				Claim
			</button>
		</a>
		{#if code}
			<p class="text-xs font-semibold text-gray-800">Or use code: {code}</p>
		{/if}
	</div>
</div>
