package server

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func updateHandler(ctx context.Context, logger zerolog.Logger, redisClient *redis.Client, userId string, projectId string, request *pb.UpdateFilesRequest) (*pb.UpdateFilesResponse, error) {
	logger.Info().Msgf("Updating... %d files", len(request.Files))

	files := make([]*models.FileInfo, 0, len(request.Files))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(request.Files))
	keysFromFileRequests(projectId, request.Files, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return models.NewFileInfo(request.Files[idx].FileId, request.Files[idx].FileHash, request.BranchName)
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

		err = parseFileInfos(result, &files, missingFileHandler, unmarshalFailHandler)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to parse file infos from Redis hash")
			return err
		}

		updateFiles(userId, request.BranchName, files, request.Files, &rejectedFiles)

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

		return &pb.UpdateFilesResponse{
			Status:        pb.Status_ERROR,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("Failed to claim files")
		return nil, err
	}

	logger.Info().Msg("Updating successful")

	return &pb.UpdateFilesResponse{
		Status: pb.Status_OK,
	}, nil
}

func updateFiles(userId string, branchName string, fileInfos []*models.FileInfo, pbFiles []*pb.UpdateFileInfo, outRejectedFiles *[]*models.FileInfo) {
	// update the files with the new file hashes
	for i := range fileInfos {
		if err := fileInfos[i].Update(userId, pbFiles[i].OldHash, pbFiles[i].FileHash, branchName); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}
