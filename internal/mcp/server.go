package mcp

import (
	"context"
	"fmt"
	"os"

	"github.com/AustinMCrane/cranestack/internal/mcp/tools"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// Server wraps the MCP server with its configuration.
type Server struct {
	mcp        *mcpserver.MCPServer
	apiBaseURL string
	transport  string
	ssePort    int
}

// NewServer initializes the MCP server, registers tools, and reads config from env.
func NewServer() *Server {
	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		apiBaseURL = "http://localhost:8080"
	}

	transport := os.Getenv("MCP_TRANSPORT")
	if transport == "" {
		transport = "sse"
	}

	s := &Server{
		mcp:        mcpserver.NewMCPServer("cranestack-mcp", "0.1.0"),
		apiBaseURL: apiBaseURL,
		transport:  transport,
		ssePort:    9090,
	}

	client := NewAuthenticatedClient()
	tools.RegisterGetUserData(s.mcp, client, apiBaseURL)

	return s
}

// Start runs the server using the transport specified by MCP_TRANSPORT.
// Accepts "sse" (default) or "stdio".
func (s *Server) Start() error {
	switch s.transport {
	case "stdio":
		return mcpserver.NewStdioServer(s.mcp).Listen(context.Background(), os.Stdin, os.Stdout)
	default:
		addr := fmt.Sprintf(":%d", s.ssePort)
		return mcpserver.NewSSEServer(s.mcp, mcpserver.WithBaseURL(fmt.Sprintf("http://localhost%s", addr))).Start(addr)
	}
}
