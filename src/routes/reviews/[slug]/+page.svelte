<script>
	import SectionCard from "$lib/SectionCard.svelte";
	import Stars from "$lib/Stars.svelte";

	export let data;

	const { review } = data;
	const {
		title,
		description,
		image,
		category,
		rating,
		licensedBy,
		availableCountries,
		games,
		welcomeOffer,
		howToRegister,
		disclaimer,
	} = review;

	console.log(review);
</script>

<div class="max-w-5xl mx-auto mt-28 flex flex-col gap-6">
	<div class="grid grid-cols-6 gap-6">
		<div class="flex flex-col gap-4 col-span-4">
			<!-- <section class="flex gap-6 shadow-primary bg-gray-100 rounded-lg"> -->
			<SectionCard style="flex gap-6">
				<img class="w-52 rounded-l-lg" src={image} alt={title} />
				<div class="p-4 w-full">
					<h1
						class="mx-auto mb-2 w-fit text-lg text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
					>
						{title}
					</h1>
					<p><span class="font-bold">Category:</span> {category}</p>
					<div class="flex w-fit">
						<p class="font-bold">Rating:</p>
						<Stars {rating} />
					</div>
					<p><span class="font-bold">Licensed By:</span> {licensedBy}</p>
					<div class="flex w-fit gap-1">
						<p class="font-bold mt-1">Countries:</p>
						{#each availableCountries as country}
							<img
								src={`https://flagicons.lipis.dev/flags/4x3/${country.toLowerCase()}.svg`}
								alt={`${country} flag`}
								class="w-8 h-8"
							/>
						{/each}
					</div>
					<!-- <p>
						<span class="font-bold">Games:</span>
						{#each games as game, i}
							{game}{#if i < games.length - 1},&nbsp;
							{/if}
						{/each}
					</p> -->
				</div>
			</SectionCard>
			<SectionCard style="p-4">
				<h2 class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold">
					Description
				</h2>
				<p>{description}</p>
			</SectionCard>
		</div>
		<SectionCard style="p-4 col-span-2">
			{welcomeOffer}
		</SectionCard>
	</div>
	<div class="flex">
		<section class="bg-orange-500">
			{welcomeOffer}
		</section>
		<section class="bg-green-400">
			{howToRegister}
		</section>
	</div>
	<section class="bg-pink-600">
		{disclaimer}
	</section>
</div>
