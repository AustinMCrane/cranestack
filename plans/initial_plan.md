# Initial Plan for System Architecture

I am building a system with three main components: a Golang REST API (backed by SQLite), a Golang MCP Server, and an iOS SwiftUI app. The auth flow relies on Apple/Google login on iOS, which passes a JWT to the Go API to create a session. The iOS app can also generate a Personal Access Token (PAT) from the Go API, which is manually given to the Go MCP server to authenticate AI requests.

Your task is to generate the boilerplate and shell only. Do not implement business logic. Create the directory structures, routing, interface definitions, and empty function stubs.

## Phase 1: The Database & Schema Shell

Generate the SQLite database initialization code and schema definitions in Golang.

Create a setup function that initializes an SQLite database file.

Write the SQL CREATE TABLE statements for two tables: users (id, email, created_at) and api_keys (id, user_id, key_hash, label, created_at).

Create a Go repository struct with empty stub methods for: GetUserByID, CreateUser, StoreAPIKeyHash, and ValidateAPIKey. Use database/sql.

## Phase 2: The Golang REST API Boilerplate

Generate the boilerplate for the Golang REST API using a standard router (like chi or the standard library net/http in Go 1.22+).

Set up the main server struct and start function.

Create an authentication middleware stub that checks for a Bearer token (leave the actual JWT validation logic empty, just parse the header).

Create empty handler stubs for: POST /auth/login (mobile login), POST /auth/generate-mcp-key (returns a PAT), and a protected GET /api/data endpoint.

Wire the routes together so the API can start and listen on a port."

## Phase 3: The Golang MCP Server Boilerplate

Generate the shell for the Golang MCP Server. It will act as a bridge between an AI agent and the REST API.

Use a standard Go MCP SDK/implementation. Set up the basic server initialization.

Read an environment variable called MCP_PAT_TOKEN.

Define an HTTP client instance that automatically attaches Authorization: Bearer {MCP_PAT_TOKEN} to all requests.

Register one mock MCP tool (e.g., get_user_data) that makes an empty GET request to the REST API and returns a dummy string. Do not implement complex logic.

## Phase 4: The iOS SwiftUI App Shell

"Generate the boilerplate for the iOS SwiftUI app. Focus on the architecture and views, not the UI styling.

Create an AuthManager class (ObservableObject) with empty stubs for loginWithApple(), loginWithGoogle(), and generateMCPAccessToken().

Create a basic App struct that routes to a LoginView if unauthenticated, or a MainTabView if authenticated.

In the MainTabView, create a settings tab that includes a button to trigger generateMCPAccessToken() and a text element to display the generated token to the user for copying.

Create an APIClient class stub that handles making authenticated requests to the Go REST API.
