package mdbclient

import (
	"context"
	"time"
)

type Cache interface {
	Incr(ctx context.Context, key string) int64
	Expire(ctx context.Context, key string, duration time.Duration)
	Close() error
}
