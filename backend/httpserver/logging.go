package httpserver

import (
	"github.com/gin-contrib/logger" // gin zero log middleware
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const LoggerKeyCtxKey = "logger"

// getLogger returns the logger from the context.
func getLogger(c *gin.Context) *zerolog.Logger {
	return c.MustGet(LoggerKeyCtxKey).(*zerolog.Logger)
}

// newGinLoggingHandler returns a new gin middleware that logs requests.
// Also adds a logger to the context for use in other handlers.
func newGinLoggingHandler() gin.HandlerFunc {
	loggerOption := logger.WithLogger(func(c *gin.Context, logger zerolog.Logger) zerolog.Logger {
		reqUUID := c.Request.Header.Get("X-Request-ID")
		updatedLogger := logger.Output(gin.DefaultWriter).With().Str("uuid", reqUUID).Logger()

		c.Set(LoggerKeyCtxKey, updatedLogger) // Add to context for use in other handlers.

		return updatedLogger
	})

	return logger.SetLogger(loggerOption)
}
