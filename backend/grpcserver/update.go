package grpcserver

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/backend/core/requests"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func updateHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, req *protos.UpdateFilesRequest) (*protos.UpdateFilesResponse, error) {
	res, err := json.Marshal(req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to marshal JSON from Redis hash")
		return nil, ErrMarshalFail
	}

	var internalReq requests.Update
	if err := json.Unmarshal(res, internalReq); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return nil, ErrUnmarshalFail
	}

	rejectedFiles, err := core.Update(ctx, logger, rc, userId, projectId, &internalReq)
	parsedErr := parseCoreError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Updating failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.UpdateFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to update files")
		return nil, parsedErr
	}

	logger.Info().Msg("Updating successful")

	return &protos.UpdateFilesResponse{
		Status: protos.Status_OK,
	}, nil
}
