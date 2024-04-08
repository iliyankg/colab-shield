package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/models"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MissingFileHandler func(idx int) *models.FileInfo
type UnmarshalFailHandler func(idx int, err error) error

var (
	ErrRejectedFiles = status.Error(codes.FailedPrecondition, "rejected files")
)

// parseFileInfos parses the file infos from the Redis hash and creates new ones where appropriate.
func parseFileInfos(toParse []any, outFileInfos *[]*models.FileInfo, missingFileHandler MissingFileHandler, unmarshalFailHandler UnmarshalFailHandler) error {
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
		var fileInfo models.FileInfo
		if err := json.Unmarshal([]byte(res.(string)), &fileInfo); err != nil {
			if unmarshalFailHandler == nil {
				continue
			}

			err = unmarshalFailHandler(i, err)
			if err != nil {
				return err
			}
		}

		*outFileInfos = append(*outFileInfos, &fileInfo)
	}

	return nil
}

func setFiles(ctx context.Context, logger zerolog.Logger, redisClient *redis.Client, keys []string, fileInfos []*models.FileInfo) error {
	// build the mset params
	mSetParams := make([]any, 0, len(fileInfos)*3)
	for i, file := range fileInfos {
		mSetParams = append(mSetParams, keys[i], "$", *file)
	}

	logger.Debug().Str("params", fmt.Sprintf("%v", mSetParams)).Msgf("Invoking JSONMSet")
	// Not pipelined because it uses MULTI/EXEC internally already and redis does not support
	// nested transactions.
	if err := redisClient.JSONMSet(ctx, mSetParams...).Err(); err != nil {
		logger.Error().Err(err).Msg("Failed to write to Redis hash")
		return err
	}

	return nil
}
