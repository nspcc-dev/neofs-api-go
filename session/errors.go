package session

import "github.com/nspcc-dev/neofs-api-go/internal"

// ErrNilCreateParamsSource is returned by functions that expect a non-nil
// CreateParamsSource, but received nil.
const ErrNilCreateParamsSource = internal.Error("create params source is nil")

// ErrNilGPRCClientConn is returned by functions that expect a non-nil
// grpc.ClientConn, but received nil.
const ErrNilGPRCClientConn = internal.Error("gRPC client connection is nil")

// ErrPrivateTokenNotFound is returned when addressed private token was
// not found in storage.
const ErrPrivateTokenNotFound = internal.Error("private token not found")
