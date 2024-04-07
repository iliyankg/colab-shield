package server

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func releaseHandler(ctx context.Context, logger zerolog.Logger, redisClient *redis.Client, userId string, projectId string, request *pb.ReleaseFilesRequest) (*pb.ReleaseFilesResponse, error) {
	logger.Info().Msgf("Releasing... %d files", len(request.FileIds))

	files := make([]*models.FileInfo, 0, len(request.FileIds))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(request.FileIds))
	for _, fileId := range request.FileIds {
		keys = append(keys, buildRedisKeyForFile(projectId, fileId))
	}

	// Handler for failed unmarshalling of JSON from the Redis hash
	unmarshalFailHandler := func(idx int, err error) error {
		logger.Error().Str("key", keys[idx]).Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return nil
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		result, err := tx.JSONMGet(ctx, ".", keys...).Result()
		if err != nil {
			logger.Error().Err(err).Msg("Failed to read keys from Redis hash")
			return err
		}

		// FIXME: &files could contain nil entries and thats not good.
		// We should probably return an error if we encounter a nil entry
		err = parseFileInfos(result, &files, nil, unmarshalFailHandler)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to parse file infos from Redis hash")
			return err
		}

		releaseFiles(userId, files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFiles(ctx, logger, redisClient, keys, files)
	}

	// Execute the watch function
	err := redisClient.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		protoRejectedFiles := make([]*pb.FileInfo, 0, len(rejectedFiles))
		models.FileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &pb.ReleaseFilesResponse{
			Status:        pb.Status_ERROR,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("Failed to claim files")
		return nil, err
	}

	logger.Info().Msg("Updating successful")

	return &pb.ReleaseFilesResponse{
		Status: pb.Status_OK,
	}, nil
}

func releaseFiles(userId string, fileInfos []*models.FileInfo, outRejectedFiles *[]*models.FileInfo) {
	// update the files with the new file hashes
	for i := range fileInfos {
		if fileInfos[i] == nil {
			continue
		}

		if err := fileInfos[i].Release(userId); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}
