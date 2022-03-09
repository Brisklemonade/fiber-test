<script lang="ts">
	import { Button } from '@svelteuidev/core';

	// fetched data from index endpoint
	export let data;
	const { message, success } = data;
</script>

<div class="container">
	<h1>Welcome to SvelteKit served by a Go server</h1>

	<Button uppercase variant="gradient" href={'/api'} external>Go to api</Button>

	{#if success}
		<p>The api says {message}</p>
	{/if}
</div>

<style>
	.container {
		display: grid;
		place-content: center;
		justify-content: center;
		align-items: center;
	}
</style>
