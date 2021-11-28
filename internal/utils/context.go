package utils

import "context"

type contextKey int

const (
	keyPrincipalID contextKey = iota
)

// ContextWithPrincipalID set the principal ID value in the context.
func ContextWithPrincipalID(ctx context.Context, principalID string) context.Context {
	newCtx := context.WithValue(ctx, keyPrincipalID, principalID)
	return newCtx
}

// GetPrincipalIDFromContext gets the principal ID value from the context.
func GetPrincipalIDFromContext(ctx context.Context) (string, bool) {
	ID, ok := ctx.Value(keyPrincipalID).(string)
	return ID, ok
}
