package api

import (
	"context"
	"fmt"
)

type ctxKey string

var ctxKey_UserID ctxKey = "userID"

// contextWithUserID augments the included context to associate a user ID with a
// well-known key
func contextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ctxKey_UserID, userID)
}

// contextUserID reads the user ID from well-known key on the input context
func contextUserID(ctx context.Context) (string, error) {
	var userIDValue = ctx.Value(ctxKey_UserID)

	if userID, ok := userIDValue.(string); !ok {
		return "", fmt.Errorf("non-string user id on context: %v", userID)
	} else if userID == "" {
		return "", fmt.Errorf("empty user id string on context")
	} else {
		return userID, nil
	}
}
