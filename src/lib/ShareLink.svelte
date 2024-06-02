<script lang="ts">
	import facebook from "$lib/images/socials/facebook.svg";
	import linkedin from "$lib/images/socials/linkedin.svg";
	import twitter from "$lib/images/socials/twitter.svg";
	import whatsapp from "$lib/images/socials/whatsapp.svg";

	const icons: Record<string, string> = {
		facebook,
		linkedin,
		twitter,
		whatsapp,
	};

	export let useTitle = true;
	export let style = "";
	export let link = "";
	export let title = "";

	const getIconPath = (title: string): string | undefined => {
		return icons[title.toLowerCase()];
	};
</script>

<a
	href={link}
	target="_blank"
	rel="noopener noreferrer"
	class={`rounded-lg flex text-white w-fill p-2 gap-2 transition-colors duration-300 ${style}`}
>
	{#if getIconPath(title)}
		<img alt={title} class="w-6 h-6" src={getIconPath(title)} />
	{/if}
	{#if useTitle}
		{title}
	{/if}
</a>
