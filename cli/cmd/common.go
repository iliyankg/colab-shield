package cmd

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/metadata"
)

var (
	ErrFileToHashMissmatch = errors.New("files and hashes must be of the same length")
)

func buildContext(projectId string, userId string) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	metaInfo := metadata.Pairs(
		"projectId", projectId,
		"userId", userId,
	)

	return metadata.NewOutgoingContext(ctx, metaInfo), cancel
}
