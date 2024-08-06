package api

import (
	"context"
	"fmt"
)

type ctxKey string

var ctxKey_UserID ctxKey = "userID"

func contextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ctxKey_UserID, userID)
}

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
