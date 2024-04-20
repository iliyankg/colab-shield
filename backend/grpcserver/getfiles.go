package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/colabom"
	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func getFilesHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, _ string, projectId string, request *protos.GetFilesRequest) (*protos.GetFilesResponse, error) {
	if len(request.FileIds) == 0 {
		logger.Warn().Msg("No files to update")
		// TODO: Consider returning an error here
		return &protos.GetFilesResponse{}, nil
	}

	logger.Info().Msgf("Getting... %d files", len(request.FileIds))

	keys := make([]string, 0, len(request.FileIds))
	for _, fileId := range request.FileIds {
		keys = append(keys, colabom.BuildRedisKeyForFile(projectId, fileId))
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return models.NewMissingFileInfo(request.FileIds[idx])
	}

	files := make([]*models.FileInfo, 0, len(request.FileIds))
	if err := colabom.GetFileInfos(ctx, logger, rc, keys, missingFileHandler, &files); err != nil {
		return nil, parseColabomError(err)
	}

	logger.Info().Msg("Getting successful")
	protoFiles := make([]*protos.FileInfo, 0, len(files))
	models.FileInfosToProto(files, &protoFiles)
	return &protos.GetFilesResponse{
		Files: protoFiles,
	}, nil
}
