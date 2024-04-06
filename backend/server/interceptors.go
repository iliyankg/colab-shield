package server

import (
	"context"
	"errors"
	"os"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

var (
	ErrMissingOrInvalidUserId    = errors.New("missing or invalid userId")
	ErrMissingOrInvalidProjectId = errors.New("missing or invalid projectId")
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return nil, ErrMissingOrInvalidUserId
	}
	projectId, ok := ctx.Value("projectId").(string)
	if !ok {
		return nil, ErrMissingOrInvalidProjectId
	}

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("userId", userId).
		Str("projectId", projectId).
		Str("method", info.FullMethod).
		Logger()

	ctx = logger.WithContext(ctx)

	resp, err := handler(ctx, req)

	return resp, err
}
