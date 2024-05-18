package domain

// FileUpdate is an interface for requests that update a claimed file.
type fileUpdate struct {
	fileId   string
	fileHash string
	oldHash  string
}

// NewFileUpdate creates a new file update request.
func NewFileUpdate(fileId string, fileHash string, oldHash string) FileUpdate {
	return &fileUpdate{
		fileId:   fileId,
		fileHash: fileHash,
		oldHash:  oldHash,
	}
}

func (f *fileUpdate) GetFileId() string {
	return f.fileId
}

func (f *fileUpdate) GetFileHash() string {
	return f.fileHash
}

func (f *fileUpdate) GetOldHash() string {
	return f.oldHash
}

// UpdateRequest is an interface for requests that files be updated.
type updateRequest struct {
	branchName string
	requests   []FileUpdate
}

// NewUpdateRequest creates a new update request.
func NewUpdateRequest(branchName string, requests []FileUpdate) UpdateRequest {
	return &updateRequest{
		branchName,
		requests,
	}
}

func (u *updateRequest) GetBranchName() string {
	return u.branchName
}

func (u *updateRequest) GetRequests() []FileUpdate {
	return u.requests
}

// Implements FilesRequest interface
// returns the file IDs from the request
func (u *updateRequest) GetFileIDs() []string {
	var filesIds []string
	for _, req := range u.requests {
		filesIds = append(filesIds, req.GetFileId())
	}
	return filesIds
}
