package domain

import (
	"context"

	"github.com/rs/zerolog"
)

// FilesRequest is an interface for requests that contain lists of files
// whos IDs can be retrieved.
type FilesRequest interface {
	GetFileIDs() []string
}

// FileUpdate is an interface for requests that update a claimed file.
type FileUpdate interface {
	GetFileId() string
	GetFileHash() string
	GetOldHash() string
}

// UpdateRequest is an interface for requests that files be updated.
type UpdateRequest interface {
	GetBranchName() string
	GetFiles() []FileUpdate
}

// FileClaim is an interface on how a file should be claimed.
type FileClaim interface {
	GetFileId() string
	GetFileHash() string
	GetClaimMode() int32
}

// ClaimRequest is an interface for requests that claim files.
type ClaimRequest interface {
	GetBranchName() string
	GetSoftClaim() bool
	GetFiles() []FileClaim
}

// ColabDatabase is the interface for the database allowing for efficient interactions with the database.
type ColabDatabase interface {
	List(ctx context.Context, logger zerolog.Logger, projectId string, cursor uint64, pageSize int64, folderPath string) ([]*FileInfo, uint64, error)
	Claim(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request ClaimRequest) ([]*FileInfo, error)
	Update(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request UpdateRequest) ([]*FileInfo, error)
	Release(ctx context.Context, logger zerolog.Logger, userId string, projectId string, branchId string, request FilesRequest) ([]*FileInfo, error)
}
