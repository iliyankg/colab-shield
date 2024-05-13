package core

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/core/requests"
	"github.com/iliyankg/colab-shield/backend/domain"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func Update(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, request *requests.Update) ([]*domain.FileInfo, error) {
	if len(request.Files) == 0 {
		logger.Warn().Msg("No files to update")
		return nil, nil
	}

	logger.Info().Msgf("Updating... %d files", len(request.Files))

	files := make([]*domain.FileInfo, 0, len(request.Files))
	rejectedFiles := make([]*domain.FileInfo, 0)

	keys := make([]string, 0, len(request.Files))
	keysFromFileRequests(projectId, request, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		rejectedFiles = append(rejectedFiles, domain.NewMissingFileInfo(request.Files[idx].FileId))
		return nil
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := getFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
			return err
		}

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		updateFiles(userId, request.BranchName, files, request.Files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFileInfos(ctx, logger, tx, keys, files)
	}

	// Execute the watch function
	err := rc.Watch(ctx, watchFn, keys...)
	switch {
	case errors.Is(err, ErrRejectedFiles):
		logger.Info().Msg("Updating failed due to rejected files")
		return rejectedFiles, err
	case err != nil:
		logger.Error().Err(err).Msg("Failed to update files")
		return nil, err
	default:
		logger.Info().Msg("Updating successful")
		return nil, nil
	}
}

func updateFiles(userId string, branchName string, fileInfos []*domain.FileInfo, pbFiles []*requests.UpdateFileInfo, outRejectedFiles *[]*domain.FileInfo) {
	// update the files with the new file hashes
	for i := range fileInfos {
		if err := fileInfos[i].Update(userId, pbFiles[i].OldHash, pbFiles[i].FileHash, branchName); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}
