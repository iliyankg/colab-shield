package redisdatabase

import (
	"context"
	"errors"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/domain"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

// RedisDatabase is a Colabdatabase implementation that uses Redis as the backend.
type RedisDatabase struct {
	rc *redis.Client
}

// NewRedisDatabase creates a new RedisDatabase instance with the
// given client.
func NewRedisDatabase(rc *redis.Client) domain.ColabDatabase {
	return &RedisDatabase{
		rc: rc,
	}
}

func (csd *RedisDatabase) Ping() error {
	if _, err := csd.rc.Ping(context.Background()).Result(); err != nil {
		return err
	}
	return nil
}

// List lists the files for the given project and folder path.
// No path will list all files for the project.
// cursor = 0 to start from the beginning
// returned cursor
func (csd *RedisDatabase) List(ctx context.Context, logger zerolog.Logger, projectId string, cursor uint64, pageSize int64, folderPath string) ([]*domain.FileInfo, uint64, error) {
	if pageSize == 0 {
		logger.Warn().Msg("No page size specified")
		return nil, 0, nil
	}

	logger.Info().Msgf("Listing %d files for project %s", pageSize, projectId)

	match := buildScanQuery(projectId, folderPath)

	keys, cursor, err := csd.rc.Scan(ctx, cursor, match, pageSize).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to scan keys")
		return nil, 0, ErrRedisError
	}

	if len(keys) == 0 {
		logger.Warn().Msgf("No files found for match: %s", match)
		return nil, 0, nil
	}

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		return nil
	}

	files := make([]*domain.FileInfo, 0, len(keys))
	if err := getFileInfos(ctx, logger, csd.rc, keys, missingFileHandler, &files); err != nil {
		return nil, 0, err
	}

	return files, cursor, nil
}

// Claim claims the files with the given file IDs.
func (csd *RedisDatabase) Claim(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request domain.ClaimRequest) ([]*domain.FileInfo, error) {
	requests := request.GetRequests()
	numRequests := len(requests)

	if numRequests == 0 {
		logger.Warn().Msg("No files to claim")
		return nil, nil
	}

	logger.Info().Msgf("Claiming... %d files", numRequests)

	files := make([]*domain.FileInfo, 0, numRequests)
	rejectedFiles := make([]*domain.FileInfo, 0)

	filesRequest, ok := request.(domain.FilesRequest)
	if !ok {
		logger.Error().Msg("Request does not implement FilesRequest")
		return nil, ErrInvalidRequest
	}

	keys := make([]string, 0, numRequests)
	keysFromFileRequests(projectId, filesRequest, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		return domain.NewFileInfo(requests[idx].GetFileId(), requests[idx].GetFileHash(), request.GetBranchName())
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := getFileInfos(ctx, logger, tx, keys, missingFileHandler, &files); err != nil {
			return err
		}

		claimFiles(userId, files, requests, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		if request.GetIsSoftClaim() {
			logger.Info().Msg("Only a soft claim, nothing saved...")
			return nil
		}

		return setFileInfos(ctx, logger, tx, keys, files)
	}

	// Execute the watch function
	err := csd.rc.Watch(ctx, watchFn, keys...)
	switch {
	case errors.Is(err, ErrRejectedFiles):
		logger.Info().Msg("Claiming failed due to rejected files")
		return rejectedFiles, err
	case err != nil:
		logger.Error().Err(err).Msg("Failed to claim files")
		return nil, err
	default:
		logger.Info().Msg("Claiming successful")
		return nil, nil
	}
}

// Update updates the files with the given file IDs.
func (csd *RedisDatabase) Update(ctx context.Context, logger zerolog.Logger, userId string, projectId string, request domain.UpdateRequest) ([]*domain.FileInfo, error) {
	requests := request.GetRequests()
	numRequests := len(requests)

	if numRequests == 0 {
		logger.Warn().Msg("No files to update")
		return nil, nil
	}

	logger.Info().Msgf("Updating... %d files", numRequests)

	files := make([]*domain.FileInfo, 0, numRequests)
	rejectedFiles := make([]*domain.FileInfo, 0)

	filesRequest, ok := request.(domain.FilesRequest)
	if !ok {
		logger.Error().Msg("Request does not implement FilesRequest")
		return nil, ErrInvalidRequest
	}

	keys := make([]string, 0, numRequests)
	keysFromFileRequests(projectId, filesRequest, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		rejectedFiles = append(rejectedFiles, domain.NewMissingFileInfo(requests[idx].GetFileId()))
		return nil
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := getFileInfos(ctx, logger, csd.rc, keys, missingFileHandler, &files); err != nil {
			return err
		}

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		updateFiles(userId, request.GetBranchName(), files, requests, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFileInfos(ctx, logger, tx, keys, files)
	}

	// Execute the watch function
	err := csd.rc.Watch(ctx, watchFn, keys...)
	switch {
	case errors.Is(err, ErrRejectedFiles):
		logger.Info().Msg("Updating failed due to rejected files")
		return rejectedFiles, err
	case err != nil:
		logger.Error().Err(err).Msg("Failed to update files")
		return nil, err
	default:
		logger.Info().Msg("Updating successful")
		return nil, nil
	}
}

// Release releases the files with the given file IDs.
func (csd *RedisDatabase) Release(ctx context.Context, logger zerolog.Logger, userId string, projectId string, branchId string, request domain.FilesRequest) ([]*domain.FileInfo, error) {
	requests := request.GetFileIDs()

	if len(requests) == 0 {
		logger.Warn().Msg("No files to release")
		return nil, nil
	}

	logger.Info().Msgf("Releasing... %d files", len(requests))

	files := make([]*domain.FileInfo, 0, len(requests))
	rejectedFiles := make([]*domain.FileInfo, 0)

	keys := make([]string, 0, len(requests))
	keysFromFileRequests(projectId, request, &keys)

	// Handler for missing files in the Redis hash
	missingFileHandler := func(idx int) *domain.FileInfo {
		rejectedFiles = append(rejectedFiles, domain.NewMissingFileInfo(requests[idx]))
		return nil
	}

	// Watch function to ensure keys do not get modified by another request while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		if err := getFileInfos(ctx, logger, csd.rc, keys, missingFileHandler, &files); err != nil {
			return err
		}

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		releaseFiles(userId, files, &rejectedFiles)

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		return setFileInfos(ctx, logger, tx, keys, files)
	}

	// Execute the watch function
	err := csd.rc.Watch(ctx, watchFn, keys...)
	switch {
	case errors.Is(err, ErrRejectedFiles):
		logger.Info().Msg("Releasing failed due to rejected files")
		return rejectedFiles, err
	case err != nil:
		logger.Error().Err(err).Msg("Failed to release files")
		return nil, err
	default:
		logger.Info().Msg("Releasing successful")
		return nil, nil
	}
}

// claimFiles claims the files with the given file IDs
func claimFiles(userId string, fileInfos []*domain.FileInfo, claimRequests []domain.FileClaim, outRejectedFiles *[]*domain.FileInfo) {
	for i := range fileInfos {
		if err := fileInfos[i].Claim(userId, claimRequests[i].GetFileHash(), claimRequests[i].GetClaimMode()); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}

// updateFiles updates the files with the given file IDs
func updateFiles(userId string, branchName string, fileInfos []*domain.FileInfo, updateRequests []domain.FileUpdate, outRejectedFiles *[]*domain.FileInfo) {
	// update the files with the new file hashes
	for i := range fileInfos {
		if err := fileInfos[i].Update(userId, branchName, updateRequests[i].GetOldHash(), updateRequests[i].GetFileHash()); err != nil {
			// we do not return the error imediately so we can build a full list of rejected files
			// and report them back all at once
			*outRejectedFiles = append(*outRejectedFiles, fileInfos[i])
		}
	}
}

// releaseFiles releases the files with the given file IDs
func releaseFiles(userId string, fileInfos []*domain.FileInfo, outRejectedFiles *[]*domain.FileInfo) {
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

// buildScanQuery builds the query string for the Redis SCAN command
func buildScanQuery(projectId string, folderPath string) string {
	// ProjectID is safe but folder path may not be.
	return fmt.Sprintf("project:%s:file:%s*", projectId, folderPath)
}
