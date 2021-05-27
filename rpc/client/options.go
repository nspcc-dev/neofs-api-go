package client

import (
	"crypto/tls"
	"net/url"
	"time"

	"google.golang.org/grpc"
)

const (
	grpcScheme    = "grpc"
	grpcTLSScheme = "grpcs"
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

// WithNetworkURIAddress combines WithNetworkAddress and WithTLSCfg options
// based on arguments.
//
// Do not use along with WithNetworkAddress and WithTLSCfg.
//
// Ignored if WithGRPCConn is provided.
func WithNetworkURIAddress(addr string, tlsCfg *tls.Config) []Option {
	uri, err := url.ParseRequestURI(addr)
	if err != nil {
		return []Option{WithNetworkAddress(addr)}
	}

	// check if passed string was parsed correctly
	// URIs that do not start with a slash after the scheme are interpreted as:
	// `scheme:opaque` => if `opaque` is not empty, then it is supposed that URI
	// is in `host:port` format
	if uri.Opaque != "" {
		return []Option{WithNetworkAddress(addr)}
	}

	switch uri.Scheme {
	case grpcScheme:
		tlsCfg = nil
	case grpcTLSScheme:
		if tlsCfg == nil {
			tlsCfg = &tls.Config{}
		}
	default:
		// not supported scheme
		return nil
	}

	return []Option{
		WithNetworkAddress(uri.Host),
		WithTLSCfg(tlsCfg),
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
		c.tlsCfg = v
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
