package grpcserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/iliyankg/colab-shield/backend/core"
	"github.com/iliyankg/colab-shield/backend/core/requests"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ColabShieldServer struct {
	protos.UnimplementedColabShieldServer
	redisClient *redis.Client
}

func Serve(port int, redisClient *redis.Client) (*grpc.Server, error) {
	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	protos.RegisterColabShieldServer(grpcServer, NewColabShieldServer(redisClient))

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return nil, err
	}
	log.Info().Msgf("Grpc listening on port: %d", port)

	redisClient.Ping(context.Background())

	// Serve gRPC server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return nil, err
	}

	return grpcServer, nil
}

func NewColabShieldServer(redisClient *redis.Client) *ColabShieldServer {
	return &ColabShieldServer{
		redisClient: redisClient,
	}
}

func (s *ColabShieldServer) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*protos.HealthCheckResponse, error) {
	return &protos.HealthCheckResponse{}, nil
}

func (s *ColabShieldServer) InitProject(ctx context.Context, request *protos.InitProjectRequest) (*protos.InitProjectResponse, error) {
	log.Error().Msg("InitProject not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListProjects(ctx context.Context, _ *emptypb.Empty) (*protos.ListProjectsResponse, error) {
	log.Error().Msg("ListProjects not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListFiles(ctx context.Context, request *protos.ListFilesRequest) (*protos.ListFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Logger()

	projectId := projectIdFromCtx(ctx)

	files, cursor, err := core.List(ctx, logger, s.redisClient, projectId, request.Cursor, request.PageSize, request.FolderPath)
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

func (s *ColabShieldServer) Claim(ctx context.Context, req *protos.ClaimFilesRequest) (*protos.ClaimFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", req.BranchName).
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	res, err := json.Marshal(req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to marshal JSON from Redis hash")
		return nil, ErrMarshalFail
	}

	var internalReq requests.Claim
	if err := json.Unmarshal(res, &internalReq); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return nil, ErrUnmarshalFail
	}

	rejectedFiles, err := core.Claim(ctx, logger, s.redisClient, userId, projectId, &internalReq)
	parsedErr := parseCoreError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Claiming failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ClaimFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to claim files")
		return nil, parsedErr
	}

	// TODO: Consider returning the files that were claimed succesfully
	return &protos.ClaimFilesResponse{
		Status: protos.Status_OK,
	}, nil
}

func (s *ColabShieldServer) Update(ctx context.Context, req *protos.UpdateFilesRequest) (*protos.UpdateFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", req.BranchName).
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	res, err := json.Marshal(req)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to marshal JSON from Redis hash")
		return nil, ErrMarshalFail
	}

	var internalReq requests.Update
	if err := json.Unmarshal(res, &internalReq); err != nil {
		logger.Error().Err(err).Msg("Failed to unmarshal JSON from Redis hash")
		return nil, ErrUnmarshalFail
	}

	rejectedFiles, err := core.Update(ctx, logger, s.redisClient, userId, projectId, &internalReq)
	parsedErr := parseCoreError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Updating failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.UpdateFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to update files")
		return nil, parsedErr
	}

	logger.Info().Msg("Updating successful")

	return &protos.UpdateFilesResponse{
		Status: protos.Status_OK,
	}, nil
}

func (s *ColabShieldServer) Release(ctx context.Context, request *protos.ReleaseFilesRequest) (*protos.ReleaseFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	rejectedFiles, err := core.Release(ctx, logger, s.redisClient, userId, projectId, request.BranchName, request.FileIds)
	parsedErr := parseCoreError(err)
	if errors.Is(parsedErr, ErrRejectedFiles) {
		logger.Info().Msg("Releasing failed due to rejected files")
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ReleaseFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if parsedErr != nil {
		logger.Error().Err(parsedErr).Msg("Failed to release files")
		return nil, parsedErr
	}

	logger.Info().Msg("Releasing successful")

	return &protos.ReleaseFilesResponse{
		Status: protos.Status_OK,
	}, nil
}
