package server

import (
	"context"
	"fmt"
	"os"

	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog"
)

func buildCollabShieldContext(ctx context.Context, userId string, projectId string, method string) context.Context {
	ctx = context.WithValue(ctx, UserIdKey, userId)
	ctx = context.WithValue(ctx, ProjectIdKey, projectId)

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("userId", userId).
		Str("projectId", projectId).
		Str("method", method).
		Logger()

	return logger.WithContext(ctx)
}

func buildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}

func keysFromFileClaimRequests(target *[]string, projectId string, files []*pb.ClaimFileInfo) {
	for _, file := range files {
		*target = append(*target, buildRedisKeyForFile(projectId, file.FileId))
	}
}

// userIdFromCtx extracts the user ID from the context
// userId is required and verified to exist in the UnaryInterceptor
func userIdFromCtx(ctx context.Context) string {
	return ctx.Value(UserIdKey).(string)
}

// projectIdFromCtx extracts the project ID from the context
// projectId is required and verified to exist in the UnaryInterceptor
func projectIdFromCtx(ctx context.Context) string {
	return ctx.Value(ProjectIdKey).(string)
}
