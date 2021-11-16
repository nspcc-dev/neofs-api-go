package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/grpc"
)

// CallOption is a messaging session option within Protobuf RPC.
type CallOption func(*callParameters)

type callParameters struct {
	callOpts []grpc.CallOption
}

func defaultCallParameters() *callParameters {
	return &callParameters{
		callOpts: make([]grpc.CallOption, 0, 1),
	}
}

// WithContext return options to specify call context.
func WithContext(ctx context.Context) CallOption {
	return func(prm *callParameters) {
		prm.callOpts = append(prm.callOpts, grpc.WithContext(ctx))
	}
}
