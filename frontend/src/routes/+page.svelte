<script lang="ts">
	import { onMount } from 'svelte';
	import { createChatClient } from '$lib/rpc';
	import { writable } from 'svelte/store';
	import ChatInput from './ChatInput.svelte';

	const client = createChatClient();

	let channel = writable([] as string[]);
	let error = '';

	async function sendMessage(message: string) {
		try {
			let resp = await client.say({
				message
			});

			console.log('say response', resp);
		} catch (e) {
			error = `ERROR: ${e}`;
		}
	}

	onMount(async () => {
		for await (const resp of client.listen({})) {
			console.log('listen', resp);

			channel.update((messages) => {
				messages.push(resp.message);
				return messages;
			});
		}
	});
</script>

<h1>Chat</h1>

{#if error}
	<p>ERROR: {error}</p>
{/if}

<ul class="channel">
	{#each $channel as message, i (i)}
		{@const k = i + 1}
		<li><b>#{k}.</b> {message}</li>
	{/each}
</ul>

<ChatInput on:send={(ev) => sendMessage(ev.detail.message)} />

<style>
	.channel {
		display: flex;
		flex-direction: column-reverse;
	}
</style>
