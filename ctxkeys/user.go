package ctxkeys

import "context"

type userKey struct{}

func WithUser(ctx context.Context, user any) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}

func User(ctx context.Context) any {
	return ctx.Value(userKey{})
}
