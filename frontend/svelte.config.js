import { default as adapterStatic } from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

let adapter = adapterStatic({
	pages: 'build',
	fallback: 'index.html',
	precompress: true
});

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter
	}
};

export default config;
