<script lang="ts">
	import Cz from "$lib/images/flags/cz.svelte";
	import Sk from "$lib/images/flags/sk.svelte";
	import Us from "$lib/images/flags/us.svelte";
	import Ie from "$lib/images/flags/ie.svelte";
	import Gb from "$lib/images/flags/gb.svelte";
	import { filters } from "$lib/stores/store";
	import { goto } from "$app/navigation";
	import { page } from "$app/stores";

	const countries = [
		{ name: "Czech Republic", code: "Cz", component: Cz },
		{ name: "Slovakia", code: "Sk", component: Sk },
		{ name: "United States", code: "Us", component: Us },
		{ name: "Ireland", code: "Ie", component: Ie },
		{ name: "Great Britain", code: "Gb", component: Gb },
	];

	$: location = $filters.location;

	function setLanguage(i: string) {
		if (i === location) return;
		filters.update((data) => ({ ...data, location: i }));

		$page.url.searchParams.set("location", i);
		const currentQuery = $page.url.searchParams.toString();

		goto(`/?${currentQuery}`, { noScroll: true });
	}
</script>

<ul class="flex justify-center items-center bg-gray-100 w-fit shadow-primary">
	{#each countries as country}
		<button
			on:click={() => setLanguage(country.code)}
			class="w-12 h-full flex p-1 items-center justify-center"
			class:bg-secondary-200={location === country.code}
			class:grayscale-[85%]={location !== country.code}
		>
			<svelte:component this={country.component} />
		</button>
	{/each}
</ul>
