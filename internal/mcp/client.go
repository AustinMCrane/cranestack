package mcp

import (
	"fmt"
	"net/http"
	"os"
)

// NewAuthenticatedClient returns an *http.Client whose transport automatically
// attaches an Authorization: Bearer header using MCP_PAT_TOKEN.
func NewAuthenticatedClient() *http.Client {
	token := os.Getenv("MCP_PAT_TOKEN")
	return &http.Client{
		Transport: &bearerTransport{
			token: token,
			inner: http.DefaultTransport,
		},
	}
}

// bearerTransport injects the Bearer token into every outgoing request.
type bearerTransport struct {
	token string
	inner http.RoundTripper
}

func (t *bearerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	clone := req.Clone(req.Context())
	clone.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
	return t.inner.RoundTrip(clone)
}
