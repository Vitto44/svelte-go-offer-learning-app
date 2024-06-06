<script lang="ts">
	import ErrorMessage from "$lib/ErrorMessage.svelte";
	import Loader from "$lib/Loader.svelte";
	import ReviewCard from "$lib/ReviewCard.svelte";

	export let data;

	const { reviewList, loading, error } = data;
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div
	class="mx-auto gap-10 p-8 mb-8 rounded-2xl flex flex-row bg-gray-100 flex-wrap justify-center shadow-primary mt-10"
>
	{#if loading}
		<Loader />
	{:else if error}
		<ErrorMessage {error} />
	{:else if reviewList && reviewList.length !== 0}
		{#each reviewList as review (review.id)}
			<ReviewCard {review} />
		{/each}
	{:else}
		<p>No reviews found</p>
	{/if}
</div>
