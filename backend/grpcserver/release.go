package grpcserver

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func releaseHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, request *pb.ReleaseFilesRequest) (*pb.ReleaseFilesResponse, error) {
	if len(request.FileIds) == 0 {
		logger.Warn().Msg("No files to release")
		// TODO: Consider returning an error here
		return &pb.ReleaseFilesResponse{
			Status: pb.Status_OK,
		}, nil
	}

	logger.Info().Msgf("Releasing... %d files", len(request.FileIds))

	files := make([]*models.FileInfo, 0, len(request.FileIds))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(request.FileIds))
	for _, fileId := range request.FileIds {
		keys = append(keys, buildRedisKeyForFile(projectId, fileId))
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		rejectedFiles = append(rejectedFiles, models.NewMissingFileInfo(request.FileIds[idx]))
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

		releaseFiles(userId, files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFiles(ctx, logger, rc, keys, files)
	}

	// Execute the watch function
	err := rc.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		logger.Info().Msg("Releasing failed due to rejected files")
		protoRejectedFiles := make([]*pb.FileInfo, 0, len(rejectedFiles))
		models.FileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &pb.ReleaseFilesResponse{
			Status:        pb.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("Failed to release files")
		return nil, err
	}

	logger.Info().Msg("Releasing successful")

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
