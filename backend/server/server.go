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

	files := make([]*models.FileInfo, 0, len(request.Files))
	rejectedFiles := make([]*models.FileInfo, 0)
	keys := make([]string, 0, len(request.Files))
	keysFromFileClaimRequests(&keys, request.ProjectId, request.Files)

	// Watch function to ensure keys do not get modified by another client while this transaction
	// is in progress
	watchFn := func(tx *redis.Tx) error {
		result, err := tx.JSONMGet(ctx, "$", keys...).Result()
		if err != nil {
			log.Error().Err(err).Msg("Failed to read keys from Redis hash")
			return err
		}

		// Try to parse and validate each result.
		for i, res := range result {
			// If res is nil then the key does not exist in the hash this it is fair game and we should construct a new entry for it.
			if res == nil {
				newFileInfo := models.NewFileInfo(request.Files[i].FileId)
				files = append(files, newFileInfo)
			}

			var fileInfo models.FileInfo
			if err := json.Unmarshal([]byte(res.(string)), &fileInfo); err != nil {
				// TODO: This may not be the best way to handle this error but should do for now.
				log.Error().Str("key", keys[i]).Err(err).Msgf("Failed to unmarshal JSON from Redis hash")
				continue
			}

			files = append(files, &fileInfo)

			// Reject if already claimed.
			// TODO: factor in claim modes and multi-user claims
			if fileInfo.ClaimMode == pb.ClaimMode_EXCLUSIVE {
				rejectedFiles = append(rejectedFiles, &fileInfo)
			}
		}

		if len(rejectedFiles) > 0 {
			return ErrRejectedFiles
		}

		mSetParams := make([]any, 0, len(request.Files)*3)
		for i, file := range request.Files {
			fileInfo := models.NewFileInfo(file.FileId)
			mSetParams = append(mSetParams, keys[i], "$", *fileInfo)
		}

		log.Debug().Str("params", fmt.Sprintf("%v", mSetParams)).Msgf("Invoking JSONMSet")
		// Not pipelined because it uses MULTI/EXEC internally already and redis does not support
		// nested transactions.
		if err := tx.JSONMSet(ctx, mSetParams...).Err(); err != nil {
			log.Error().Err(err).Msg("Failed to write to Redis hash")
			return err
		}

		return nil
	}

	// Execute the watch function
	err := s.redisClient.Watch(ctx, watchFn, keys...)

	if errors.Is(err, ErrRejectedFiles) {
		protoRejectedFiles := make([]*pb.FileInfo, 0, len(rejectedFiles))
		models.FileInfosToProto(&protoRejectedFiles, rejectedFiles)

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
