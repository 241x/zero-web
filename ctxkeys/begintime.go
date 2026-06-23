package ctxkeys

import (
	"context"
	"time"
)

type beginTimeKey struct{}

func WithBeginTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, beginTimeKey{}, t)
}

func BeginTime(ctx context.Context) (time.Time, bool) {
	v, ok := ctx.Value(beginTimeKey{}).(time.Time)
	return v, ok
}
