package httpserver

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/backend/httpserver/protocol"
	"github.com/rs/zerolog"
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
	ginEngine.Use(zerologGinHandler())
	ginEngine.Use(gin.Recovery())

	return &ColabShieldServer{
		ginEngine,
		redisClient,
	}
}

func (css *ColabShieldServer) Serve(port int) error {
	css.ginEngine.GET("/health", css.createHealthHandler())

	private := css.ginEngine.Group("/project")
	private.Use(annonAuthHandler()) // FIXME: Should not be using annon auth.
	private.Use(zerologAuthRequestHandler())
	{
		private.GET("/:projectId/files", css.listHandler)
		private.POST("/:projectId/files/claim", css.claimHandler)
		private.PATCH("/:projectId/files/update", css.updateHandler)
		private.PATCH("/:projectId/files/release", css.releaseHandler)
	}

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
	projectId := ctx.Param("projectId") // TODO: Validate projectId with DB.
	userId := ctx.GetString("userId")

	logger := zerolog.Ctx(ctx).
		With().
		Str("projectId", projectId).
		Str("branchName", ctx.Query("branchName")).
		Logger()

	// TODO: Look into binding.
	var protoRequest protocol.Claim
	if err := json.NewDecoder(ctx.Request.Body).Decode(&protoRequest); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from request body")
		ctx.JSON(400, gin.H{
			"error": "Failed to unmarshal JSON from request body",
		})
		return
	}

	coreReq := newCoreClaimRequest(&protoRequest)
	rejectedFiles, err := core.Claim(ctx, logger, css.redisClient, userId, projectId, coreReq)
	switch {
	case errors.Is(err, core.ErrRejectedFiles):
		protoRejectedFiles := make([]*protocol.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)
		ctx.JSON(409, protoRejectedFiles) // 409 Conflict
	case err != nil:
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	default:
		ctx.JSON(200, gin.H{})
	}
}

func (css *ColabShieldServer) updateHandler(ctx *gin.Context) {
	projectId := ctx.Param("projectId") // TODO: Validate projectId with DB.
	userId := ctx.GetString("userId")

	logger := zerolog.Ctx(ctx).
		With().
		Str("projectId", projectId).
		Str("branchName", ctx.Query("branchName")).
		Logger()

	// TODO: Look into binding
	var protoRequest protocol.Update
	if err := json.NewDecoder(ctx.Request.Body).Decode(&protoRequest); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from request body")
		ctx.JSON(400, gin.H{
			"error": "Failed to unmarshal JSON from request body",
		})
		return
	}

	coreReq := newCoreUpdateRequest(&protoRequest)
	rejectedFiles, err := core.Update(ctx, logger, css.redisClient, userId, projectId, coreReq)
	switch {
	case errors.Is(err, core.ErrRejectedFiles):
		protoRejectedFiles := make([]*protocol.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)
		ctx.JSON(409, protoRejectedFiles) // 409 Conflict
	case err != nil:
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	default:
		ctx.JSON(200, gin.H{})
	}
}

func (css *ColabShieldServer) releaseHandler(ctx *gin.Context) {
	projectId := ctx.Param("projectId") // TODO: Validate projectId with DB.
	userId := ctx.GetString("userId")

	logger := zerolog.Ctx(ctx).
		With().
		Str("projectId", projectId).
		Str("branchName", ctx.Query("branchName")).
		Logger()

	body := ctx.Request.Body
	var protoRequest protocol.Release
	if err := json.NewDecoder(body).Decode(&protoRequest); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from request body")
		ctx.JSON(400, gin.H{
			"error": "Failed to unmarshal JSON from request body",
		})
		return
	}

	rejectedFiles, err := core.Release(ctx, logger, css.redisClient, userId, projectId, protoRequest.BranchName, protoRequest.FileIds)
	switch {
	case errors.Is(err, core.ErrRejectedFiles):
		protoRejectedFiles := make([]*protocol.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)
		ctx.JSON(409, protoRejectedFiles) // 409 Conflict
	case err != nil:
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	default:
		ctx.JSON(200, gin.H{})
	}
}

func (css *ColabShieldServer) listHandler(ctx *gin.Context) {
	projectId := ctx.Param("projectId") // TODO: Validate projectId with DB.

	logger := zerolog.Ctx(ctx).
		With().
		Str("projectId", projectId).
		Str("branchName", ctx.Query("branchName")).
		Logger()

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
