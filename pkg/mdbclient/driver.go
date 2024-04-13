package mdbclient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type client struct {
	cli *redis.Client
}

// NewClient initializes a new DB driver using the provided database configuration.
func NewClient(host string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

// Incr implements Cache.Incr
func (c *client) Incr(ctx context.Context, key string) int64 {
	return c.cli.Incr(ctx, key).Val()
}

// Expire implements Cache.Expire
func (c *client) Expire(ctx context.Context, key string, duration time.Duration) {
	c.cli.Expire(ctx, key, duration)
}

// Close closes the database client
func (c *client) Close() error {
	return c.cli.Close()
}
