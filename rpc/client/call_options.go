package client

import (
	"context"
)

// CallOption is a messaging session option within Protobuf RPC.
type CallOption func(*callParameters)

type callParameters struct {
	ctx context.Context

	allowBinarySendingOnly bool
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

// AllowBinarySendingOnly allows only [MessageWriter.WriteMessage] method's
// arguments that are convertible to binary gRPC messages ([]byte). For example,
// [BinaryMessage] may be used for such write. By default, only arguments
// convertible to [proto.Message] may be used. Use this option when binary
// message transmission is needed. Note that only [proto.Message] convertible
// response messages are supported even with this option.
func AllowBinarySendingOnly() CallOption {
	return func(prm *callParameters) {
		prm.allowBinarySendingOnly = true
	}
}
