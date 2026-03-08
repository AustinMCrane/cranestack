package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

// RegisterGetUserData registers the get_user_data tool on the MCP server.
func RegisterGetUserData(s *mcpserver.MCPServer, client *http.Client, apiBaseURL string) {
	tool := mcp.NewTool("get_user_data",
		mcp.WithDescription("Fetches user data from the CraneStack REST API."),
	)

	s.AddTool(tool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// TODO: replace with real endpoint and parse the response
		resp, err := client.Get(fmt.Sprintf("%s/api/data", apiBaseURL))
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		defer resp.Body.Close()

		return mcp.NewToolResultText("get_user_data stub: received response"), nil
	})
}
