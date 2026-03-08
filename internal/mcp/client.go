package mcp

import (
	"net/http"
	"os"

	cranekit "github.com/AustinMCrane/cranekit/mcp"
)

// NewAuthenticatedClient returns an *http.Client that injects
// Authorization: Bearer {MCP_PAT_TOKEN} on every request.
func NewAuthenticatedClient() *http.Client {
	return cranekit.NewAuthenticatedClient(os.Getenv("MCP_PAT_TOKEN"))
}
