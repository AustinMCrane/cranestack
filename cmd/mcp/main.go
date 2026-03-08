package main

import (
	"log/slog"
	"os"

	"github.com/AustinMCrane/cranekit/server"
	"github.com/AustinMCrane/toedoe/internal/mcp"
)

func main() {
	slog.SetDefault(server.NewLogger(os.Getenv("LOG_FORMAT")))

	if os.Getenv("MCP_PAT_TOKEN") == "" {
		slog.Error("MCP_PAT_TOKEN environment variable is required")
		os.Exit(1)
	}

	transport := os.Getenv("MCP_TRANSPORT")
	if transport == "" {
		transport = "sse"
	}

	srv := mcp.NewServer()
	slog.Info("MCP server starting", "transport", transport)

	if err := srv.Start(); err != nil {
		slog.Error("mcp server error", "err", err)
		os.Exit(1)
	}
}
