package mcp

import (
	"os"

	cranemcp "github.com/AustinMCrane/cranekit/mcp"
	"github.com/AustinMCrane/cranestack/internal/mcp/tools"
)

// NewServer initializes the MCP server, registers tools, and reads config from env.
func NewServer() *cranemcp.MCPServer {
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		apiBaseURL = "http://localhost:8080"
	}

	s := cranemcp.NewServer(cranemcp.ServerConfig{
		Name:       "cranestack-mcp",
		Version:    "0.1.0",
		APIBaseURL: apiBaseURL,
		PATToken:   os.Getenv("MCP_PAT_TOKEN"),
		Transport:  os.Getenv("MCP_TRANSPORT"),
		SSEPort:    9090,
	})

	tools.Register(s)

	return s
}
