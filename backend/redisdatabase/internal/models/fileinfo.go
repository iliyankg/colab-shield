package models

import (
	"encoding/json"
	"errors"

	"github.com/iliyankg/colab-shield/backend/domain"
)

var (
	ErrFileInfoCorrupted = errors.New("file info in database does not meet the invariants of the domain")
)

// FileInfo represents a file in the system
type FileInfo struct {
	FileId     string   `json:"fileId"`
	FileHash   string   `json:"fileHash"`
	UserIds    []string `json:"userIds"`
	BranchName string   `json:"branchName"`
	ClaimMode  int32    `json:"claimMode"`
}

// NewFileInfo creates a new FileInfo that meets the invariants of the struct
func NewFileInfo(fi *domain.FileInfo) *FileInfo {
	return &FileInfo{
		FileId:     fi.GetFileId(),
		FileHash:   fi.GetFileHash(),
		UserIds:    fi.GetUserIds(),
		BranchName: fi.GetBranchName(),
		ClaimMode:  int32(fi.GetClaimMode()),
	}
}

// NewBlankFileInfo creates a new blank without the need for parameters.
//
// Ideal for creating a new FileInfo that will be populated later on such as
// unmarshalling from a database.
func NewBlankFileInfo() *FileInfo {
	return &FileInfo{
		FileId:     "",
		FileHash:   "",
		UserIds:    []string{},
		BranchName: "",
		ClaimMode:  0,
	}
}

// ToDomain converts the FileInfo to a domain.FileInfo
func (fi *FileInfo) ToDomain() (*domain.FileInfo, error) {
	toReturn, err := domain.NewFullFileInfo(fi.FileId, fi.FileHash, fi.UserIds, fi.BranchName, domain.ClaimMode(fi.ClaimMode), domain.RejectReason_None)
	if err != nil { // Can only return one kind of error here so its a straight forward substitution
		return nil, ErrFileInfoCorrupted
	}
	return toReturn, nil
}

// Implements encoding.BinaryMarshaler interface for FileInfo
// Needed for redis client to marshal FileInfo
// Deliberately pass by value not by pointer!
func (fi FileInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(fi)
}
