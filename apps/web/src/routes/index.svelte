<script>
	import { Button } from '@svelteuidev/core';
</script>

<div class="container">
	<h1>Welcome to SvelteKit served by a Go server</h1>

	<Button uppercase variant="gradient" href="/api" external>Go to api</Button>
</div>

<style>
	.container {
		display: grid;
		place-content: center;
		justify-content: center;
		align-items: center;
	}
</style>
