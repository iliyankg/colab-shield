package grpcserver

import (
	"errors"

	"github.com/iliyankg/colab-shield/backend/redisdatabase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Common status error for rejected files regardless of internal reason.
	ErrRejectedFiles = status.Error(codes.FailedPrecondition, "rejected files")
	ErrMarshalFail   = status.Error(codes.Internal, "failed to marshal JSON")
	ErrUnmarshalFail = status.Error(codes.Internal, "failed to unmarshal JSON")
	ErrRedisError    = status.Error(codes.Internal, "encountered an error with Redis")
	ErrUnknown       = status.Error(codes.Unknown, "unknown error")
)

// parseCoreErrorToGrpc converts colabom errors to gRPC status errors.
func parseCoreErrorToGrpc(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, redisdatabase.ErrUnmarshalFail) {
		return ErrUnmarshalFail
	}

	if errors.Is(err, redisdatabase.ErrRedisError) {
		return ErrRedisError
	}

	return ErrUnknown
}
