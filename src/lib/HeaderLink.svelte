<script>
	/**
	 * @type {string}
	 */
	export let url;
	/**
	 * @type {any}
	 */
	export let title;

	import { page } from "$app/stores";
	let isCurrentPage = $page.url.pathname === url;
</script>

<li
	aria-current={isCurrentPage ? "page" : undefined}
	class={`skew-r-12 py-3 border-b-4 ${isCurrentPage ? "bg-primary text-text-primary border-gray-100" : "bg-gray-100 border-gray-200"}`}
>
	<a href={url}>{title}</a>
</li>
