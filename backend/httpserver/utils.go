package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/iliyankg/colab-shield/backend/core/requests"
	"github.com/iliyankg/colab-shield/backend/httpserver/protocol"
	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/rs/zerolog"
)

// getLogger returns the logger from the context.
func getLogger(ctx *gin.Context) zerolog.Logger {
	return ctx.MustGet(LoggerCtxKey).(zerolog.Logger)
}

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*models.FileInfo, outTarget *[]*protocol.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, protocol.NewFileInfoFromModel(fi))
	}
}

func newCoreClaimRequest(claimRequest *protocol.Claim) *requests.Claim {
	files := make([]*requests.ClaimFileInfo, 0, len(claimRequest.Files))
	for i, file := range claimRequest.Files {
		files[i] = &requests.ClaimFileInfo{
			FileId:    file.FileId,
			FileHash:  file.FileHash,
			ClaimMode: requests.ClaimMode(file.ClaimMode),
		}
	}
	return &requests.Claim{
		BranchName: claimRequest.BranchName,
		SoftClaim:  claimRequest.SoftClaim,
		Files:      files,
	}
}

func newCoreUpdateRequest(claimRequest *protocol.Update) *requests.Update {
	files := make([]*requests.UpdateFileInfo, 0, len(claimRequest.Files))
	for i, file := range claimRequest.Files {
		files[i] = &requests.UpdateFileInfo{
			FileId:   file.FileId,
			FileHash: file.FileHash,
		}
	}
	return &requests.Update{
		BranchName: claimRequest.BranchName,
		Files:      files,
	}
}
