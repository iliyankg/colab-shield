package models

import pb "github.com/iliyankg/colab-shield/protos"

type FileInfo struct {
	FileId     string `json:"fileId"`
	FileHash   string `json:"fileHash"`
	UserId     string `json:"userId"`
	BranchName string `json:"branchName"`
	Claimed    bool   `json:"claimed"`
}

func NewFileInfoFromProto(fileId string, fileHash string, userId string, branchName string, claimed bool) *FileInfo {
	return &FileInfo{
		FileId:     fileId,
		FileHash:   fileHash,
		UserId:     userId,
		BranchName: branchName,
		Claimed:    claimed,
	}
}

func (fi *FileInfo) ToProto() *pb.FileInfo {
	return &pb.FileInfo{
		FileId:     fi.FileId,
		FileHash:   fi.FileHash,
		UserId:     fi.UserId,
		BranchName: fi.BranchName,
		Claimed:    fi.Claimed,
	}
}
