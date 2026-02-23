package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	messagingv1 "OlympusGCP-Messaging/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/messaging/v1"
	"OlympusGCP-Messaging/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/messaging/v1/messagingv1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type MessagingServer struct{}

func (s *MessagingServer) SendMessage(ctx context.Context, req *connect.Request[messagingv1.SendMessageRequest]) (*connect.Response[messagingv1.SendMessageResponse], error) {
	slog.Info("SendMessage", "to", req.Msg.Recipient)
	return connect.NewResponse(&messagingv1.SendMessageResponse{MessageId: "msg-999"}), nil
}

func main() {
	server := &MessagingServer{}
	mux := http.NewServeMux()
	path, handler := messagingv1connect.NewMessagingServiceHandler(server)
	mux.Handle(path, handler)

	port := "8090" // From genesis.json (MessagingManager replaces MeshHub)
	slog.Info("MessagingManager starting", "port", port)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Server failed", "error", err)
	}
}
