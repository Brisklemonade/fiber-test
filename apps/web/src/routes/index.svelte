<script lang="ts">
	import { Button } from '@svelteuidev/core';

	interface GoAPIResponse {
		data: { success: true; message: string };
	}
	async function load(): Promise<GoAPIResponse> {
		const url = 'http://127.0.0.1:8080/api';
		const response = await fetch(url);
		const data = await response.json();

		return {
			data
		};
	}
</script>

<div class="container">
	<h1>Welcome to SvelteKit served by a Go server</h1>

	<Button uppercase variant="gradient" href="/api" external>Go to api</Button>

	{#await load()}
		<p>The api is loading...</p>
	{:then value}
		<p>The api says {value.data.message}</p>
	{:catch error}
		<p>There was an unexpected error {error}</p>
	{/await}
</div>

<style>
	.container {
		display: grid;
		place-content: center;
		justify-content: center;
		align-items: center;
	}
</style>
