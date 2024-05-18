package domain

type filesRequest struct {
	fileIds []string
}

// NewFilesRequest creates a new files request.
func NewFilesRequest(fileIds []string) FilesRequest {
	return &filesRequest{
		fileIds,
	}
}

func (f *filesRequest) GetFileIDs() []string {
	return f.fileIds
}
