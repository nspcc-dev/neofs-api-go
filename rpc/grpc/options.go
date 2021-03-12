package grpc

import (
	"google.golang.org/grpc"
)

type cfg struct {
	con *grpc.ClientConn
}

func defaultCfg() *cfg {
	return new(cfg)
}

// WithClientConnection returns option to set gRPC connection
// to the remote server.
func WithClientConnection(con *grpc.ClientConn) Option {
	return func(c *cfg) {
		c.con = con
	}
}
