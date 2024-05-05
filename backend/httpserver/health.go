package httpserver

import (
	"context"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/redis/go-redis/v9"
)

const version = "0.1.0" // TODO: Get this from the build system

// newHealthChecker builds the health checker for the HTTP server.
func newHealthChecker(redisClient *redis.Client) health.Checker {
	return health.NewChecker(
		health.WithInfo(map[string]any{
			"version": version,
		}),
		health.WithCheck(health.Check{
			Name:    "redis",
			Timeout: 3 * time.Second,
			Check: func(ctx context.Context) error {
				return redisClient.Ping(ctx).Err()
			},
		}),
	)
}
