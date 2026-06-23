package ctxkeys

import "context"

type storeIdKey struct{}

func WithStoreID(ctx context.Context, id uint32) context.Context {
	return context.WithValue(ctx, storeIdKey{}, id)
}

func StoreID(ctx context.Context) uint32 {
	v, ok := ctx.Value(storeIdKey{}).(uint32)
	if !ok {
		return 0
	}
	return v
}
