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

func (svc *chatService) Init() error {
	*svc = chatService{
		srv: svc.srv,
	}

	return nil
}

func (svc *chatService) asHandler() (string, http.Handler) {
	return chatv1connect.NewChatHandler(svc)
}

func (svc *chatService) Say(ctx context.Context, req *connect.Request[proto.Chat_SayRequest]) (*connect.Response[proto.Chat_SayResponse], error) {
	log.Printf("Say: %s", req.Msg.Message)

	svc.srv.broadcastc <- req.Msg.Message

	var resp = proto.Chat_SayResponse{
		Ok: true,
	}

	return resp.ConnectResponse(), nil
}

func (svc *chatService) Listen(ctx context.Context, req *connect.Request[proto.Chat_ListenRequest], resp *connect.ServerStream[proto.Chat_ListenResponse]) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := resp.Send(&proto.Chat_ListenResponse{
		Message: "/joined",
	})
	if err != nil {
		return err
	}

	var msgc = make(chan string, 100)

	svc.srv.subc <- msgc

	defer func() {
		svc.srv.unsubc <- msgc
	}()

	for {
		select {
		case <-ctx.Done():
			return nil

		case msg, more := <-msgc:
			if msg != "" {
				err := resp.Send(&proto.Chat_ListenResponse{
					Message: msg,
				})
				if err != nil {
					return err
				}
			}

			if !more {
				return nil
			}
		}
	}

	return nil
}
