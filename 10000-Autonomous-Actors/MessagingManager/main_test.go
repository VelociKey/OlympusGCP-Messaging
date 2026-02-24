package main

import (
	"context"
	"testing"

	messagingv1 "OlympusGCP-Messaging/40000-Communication-Contracts/430-Protocol-Definitions/000-gen/messaging/v1"
	"connectrpc.com/connect"
)

func TestMessagingServer(t *testing.T) {
	server := &MessagingServer{}
	ctx := context.Background()

	// Test SendMessage
	req := connect.NewRequest(&messagingv1.SendMessageRequest{
		Recipient: "user-1",
		Content:   "Hello",
	})
	res, err := server.SendMessage(ctx, req)
	if err != nil {
		t.Fatalf("SendMessage failed: %v", err)
	}
	if res.Msg.MessageId == "" {
		t.Error("Expected message ID, got empty string")
	}
}
