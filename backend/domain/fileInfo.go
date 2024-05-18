package domain

import (
	"errors"
)

var (
	ErrFileAlreadyClaimed      = errors.New("file already claimed")
	ErrFileOutOfDate           = errors.New("file out of date")
	ErrUserNotOwner            = errors.New("user not owner")
	ErrInvalidClaimMode        = errors.New("invalid claim mode")
	ErrorFileNotMissing        = errors.New("file not missing")
	ErrFileInfoInvariantBroken = errors.New("file info invariant broken")
)

// Enum for claim mode
// Should be kept in sync with the protobuf file.
type ClaimMode int32

const (
	ClaimMode_Unclaimed ClaimMode = iota
	ClaimMode_Shared
	ClaimMode_Exclusive
)

// Enum for reject reason
// Should be kept in sync with the protobuf file.
type RejectReason int32

const (
	RejectReason_None RejectReason = iota
	RejectReason_AlreadyClaimed
	RejectReason_OutOfDate
	RejectReason_NotOwner
	RejectReason_InvalidClaimMode
	RejectReason_Missing
)

// FileInfo represents a file in the system
//
// TODO: Look into: https://pkg.go.dev/github.com/redis/rueidis/om#section-readme
// TODO: Look into: https://stackoverflow.com/questions/11126793/json-and-dealing-with-unexported-fields
type FileInfo struct {
	fileId       string
	fileHash     string
	userIds      []string
	branchName   string
	claimMode    ClaimMode
	rejectReason RejectReason
}

// NewFileInfo creates a new FileInfo that meets the invariants of the struct
func NewFileInfo(fileId string, fileHash string, branchName string) *FileInfo {
	return &FileInfo{
		fileId:       fileId,
		fileHash:     fileHash,
		userIds:      []string{},
		branchName:   branchName,
		claimMode:    ClaimMode_Unclaimed,
		rejectReason: RejectReason_None,
	}
}

// NewBlankFileInfo creates a new blank without the need for parameters.
//
// Ideal for creating a new FileInfo that will be populated later on such as
// unmarshalling from a database.
func NewBlankFileInfo() *FileInfo {
	return &FileInfo{
		fileId:       "",
		fileHash:     "",
		userIds:      []string{},
		branchName:   "",
		claimMode:    ClaimMode_Unclaimed,
		rejectReason: RejectReason_None,
	}
}

// NewMissingFileInfo creates a new FileInfo for a file that is missing.
func NewMissingFileInfo(fileId string) *FileInfo {
	return &FileInfo{
		fileId:       fileId,
		fileHash:     "",
		userIds:      []string{},
		branchName:   "",
		claimMode:    ClaimMode_Unclaimed,
		rejectReason: RejectReason_Missing,
	}
}

// NewFullFileInfo creates a new FileInfo with all the fields set.
// Can error if the claim mode is not Unclaimed and there are no user IDs.
func NewFullFileInfo(fileId string, fileHash string, userIds []string, branchName string, claimMode ClaimMode, rejectReason RejectReason) (*FileInfo, error) {
	if claimMode != ClaimMode_Unclaimed && len(userIds) == 0 {
		return nil, ErrFileInfoInvariantBroken
	}

	return &FileInfo{
		fileId:       fileId,
		fileHash:     fileHash,
		userIds:      userIds,
		branchName:   branchName,
		claimMode:    claimMode,
		rejectReason: rejectReason,
	}, nil
}

// GetFileId returns the fileId of the FileInfo
func (fi *FileInfo) GetFileId() string {
	return fi.fileId
}

// GetFileHash returns the fileHash of the FileInfo
func (fi *FileInfo) GetFileHash() string {
	return fi.fileHash
}

// GetUserIds returns the userIds of the FileInfo
func (fi *FileInfo) GetUserIds() []string {
	return fi.userIds
}

// GetBranchName returns the branchName of the FileInfo
func (fi *FileInfo) GetBranchName() string {
	return fi.branchName
}

// GetClaimMode returns the claimMode of the FileInfo
func (fi *FileInfo) GetClaimMode() ClaimMode {
	return fi.claimMode
}

// GetRejectReason returns the rejectReason of the FileInfo
func (fi *FileInfo) GetRejectReason() RejectReason {
	return fi.rejectReason
}

// UpgradeMissingToNew upgrades a missing file to a new file setting up the object invariants.
// Can error if the file is not missing.
func (fi *FileInfo) UpgradeMissingToNew(fileHash string, branchName string) error {
	if fi.rejectReason != RejectReason_Missing {
		return ErrorFileNotMissing
	}

	fi.fileHash = fileHash
	fi.branchName = branchName
	fi.rejectReason = RejectReason_None

	return nil
}

// Claim claims a file for a user
func (fi *FileInfo) Claim(userId string, fileHash string, claimMode ClaimMode) error {
	if claimMode == ClaimMode_Unclaimed {
		return fi.invalidClaimMode()
	}

	// Already claimed by someone else.
	if fi.claimMode == ClaimMode_Exclusive {
		return fi.alreadyClaimed()
	}

	if fi.fileHash != fileHash {
		return fi.fileOutOfDate()
	}

	if fi.claimMode == ClaimMode_Shared && claimMode == ClaimMode_Exclusive {
		return fi.invalidClaimMode()
	}

	fi.addOwner(userId)
	fi.claimMode = claimMode

	return nil
}

// Update updates the file hash for a file only if the user is an owner
func (fi *FileInfo) Update(userId string, branchName string, oldHash string, fileHash string) error {
	if !fi.CheckOwner(userId) {
		return fi.userNotOwner()
	}

	if fi.fileHash != oldHash {
		return fi.fileOutOfDate()
	}

	fi.fileHash = fileHash
	fi.branchName = branchName
	return nil
}

// Release releases a file from a user if they are an owner
func (fi *FileInfo) Release(userId string) error {
	if err := fi.removeOwner(userId); err != nil {
		return err
	}

	if len(fi.userIds) == 0 {
		fi.claimMode = ClaimMode_Unclaimed
	}

	return nil
}

// CheckOwner checks if a user is an owner of a file
func (fi *FileInfo) CheckOwner(userId string) bool {
	for _, id := range fi.userIds {
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

	fi.userIds = append(fi.userIds, userId)
}

// removeOwner removes a userId from the FileInfo
func (fi *FileInfo) removeOwner(userId string) error {
	for i, id := range fi.userIds {
		if id == userId {
			fi.userIds = append(fi.userIds[:i], fi.userIds[i+1:]...)
			return nil
		}
	}

	return fi.userNotOwner()
}

func (fi *FileInfo) invalidClaimMode() error {
	fi.rejectReason = RejectReason_InvalidClaimMode
	return ErrInvalidClaimMode
}

func (fi *FileInfo) alreadyClaimed() error {
	fi.rejectReason = RejectReason_AlreadyClaimed
	return ErrFileAlreadyClaimed
}

func (fi *FileInfo) fileOutOfDate() error {
	fi.rejectReason = RejectReason_OutOfDate
	return ErrFileOutOfDate
}

func (fi *FileInfo) userNotOwner() error {
	fi.rejectReason = RejectReason_NotOwner
	return ErrUserNotOwner
}
