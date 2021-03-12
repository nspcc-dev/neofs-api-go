package grpc

import (
	"context"
)

// CallOption is a messaging session option within RPC.
type CallOption func(*callParameters)

type callParameters struct {
	ctx context.Context
}

func defaultCallParameters() *callParameters {
	return &callParameters{
		ctx: context.Background(),
	}
}

// WithContext returns option to set RPC context.
func WithContext(ctx context.Context) CallOption {
	return func(prm *callParameters) {
		prm.ctx = ctx
	}
}
