package main

import (
	"context"
	"dagger/olympusgcp-messaging/internal/dagger"
)

type OlympusGCPMessaging struct{}

func (m *OlympusGCPMessaging) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Messaging!"
}

func main() {
	dagger.Serve()
}
