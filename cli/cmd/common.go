package cmd

import (
	"context"
	"errors"
	"time"

	"github.com/iliyankg/colab-shield/protos"
	"google.golang.org/grpc/metadata"
)

var (
	ErrFileToHashMissmatch = errors.New("files and hashes must be of the same length")
)

func buildContext(projectId string, userId string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	metaInfo := metadata.Pairs(
		"projectId", projectId,
		"userId", userId,
	)

	return metadata.NewOutgoingContext(ctx, metaInfo), cancel
}

// newClaimFilesRequest creates a new ClaimFilesRequest from the given files and hashes
// Same claim mode is applied to all files.
// TODO: Per file claim mode support?
func newClaimFilesRequest(files []string, hashes []string, claimMode protos.ClaimMode, softClaim bool) (*protos.ClaimFilesRequest, error) {
	if len(files) != len(hashes) {
		return nil, ErrFileToHashMissmatch
	}

	claimFileInfos := make([]*protos.ClaimFileInfo, 0, len(files))
	for i, file := range files {
		claimFileInfos = append(claimFileInfos, &protos.ClaimFileInfo{
			FileId:    file,
			FileHash:  hashes[i],
			ClaimMode: claimMode,
		})
	}

	return &protos.ClaimFilesRequest{
		BranchName: gitBranch,
		Files:      claimFileInfos,
		SoftClaim:  softClaim,
	}, nil
}
