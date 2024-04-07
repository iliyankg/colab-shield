package server

import (
	"context"
	"fmt"

	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrRejectedFiles = fmt.Errorf("rejected files")
)

type ColabShieldServer struct {
	pb.UnimplementedColabShieldServer
	redisClient *redis.Client
}

func NewColabShieldServer(redisClient *redis.Client) *ColabShieldServer {
	return &ColabShieldServer{
		redisClient: redisClient,
	}
}

func (s *ColabShieldServer) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.Status_OK,
	}, nil
}

func (s *ColabShieldServer) InitProject(ctx context.Context, request *pb.InitProjectRequest) (*pb.InitProjectResponse, error) {
	log.Error().Msg("InitProject not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListProjects(ctx context.Context, _ *emptypb.Empty) (*pb.ListProjectsResponse, error) {
	log.Error().Msg("ListProjects not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListFiles(ctx context.Context, request *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	log.Error().Msg("ListFiles not implemented")
	return nil, nil
}

func (s *ColabShieldServer) Claim(ctx context.Context, request *pb.ClaimFilesRequest) (*pb.ClaimFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", request.BranchName).
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	return claimHandler(ctx, logger, s.redisClient, userId, projectId, request)
}

func (s *ColabShieldServer) Update(ctx context.Context, request *pb.UpdateFilesRequest) (*pb.UpdateFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", request.BranchName).
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	return updateHandler(ctx, logger, s.redisClient, userId, projectId, request)
}
