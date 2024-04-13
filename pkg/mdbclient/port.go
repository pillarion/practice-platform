package mdbclient

import "context"

type Cache interface {
	Incr(ctx context.Context, key string) int64
	Expire(ctx context.Context, key string, duration int)
	Close() error
}
