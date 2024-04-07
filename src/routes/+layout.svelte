<script>
	import Footer from "$lib/Footer.svelte";
	import Header from "$lib/Header.svelte";
	import "./styles.css";
</script>

<div class="app bg-stone-200 font-sans min-h-screen flex flex-col">
	<Header />

	<main class="mx-5">
		<slot />
	</main>
	<Footer />
</div>
