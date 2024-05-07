package grpcserver

import (
	"context"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/colabom"
	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func listHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, _ string, projectId string, request *protos.ListFilesRequest) (*protos.ListFilesResponse, error) {
	if request.PageSize == 0 {
		logger.Warn().Msg("No page size specified")
		return &protos.ListFilesResponse{}, nil
	}

	logger.Info().Msgf("Listing %d files for project %s", request.PageSize, projectId)

	match := buildScanQuery(projectId, request.FolderPath)

	keys, cursor, err := rc.Scan(ctx, request.Cursor, match, request.PageSize).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to scan keys")
		return nil, ErrRedisError
	}

	if len(keys) == 0 {
		logger.Warn().Msgf("No files found for match: %s", match)
		return &protos.ListFilesResponse{}, nil
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return nil
	}

	files := make([]*models.FileInfo, 0, len(keys))
	if err := colabom.GetFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
		return nil, parseCoreError(err)
	}

	protoFiles := make([]*protos.FileInfo, 0, len(files))
	fileInfosToProto(files, &protoFiles)
	return &protos.ListFilesResponse{
		NextCursor: cursor,
		Files:      protoFiles,
	}, nil
}

func buildScanQuery(projectId string, folderPath string) string {
	// ProjectID is safe but folder path may not be.
	return fmt.Sprintf("project:%s:file:%s*", projectId, folderPath)
}
