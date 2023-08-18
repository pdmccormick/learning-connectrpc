<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	let dispatch = createEventDispatcher<{
		send: {
			message: string;
		};
	}>();

	let message: string;

	function onSend() {
		dispatch('send', { message });
		message = '';
	}
</script>

<div>
	<form on:submit|preventDefault={onSend}>
		<!-- svelte-ignore a11y-autofocus -->
		<input type="text" bind:value={message} autofocus={true} />
		<button type="submit">Send!</button>
	</form>
</div>
