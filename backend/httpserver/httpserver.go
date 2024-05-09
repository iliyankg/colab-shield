package httpserver

import (
	"context"
	"fmt"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

const version = "0.1.0" // TODO: Get this from the build system

// ColabShieldServer is the main HTTP server for Colab Shield
type ColabShieldServer struct {
	ginEngine   *gin.Engine
	redisClient *redis.Client
}

func NewColabShieldServer(redisClient *redis.Client) *ColabShieldServer {
	ginEngine := gin.New()
	ginEngine.Use(requestid.New())
	ginEngine.Use(createGinLoggingHandler())
	ginEngine.Use(gin.Recovery())

	return &ColabShieldServer{
		ginEngine,
		redisClient,
	}
}

func (css *ColabShieldServer) Serve(port int) error {
	css.ginEngine.GET("/health", css.createHealthHandler())
	css.ginEngine.POST("/files/claim", css.claimHandler)

	log.Info().Msgf("Http listening on port: %d", port)
	return css.ginEngine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}

func (css *ColabShieldServer) claimHandler(ctx *gin.Context) {
	// TODO: Implement
}

func (css *ColabShieldServer) createHealthHandler() gin.HandlerFunc {
	checker := health.NewChecker(
		health.WithInfo(map[string]any{
			"version": version,
		}),
		health.WithCheck(health.Check{
			Name:    "redis",
			Timeout: 3 * time.Second,
			Check: func(ctx context.Context) error {
				return css.redisClient.Ping(ctx).Err()
			},
		}),
	)

	return func(c *gin.Context) {
		httpHandler := health.NewHandler(checker)
		httpHandler(c.Writer, c.Request)
	}
}
