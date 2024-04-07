<script lang="ts">
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	$: currentPage = $page.route.id;

	const navigate = (/** @type {string | URL} */ url: string) => {
		if (url === currentPage) return;
		goto(url);
	};
</script>

<footer class="max-w-5xl mx-auto gap-8 my-8">
	<div
		class="flex flex-wrap justify-center gap-4 bg-gray-100 px-8 py-4 mx-4 rounded-2xl shadow-primary"
	>
		<p class="text-text-primary bg-gray-800 px-4 py-2 rounded-md">© Certain Wager</p>
		<ul class="flex gap-3 py-2">
			<a href="/reviews" on:click|preventDefault={() => navigate("/reviews")}>Reviews</a>

			<a href="/" on:click|preventDefault={() => navigate("/")}>Home</a>

			<a href="/blog" on:click|preventDefault={() => navigate("/blog")}>Blog</a>
		</ul>
	</div>
</footer>
