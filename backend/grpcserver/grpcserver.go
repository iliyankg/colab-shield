package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/iliyankg/colab-shield/backend/domain"
	"github.com/iliyankg/colab-shield/protos"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ColabShieldServer is serves a gRPC endpoint.
type ColabShieldServer struct {
	protos.UnimplementedColabShieldServer
	db domain.ColabDatabase
}

func NewColabShieldServer(db domain.ColabDatabase) *ColabShieldServer {
	return &ColabShieldServer{
		db: db,
	}
}

func (css *ColabShieldServer) Serve(port int) (*grpc.Server, error) {
	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
	)
	protos.RegisterColabShieldServer(grpcServer, css)

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
		return nil, err
	}
	log.Info().Msgf("Grpc listening on port: %d", port)

	if err := css.db.Ping(); err != nil {
		return nil, err
	}

	// Serve gRPC server
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error().Err(err).Msg("failed to serve")
		return nil, err
	}

	return grpcServer, nil
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

	files, cursor, err := s.db.List(ctx, logger, projectId, request.Cursor, request.PageSize, request.FolderPath)
	if err != nil {
		return nil, parseCoreErrorToGrpc(err)
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
	domainReq := newClaimRequest(req)

	rejectedFiles, err := s.db.Claim(ctx, logger, userId, projectId, domainReq)

	parsedErr := parseCoreErrorToGrpc(err)
	switch {
	case errors.Is(parsedErr, ErrRejectedFiles):
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ClaimFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	case parsedErr != nil:
		return nil, parsedErr
	default:
		return &protos.ClaimFilesResponse{
			Status: protos.Status_OK,
		}, nil
	}
}

func (s *ColabShieldServer) Update(ctx context.Context, req *protos.UpdateFilesRequest) (*protos.UpdateFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Str("branchName", req.BranchName).
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)
	domainReq := newUpdateRequest(req)

	rejectedFiles, err := s.db.Update(ctx, logger, userId, projectId, domainReq)
	parsedErr := parseCoreErrorToGrpc(err)
	switch {
	case errors.Is(parsedErr, ErrRejectedFiles):
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.UpdateFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	case parsedErr != nil:
		return nil, parsedErr
	default:
		return &protos.UpdateFilesResponse{
			Status: protos.Status_OK,
		}, nil
	}
}

func (s *ColabShieldServer) Release(ctx context.Context, request *protos.ReleaseFilesRequest) (*protos.ReleaseFilesResponse, error) {
	logger := zerolog.Ctx(ctx).
		With().
		Logger()

	userId := userIdFromCtx(ctx)
	projectId := projectIdFromCtx(ctx)

	rejectedFiles, err := s.db.Release(ctx, logger, userId, projectId, request.BranchName, domain.NewFilesRequest(request.FileIds))

	parsedErr := parseCoreErrorToGrpc(err)
	switch {
	case errors.Is(parsedErr, ErrRejectedFiles):
		protoRejectedFiles := make([]*protos.FileInfo, 0, len(rejectedFiles))
		fileInfosToProto(rejectedFiles, &protoRejectedFiles)

		return &protos.ReleaseFilesResponse{
			Status:        protos.Status_REJECTED,
			RejectedFiles: protoRejectedFiles,
		}, nil
	case parsedErr != nil:
		return nil, parsedErr
	default:
		return &protos.ReleaseFilesResponse{
			Status: protos.Status_OK,
		}, nil
	}
}
