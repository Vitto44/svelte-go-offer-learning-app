<script>
	import ArticleCard from "$lib/ArticleCard.svelte";
	import Card from "$lib/Card.svelte";
	import SectionCard from "$lib/SectionCard.svelte";

	export let data;

	const { blog } = data;
	const { content, title, articles, promotions, image } = blog;
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="">
	<div class="grid grid-cols-6 gap-3 md:gap-6 mt-12 max-w-7xl mx-auto">
		<SectionCard style="col-span-6 md:col-span-5">
			<div class=" flex justify-between flex-col sm:flex-row bg-secondary-200 rounded-lg">
				<h1 class=" m-4 sm:my-auto ml-4 text-3xl lg:text-4xl font-bold bg-white p-2 rounded-lg">
					{title}
				</h1>
				<img
					class="w-full sm:h-72 sm:w-auto rounded-t-full sm:rounded-r-lg sm:rounded-full"
					src={image}
					alt={title}
				/>
			</div>
		</SectionCard>
		<SectionCard style="col-span-1 flex">
			<h1 class="">share</h1>
		</SectionCard>

		<SectionCard style="col-span-6 md:col-span-4 p-4">
			{#if content}
				{@html content}
			{:else}
				<p>Fuck</p>
			{/if}
		</SectionCard>

		<div class="col-span-6 md:col-span-2 flex flex-col gap-8">
			<SectionCard style="py-12 flex flex-col gap-12">
				{#each promotions as promotion}
					<Card {promotion} useShadow />
				{/each}
			</SectionCard>
			<SectionCard style="py-12 flex flex-col gap-12">
				{#each articles as article}
					<ArticleCard {article} />
				{/each}
			</SectionCard>
		</div>
	</div>
</div>
