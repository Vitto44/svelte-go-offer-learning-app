<script>
	import heroImage from "$lib/images/hero-image.png";
</script>

<header>
	<div
		class="bg-gradient-to-r from-brand to-brandLight text-stone-100 max-w-5xl mx-auto rounded-xl p-8 grid grid-cols-3 relative mt-28 shadow-[4.0px_8.0px_8.0px_rgba(0,0,0,0.38)]"
	>
		<div class=" flex gap-4 md:gap-8 flex-col col-span-3 sm:col-span-2">
			<h1 class="font-semibold text-2xl md:text-5xl">
				Shortcut to <span class="text-secondary-200 font-bold">Free Bots</span>
				<br class="hidden lg:inline-block" />
				and
				<span class="text-secondary-200 font-bold">Clams</span>!
			</h1>
			<p class="text-sm md:text-base lg:text-lg sm:mr-12 lg:mr-0">What about this desc?</p>
		</div>
		<img
			src={heroImage}
			alt="hero"
			class="absolute max-w-36 sm:max-w-64 md:max-w-80 right-0 -top-24 sm:-top-16"
		/>
	</div>
</header>
