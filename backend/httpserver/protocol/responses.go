package protocol

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
