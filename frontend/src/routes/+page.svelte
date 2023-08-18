<script lang="ts">
	import { onMount } from 'svelte';
	import { createChatClient } from '$lib/rpc';

	const client = createChatClient();

	let message = '';

	onMount(async () => {
		try {
			let response = await client.say({
				message: 'Hello World!'
			});

			message = JSON.stringify(response);
		} catch (e) {
			message = `ERROR: ${e}`;
		}
	});
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<h1>Response</h1>
<pre>{message} </pre>
