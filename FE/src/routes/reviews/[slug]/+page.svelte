<script>
	import Card from "$lib/Card.svelte";
	import ErrorMessage from "$lib/ErrorMessage.svelte";
	import SectionCard from "$lib/SectionCard.svelte";
	import Stars from "$lib/Stars.svelte";

	export let data;

	const { review, error } = data;
</script>

<div class="max-w-5xl mx-auto mt-8 md:mt-28">
	{#if error}
		{#if error === "404"}
			<div class="bg-white mx-auto text-center p-4 rounded-lg">
				<ErrorMessage error="Review you are looking for does not exist." />
			</div>
		{:else}
			<div class="bg-white p-4 rounded-lg shadow-lg">
				<ErrorMessage {error} />
			</div>
		{/if}
	{:else}
		<div class="grid grid-cols-6 gap-6">
			<div class="flex flex-col gap-4 col-span-6 lg:col-span-4">
				<SectionCard style="flex gap-0 sm:gap-6 flex-col sm:flex-row bg-secondary-200">
					<img
						class="w-52 h-52 sm:rounded-l-lg mx-auto rounded-full mt-6 sm:mt-0"
						src={review.image}
						alt={review.title}
					/>
					<div class="p-4 w-full">
						<h1
							class="mx-auto mb-2 w-fit text-lg text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
						>
							{review.title}
						</h1>
						<div class="bg-gray-100 p-2 rounded-lg">
							<p><span class="font-bold">Category:</span> {review.category}</p>
							<div class="flex w-fit gap-1">
								<p class="font-bold mt-1">Countries:</p>
								{#each review.availableCountries as country}
									<img
										src={`https://flagicons.lipis.dev/flags/4x3/${country.toLowerCase()}.svg`}
										alt={`${country} flag`}
										class="w-8 h-8"
									/>
								{/each}
							</div>
						</div>
					</div>
				</SectionCard>
				{#if review.welcomeOffer.howToRegister}
					<SectionCard style="p-4">
						<h2
							class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
						>
							How To Register
						</h2>
						{@html review.welcomeOffer.howToRegister}
					</SectionCard>
				{/if}
				{#if review.games && review.games.length > 0}
					<SectionCard style="p-4">
						<h2
							class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
						>
							Games
						</h2>
						<ul class="flex flex-wrap gap-x-2">
							{#each review.games as game, i}
								<li>
									{game}{#if i < review.games.length - 1},{/if}
								</li>
							{/each}
						</ul>
					</SectionCard>
				{/if}
				{#if review.description}
					<SectionCard style="p-4">
						<h2
							class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
						>
							review.Description
						</h2>
						<p>{review.description}</p>
					</SectionCard>
				{/if}
			</div>
			<div class="col-span-6 lg:col-span-2 flex gap-4 flex-col sm:flex-row lg:flex-col">
				<SectionCard style="p-4 bg-gray-800 w-full">
					<h2 class="mb-10 w-fit text-md bg-gray-100 py-1.5 px-6 rounded-md font-bold">
						Welcome Offer
					</h2>
					<Card isCard promotion={review.welcomeOffer} />
				</SectionCard>
				{#if review.prosCons}
					<SectionCard style="p-4 mt-4 w-full">
						<h2
							class="mb-2 w-fit text-md text-gray-100 bg-gray-800 py-1.5 px-6 rounded-md font-bold"
						>
							Pros & Cons
						</h2>
						<h3 class="mt-2 font-bold text-lg">Pros</h3>
						<hr class="h-2 w-20 mb-2 bg-green-600 rounded-xl" />
						<ul class="list-disc ml-5">
							{#each review.prosCons.pros as pro}
								<li>{pro}</li>
							{/each}
						</ul>
						<h3 class="mt-2 font-bold text-lg">Cons</h3>
						<hr class="h-2 w-20 mb-2 bg-red-600 rounded-xl" />
						<ul class="list-disc ml-5">
							{#each review.prosCons.cons as c}
								<li>{c}</li>
							{/each}
						</ul>
					</SectionCard>
				{/if}
			</div>
		</div>
	{/if}
</div>
