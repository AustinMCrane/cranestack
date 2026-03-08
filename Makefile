.PHONY: build run-api run-mcp clean

build:
	go build -o bin/api ./cmd/api
	go build -o bin/mcp ./cmd/mcp

run-api:
	DEV_SESSION_TOKEN=dev-session-token go run ./cmd/api

# Usage: make run-mcp PAT=<your-pat-here>
run-mcp:
	MCP_PAT_TOKEN=$(PAT) MCP_TRANSPORT=sse go run ./cmd/mcp

clean:
	rm -rf bin/
