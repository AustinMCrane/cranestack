# ToeDoe — Todo List App Plan

## Problem
Add minimal todo list functionality on top of the existing boilerplate platform (REST API + MCP server + iOS app). The boilerplate already has auth scaffolding, routing, middleware, SQLite, and MCP transport. We need to wire up the missing auth pieces and add the actual todo domain.

## Approach
Build in layers: finish auth → add todos to DB → add API routes → add MCP tools → build iOS UI. Keep each layer thin — only the code needed to do the job.

---

## Todos

### 1. Finish Auth Wiring
The boilerplate has session/token middleware but several key pieces are TODOs:
- Implement `ValidateAPIKey` in `internal/db/repository.go` (SHA-256 hash raw key, compare to stored hash)
- Implement `Login` handler in `internal/api/handlers/auth.go` (parse Apple identity token, create/lookup user, return session token)
- Implement `loginWithApple()` in `ios/ToeDoe/AuthManager.swift` (Apple Sign-In credential flow → POST to /auth/login)

### 2. Add Todos Table & Repository
- Add `todos` table to `internal/db/db.go`:
  ```sql
  CREATE TABLE IF NOT EXISTS todos (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id),
    title TEXT NOT NULL,
    completed INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
  );
  ```
- Add repository methods in `internal/db/repository.go`:
  - `CreateTodo(userID, title string) (*Todo, error)`
  - `ListTodos(userID string) ([]*Todo, error)`
  - `UpdateTodo(id, userID, title string, completed bool) (*Todo, error)`
  - `DeleteTodo(id, userID string) error`
  - `GetTodoByID(id, userID string) (*Todo, error)`
- Add `Todo` model struct

### 3. Add Todo API Routes & Handlers
New protected routes under `/api/todos`:
- `GET    /api/todos`         — list user's todos
- `POST   /api/todos`         — create todo `{ "title": "..." }`
- `PUT    /api/todos/{id}`    — update todo `{ "title": "...", "completed": true }`
- `DELETE /api/todos/{id}`    — delete todo

New files:
- `internal/api/handlers/todos.go` — handler funcs
- Wire routes in `internal/api/routes.go`

### 4. Add MCP Tools for Todos
Replace the `get_user_data` stub with focused todo tools:
- `list_todos`    — list the user's todos (no input)
- `create_todo`  — create a todo (`title` string param)
- `complete_todo` — mark a todo complete (`id` string param)
- `delete_todo`  — delete a todo (`id` string param)

Each tool hits the corresponding API route using the authenticated HTTP client.

New files:
- `internal/mcp/tools/list_todos.go`
- `internal/mcp/tools/create_todo.go`
- `internal/mcp/tools/complete_todo.go`
- `internal/mcp/tools/delete_todo.go`
- Remove/replace `internal/mcp/tools/get_user_data.go`
- Update `internal/mcp/server.go` to register the new tools

### 5. iOS Todo UI
Replace the placeholder Home tab with a real todo list:
- `ios/ToeDoe/Models/Todo.swift` — Codable Todo model matching API
- `ios/ToeDoe/ViewModels/TodoViewModel.swift` — fetch/create/complete/delete, holds `[Todo]` state
- `ios/ToeDoe/Views/HomeView.swift` — SwiftUI list with:
  - Rows: checkbox + title, swipe-to-delete
  - Toolbar "+" button → sheet or inline field for new todo
  - Pull-to-refresh
- Wire `HomeView` into `MainTabView`

---

## Constraints / Notes
- Keep auth simple for dev: `DEV_SESSION_TOKEN` bypass stays, Apple Sign-In only in prod
- No pagination needed yet — todos list is unbounded for now
- No due dates, priorities, or labels in v1 — title + completed only
- MCP tools call the REST API (not DB directly) to stay DRY
- iOS base URL stays `http://localhost:8080` for local dev
