package httpserver

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/backend/httpserver/protocol"
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
	css.ginEngine.GET("/project/:projectId/files", css.listHandler)
	css.ginEngine.POST("/project/:projectId/files/claim", css.claimHandler)
	css.ginEngine.PATCH("/project/:projectId/files/update", css.updateHandler)
	css.ginEngine.PATCH("/project/:projectId/files/release", css.releaseHandler)

	log.Info().Msgf("Http listening on port: %d", port)
	return css.ginEngine.Run(fmt.Sprintf("0.0.0.0:%d", port))
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

func (css *ColabShieldServer) claimHandler(ctx *gin.Context) {
	// TODO: Implement
}

func (css *ColabShieldServer) updateHandler(ctx *gin.Context) {
	// TODO: Implement
}

func (css *ColabShieldServer) releaseHandler(ctx *gin.Context) {
	// TODO: Implement
}

func (css *ColabShieldServer) listHandler(ctx *gin.Context) {
	logger := getLogger(ctx)

	// TODO: Validate projectId with DB.
	projectId := ctx.Param("projectId")

	cursor, err := strconv.ParseUint(ctx.Query("cursor"), 10, 64)
	if err != nil {
		cursor = 0
	}

	pageSize, err := strconv.ParseInt(ctx.Query("pageSize"), 10, 64)
	if err != nil {
		pageSize = 100
	}

	pathStr := ""
	path, err := base64.StdEncoding.DecodeString(ctx.Query("path"))
	if err == nil {
		pathStr = string(path)
	}

	files, cursor, err := core.List(ctx, logger, css.redisClient, projectId, cursor, pageSize, pathStr)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	protoFiles := make([]*protocol.FileInfo, 0, len(files))
	fileInfosToProto(files, &protoFiles)
	ctx.JSON(200, gin.H{
		"nextCursor": cursor,
		"files":      protoFiles,
	})
}
