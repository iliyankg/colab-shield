package httpserver

import (
	"fmt"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

// ColabShieldServer is the main HTTP server for Colab Shield
type ColabShieldServer struct {
	ginEngine   *gin.Engine
	redisClient *redis.Client
}

func NewColabShieldServer(redisClient *redis.Client) *ColabShieldServer {
	ginEngine := gin.New()
	ginEngine.Use(requestid.New())
	ginEngine.Use(newGinLoggingHandler())
	ginEngine.Use(gin.Recovery())

	ginEngine.GET("/health", newGinHealthHandler(redisClient))
	ginEngine.POST("/files/claim")

	return &ColabShieldServer{
		ginEngine,
		redisClient,
	}
}

func (css *ColabShieldServer) Serve(port int) error {
	log.Info().Msgf("Http listening on port: %d", port)
	return css.ginEngine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
