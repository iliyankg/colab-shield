package httpserver

import (
	"github.com/iliyankg/colab-shield/backend/httpserver/protocol"
	"github.com/iliyankg/colab-shield/backend/models"
)

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*models.FileInfo, outTarget *[]*protocol.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, toProto(fi))
	}
}

// ToProto converts a FileInfo to a protos.FileInfo
func toProto(fi *models.FileInfo) *protocol.FileInfo {
	return &protocol.FileInfo{
		FileId:       fi.FileId,
		FileHash:     fi.FileHash,
		UserIds:      fi.UserIds,
		BranchName:   fi.BranchName,
		ClaimMode:    protocol.ClaimMode(fi.ClaimMode),
		RejectReason: protocol.RejectReason(fi.RejectReason),
	}
}
