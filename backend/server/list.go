package server

import (
	"context"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func listHandler(ctx context.Context, logger zerolog.Logger, redisClient *redis.Client, _ string, projectId string, request *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	if request.PageSize == 0 {
		logger.Warn().Msg("No page size specified")
		return &pb.ListFilesResponse{}, nil
	}

	logger.Info().Msgf("Listing %d files for project %s", request.PageSize, projectId)

	match := buildScanQuery(projectId, request.FolderPath)

	keys, cursor, err := redisClient.Scan(ctx, request.Cursor, match, request.PageSize).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to scan keys")
		return nil, err // TODO: Better and/or more uniform error handling.
	}

	if len(keys) == 0 {
		logger.Warn().Msg("No files found")
		return &pb.ListFilesResponse{}, nil
	}

	result, err := redisClient.JSONMGet(ctx, ".", keys...).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get files from Redis")
		return nil, err // TODO: Better and/or more uniform error handling.
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return nil
	}

	// Handler for failed unmarshalling of JSON from the Redis hash
	unmarshalFailHandler := func(idx int, err error) error {
		logger.Error().Str("key", keys[idx]).Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return ErrUnmarshalFail
	}

	files := make([]*models.FileInfo, 0, len(keys))
	err = parseFileInfos(result, &files, missingFileHandler, unmarshalFailHandler)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse file infos from Redis hash")
		return nil, err // TODO: Better and/or more uniform error handling.
	}

	protoFiles := make([]*pb.FileInfo, 0, len(files))
	models.FileInfosToProto(files, &protoFiles)
	return &pb.ListFilesResponse{
		NextCursor: cursor,
		Files:      protoFiles,
	}, nil
}

func buildScanQuery(projectId string, folderPath string) string {
	// ProjectID is safe but folder path may not be.
	return fmt.Sprintf("project:%s:file:%s*", projectId, folderPath)
}
