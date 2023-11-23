package grpc

import (
	"context"
)

// CallOption is a messaging session option within RPC.
type CallOption func(*callParameters)

type callParameters struct {
	ctx context.Context

	allowBinarySendingOnly bool

	syncWrite bool
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

// AllowBinarySendingOnly allows to pass []byte argument only to
// [MessageReadWriter.WriteMessage] method. By default, only [proto.Message]
// instances may be used. Use this option when binary message transmission is
// needed. Note that only [proto.Message] response messages are supported even
// with this option.
func AllowBinarySendingOnly() CallOption {
	return func(prm *callParameters) {
		prm.allowBinarySendingOnly = true
	}
}

// SyncWrite makes each [MessageReadWriter.WriteMessage] call to wait for
// message to be completely written out to the underlying client network
// connection. By default, the method may return before writing to the wire.
func SyncWrite() CallOption {
	return func(prm *callParameters) {
		prm.syncWrite = true
	}
}
