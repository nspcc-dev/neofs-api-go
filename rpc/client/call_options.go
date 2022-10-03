package client

import (
	"context"
)

// CallOption is a messaging session option within Protobuf RPC.
type CallOption func(*callParameters)

type callParameters struct {
	ctx context.Context
}

func defaultCallParameters() *callParameters {
	return &callParameters{
		ctx: context.Background(),
	}
}

// WithContext returns option to specify call context. If provided, all network
// communications will be based on this context. Otherwise, context.Background()
// is used.
//
// Context SHOULD NOT be nil.
func WithContext(ctx context.Context) CallOption {
	return func(prm *callParameters) {
		prm.ctx = ctx
	}
}
