package core

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func Release(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, branchId string, fileIds []string) ([]*models.FileInfo, error) {
	if len(fileIds) == 0 {
		logger.Warn().Msg("No files to release")
		return nil, nil
	}

	logger.Info().Msgf("Releasing... %d files", len(fileIds))

	files := make([]*models.FileInfo, 0, len(fileIds))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(fileIds))
	for _, fileId := range fileIds {
		keys = append(keys, buildRedisKeyForFile(projectId, fileId))
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		rejectedFiles = append(rejectedFiles, models.NewMissingFileInfo(fileIds[idx]))
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

		releaseFiles(userId, files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFileInfos(ctx, logger, tx, keys, files)
	}

	// Execute the watch function
	err := rc.Watch(ctx, watchFn, keys...)
	switch {
	case errors.Is(err, ErrRejectedFiles):
		logger.Info().Msg("Releasing failed due to rejected files")
		return rejectedFiles, nil
	case err != nil:
		logger.Error().Err(err).Msg("Failed to release files")
		return nil, err
	default:
		logger.Info().Msg("Releasing successful")
		return nil, nil
	}
}

func releaseFiles(userId string, fileInfos []*models.FileInfo, outRejectedFiles *[]*models.FileInfo) {
	// update the files with the new file hashes
	for i := range fileInfos {
		if fileInfos[i] == nil {
			continue
		}

		if err := fileInfos[i].Release(userId); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}
