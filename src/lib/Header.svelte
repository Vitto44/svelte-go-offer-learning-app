<script lang="ts">
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";
	import Cz from "$lib/images/flags/cz.svelte";
	import Sk from "$lib/images/flags/sk.svelte";
	import Us from "$lib/images/flags/us.svelte";
	import Ie from "$lib/images/flags/ie.svelte";
	import Gb from "$lib/images/flags/gb.svelte";

	$: currentPage = $page.route.id;

	const navigate = (/** @type {string | URL} */ url: string) => {
		if (url === currentPage) return;
		goto(url);
	};
</script>

<header>
	<nav class="bg-gray-100 text-text-secondary text-xl font-bold shadow-[0_3px_10px_rgb(0,0,0,0.2)]">
		<ul class="grid grid-cols-3 text-center">
			<a
				href="/reviews"
				on:click|preventDefault={() => navigate("/reviews")}
				class={`py-3`}
				class:bg-gray-300={$page.route.id?.match("reviews")}>Reviews</a
			>

			<a
				href="/"
				on:click|preventDefault={() => navigate("/")}
				class="py-3"
				class:bg-gray-300={$page.route.id === "/"}>Home</a
			>

			<a
				href="/blog"
				on:click|preventDefault={() => navigate("/blog")}
				class={`py-3`}
				class:bg-gray-300={$page.route.id?.match("blog")}>Blog</a
			>
			<ul class="flex justify-center items-center">
				<li class="w-12">
					<Cz />
				</li>
				<li class="w-12 grayscale">
					<Sk />
				</li>
				<li class="w-12 grayscale">
					<Ie />
				</li>
				<li class="w-12 grayscale">
					<Us />
				</li>
				<li class="w-12 grayscale">
					<Gb />
				</li>
			</ul>
		</ul>
	</nav>
</header>
