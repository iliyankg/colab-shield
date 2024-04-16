package grpcserver

import (
	"context"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func listHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, _ string, projectId string, request *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	if request.PageSize == 0 {
		logger.Warn().Msg("No page size specified")
		return &pb.ListFilesResponse{}, nil
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
		return &pb.ListFilesResponse{}, nil
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return nil
	}

	files := make([]*models.FileInfo, 0, len(keys))
	if err := getFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
		return nil, err
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
