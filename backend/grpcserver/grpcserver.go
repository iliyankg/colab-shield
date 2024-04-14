package grpcserver

import (
	"context"
	"fmt"
	"net"

	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ColabShieldServer struct {
	pb.UnimplementedColabShieldServer
	redisClient *redis.Client
}

func Serve(port int, redisClient *redis.Client) (*grpc.Server, error) {
	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	pb.RegisterColabShieldServer(grpcServer, NewColabShieldServer(redisClient))

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return nil, err
	}
	log.Info().Msgf("Listening on port: %d", port)

	// Serve gRPC server
	log.Info().Msg("Serving gRPC")
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

func (s *ColabShieldServer) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{}, nil
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
	logger := zerolog.Ctx(ctx).
		With().
		Logger()

	return listHandler(ctx, logger, s.redisClient, userIdFromCtx(ctx), projectIdFromCtx(ctx), request)
}

func (s *ColabShieldServer) Claim(ctx context.Context, request *pb.ClaimFilesRequest) (*pb.ClaimFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", request.BranchName).
		Logger()

	return claimHandler(ctx, logger, s.redisClient, userIdFromCtx(ctx), projectIdFromCtx(ctx), request)
}

func (s *ColabShieldServer) Update(ctx context.Context, request *pb.UpdateFilesRequest) (*pb.UpdateFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", request.BranchName).
		Logger()

	return updateHandler(ctx, logger, s.redisClient, userIdFromCtx(ctx), projectIdFromCtx(ctx), request)
}

func (s *ColabShieldServer) Release(ctx context.Context, request *pb.ReleaseFilesRequest) (*pb.ReleaseFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Logger()

	return releaseHandler(ctx, logger, s.redisClient, userIdFromCtx(ctx), projectIdFromCtx(ctx), request)
}
