package tools

import (
"context"
"fmt"

cranemcp "github.com/AustinMCrane/cranekit/mcp"
)

func registerGetUserData(s *cranemcp.MCPServer) {
s.RegisterTool(
"get_user_data",
"Fetches data from the CraneStack REST API.",
nil,
func(ctx context.Context, args map[string]string) (string, error) {
var result any
if err := getJSON(s.Client(), fmt.Sprintf("%s/api/data", s.APIBaseURL()), &result); err != nil {
return "", err
}
return fmt.Sprintf("%v", result), nil
},
)
}
