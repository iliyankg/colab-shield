package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/iliyankg/colab-shield/protos"
)

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

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*models.FileInfo, outTarget *[]*protos.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, toProto(fi))
	}
}

// ToProto converts a FileInfo to a protos.FileInfo
func toProto(fi *models.FileInfo) *protos.FileInfo {
	return &protos.FileInfo{
		FileId:       fi.FileId,
		FileHash:     fi.FileHash,
		UserIds:      fi.UserIds,
		BranchName:   fi.BranchName,
		ClaimMode:    protos.ClaimMode(fi.ClaimMode),
		RejectReason: protos.RejectReason(fi.RejectReason),
	}
}
