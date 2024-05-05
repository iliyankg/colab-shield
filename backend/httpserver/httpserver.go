package httpserver

import (
	"fmt"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/redis/go-redis/v9"
)

func Serve(port int, redisClient *redis.Client) error {
	ginEngine := gin.New()
	ginEngine.Use(requestid.New())
	ginEngine.Use(newGinLoggingHandler())
	ginEngine.Use(gin.Recovery())

	ginEngine.GET("/health", newGinHealthHandler(redisClient))

	log.Info().Msgf("Http listening on port: %d", port)
	return ginEngine.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
