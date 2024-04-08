package server

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type colabShieldContextKey string

const (
	UserIdKey    colabShieldContextKey = "userId"
	ProjectIdKey colabShieldContextKey = "projectId"
)

var (
	// TODO: Unauthenticated may not be the best code for these errors but should do for now.
	ErrMissingMetadata           = status.Error(codes.Unauthenticated, "missing metadata")
	ErrMissingOrInvalidUserId    = status.Error(codes.Unauthenticated, "missing or invalid userId")
	ErrMissingOrInvalidProjectId = status.Error(codes.Unauthenticated, "missing or invalid projectId")
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}

	// Get userId and projectId from metadata
	userIds := md.Get("userId")
	if len(userIds) != 1 {
		return nil, ErrMissingOrInvalidUserId
	}
	projectIds := md.Get("projectId")
	if len(projectIds) != 1 {
		return nil, ErrMissingOrInvalidProjectId
	}

	if userIds[0] == "" {
		return nil, ErrMissingOrInvalidUserId
	}

	if projectIds[0] == "" {
		return nil, ErrMissingOrInvalidProjectId
	}

	ctx = buildCollabShieldContext(ctx, userIds[0], projectIds[0], info.FullMethod)

	resp, err := handler(ctx, req)

	return resp, err
}

func buildCollabShieldContext(ctx context.Context, userId string, projectId string, method string) context.Context {
	ctx = context.WithValue(ctx, UserIdKey, userId)
	ctx = context.WithValue(ctx, ProjectIdKey, projectId)

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("userId", userId).
		Str("projectId", projectId).
		Str("method", method).
		Logger()

	return logger.WithContext(ctx)
}
