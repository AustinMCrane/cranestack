package main

import (
	"log"
	"os"

	"github.com/AustinMCrane/toedoe/internal/mcp"
)

func main() {
	if os.Getenv("MCP_PAT_TOKEN") == "" {
		log.Fatal("MCP_PAT_TOKEN environment variable is required")
	}

	srv := mcp.NewServer()

	transport := os.Getenv("MCP_TRANSPORT")
	if transport == "" {
		transport = "sse"
	}
	log.Printf("MCP server starting with %s transport", transport)

	if err := srv.Start(); err != nil {
		log.Fatalf("mcp server: %v", err)
	}
}
