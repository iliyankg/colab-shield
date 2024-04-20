package grpcserver

import (
	"context"
	"errors"

	"github.com/iliyankg/colab-shield/backend/colabom"
	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func claimHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, userId string, projectId string, request *protos.ClaimFilesRequest) (*protos.ClaimFilesResponse, error) {
	if len(request.Files) == 0 {
		logger.Warn().Msg("No files to claim")
		// TODO: Consider returning an error here
		return &protos.ClaimFilesResponse{
			Status: protos.Status_OK,
		}, nil
	}

	logger.Info().Msgf("Claiming... %d files", len(request.Files))

	files := make([]*models.FileInfo, 0, len(request.Files))
	rejectedFiles := make([]*models.FileInfo, 0)

	keys := make([]string, 0, len(request.Files))
	keysFromFileRequests(projectId, request.Files, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *models.FileInfo {
		return models.NewFileInfo(request.Files[idx].FileId, request.Files[idx].FileHash, request.BranchName)
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := colabom.GetFileInfos(ctx, logger, tx, keys, missingFileHandler, &files); err != nil {
			return parseColabomError(err)
		}

		claimFiles(userId, files, request.Files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		if request.SoftClaim {
			logger.Info().Msg("Only a soft claim, nothing saved...")
			return nil
		}

		return parseColabomError(colabom.SetFileInfos(ctx, logger, tx, keys, files))
	}

	// Execute the watch function
	err := rc.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		logger.Info().Msg("Claiming failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ClaimFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		logger.Error().Err(err).Msg("Failed to claim files")
		return nil, err
	}

	logger.Info().Msg("Claiming successful")

	// TODO: Consider returning the files that were claimed succesfully
	return &protos.ClaimFilesResponse{
		Status: protos.Status_OK,
	}, nil
}

func claimFiles(userId string, fileInfos []*models.FileInfo, claimRequests []*protos.ClaimFileInfo, outRejectedFiles *[]*models.FileInfo) {
	for i := range fileInfos {
		reqFile := claimRequests[i]

		// try claiming the file
		if err := fileInfos[i].Claim(userId, reqFile.FileHash, models.ClaimMode(reqFile.ClaimMode)); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}
