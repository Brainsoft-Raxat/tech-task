package ctxconst

import (
	"context"
)

type CtxKey string

const (
	UserIDKey CtxKey = "user_id"
)

func GetUserID(ctx context.Context) string {
	return ctx.Value(UserIDKey).(string)
}

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}
