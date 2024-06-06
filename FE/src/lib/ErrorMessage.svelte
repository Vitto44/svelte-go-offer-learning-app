<script>
	import errorImage from "$lib/images/error.svg";
	export let error = "Error";
</script>

<div class="mx-auto flex flex-col gap-2 text-center">
	<img src={errorImage} alt="Error" class="h-14" />
	<p class="text-red-600 font-bold text-3xl">{error}</p>
</div>
