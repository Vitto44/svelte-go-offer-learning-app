<script>
	import HeaderLink from "./HeaderLink.svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	$: currentPage = $page.route.id;

	const navigate = (/** @type {string | URL} */ url) => {
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
				class={`border-b-4 py-3 ${currentPage === "/reviews" ? " border-secondary-100 bg-gray-200" : " border-gray-400"} `}
				>Reviews</a
			>

			<a
				href="/"
				on:click|preventDefault={() => navigate("/")}
				class={`border-b-4 py-3 ${currentPage === "/" ? " border-secondary-100 bg-gray-200" : " border-gray-400"} `}
				>Home</a
			>

			<a
				href="/blog"
				on:click|preventDefault={() => navigate("/blog")}
				class={`border-b-4 py-3 ${currentPage === "/blog" ? " border-secondary-100 bg-gray-200" : " border-gray-400"} `}
				>Blog</a
			>
			<!-- <HeaderLink url="/blog" title="Blog" />
			<HeaderLink url="/" title="Home" />
			<HeaderLink url="/reviews" title="Reviews" /> -->
		</ul>
	</nav>
</header>
