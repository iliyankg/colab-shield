package server

import (
	"context"
	"fmt"

	pb "github.com/iliyankg/colab-shield/protos"
)

type FileIdParser func(any) string

func buildRedisKeyForFile(projectId string, fileId string) string {
	return fmt.Sprintf("project:%s:file:%s", projectId, fileId)
}

func keysFromFileRequests(projectId string, pbFiles any, outKeys *[]string) {
	claimFileInfoParser := func(file any) string {
		return file.(*pb.ClaimFileInfo).FileId
	}

	updateFileInfoParser := func(file any) string {
		return file.(*pb.UpdateFileInfo).FileId
	}

	var parser FileIdParser = nil
	switch (pbFiles).(type) {
	case []*pb.ClaimFileInfo:
		parser = claimFileInfoParser
	case []*pb.UpdateFileInfo:
		parser = updateFileInfoParser
	default:
		return
	}

	castPbFiles := pbFiles.([]*any)
	if len(castPbFiles) == 0 {
		return
	}

	for _, file := range castPbFiles {
		fileId := parser(file)
		*outKeys = append(*outKeys, buildRedisKeyForFile(projectId, fileId))
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
