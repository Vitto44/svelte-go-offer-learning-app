<script lang="ts">
	import { goto } from "$app/navigation";
	import { page } from "$app/stores";
	import Casino from "$lib/images/niches/casino.svelte";
	import Esports from "$lib/images/niches/esports.svelte";
	import Sports from "$lib/images/niches/sports.svelte";
	import { filters } from "$lib/stores/store";

	const types = [
		{ name: "Sport", code: "sport", component: Sports },
		{ name: "Casino", code: "casino", component: Casino },
		{ name: "E-sport", code: "esport", component: Esports },
	];

	$: niche = $filters.niche;

	function setNiche(i: string) {
		if (i === niche) return;
		filters.update((data) => ({ ...data, niche: i }));

		$page.url.searchParams.set("niche", i);
		const currentQuery = $page.url.searchParams.toString();

		goto(`/?${currentQuery}`, { noScroll: true });
	}
</script>

<ul class="flex justify-center items-center bg-gray-100 w-fit shadow-primary">
	{#each types as category}
		<button
			on:click={() => setNiche(category.code)}
			class="p-2 h-full flex gap-2 items-center"
			class:bg-secondary-200={niche === category.code}
		>
			<span class="text-sm font-semibold">{category.name}</span>
			<div class="w-8">
				<svelte:component this={category.component} />
			</div>
		</button>
	{/each}
</ul>
