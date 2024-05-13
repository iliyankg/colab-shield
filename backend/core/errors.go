package core

import "errors"

var (
	// Common status error for rejected files regardless of internal reason.
	ErrInvalidRequest = errors.New("invalid request")
	ErrUnmarshalFail  = errors.New("failed to unmarshal JSON from Redis hash")
	ErrRedisError     = errors.New("encountered an error with Redis")
	ErrRejectedFiles  = errors.New("rejected files")
)
