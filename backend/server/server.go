package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/iliyankg/colab-shield/backend/models"
	pb "github.com/iliyankg/colab-shield/protos"
	"github.com/redis/go-redis/v9"
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
	log.Info().Msgf("Claiming files for project %s, branch %s, user %s", request.ProjectId, request.BranchName, request.UserId)

	pipelinedFn := func(pipe redis.Pipeliner) error {
		for _, file := range request.Files {
			key := newRedisKeyFile(request.ProjectId, file.FileId)
			fileInfo := models.NewFileInfoFromProto(file.FileId, file.FileHash, request.UserId, request.BranchName, true)
			err := pipe.JSONSet(ctx, key, "$", fileInfo).Err()
			if err != nil {
				log.Error().Err(err).Msg("Failed to write to Redis hash")
				return err
			}
		}
		return nil
	}

	keys := make([]string, 0, len(request.Files))
	keysFromFileClaimRequests(&keys, request.ProjectId, request.Files)
	rejectedFiles := make([]*models.FileInfo, 0)
	watchFn := func(tx *redis.Tx) error {
		for _, key := range keys {
			result, err := tx.JSONGet(ctx, key, "$").Result()
			if err != nil {
				log.Error().Err(err).Msg("Failed to read keys from Redis hash")
				continue
			}

			var fileInfo models.FileInfo
			if err := json.Unmarshal([]byte(result), &fileInfo); err != nil {
				log.Error().Err(err).Msg("Failed to unmarshal JSON from Redis hash")
				continue
			}

			if fileInfo.Claimed {
				rejectedFiles = append(rejectedFiles, &fileInfo)
				continue
			}
		}

		_, err := tx.TxPipelined(ctx, pipelinedFn)
		return err
	}

	// Execute the watch function
	err := s.redisClient.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		protoRejectedFiles := make([]*pb.FileInfo, 0, len(rejectedFiles))
		for _, file := range rejectedFiles {
			protoRejectedFiles = append(protoRejectedFiles, file.ToProto())
		}

		return &pb.ClaimFilesResponse{
			Status:        pb.Status_ERROR,
			RejectedFiles: protoRejectedFiles,
		}, nil
	} else if err != nil {
		log.Error().Err(err).Msg("Failed to claim files")
		return nil, err
	}

	return &pb.ClaimFilesResponse{
		Status: pb.Status_OK,
	}, nil
}
