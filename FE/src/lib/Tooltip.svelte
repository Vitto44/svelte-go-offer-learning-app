<script>
	import Info from "$lib/images/info.svg?raw";

	let show = false;
</script>

<div class="text-gray-600 text-xs absolute top-16 left-4">
	<button
		on:mouseleave={() => (show = false)}
		on:mouseenter={() => (show = true)}
		on:keydown={(e) => {
			if (e.key === "Enter") show = !show;
		}}
		class="flex gap-1"
	>
		<p>Advertiser Disclosure</p>
		<p class="fill-gray-600 w-4 h-4">
			{@html Info}
		</p>
	</button>
	<div
		class={`absolute z-10 p-2 rounded-lg mt-1 w-64 transition duration-500 bg-gray-100 ${show ? "opacity-1" : "opacity-0"}`}
	>
		<p>Content of this will contain refference to products and whatnot!</p>
	</div>
</div>
