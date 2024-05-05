package httpserver

import (
	"context"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const version = "0.1.0" // TODO: Get this from the build system

func newGinHealthHandler(redisClient *redis.Client) gin.HandlerFunc {
	checker := health.NewChecker(
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

	return func(c *gin.Context) {
		httpHandler := health.NewHandler(checker)
		httpHandler(c.Writer, c.Request)
	}
}
