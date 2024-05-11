package httpserver

import (
	"github.com/iliyankg/colab-shield/backend/models"
)

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*models.FileInfo, outTarget *[]*FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, toProto(fi))
	}
}

// ToProto converts a FileInfo to a protos.FileInfo
func toProto(fi *models.FileInfo) *FileInfo {
	return &FileInfo{
		FileId:       fi.FileId,
		FileHash:     fi.FileHash,
		UserIds:      fi.UserIds,
		BranchName:   fi.BranchName,
		ClaimMode:    ClaimMode(fi.ClaimMode),
		RejectReason: RejectReason(fi.RejectReason),
	}
}
