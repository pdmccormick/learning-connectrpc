package proto

import (
	chatv1 "github.com/pdmccormick/learning-connectrpc/proto/gen/chat/v1"
	chatv1connect "github.com/pdmccormick/learning-connectrpc/proto/gen/chat/v1/chatv1connect"
)

// Chat
type (
	ChatClient  = chatv1connect.ChatClient
	ChatHandler = chatv1connect.ChatHandler

	Chat_SayRequest     = chatv1.SayRequest
	Chat_SayResponse    = chatv1.SayResponse
	Chat_ListenRequest  = chatv1.ListenRequest
	Chat_ListenResponse = chatv1.ListenResponse
)
