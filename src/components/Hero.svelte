<script>
	import { page } from "$app/stores";
	import logo from "$lib/images/svelte-logo.svg";
	import heroImage from "$lib/images/hero-image.png";
</script>

<header>
	<div
		class="bg-gradient-to-r from-blue-700 to-cyan-600 text-stone-100 max-w-5xl mx-auto rounded-xl p-8 grid grid-cols-3 gap-4 relative mt-24"
	>
		<div class=" flex gap-8 flex-col col-span-2">
			<h1 class="font-semibold text-5xl">
				Shortcut to <span class="text-orange-400 font-bold">Free Bets</span> <br /> and
				<span class="text-orange-400 font-bold">Wins</span>!
			</h1>
			<p class="text-lg">
				Get the latest free bets and bonuses from the best bookmakers and casinos. We have the best
				free bets and bonuses for you to claim and win real money.
			</p>
		</div>
		<img src={heroImage} alt="hero" class="absolute max-w-80 right-0 -bottom-3" />
	</div>
</header>
``
