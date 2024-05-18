package domain

// FileClaim is an interface on how a file should be claimed.
type fileClaim struct {
	fileId    string
	fileHash  string
	claimMode ClaimMode
}

// NewFileClaim creates a new file claim.
func NewFileClaim(fileId string, fileHash string, claimMode ClaimMode) FileClaim {
	return &fileClaim{
		fileId,
		fileHash,
		claimMode,
	}
}

func (f *fileClaim) GetFileId() string {
	return f.fileId
}

func (f *fileClaim) GetFileHash() string {
	return f.fileHash
}

func (f *fileClaim) GetClaimMode() ClaimMode {
	return f.claimMode
}

// ClaimRequest is an interface for requests that claim files.
type claimRequest struct {
	branchName  string
	isSoftClaim bool
	requests    []FileClaim
}

// NewClaimRequest creates a new claim request.
func NewClaimRequest(branchName string, softClaim bool, requests []FileClaim) ClaimRequest {
	return &claimRequest{
		branchName,
		softClaim,
		requests,
	}
}

func (c *claimRequest) GetBranchName() string {
	return c.branchName
}

func (c *claimRequest) GetRequests() []FileClaim {
	return c.requests
}

func (c *claimRequest) GetIsSoftClaim() bool {
	return c.isSoftClaim
}

// Implements FilesRequest interface
// returns the file IDs from the request
func (c *claimRequest) GetFileIDs() []string {
	var filesIds []string
	for _, req := range c.requests {
		filesIds = append(filesIds, req.GetFileId())
	}
	return filesIds
}
