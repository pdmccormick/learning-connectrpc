import { createPromiseClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';

import { Chat } from './gen/chat/v1/chat_connect';

export const createChatClient = () => {
	return createPromiseClient(
		Chat,
		createConnectTransport({
			baseUrl: '/api'
		})
	);
};
