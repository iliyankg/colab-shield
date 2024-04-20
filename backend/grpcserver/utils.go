package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/colabom"
	"github.com/iliyankg/colab-shield/protos"
)

// Represents a type constraint for FileId field and getter
// in the respective proto messages. Used for generics.
type protoFileId interface {
	*protos.ClaimFileInfo | *protos.UpdateFileInfo
	GetFileId() string
}

// keysFromFileRequests extracts the file IDs from the given file requests
// and builds the Redis keys for them
func keysFromFileRequests[T protoFileId](projectId string, pbFiles []T, outKeys *[]string) {
	for _, file := range pbFiles {
		fileId := file.GetFileId()
		*outKeys = append(*outKeys, colabom.BuildRedisKeyForFile(projectId, fileId))
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
