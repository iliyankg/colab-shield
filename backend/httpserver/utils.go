package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/iliyankg/colab-shield/backend/domain"
	"github.com/iliyankg/colab-shield/backend/httpserver/internal/protocol"
	"github.com/rs/zerolog"
)

// getLogger returns the logger from the context.
func getLogger(ctx *gin.Context) zerolog.Logger {
	return ctx.MustGet(LoggerCtxKey).(zerolog.Logger)
}

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*domain.FileInfo, outTarget *[]*protocol.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, protocol.NewFileInfoFromModel(fi))
	}
}

// newDomainClaimRequest converts a protocol.Claim to a domain.ClaimRequest
func newDomainClaimRequest(cr *protocol.Claim) domain.ClaimRequest {
	files := make([]domain.FileClaim, 0, len(cr.Files))
	for _, file := range cr.Files {
		files = append(files, domain.NewFileClaim(file.FileId, file.FileHash, domain.ClaimMode(file.ClaimMode)))
	}
	return domain.NewClaimRequest(cr.BranchName, cr.SoftClaim, files)
}

// newDomainUpdateRequest converts a protocol.Update to a domain.UpdateRequest
func newDomainUpdateRequest(cr *protocol.Update) domain.UpdateRequest {
	files := make([]domain.FileUpdate, 0, len(cr.Files))
	for _, file := range cr.Files {
		files = append(files, domain.NewFileUpdate(file.FileId, file.FileHash, file.OldHash))
	}
	return domain.NewUpdateRequest(cr.BranchName, files)
}
