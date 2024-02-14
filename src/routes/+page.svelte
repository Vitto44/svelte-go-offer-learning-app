<script lang="ts">
	import type { Promotion } from "../types";
	import Hero from "$lib/Hero.svelte";
	import Card from "$lib/Card.svelte";

	export let data: { promotions: Promotion[] };
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Hero />

	<section
		class="max-w-5xl mx-auto gap-10 p-8 my-8 rounded-2xl flex flex-row bg-gray-100 flex-wrap justify-center shadow-[rgba(0,_0,_0,_0.24)_0px_3px_8px]"
	>
		{#each data.promotions as promotion}
			<Card {promotion} />
		{/each}
	</section>

	<!-- <h1>
    <span class="welcome">
      <picture>
        <source srcset={welcome} type="image/webp" />
        <img src={welcome_fallback} alt="Welcome" />
      </picture>
    </span>

    to your new<br />SvelteKit app
  </h1>

  <h2>
    try editing <strong>src/routes/+page.svelte</strong>
  </h2> -->
</section>
