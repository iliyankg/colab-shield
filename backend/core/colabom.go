package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type MissingFileHandler func(idx int) *models.FileInfo

// getFileInfos reads the file infos from the Redis hash and populates the outFiles slice.
// Using redis.Cmdable to allow for both a client and a transaction to be passed in.
func getFileInfos(ctx context.Context, logger zerolog.Logger, rc redis.Cmdable, keys []string, missingFileHandler MissingFileHandler, outFiles *[]*models.FileInfo) error {
	result, err := rc.JSONMGet(ctx, ".", keys...).Result()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read keys from Redis hash")
		return ErrRedisError
	}

	err = parseFileInfos(logger, keys, result, missingFileHandler, outFiles)
	if err != nil {
		return err
	}

	return nil
}

// setFileInfos writes the file infos to the Redis JSON.
// Only redis client is used because JSON MSet uses  MULTI/EXEC internally already and redis does not support nested transactions.
func setFileInfos(ctx context.Context, logger zerolog.Logger, rc redis.Cmdable, keys []string, fileInfos []*models.FileInfo) error {
	// build the mset params
	mSetParams := make([]any, 0, len(fileInfos)*3)
	for i, file := range fileInfos {
		mSetParams = append(mSetParams, keys[i], ".", *file)
	}

	logger.Debug().Str("params", fmt.Sprintf("%v", mSetParams)).Msgf("Invoking JSONMSet")
	// Not pipelined because it uses MULTI/EXEC internally already and redis does not support
	// nested transactions.
	if err := rc.JSONMSet(ctx, mSetParams...).Err(); err != nil {
		logger.Error().Err(err).Msg("Failed to write to Redis hash")
		return ErrRedisError
	}

	return nil
}

// parseFileInfos parses the file infos from the Redis hash and creates new ones where appropriate.
func parseFileInfos(logger zerolog.Logger, keys []string, toParse []any, missingFileHandler MissingFileHandler, outFileInfos *[]*models.FileInfo) error {
	// parse all files from the Redis and create new where appropriate
	for i, res := range toParse {
		// key does not exist in the DB so assume brand new
		if res == nil {
			if missingFileHandler != nil {
				*outFileInfos = append(*outFileInfos, missingFileHandler(i))
			} else {
				*outFileInfos = append(*outFileInfos, nil)
			}
			continue
		}

		// unmarshal the JSON from the Redis hash
		fileInfo := models.NewBlankFileInfo()
		if err := json.Unmarshal([]byte(res.(string)), fileInfo); err != nil {
			logger.Error().Str("key", keys[i]).Err(err).Msg("Failed to unmarshal JSON from Redis")
			return ErrUnmarshalFail
		}

		*outFileInfos = append(*outFileInfos, fileInfo)
	}

	return nil
}
