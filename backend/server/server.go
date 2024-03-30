package server

import (
	"context"

	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *ColabShieldServer) HealthCheck(context.Context, *emptypb.Empty) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.Status_OK,
	}, nil
}

func (s *ColabShieldServer) InitProject(context.Context, *pb.InitProjectRequest) (*pb.InitProjectResponse, error) {
	log.Error().Msg("InitProject not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListProjects(context.Context, *emptypb.Empty) (*pb.ListProjectsResponse, error) {
	log.Error().Msg("ListProjects not implemented")
	return nil, nil
}

func (s *ColabShieldServer) ListFiles(context.Context, *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	log.Error().Msg("ListFiles not implemented")
	return nil, nil
}

func (s *ColabShieldServer) Claim(context.Context, *pb.ClaimFilesRequest) (*pb.ClaimFilesResponse, error) {
	log.Error().Msg("Claim not implemented")
	return nil, nil
}
