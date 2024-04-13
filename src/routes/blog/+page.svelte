<script>
	import SectionCard from "$lib/SectionCard.svelte";
</script>

<svelte:head>
	<title>About</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="">
	<div class="grid grid-cols-6 gap-6 mt-12 max-w-7xl mx-auto">
		<SectionCard style="col-span-6 p-4">
			<h1 class="bg-brand text-white">About this blog</h1>
		</SectionCard>

		<SectionCard style="col-span-6 md:col-span-4">
			<p>
				This is a <a href="https://kit.svelte.dev">SvelteKit</a> app. You can make your own by typing
				the following into your command line and following the prompts:
			</p>
		</SectionCard>

		<SectionCard style="col-span-6 md:col-span-2">
			<p>
				The page you're looking at is purely static HTML, with no client-side interactivity needed.
				Because of that, we don't need to load any JavaScript. Try viewing the page's source, or
				opening the devtools network panel and reloading.
			</p>
		</SectionCard>
	</div>
</div>
