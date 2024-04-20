package models

import (
	"encoding/json"
	"errors"

	"github.com/iliyankg/colab-shield/protos"
)

var (
	ErrFileAlreadyClaimed = errors.New("file already claimed")
	ErrFileOutOfDate      = errors.New("file out of date")
	ErrUserNotOwner       = errors.New("user not owner")
	ErrInvalidClaimMode   = errors.New("invalid claim mode")
	ErrorFileNotMissing   = errors.New("file not missing")
)

// FileInfo represents a file in the system
//
// TODO: Look into: https://pkg.go.dev/github.com/redis/rueidis/om#section-readme
// TODO: Look into: https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields
type FileInfo struct {
	Key          string              `json:"key" redis:",key"` // the redis:",key" is required to indicate which field is the ULID key
	Ver          int64               `json:"ver" redis:",ver"` // the redis:",ver" is required to do optimistic locking to prevent lost update
	FileId       string              `json:"fileId"`
	FileHash     string              `json:"fileHash"`
	UserIds      []string            `json:"userIds"`
	BranchName   string              `json:"branchName"`
	ClaimMode    protos.ClaimMode    `json:"mode"`
	RejectReason protos.RejectReason `json:"-"` // RejectReason is not stored in db
}

// NewFileInfo creates a new FileInfo that meets the invariants of the struct
func NewFileInfo(fileId string, fileHash string, branchName string) *FileInfo {
	return &FileInfo{
		FileId:       fileId,
		FileHash:     fileHash,
		UserIds:      []string{},
		BranchName:   branchName,
		ClaimMode:    protos.ClaimMode_UNCLAIMED,
		RejectReason: protos.RejectReason_NONE,
	}
}

// NewBlankFileInfo creates a new blank without the need for parameters.
//
// Ideal for creating a new FileInfo that will be populated later on such as
// unmarshalling from a database.
func NewBlankFileInfo() *FileInfo {
	return &FileInfo{
		FileId:       "",
		FileHash:     "",
		UserIds:      []string{},
		BranchName:   "",
		ClaimMode:    protos.ClaimMode_UNCLAIMED,
		RejectReason: protos.RejectReason_NONE,
	}
}

// NewMissingFileInfo creates a new FileInfo for a file that is missing.
func NewMissingFileInfo(fileId string) *FileInfo {
	return &FileInfo{
		FileId:       fileId,
		FileHash:     "",
		UserIds:      []string{},
		BranchName:   "",
		ClaimMode:    protos.ClaimMode_UNCLAIMED,
		RejectReason: protos.RejectReason_MISSING,
	}
}

// UpgradeMissingToNew upgrades a missing file to a new file setting up the object invariants.
// Can error if the file is not missing.
func (fi *FileInfo) UpgradeMissingToNew(fileHash string, branchName string) error {
	if fi.RejectReason != protos.RejectReason_MISSING {
		return ErrorFileNotMissing
	}

	fi.FileHash = fileHash
	fi.BranchName = branchName
	fi.RejectReason = protos.RejectReason_NONE

	return nil
}

// Claim claims a file for a user
func (fi *FileInfo) Claim(userId string, fileHash string, claimMode protos.ClaimMode) error {
	if claimMode == protos.ClaimMode_UNCLAIMED {
		return fi.invalidClaimMode()
	}

	// Already claimed by someone else.
	if fi.ClaimMode == protos.ClaimMode_EXCLUSIVE {
		return fi.alreadyClaimed()
	}

	if fi.FileHash != fileHash {
		return fi.fileOutOfDate()
	}

	if fi.ClaimMode == protos.ClaimMode_SHARED && claimMode == protos.ClaimMode_EXCLUSIVE {
		return fi.invalidClaimMode()
	}

	fi.addOwner(userId)
	fi.ClaimMode = claimMode

	return nil
}

// Update updates the file hash for a file only if the user is an owner
func (fi *FileInfo) Update(userId string, oldHash string, fileHash string, branchName string) error {
	if !fi.CheckOwner(userId) {
		return fi.userNotOwner()
	}

	if fi.FileHash != oldHash {
		return fi.fileOutOfDate()
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
		fi.ClaimMode = protos.ClaimMode_UNCLAIMED
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
// adding an owner should only be done through claiming
func (fi *FileInfo) addOwner(userId string) {
	if fi.CheckOwner(userId) {
		return
	}

	fi.UserIds = append(fi.UserIds, userId)
}

// removeOwner removes a userId from the FileInfo
func (fi *FileInfo) removeOwner(userId string) error {
	for i, id := range fi.UserIds {
		if id == userId {
			fi.UserIds = append(fi.UserIds[:i], fi.UserIds[i+1:]...)
			return nil
		}
	}

	return fi.userNotOwner()
}

func (fi *FileInfo) invalidClaimMode() error {
	fi.RejectReason = protos.RejectReason_INVALID_CLAIM_MODE
	return ErrInvalidClaimMode
}

func (fi *FileInfo) alreadyClaimed() error {
	fi.RejectReason = protos.RejectReason_ALREADY_CLAIMED
	return ErrFileAlreadyClaimed
}

func (fi *FileInfo) fileOutOfDate() error {
	fi.RejectReason = protos.RejectReason_OUT_OF_DATE
	return ErrFileOutOfDate
}

func (fi *FileInfo) userNotOwner() error {
	fi.RejectReason = protos.RejectReason_NOT_OWNER
	return ErrUserNotOwner
}

// Implements encoding.BinaryMarshaler interface for FileInfo
// Needed for redis client to marshal FileInfo
// Deliberately pass by value not by pointer!
func (fi FileInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(fi)
}

// ToProto converts a FileInfo to a protos.FileInfo
func (fi *FileInfo) ToProto() *protos.FileInfo {
	return &protos.FileInfo{
		FileId:       fi.FileId,
		FileHash:     fi.FileHash,
		UserIds:      fi.UserIds,
		BranchName:   fi.BranchName,
		ClaimMode:    fi.ClaimMode,
		RejectReason: fi.RejectReason,
	}
}

// FileInfosToProto converts a slice of FileInfo to a slice of protos.FileInfo
func FileInfosToProto(fileInfos []*FileInfo, outTarget *[]*protos.FileInfo) {
	for _, fi := range fileInfos {
		*outTarget = append(*outTarget, fi.ToProto())
	}
}
