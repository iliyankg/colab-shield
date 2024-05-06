package request

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
