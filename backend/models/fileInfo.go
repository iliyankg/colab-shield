package models

import (
	"encoding/json"
	"errors"

	pb "github.com/iliyankg/colab-shield/protos"
)

var (
	ErrFileAlreadyClaimed = errors.New("file already claimed")
	ErrFileOutOfDate      = errors.New("file out of date")
	ErrUserNotOwner       = errors.New("user not owner")
	ErrInvalidClaimMode   = errors.New("invalid claim mode")
)

// FileInfo represents a file in the system
type FileInfo struct {
	FileId     string       `json:"fileId"`
	FileHash   string       `json:"fileHash"`
	UserIds    []string     `json:"userIds"`
	BranchName string       `json:"branchName"`
	ClaimMode  pb.ClaimMode `json:"mode"`
}

// NewFileInfo creates a new blank FileInfo with just a File ID
func NewFileInfo(fileId string, fileHash string, branchName string) *FileInfo {
	return &FileInfo{
		FileId:     fileId,
		FileHash:   fileHash,
		UserIds:    []string{},
		BranchName: branchName,
		ClaimMode:  pb.ClaimMode_UNCLAIMED,
	}
}

// Claim claims a file for a user
func (fi *FileInfo) Claim(userId string, fileHash string, claimMode pb.ClaimMode) error {
	if claimMode == pb.ClaimMode_UNCLAIMED {
		return ErrInvalidClaimMode
	}

	// Already claimed by someone else.
	if fi.ClaimMode == pb.ClaimMode_EXCLUSIVE {
		return ErrFileAlreadyClaimed
	}

	if fi.FileHash != fileHash {
		return ErrFileOutOfDate
	}

	if fi.ClaimMode == pb.ClaimMode_SHARED && claimMode == pb.ClaimMode_EXCLUSIVE {
		return ErrInvalidClaimMode
	}

	if err := fi.addOwner(userId); err != nil {
		return err
	}

	fi.ClaimMode = claimMode

	return nil
}

// Update updates the file hash for a file only if the user is an owner
func (fi *FileInfo) Update(userId string, oldHash string, fileHash string, branchName string) error {
	if !fi.CheckOwner(userId) {
		return ErrUserNotOwner
	}

	if fi.FileHash != oldHash {
		return ErrFileOutOfDate
	}

	fi.FileHash = fileHash
	fi.BranchName = branchName
	return nil
}

// Release releases a file from a user if they are an owner
func (fi *FileInfo) Release(userId string) error {
	if err := fi.removeOwner(userId); err != nil {
		return err
	}

	if len(fi.UserIds) == 0 {
		fi.ClaimMode = pb.ClaimMode_UNCLAIMED
	}

	return nil
}

// CheckOwner checks if a user is an owner of a file
func (fi *FileInfo) CheckOwner(userId string) bool {
	for _, id := range fi.UserIds {
		if id == userId {
			return true
		}
	}
	return false
}

// addOwner adds a userId to the FileInfo
// Adding an owner can only happen through claiming.
func (fi *FileInfo) addOwner(userId string) error {
	if fi.CheckOwner(userId) {
		return nil
	}

	fi.UserIds = append(fi.UserIds, userId)

	return nil
}

// removeOwner removes a userId from the FileInfo
func (fi *FileInfo) removeOwner(userId string) error {
	for i, id := range fi.UserIds {
		if id == userId {
			fi.UserIds = append(fi.UserIds[:i], fi.UserIds[i+1:]...)
			return nil
		}
	}

	return ErrUserNotOwner
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
		UserIds:    fi.UserIds,
		BranchName: fi.BranchName,
		ClaimMode:  fi.ClaimMode,
	}
}

// FileInfosToProto converts a slice of FileInfo to a slice of pb.FileInfo
func FileInfosToProto(fileInfos []*FileInfo, outTarget *[]*pb.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, fi.ToProto())
	}
}
