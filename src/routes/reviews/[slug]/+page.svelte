<script>
	import Card from "$lib/Card.svelte";
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
		prosCons,
	} = review;

	console.log(welcomeOffer);
</script>

<div class="max-w-5xl mx-auto mt-8 md:mt-28">
	<div class="grid grid-cols-6 gap-6">
		<div class="flex flex-col gap-4 col-span-6 lg:col-span-4">
			<SectionCard style="flex gap-0 sm:gap-6 flex-col sm:flex-row">
				<img
					class="w-52 h-52 sm:rounded-l-lg mx-auto rounded-full mt-6 sm:mt-0"
					src={image}
					alt={title}
				/>
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
				</div>
			</SectionCard>
			{#if welcomeOffer.howToRegister}
				<SectionCard style="p-4">
					<h2 class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold">
						How To Register
					</h2>
					{@html welcomeOffer.howToRegister}
				</SectionCard>
			{/if}
			{#if games && games.length > 0}
				<SectionCard style="p-4">
					<h2 class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold">
						Games
					</h2>
					<ul class="flex flex-wrap gap-x-2">
						{#each games as game, i}
							<li>
								{game}{#if i < games.length - 1},{/if}
							</li>
						{/each}
					</ul>
				</SectionCard>
			{/if}
			{#if description}
				<SectionCard style="p-4">
					<h2 class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold">
						Description
					</h2>
					<p>{description}</p>
				</SectionCard>
			{/if}
		</div>
		<div class="col-span-6 lg:col-span-2 flex gap-4 flex-col sm:flex-row lg:flex-col">
			<SectionCard style="p-4 bg-gray-800 w-full">
				<h2 class="mb-10 w-fit text-md bg-gray-100 py-1.5 px-6 rounded-md font-bold">
					Welcome Offer
				</h2>
				<Card promotion={welcomeOffer} />
			</SectionCard>
			{#if prosCons}
				<SectionCard style="p-4 mt-4 w-full">
					<h2 class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold">
						Pros & Cons
					</h2>
					<h3 class="mt-2 font-bold text-lg">Pros</h3>
					<hr class="h-1 w-8 mb-2 bg-green-600 rounded-xl" />
					<ul class="list-disc ml-5">
						{#each prosCons.pros as pro}
							<li>{pro}</li>
						{/each}
					</ul>
					<h3 class="mt-2 font-bold text-lg">Cons</h3>
					<hr class="h-1 w-8 mb-2 bg-red-600 rounded-xl" />
					<ul class="list-disc ml-5">
						{#each prosCons.cons as c}
							<li>{c}</li>
						{/each}
					</ul>
				</SectionCard>
			{/if}
		</div>
	</div>
</div>
