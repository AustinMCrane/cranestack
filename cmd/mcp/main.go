package main

import (
	"log/slog"
	"os"

	"github.com/AustinMCrane/toedoe/internal/mcp"
)

func main() {
	var handler slog.Handler
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}
	if os.Getenv("LOG_FORMAT") == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}
	slog.SetDefault(slog.New(handler))

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
