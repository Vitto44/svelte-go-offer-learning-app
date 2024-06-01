<script>
	import ArticleCard from "$lib/ArticleCard.svelte";
	import ErrorMessage from "$lib/ErrorMessage.svelte";
	import Loader from "$lib/Loader.svelte";
	export let data;

	const { blogList, loading, error } = data;
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div
	class="mx-auto gap-10 p-8 mb-8 rounded-2xl flex bg-gray-100 flex-wrap shadow-primary mt-10 justify-center"
>
	{#if loading}
		<Loader />
	{:else if error}
		<ErrorMessage {error} />
	{:else if blogList && blogList.length !== 0}
		{#each blogList as article (article.id)}
			<ArticleCard {article} />
		{/each}
	{:else}
		<p>No reviews found</p>
	{/if}
</div>
