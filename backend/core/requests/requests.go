package requests

type ClaimMode int32

const (
	UNCLAIMED ClaimMode = 0
	EXCLUSIVE ClaimMode = 1
	SHARED    ClaimMode = 2
)

type Claim struct {
	SoftClaim  bool         `json:"softClaim"`
	BranchName string       `json:"branchName"`
	Files      []*ClaimInfo `json:"files"`
}

func (c *Claim) GetFilesIds() []string {
	var filesIds []string
	for _, file := range c.Files {
		filesIds = append(filesIds, file.FileId)
	}
	return filesIds
}

type ClaimInfo struct {
	FileId    string    `json:"fileId"`
	FileHash  string    `json:"fileHash"`
	ClaimMode ClaimMode `json:"claimMode"`
}

type Update struct {
	BranchName string            `json:"branchName"`
	Files      []*UpdateFileInfo `json:"files"`
}

func (u *Update) GetFilesIds() []string {
	var filesIds []string
	for _, file := range u.Files {
		filesIds = append(filesIds, file.FileId)
	}
	return filesIds
}

type UpdateFileInfo struct {
	FileId   string `json:"fileId"`
	OldHash  string `json:"oldHash"`
	FileHash string `json:"fileHash"`
}
