<script>
	import { page } from "$app/stores";
	import logo from "$lib/images/svelte-logo.svg";
</script>

<header>
	<div class="corner bg-red-400">
		<a href="https://kit.svelte.dev">
			<img src={logo} alt="SvelteKit" />
		</a>
	</div>

	<nav>
		<ul>
			<li aria-current={$page.url.pathname === "/" ? "page" : undefined}>
				<a href="/">Home</a>
			</li>
			<li aria-current={$page.url.pathname === "/about" ? "page" : undefined}>
				<a href="/about">About</a>
			</li>
		</ul>
	</nav>
</header>
