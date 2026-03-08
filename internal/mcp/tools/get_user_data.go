package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	cranemcp "github.com/AustinMCrane/cranekit/mcp"
)

// Register wires all tools onto the given server.
func Register(s *cranemcp.MCPServer) {
	registerGetUserData(s)
}

// getJSON performs an authenticated GET and decodes JSON into out.
func getJSON(client *http.Client, url string, out any) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("GET %s: %d %s", url, resp.StatusCode, string(body))
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
