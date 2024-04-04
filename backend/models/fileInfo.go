package models

import (
	"encoding/json"

	pb "github.com/iliyankg/colab-shield/protos"
)

type FileInfo struct {
	FileId     string `json:"fileId"`
	FileHash   string `json:"fileHash"`
	UserId     string `json:"userId"`
	BranchName string `json:"branchName"`
	Claimed    bool   `json:"claimed"`
}

// NewFileInfoFromProto creates a new FileInfo from a pb.FileInfo
func NewFileInfoFromProto(fileId string, fileHash string, userId string, branchName string, claimed bool) *FileInfo {
	return &FileInfo{
		FileId:     fileId,
		FileHash:   fileHash,
		UserId:     userId,
		BranchName: branchName,
		Claimed:    claimed,
	}
}

// Implements encoding.BinaryMarshaler interface for FileInfo
// Needed for redis client to marshal FileInfo
// Deliberately pass by value not by pointer!
func (fi FileInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(fi)
}

// ToProto converts a FileInfo to a pb.FileInfo
func (fi *FileInfo) ToProto() *pb.FileInfo {
	return &pb.FileInfo{
		FileId:     fi.FileId,
		FileHash:   fi.FileHash,
		UserId:     fi.UserId,
		BranchName: fi.BranchName,
		Claimed:    fi.Claimed,
	}
}

// FileInfosToProto converts a slice of FileInfo to a slice of pb.FileInfo
func FileInfosToProto(target *[]*pb.FileInfo, fileInfos []*FileInfo) {
	for _, fi := range fileInfos {
		*target = append(*target, fi.ToProto())
	}
}
