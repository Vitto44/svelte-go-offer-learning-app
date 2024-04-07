<script lang="ts">
	import { goto } from "$app/navigation";
	import { page } from "$app/stores";
	import AllTypes from "$lib/images/types/allTypes.svelte";
	import Money from "$lib/images/types/money.svelte";
	import NoMoney from "$lib/images/types/no-money.svelte";
	import { filters } from "$lib/stores/store";

	const types = [
		{ name: "No Deposit", code: "noDeposit", component: NoMoney },
		{ name: "Deposit", code: "deposit", component: Money },
		{ name: "All", code: "all", component: AllTypes },
	];

	$: storeType = $filters.type;

	function setType(i: string) {
		if (i === storeType) return;
		filters.update((data) => ({ ...data, type: i }));

		$page.url.searchParams.set("type", i);
		const currentQuery = $page.url.searchParams.toString();

		goto(`/?${currentQuery}`, { noScroll: true });
	}
</script>

<ul class="flex justify-center items-center bg-gray-100 w-fit shadow-primary">
	{#each types as type}
		<button
			on:click={() => setType(type.code)}
			class="p-2 h-full flex gap-2 items-center"
			class:bg-secondary-200={storeType === type.code}
		>
			<span class="text-sm font-semibold">{type.name}</span>
			<div class="w-8">
				<svelte:component this={type.component} />
			</div>
		</button>
	{/each}
</ul>
