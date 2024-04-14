package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func getFilesHandler(ctx context.Context, logger zerolog.Logger, redisClient *redis.Client, userId string, projectId string, request *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	if len(request.FileIds) == 0 {
		logger.Warn().Msg("No files to update")
		// TODO: Consider returning an error here
		return &pb.GetFilesResponse{}, nil
	}

	logger.Info().Msgf("Getting... %d files", len(request.FileIds))

	keys := make([]string, 0, len(request.FileIds))
	for _, fileId := range request.FileIds {
		keys = append(keys, buildRedisKeyForFile(projectId, fileId))
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return models.NewMissingFileInfo(request.FileIds[idx])
	}

	// Handler for failed unmarshalling of JSON from the Redis hash
	unmarshalFailHandler := func(idx int, err error) error {
		logger.Error().Str("key", keys[idx]).Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return ErrUnmarshalFail
	}

	result, err := redisClient.JSONMGet(ctx, ".", keys...).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read keys from Redis hash")
		return nil, err
	}

	files := make([]*models.FileInfo, 0, len(request.FileIds))
	err = parseFileInfos(result, &files, missingFileHandler, unmarshalFailHandler)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to parse file infos from Redis hash")
		return nil, err
	}

	logger.Info().Msg("Getting successful")
	protoFiles := make([]*pb.FileInfo, 0, len(files))
	models.FileInfosToProto(files, &protoFiles)
	return &pb.GetFilesResponse{
		Files: protoFiles,
	}, nil
}
