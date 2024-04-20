<script lang="ts">
	import { goto } from "$app/navigation";

	export let style = "";
	export let article: BlogCard;
	const { title, image, cardText, id } = article;

	function goToReview() {
		goto(`/blog/${id}`);
	}
</script>

<div
	class={`flex flex-col align-center w-72 rounded-xl pb-2 border-2 border-gray-900 shadow-[5px_5px_rgba(0,_64,_128,_0.4),_10px_10px_rgba(0,_64,_128,_0.3),_15px_15px_rgba(0,_64,_128,_0.2),_20px_20px_rgba(0,_64,_128,_0.1),_25px_25px_rgba(0,_64,_128,_0.05)]
${style}`}
>
	<img src={image} alt={`${title} logo`} class="rounded-t-lg mx-auto w-full" />
	<h3 class="mt-2 mx-2 w-fit text-gray-100 bg-gray-800 py-1 px-2 rounded-md font-bold">
		{title}
	</h3>
	<p class="px-2 my-2 text-sm line-clamp-3">
		{cardText}
	</p>
	<button
		class=" mx-2 bg-secondary-200 w-fit px-6 py-1 mb-1 font-bold rounded-md hover:bg-secondary-100 hover:scale-105 transition-all duration-200 ease-in-out"
		on:click={goToReview}
	>
		Read Post
	</button>
</div>
