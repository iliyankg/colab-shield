package grpcserver

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/iliyankg/colab-shield/backend/common"
	"github.com/iliyankg/colab-shield/backend/common/request"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func claimHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, req *protos.ClaimFilesRequest) (*protos.ClaimFilesResponse, error) {
	res, err := json.Marshal(req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to marshal JSON from Redis hash")
		return nil, ErrMarshalFail
	}

	var internalReq request.Claim
	if err := json.Unmarshal(res, internalReq); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return nil, ErrUnmarshalFail
	}

	rejectedFiles, err := common.Claim(ctx, logger, rc, userId, projectId, &internalReq)
	parsedErr := parseColabomError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Claiming failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ClaimFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to claim files")
		return nil, parsedErr
	}

	// TODO: Consider returning the files that were claimed succesfully
	return &protos.ClaimFilesResponse{
		Status: protos.Status_OK,
	}, nil
}
