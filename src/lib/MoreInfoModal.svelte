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
	class="fixed top-0 left-0 w-full h-full flex justify-center items-center bg-black bg-opacity-50"
>
	F
	<div class="absolute mx-5" on:click={stopPropagation} on:keydown={handleKeyDown}>
		<SectionCard style="p-6 pt-14 max-w-5xl">
			<button class="absolute top-0 right-0 p-4 text-2xl text-red-600" on:click={closeModal}
				>❌</button
			>

			{@html card}
		</SectionCard>
	</div>
</div>
