<script lang="ts">
	import ArticleCard from "$lib/ArticleCard.svelte";
	import Card from "$lib/Card.svelte";
	import ErrorMessage from "$lib/ErrorMessage.svelte";
	import Loader from "$lib/Loader.svelte";
	import NewsLetter from "$lib/NewsLetter.svelte";
	import SectionCard from "$lib/SectionCard.svelte";
	import ShareLink from "$lib/ShareLink.svelte";

	export let data;

	const { blog, loading, error } = data;
</script>

<svelte:head>
	<title>{blog ? blog.title : "Blog Post"}</title>
	<meta name="description" content={blog ? blog.description : "Blog post"} />
</svelte:head>

<div class="container mx-auto xl:p-8">
	{#if loading}
		<div>
			<Loader />
		</div>
	{:else if error}
		{#if error === "404"}
			<div class="bg-white mx-auto text-center p-4 rounded-lg">
				<ErrorMessage error="Post you are looking for does not exist." />
			</div>
		{:else}
			<div class="bg-white p-4 rounded-lg shadow-lg">
				<ErrorMessage {error} />
			</div>
		{/if}
	{:else}
		<div class="grid grid-cols-6 gap-3 lg:gap-6 mt-12 max-w-7xl mx-auto">
			<SectionCard style="col-span-6 lg:col-span-5">
				<div class="flex justify-between flex-col sm:flex-row bg-secondary-200 rounded-lg">
					<h1 class="m-4 sm:my-auto ml-4 text-3xl lg:text-4xl font-bold bg-white p-2 rounded-lg">
						{blog.title}
					</h1>
					<img
						class="max-w-72 mx-auto sm:max-w-full sm:h-72 sm:w-auto rounded-t-full sm:rounded-r-lg sm:rounded-full"
						src={blog.image}
						alt={blog.title}
					/>
				</div>
			</SectionCard>
			<SectionCard style="col-span-6 lg:col-span-1 shadow-primary bg-gray-100">
				<h1
					class="w-full p-2 pl-6 lg:text-center lg:p-2 bg-gray-800 text-white font-bold text-xl rounded-t-lg"
				>
					Share
				</h1>
				<ul class="grid grid-cols-4 lg:grid-cols-1 gap-4 py-4 px-2 justify-between my-auto">
					<li>
						<ShareLink
							link={`https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(
								`https://www.example.com/blog/${blog.slug}`,
							)}`}
							title="Facebook"
							style="bg-blue-500 hover:bg-blue-600"
						/>
					</li>
					<li>
						<ShareLink
							link={`https://twitter.com/intent/tweet?text=${encodeURIComponent(
								blog.title,
							)}&url=${encodeURIComponent(`https://www.example.com/blog/${blog.slug}`)}`}
							title="Twitter"
							style="bg-gray-900 hover:bg-gray-800"
						/>
					</li>
					<li>
						<ShareLink
							link={`https://www.linkedin.com/shareArticle?mini=true&url=${encodeURIComponent(
								`https://www.example.com/blog/${blog.slug}`,
							)}`}
							title="LinkedIn"
							style="bg-blue-800 hover:bg-blue-900"
						/>
					</li>
					<li>
						<ShareLink
							link={`https://wa.me/?text=${encodeURIComponent(
								`https://www.example.com/blog/${blog.slug}`,
							)}`}
							title="WhatsApp"
							style="bg-green-500 hover:bg-green-600"
						/>
					</li>
				</ul>
			</SectionCard>

			<SectionCard style="col-span-6 lg:col-span-4 p-4 h-fit shadow-primary bg-gray-100">
				{@html blog.content}
			</SectionCard>

			<div class="col-span-6 lg:col-span-2 flex flex-col gap-8">
				<SectionCard style="py-12 flex flex-col gap-12">
					{#each blog.promotions as promotion}
						<Card {promotion} isCard />
					{/each}
				</SectionCard>
				<NewsLetter />
				<SectionCard style="py-4 flex flex-col gap-12 ">
					{#each blog.articles as article}
						<ArticleCard {article} style="mx-auto" />
					{/each}
				</SectionCard>
			</div>
		</div>
	{/if}
</div>
