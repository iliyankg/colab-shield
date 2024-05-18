package domain

import (
	"context"

	"github.com/rs/zerolog"
)

// ColabDatabase is the interface for the database.
type ColabDatabase interface {
	Ping() error
	List(ctx context.Context, logger zerolog.Logger, projectId string, cursor uint64, pageSize int64, folderPath string) ([]*FileInfo, uint64, error)
	Claim(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request ClaimRequest) ([]*FileInfo, error)
	Update(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request UpdateRequest) ([]*FileInfo, error)
	Release(ctx context.Context, logger zerolog.Logger, userId string, projectId string, branchId string, request FilesRequest) ([]*FileInfo, error)
}

// FilesRequest is an interface for requests that contain lists of files
// whos IDs can be retrieved.
type FilesRequest interface {
	GetFileIDs() []string
}

type FileUpdate interface {
	GetFileId() string
	GetFileHash() string
	GetOldHash() string
}

type UpdateRequest interface {
	FilesRequest
	GetBranchName() string
	GetRequests() []FileUpdate
}

// FileClaim is an interface on how a file should be claimed.
type FileClaim interface {
	GetFileId() string
	GetFileHash() string
	GetClaimMode() ClaimMode
}

// ClaimRequest is an interface for requests that claim files.
type ClaimRequest interface {
	FilesRequest
	GetBranchName() string
	GetIsSoftClaim() bool
	GetRequests() []FileClaim
}
