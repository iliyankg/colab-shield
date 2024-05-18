package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/domain"
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
func fileInfosToProto(fileInfos []*domain.FileInfo, outTarget *[]*protos.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, toProto(fi))
	}
}

// ToProto converts a FileInfo to a protos.FileInfo
func toProto(fi *domain.FileInfo) *protos.FileInfo {
	return &protos.FileInfo{
		FileId:       fi.GetFileId(),
		FileHash:     fi.GetFileHash(),
		UserIds:      fi.GetUserIds(),
		BranchName:   fi.GetBranchName(),
		ClaimMode:    protos.ClaimMode(fi.GetClaimMode()),
		RejectReason: protos.RejectReason(fi.GetRejectReason()),
	}
}

func newClaimRequest(claimRequest *protos.ClaimFilesRequest) domain.ClaimRequest {
	files := make([]domain.FileClaim, 0, len(claimRequest.Files))
	for _, file := range claimRequest.Files {
		files = append(files, domain.NewFileClaim(file.FileId, file.FileHash, domain.ClaimMode(file.ClaimMode)))
	}

	return domain.NewClaimRequest(claimRequest.BranchName, claimRequest.SoftClaim, files)
}

func newUpdateRequest(req *protos.UpdateFilesRequest) domain.UpdateRequest {
	files := make([]domain.FileUpdate, 0, len(req.Files))
	for _, file := range req.Files {
		files = append(files, domain.NewFileUpdate(file.FileId, file.FileHash, file.OldHash))
	}

	return domain.NewUpdateRequest(req.BranchName, files)
}
