package server

import (
	"context"
	"fmt"

	pb "github.com/iliyankg/colab-shield/protos"
)

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
	return ctx.Value("userId").(string)
}

// projectIdFromCtx extracts the project ID from the context
// projectId is required and verified to exist in the UnaryInterceptor
func projectIdFromCtx(ctx context.Context) string {
	return ctx.Value("projectId").(string)
}
