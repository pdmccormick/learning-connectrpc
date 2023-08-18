package backend

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"

	"github.com/pdmccormick/learning-connectrpc/proto"
	"github.com/pdmccormick/learning-connectrpc/proto/gen/chat/v1/chatv1connect"
)

type chatService struct {
	srv *Server
}

var _ proto.ChatHandler = (*chatService)(nil)

func NewChatService() *chatService {
	return &chatService{}
}

func (svc *chatService) asHandler() (string, http.Handler) {
	return chatv1connect.NewChatHandler(svc)
}

func (svc *chatService) Say(ctx context.Context, req *connect.Request[proto.Chat_SayRequest]) (*connect.Response[proto.Chat_SayResponse], error) {
	log.Printf("Say: %s", req.Msg.Message)

	var resp = proto.Chat_SayResponse{
		Ok: true,
	}

	return resp.ConnectResponse(), nil
}

func (svc *chatService) Listen(ctx context.Context, req *connect.Request[proto.Chat_ListenRequest], resp *connect.ServerStream[proto.Chat_ListenResponse]) error {
	err := resp.Send(&proto.Chat_ListenResponse{
		Person:  "admin",
		Message: "Hello!",
	})
	if err != nil {
		return err
	}

	return nil
}
