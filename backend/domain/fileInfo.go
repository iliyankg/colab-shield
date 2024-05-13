package domain

import (
	"encoding/json"
	"errors"
)

var (
	ErrFileAlreadyClaimed = errors.New("file already claimed")
	ErrFileOutOfDate      = errors.New("file out of date")
	ErrUserNotOwner       = errors.New("user not owner")
	ErrInvalidClaimMode   = errors.New("invalid claim mode")
	ErrorFileNotMissing   = errors.New("file not missing")
)

// Enum for claim mode
// Should be kept in sync with the protobuf file.
type ClaimMode int32

const (
	Unclaimed ClaimMode = iota
	Shared
	Exclusive
)

// Enum for reject reason
// Should be kept in sync with the protobuf file.
type RejectReason int32

const (
	None RejectReason = iota
	AlreadyClaimed
	OutOfDate
	NotOwner
	InvalidClaimMode
	Missing
)

// FileInfo represents a file in the system
//
// TODO: Look into: https://pkg.go.dev/github.com/redis/rueidis/om#section-readme
// TODO: Look into: https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields
type FileInfo struct {
	FileId       string       `json:"fileId"`
	FileHash     string       `json:"fileHash"`
	UserIds      []string     `json:"userIds"`
	BranchName   string       `json:"branchName"`
	ClaimMode    ClaimMode    `json:"mode"`
	RejectReason RejectReason `json:"-"` // RejectReason is not stored in db
}

// NewFileInfo creates a new FileInfo that meets the invariants of the struct
func NewFileInfo(fileId string, fileHash string, branchName string) *FileInfo {
	return &FileInfo{
		FileId:       fileId,
		FileHash:     fileHash,
		UserIds:      []string{},
		BranchName:   branchName,
		ClaimMode:    Unclaimed,
		RejectReason: None,
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
		ClaimMode:    Unclaimed,
		RejectReason: None,
	}
}

// NewMissingFileInfo creates a new FileInfo for a file that is missing.
func NewMissingFileInfo(fileId string) *FileInfo {
	return &FileInfo{
		FileId:       fileId,
		FileHash:     "",
		UserIds:      []string{},
		BranchName:   "",
		ClaimMode:    Unclaimed,
		RejectReason: Missing,
	}
}

// UpgradeMissingToNew upgrades a missing file to a new file setting up the object invariants.
// Can error if the file is not missing.
func (fi *FileInfo) UpgradeMissingToNew(fileHash string, branchName string) error {
	if fi.RejectReason != Missing {
		return ErrorFileNotMissing
	}

	fi.FileHash = fileHash
	fi.BranchName = branchName
	fi.RejectReason = None

	return nil
}

// Claim claims a file for a user
func (fi *FileInfo) Claim(userId string, fileHash string, claimMode ClaimMode) error {
	if claimMode == Unclaimed {
		return fi.invalidClaimMode()
	}

	// Already claimed by someone else.
	if fi.ClaimMode == Exclusive {
		return fi.alreadyClaimed()
	}

	if fi.FileHash != fileHash {
		return fi.fileOutOfDate()
	}

	if fi.ClaimMode == Shared && claimMode == Exclusive {
		return fi.invalidClaimMode()
	}

	fi.addOwner(userId)
	fi.ClaimMode = claimMode

	return nil
}

// Update updates the file hash for a file only if the user is an owner
func (fi *FileInfo) Update(userId string, branchName string, oldHash string, fileHash string) error {
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
		fi.ClaimMode = Unclaimed
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
	fi.RejectReason = InvalidClaimMode
	return ErrInvalidClaimMode
}

func (fi *FileInfo) alreadyClaimed() error {
	fi.RejectReason = AlreadyClaimed
	return ErrFileAlreadyClaimed
}

func (fi *FileInfo) fileOutOfDate() error {
	fi.RejectReason = OutOfDate
	return ErrFileOutOfDate
}

func (fi *FileInfo) userNotOwner() error {
	fi.RejectReason = NotOwner
	return ErrUserNotOwner
}

// Implements encoding.BinaryMarshaler interface for FileInfo
// Needed for redis client to marshal FileInfo
// Deliberately pass by value not by pointer!
func (fi FileInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(fi)
}
