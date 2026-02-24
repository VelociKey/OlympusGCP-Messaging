package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"mcp-go/mcp"

	messagingv1connect "OlympusGCP-Messaging/gen/v1/messaging/messagingv1connect"
	messagingv1 "OlympusGCP-Messaging/gen/v1/messaging"
	"Olympus2/90000-Enablement-Labs/P0000-pkg/000-mcp-bridge"
)

func main() {
	s := mcpbridge.NewBridgeServer("OlympusMessagingBridge", "1.0.0")

	client := messagingv1connect.NewMessagingServiceClient(
		http.DefaultClient,
		"http://localhost:8090",
	)

	s.AddTool(mcp.NewTool("messaging_send",
		mcp.WithDescription("Send a message to another agent or user. Args: {recipient: string, content: string}"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		m, err := mcpbridge.ExtractMap(request)
		if err != nil {
			return mcpbridge.HandleError(err)
		}

		recipient, _ := m["recipient"].(string)
		content, _ := m["content"].(string)

		resp, err := client.SendMessage(ctx, connect.NewRequest(&messagingv1.SendMessageRequest{
			Recipient: recipient,
			Content:   content,
		}))
		if err != nil {
			return mcpbridge.HandleError(err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Message sent successfully. ID: %s", resp.Msg.MessageId)), nil
	})

	s.Run()
}
