package grpcserver

import (
	"context"

	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

func getFilesHandler(ctx context.Context, logger zerolog.Logger, rc *redis.Client, _ string, projectId string, request *protos.GetFilesRequest) (*protos.GetFilesResponse, error) {
	files, err := core.GetFiles(ctx, logger, rc, projectId, request.FileIds)
	if err != nil {
		return nil, parseCoreErrorToGrpc(err)
	}

	protoFiles := make([]*protos.FileInfo, 0, len(files))
	fileInfosToProto(files, &protoFiles)
	return &protos.GetFilesResponse{
		Files: protoFiles,
	}, nil
}
