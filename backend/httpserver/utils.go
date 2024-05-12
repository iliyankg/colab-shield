package httpserver

import (
	"github.com/iliyankg/colab-shield/backend/httpserver/protocol"
	"github.com/iliyankg/colab-shield/backend/models"
)

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func fileInfosToProto(fileInfos []*models.FileInfo, outTarget *[]*protocol.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, protocol.NewFileInfoFromModel(fi))
	}
}
