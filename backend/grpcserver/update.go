package grpcserver

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func updateHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, request *pb.UpdateFilesRequest) (*pb.UpdateFilesResponse, error) {
	if len(request.Files) == 0 {
		logger.Warn().Msg("No files to update")
		// TODO: Consider returning an error here
		return &pb.UpdateFilesResponse{
			Status: pb.Status_OK,
		}, nil
	}

	logger.Info().Msgf("Updating... %d files", len(request.Files))

	files := make([]*models.FileInfo, 0, len(request.Files))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(request.Files))
	keysFromFileRequests(projectId, request.Files, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		rejectedFiles = append(rejectedFiles, models.NewMissingFileInfo(request.Files[idx].FileId))
		return nil
	}

	// Handler for failed unmarshalling of JSON from the Redis hash
	unmarshalFailHandler := func(idx int, err error) error {
		logger.Error().Str("key", keys[idx]).Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return ErrUnmarshalFail
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := getFileInfos(ctx, logger, rc, keys, missingFileHandler, unmarshalFailHandler, &files); err != nil {
			return err
		}

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		updateFiles(userId, request.BranchName, files, request.Files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFiles(ctx, logger, rc, keys, files)
	}

	// Execute the watch function
	err := rc.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		logger.Info().Msg("Updating failed due to rejected files")
		protoRejectedFiles := make([]*pb.FileInfo, 0, len(rejectedFiles))
		models.FileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &pb.UpdateFilesResponse{
			Status:        pb.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("Failed to update files")
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
