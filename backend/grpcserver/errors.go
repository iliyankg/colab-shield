package grpcserver

import (
	"errors"

	"github.com/iliyankg/colab-shield/backend/colabom"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Common status error for rejected files regardless of internal reason.
	ErrRejectedFiles = status.Error(codes.FailedPrecondition, "rejected files")
	ErrUnmarshalFail = status.Error(codes.Internal, "failed to unmarshal JSON from Redis hash")
	ErrRedisError    = status.Error(codes.Internal, "encountered an error with Redis")
	ErrUnknown       = status.Error(codes.Unknown, "unknown error")
)

// parseColabomError converts colabom errors to gRPC status errors.
func parseColabomError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, colabom.ErrUnmarshalFail) {
		return ErrUnmarshalFail
	}

	if errors.Is(err, colabom.ErrRedisError) {
		return ErrRedisError
	}

	return ErrUnknown
}
