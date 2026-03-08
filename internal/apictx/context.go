package apictx

import "context"

type contextKey string

const userIDContextKey contextKey = "userID"

// WithUserID stores a user ID in the context.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

// UserIDFromContext retrieves the authenticated user's ID from the context.
func UserIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userIDContextKey).(string)
	return id, ok
}
