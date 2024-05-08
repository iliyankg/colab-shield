package grpcserver

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func releaseHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, request *protos.ReleaseFilesRequest) (*protos.ReleaseFilesResponse, error) {
	rejectedFiles, err := core.Release(ctx, logger, rc, userId, projectId, request.BranchName, request.FileIds)
	parsedErr := parseCoreError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Releasing failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ReleaseFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to release files")
		return nil, parsedErr
	}

	logger.Info().Msg("Releasing successful")

	return &protos.ReleaseFilesResponse{
		Status: protos.Status_OK,
	}, nil
}
