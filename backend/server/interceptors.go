package server

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type colabShieldContextKey string

const (
	UserIdKey    colabShieldContextKey = "userId"
	ProjectIdKey colabShieldContextKey = "projectId"
)

var (
	ErrMissingMetadata           = errors.New("missing metadata")
	ErrMissingOrInvalidUserId    = errors.New("missing or invalid userId")
	ErrMissingOrInvalidProjectId = errors.New("missing or invalid projectId")
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}

	// Get userId and projectId from metadata
	userIds := md.Get("userId")
	if len(userIds) == 0 {
		return nil, ErrMissingOrInvalidUserId
	}
	projectIds := md.Get("projectId")
	if len(projectIds) == 0 {
		return nil, ErrMissingOrInvalidProjectId
	}

	userId := userIds[0]
	projectId := projectIds[0]

	ctx = buildCollabShieldContext(ctx, userId, projectId, info.FullMethod)

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
