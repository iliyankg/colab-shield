package protocol

import "github.com/iliyankg/colab-shield/backend/models"

// FilesRequest is an interface for requests that contain files
type FilesRequest interface {
	GetFilesIds() []string
}

// ClaimMode is the mode in which a file is claimed
type ClaimMode int32

const (
	UNCLAIMED ClaimMode = 0
	EXCLUSIVE ClaimMode = 1
	SHARED    ClaimMode = 2
)

// RejectReason is the reason a claim was rejected
type RejectReason int32

const (
	NONE               RejectReason = 0
	ALREADY_CLAIMED    RejectReason = 1
	OUT_OF_DATE        RejectReason = 2
	NOT_OWNER          RejectReason = 3
	INVALID_CLAIM_MODE RejectReason = 4
	MISSING            RejectReason = 5
)

// FileInfo represents the information sent to the client about a file
type FileInfo struct {
	FileId       string       `json:"fileId"`
	FileHash     string       `json:"fileHash"`
	UserIds      []string     `json:"userIds"`
	BranchName   string       `json:"branchName"`
	ClaimMode    ClaimMode    `json:"claimMode"`
	RejectReason RejectReason `json:"rejectReason"`
}

func NewFileInfoFromModel(fi *models.FileInfo) *FileInfo {
	return &FileInfo{
		FileId:       fi.FileId,
		FileHash:     fi.FileHash,
		UserIds:      fi.UserIds,
		BranchName:   fi.BranchName,
		ClaimMode:    ClaimMode(fi.ClaimMode),
		RejectReason: RejectReason(fi.RejectReason),
	}
}

// ClaimInfo contains the information needed to claim a file
type ClaimInfo struct {
	FileId    string    `json:"fileId"`
	FileHash  string    `json:"fileHash"`
	ClaimMode ClaimMode `json:"claimMode"`
}

// Claim is a request to claim files
// Implements FileRequests interface
type Claim struct {
	SoftClaim  bool         `json:"softClaim"`
	BranchName string       `json:"branchName"`
	Files      []*ClaimInfo `json:"files"`
}

// GetFilesIds returns the file IDs from the request
func (c *Claim) GetFilesIds() []string {
	var filesIds []string
	for _, file := range c.Files {
		filesIds = append(filesIds, file.FileId)
	}
	return filesIds
}

// UpdateFileInfo contains the information needed to update a file
type UpdateFileInfo struct {
	FileId   string `json:"fileId"`
	OldHash  string `json:"oldHash"`
	FileHash string `json:"fileHash"`
}

// Update is a request to update files
// Implements FileRequests interface
type Update struct {
	BranchName string            `json:"branchName"`
	Files      []*UpdateFileInfo `json:"files"`
}

// GetFilesIds returns the file IDs from the request
func (u *Update) GetFilesIds() []string {
	var filesIds []string
	for _, file := range u.Files {
		filesIds = append(filesIds, file.FileId)
	}
	return filesIds
}

// Release is a request to release files
type Release struct {
	BranchName string   `json:"branchName"`
	FileIds    []string `json:"fileIds"`
}
