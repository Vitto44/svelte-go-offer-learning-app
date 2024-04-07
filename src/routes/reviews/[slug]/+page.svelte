<script>
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

<div class="max-w-5xl mx-auto">
	<section class="flex max-h-44 bg-red-400">
		<div>
			<h1>{title}</h1>
			<p>{category}</p>
		</div>
		<img class="max-h-44" src={image} alt={title} />
	</section>
	<div class="flex">
		<section class="bg-blue-500">
			<h2>Description</h2>
			<p>{description}</p>
		</section>
		<section class="bg-yellow-400">
			<p>{rating}</p>
			{games}
			<p>{availableCountries}</p>
			<p>{licensedBy}</p>
		</section>
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
