package ratelimiter

import (
	"github.com/redis/go-redis/v9"
)

type limiter struct {
	MaxHit  int
	counter *redis.Client
}
