<script lang="ts">
	import Hero from "$lib/Hero.svelte";
	import Card from "$lib/Card.svelte";
	import NationPicker from "$lib/filters/NationPicker.svelte";
	import TypePicker from "$lib/filters/TypePicker.svelte";
	import NichePicker from "$lib/filters/NichePicker.svelte";
	import { pomos } from "$lib/stores/store";
	import { onDestroy } from "svelte";
	import Tooltip from "$lib/Tooltip.svelte";

	export let promotions: Promotion[] = [];

	const unsubscribePromos = pomos.subscribe(async (p) => {
		promotions = await p;
		console.log(promotions);
	});

	// Cleanup the subscription
	onDestroy(() => {
		unsubscribePromos();
	});
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section class="max-w-5xl mx-auto">
	<Tooltip />

	<Hero />

	<section class="flex justify-center lg:justify-between flex-wrap gap-2 my-8">
		<NationPicker />
		<NichePicker />
		<TypePicker />
	</section>

	<section
		class=" mx-auto gap-10 p-8 mb-8 rounded-2xl flex flex-row bg-gray-100 flex-wrap justify-center shadow-primary"
	>
		{#each promotions as promotion (promotion.id)}
			<Card {promotion} />
		{/each}
	</section>
</section>
