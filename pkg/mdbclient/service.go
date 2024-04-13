package mdbclient

import (
	"context"
	"time"
)

type mdbClient struct {
	cli Cache
}

// NewMdbClient initializes a new DB driver using the provided database configuration.
func NewMdbClient(cli Cache) *mdbClient {
	return &mdbClient{
		cli: cli,
	}
}

// Incr implements Cache.Incr
func (m *mdbClient) Incr(ctx context.Context, key string) int64 {
	return m.cli.Incr(ctx, key)
}

// Expire implements Cache.Expire
func (m *mdbClient) Expire(ctx context.Context, key string, duration time.Duration) {
	m.cli.Expire(ctx, key, duration)
}

// Close closes the database client
func (m *mdbClient) Close() error {
	return m.cli.Close()
}
