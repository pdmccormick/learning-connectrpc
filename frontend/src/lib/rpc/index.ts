import { createPromiseClient } from '@bufbuild/connect';
import { createConnectTransport } from '@bufbuild/connect-web';

import { Chat } from './gen/chat/v1/chat_connect';

export const createChatClient = () => {
	return createPromiseClient(
		Chat,
		createConnectTransport({
			baseUrl: '/api'
		})
	);
};
