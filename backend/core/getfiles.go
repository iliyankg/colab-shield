package core

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func GetFiles(ctx context.Context, logger zerolog.Logger, rc *redis.Client, projectId string, fileIds []string) ([]*models.FileInfo, error) {
	if len(fileIds) == 0 {
		logger.Warn().Msg("No files to update")
		return nil, nil
	}

	logger.Info().Msgf("Getting... %d files", len(fileIds))

	keys := make([]string, 0, len(fileIds))
	for _, fileId := range fileIds {
		keys = append(keys, buildRedisKeyForFile(projectId, fileId))
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return models.NewMissingFileInfo(fileIds[idx])
	}

	files := make([]*models.FileInfo, 0, len(fileIds))
	if err := getFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
		return nil, err
	}

	logger.Info().Msg("Getting successful")
	return files, nil
}
