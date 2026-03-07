package inference

import (
	"context"
	"testing"

	messagingv1 "olympus.fleet/00SDLC/OlympusGCP-Messaging/40000-Communication-Contracts/40400-Protocol-Synthetics/connect-rpc/gen/v1/messaging"
	"connectrpc.com/connect"
)

func TestMessagingServer_CoverageExpansion(t *testing.T) {
	server := &MessagingServer{}
	ctx := context.Background()

	// 1. Test SendMessage
	res, err := server.SendMessage(ctx, connect.NewRequest(&messagingv1.SendMessageRequest{
		Recipient: "user1",
		Content: "msg",
	}))
	if err != nil || res.Msg.MessageId == "" {
		t.Error("SendMessage failed")
	}
}
