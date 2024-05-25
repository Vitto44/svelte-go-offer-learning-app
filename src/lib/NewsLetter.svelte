<script lang="ts">
	import heroImage from "$lib/images/notification.png";
</script>

<div class="bg-gradient-to-r from-brand to-brandLight rounded-lg p-4 shadow-lg w-full col-span-6">
	<form class="">
		<div class="flex gap-4">
			<h2 class="text-3xl font-bold text-secondary-100">Hear it first!</h2>
			<img src={heroImage} alt="hero" class="col-span-1 h-10 -mt-3" />
		</div>
		<p class="text-white">
			Get the latest free bets and bonuses from the best bookmakers and casinos.
		</p>
		<div class="mt-4 sm:grid grid-cols-7 gap-4">
			<input
				type="email"
				placeholder="Enter your email"
				class="w-full p-2 border border-gray-300 rounded-lg col-span-5"
			/>
			<button
				type="submit"
				class="w-full bg-secondary-200 font-bold py-2 col-span-2 rounded-lg mt-2 sm:mt-0"
				>Subscribe</button
			>
		</div>
	</form>
</div>
