package client

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Option func(*cfg)

type cfg struct {
	addr string

	conn net.Conn

	gRPC cfgGRPC

	dialTimeout time.Duration
}

type cfgGRPC struct {
	dialOpts []grpc.DialOption

	conn *grpc.ClientConn
}

const defaultDialTimeout = 5 * time.Second

func defaultCfg() *cfg {
	return &cfg{
		gRPC: cfgGRPC{
			dialOpts: []grpc.DialOption{
				grpc.WithInsecure(),
			},
		},
		dialTimeout: defaultDialTimeout,
	}
}

func NewGRPCClientConn(opts ...Option) (*grpc.ClientConn, error) {
	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	var err error

	if cfg.gRPC.conn == nil {
		if cfg.conn != nil {
			if cfg.addr == "" {
				cfg.addr = cfg.conn.RemoteAddr().String()
			}

			cfg.gRPC.dialOpts = append(cfg.gRPC.dialOpts,
				grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
					return cfg.conn, nil
				}),
			)
		}

		dialCtx, cancel := context.WithTimeout(context.Background(), cfg.dialTimeout)

		cfg.gRPC.conn, err = grpc.DialContext(dialCtx, cfg.addr, cfg.gRPC.dialOpts...)

		cancel()
		if err != nil {
			return nil, err
		}
	}

	return cfg.gRPC.conn, err
}

func WithNetworkAddress(v string) Option {
	return func(c *cfg) {
		if v != "" {
			c.addr = v
		}
	}
}

func WithNetConn(v net.Conn) Option {
	return func(c *cfg) {
		if v != nil {
			c.conn = v
		}
	}
}

func WithDialTimeout(v time.Duration) Option {
	return func(c *cfg) {
		if v > 0 {
			c.dialTimeout = v
		}
	}
}

func WithGRPCDialOpts(v []grpc.DialOption) Option {
	return func(c *cfg) {
		if len(v) > 0 {
			c.gRPC.dialOpts = v
		}
	}
}

func WithGRPCConn(v *grpc.ClientConn) Option {
	return func(c *cfg) {
		if v != nil {
			c.gRPC.conn = v
		}
	}
}
