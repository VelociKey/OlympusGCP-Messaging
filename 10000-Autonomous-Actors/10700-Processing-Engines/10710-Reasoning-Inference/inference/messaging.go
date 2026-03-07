package inference

import (
	"context"
	"log/slog"

	messagingv1 "olympus.fleet/00SDLC/OlympusGCP-Messaging/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen/v1/messaging"
	"connectrpc.com/connect"
)

type MessagingServer struct{}

func (s *MessagingServer) SendMessage(ctx context.Context, req *connect.Request[messagingv1.SendMessageRequest]) (*connect.Response[messagingv1.SendMessageResponse], error) {
	slog.Info("SendMessage", "to", req.Msg.Recipient)
	return connect.NewResponse(&messagingv1.SendMessageResponse{MessageId: "msg-999"}), nil
}
