package server

import (
	"context"
	"errors"

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
