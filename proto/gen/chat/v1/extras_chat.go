package chatv1

import "connectrpc.com/connect"

func (msg *SayResponse) ConnectResponse() *connect.Response[SayResponse] {
	return connect.NewResponse(msg)
}
