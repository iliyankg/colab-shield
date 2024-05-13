package core

import (
	"context"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/domain"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func List(ctx context.Context, logger zerolog.Logger, rc *redis.Client, projectId string, cursor uint64, pageSize int64, folderPath string) ([]*domain.FileInfo, uint64, error) {
	if pageSize == 0 {
		logger.Warn().Msg("No page size specified")
		return nil, 0, nil
	}

	logger.Info().Msgf("Listing %d files for project %s", pageSize, projectId)

	match := buildScanQuery(projectId, folderPath)

	keys, cursor, err := rc.Scan(ctx, cursor, match, pageSize).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to scan keys")
		return nil, 0, ErrRedisError
	}

	if len(keys) == 0 {
		logger.Warn().Msgf("No files found for match: %s", match)
		return nil, 0, nil
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		return nil
	}

	files := make([]*domain.FileInfo, 0, len(keys))
	if err := getFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
		return nil, 0, err
	}

	return files, cursor, nil
}

func buildScanQuery(projectId string, folderPath string) string {
	// ProjectID is safe but folder path may not be.
	return fmt.Sprintf("project:%s:file:%s*", projectId, folderPath)
}
