package httpserver

import (
	"github.com/gin-contrib/logger" // gin zero log middleware
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const (
	LoggerCtxKey = "logger"
	UserIdCtxKey = "userId"
)

// annonAuthHandler is a gin middleware that checks for a userId in the request header.
func annonAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("userId")
		if userId == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing userId"})
		}

		c.Set(UserIdCtxKey, userId)
		c.Next()
	}
}

// zerologGinHandler returns a new gin middleware that logs requests.
// Also adds a logger to the context for use in other handlers.
func zerologGinHandler() gin.HandlerFunc {
	loggerOption := logger.WithLogger(func(c *gin.Context, logger zerolog.Logger) zerolog.Logger {
		reqUUID := c.Request.Header.Get("X-Request-ID")
		updatedLogger := logger.
			Output(gin.DefaultWriter).
			With().
			Str("uuid", reqUUID).
			Logger()

		c.Set(LoggerCtxKey, updatedLogger) // Add to context for use in other handlers.
		return updatedLogger
	})

	return logger.SetLogger(loggerOption)
}

// zerologAuthRequestHandler is a gin middleware that updates the logger in
// the context to have an extra string field for the userId.
func zerologAuthRequestHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := getLogger(ctx)
		userId := ctx.GetString(UserIdCtxKey)
		logger = logger.With().Str("userId", userId).Logger()

		ctx.Set(LoggerCtxKey, logger)
		ctx.Next()
	}
}
