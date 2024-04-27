package cmd

import (
	"context"
	"errors"
	"path"
	"time"

	"github.com/iliyankg/colab-shield/cli/config"
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

// filterToFilesOfInterest filters the given files to only those that are of interest basedon
// the extensions and ignore paths in the config.
func filterToFilesOfInterest(files []string) ([]string, error) {
	mappedExtensions := strSliceToHashMap(config.Extensions())
	excludedPaths := config.IgnorePaths()

	toReturn := make([]string, 0)
	for _, file := range files {
		extension := path.Ext(file)
		if _, ok := mappedExtensions[extension]; !ok {
			continue
		}

		if match, err := matchPathToExcludedPaths(file, excludedPaths); err != nil {
			return nil, err
		} else if match {
			continue
		}

		toReturn = append(toReturn, file)
	}

	return toReturn, nil
}

func matchPathToExcludedPaths(filePath string, excludedPaths []string) (bool, error) {
	for _, excludedPath := range excludedPaths {
		if match, err := path.Match(excludedPath, filePath); err != nil {
			return false, err
		} else if match {
			return true, nil
		}
	}
	return false, nil
}

func strSliceToHashMap(slice []string) map[string]any {
	hashMap := make(map[string]any)
	for _, s := range slice {
		hashMap[s] = nil
	}
	return hashMap
}
