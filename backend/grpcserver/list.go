package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func listHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, _ string, projectId string, request *protos.ListFilesRequest) (*protos.ListFilesResponse, error) {
	files, cursor, err := core.List(ctx, logger, rc, projectId, request.Cursor, request.PageSize, request.FolderPath)
	if err != nil {
		return nil, parseCoreError(err)
	}

	protoFiles := make([]*protos.FileInfo, 0, len(files))
	fileInfosToProto(files, &protoFiles)
	return &protos.ListFilesResponse{
		NextCursor: cursor,
		Files:      protoFiles,
	}, nil
}
