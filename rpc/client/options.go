package client

import (
	"crypto/tls"
	"time"

	"google.golang.org/grpc"
)

// Option is a Client's option.
type Option func(*cfg)

type cfg struct {
	addr string

	dialTimeout time.Duration

	tlsCfg *tls.Config

	conn *grpc.ClientConn
}

const defaultDialTimeout = 5 * time.Second

func defaultCfg() *cfg {
	return &cfg{
		dialTimeout: defaultDialTimeout,
	}
}

// WithNetworkAddress returns option to specify
// network address of the remote server.
//
// Ignored if WithGRPCConn is provided.
func WithNetworkAddress(v string) Option {
	return func(c *cfg) {
		if v != "" {
			c.addr = v
		}
	}
}

// WithDialTimeout returns option to specify
// dial timeout of the remote server connection.
//
// Ignored if WithGRPCConn is provided.
func WithDialTimeout(v time.Duration) Option {
	return func(c *cfg) {
		if v > 0 {
			c.dialTimeout = v
		}
	}
}

// WithTLSCfg returns option to specify
// TLS configuration.
//
// Ignored if WithGRPCConn is provided.
func WithTLSCfg(v *tls.Config) Option {
	return func(c *cfg) {
		if v != nil {
			c.tlsCfg = v
		}
	}
}

// WithGRPCConn returns option to specify
// gRPC virtual connection.
func WithGRPCConn(v *grpc.ClientConn) Option {
	return func(c *cfg) {
		if v != nil {
			c.conn = v
		}
	}
}
