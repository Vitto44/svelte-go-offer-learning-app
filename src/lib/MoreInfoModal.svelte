<script>
	import SectionCard from "./SectionCard.svelte";

	export let card = "";
	export let closeModal = () => {};

	function bgClose() {
		closeModal();
	}

	/**
	 * @param {{ stopPropagation: () => void; }} event
	 */
	function stopPropagation(event) {
		event.stopPropagation();
	}

	/**
	 * @param {{ key: string; }} event
	 */
	function handleKeyDown(event) {
		if (event.key === "Escape") {
			closeModal();
		}
	}
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	on:click={bgClose}
	on:keydown={handleKeyDown}
	class="fixed top-0 left-0 w-full h-full flex justify-center items-center bg-black bg-opacity-80"
>
	<div class="absolute mx-5" on:click={stopPropagation} on:keydown={handleKeyDown}>
		<SectionCard style="px-6 py-16 max-w-5xl max-h-[90vh] overflow-y-scroll">
			<button
				class="absolute top-2 right-2 md:right-4 p-2 text-lg text-red-600 bg-white rounded-full shadow-lg hover:shadow-xl transition-all duration-200 ease-in-out hover:scale-105"
				on:click={closeModal}>❌</button
			>

			{@html card}
		</SectionCard>
	</div>
</div>
