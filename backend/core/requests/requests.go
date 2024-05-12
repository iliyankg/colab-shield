package requests

// FilesRequest is an interface for requests that contain files
type FilesRequest interface {
	GetFilesIds() []string
}

// ClaimMode is the mode in which a file can be claimed
type ClaimMode int32

const (
	UNCLAIMED ClaimMode = 0
	EXCLUSIVE ClaimMode = 1
	SHARED    ClaimMode = 2
)

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

// ClaimInfo contains the information needed to claim a file
type ClaimInfo struct {
	FileId    string    `json:"fileId"`
	FileHash  string    `json:"fileHash"`
	ClaimMode ClaimMode `json:"claimMode"`
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

// UpdateFileInfo contains the information needed to update a file
type UpdateFileInfo struct {
	FileId   string `json:"fileId"`
	OldHash  string `json:"oldHash"`
	FileHash string `json:"fileHash"`
}
