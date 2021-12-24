package grpc

import (
	"time"

	"google.golang.org/grpc"
)

const defaultRWTimeout = 1 * time.Minute

type cfg struct {
	con       *grpc.ClientConn
	rwTimeout time.Duration
}

func defaultCfg() *cfg {
	return &cfg{
		rwTimeout: defaultRWTimeout,
	}
}

// WithClientConnection returns option to set gRPC connection
// to the remote server.
func WithClientConnection(con *grpc.ClientConn) Option {
	return func(c *cfg) {
		c.con = con
	}
}

// WithRWTimeout returns option to specify rwTimeout
// for reading and writing single gRPC message.
func WithRWTimeout(t time.Duration) Option {
	return func(c *cfg) {
		c.rwTimeout = t
	}
}
