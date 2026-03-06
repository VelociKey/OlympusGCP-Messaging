package main

import (
	"context"
	"olympus.fleet/00SDLC/OlympusForge/70000-Environmental-Harness/dagger/olympusgcp-messaging/internal/dagger"
)

type OlympusGCPMessaging struct{}

func (m *OlympusGCPMessaging) HelloWorld(ctx context.Context) string {
	return "Hello from OlympusGCP-Messaging!"
}

func main() {
	dagger.Serve()
}
