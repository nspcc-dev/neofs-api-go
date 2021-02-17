package netmap

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client wraps NetmapServiceClient
// with pre-defined configurations.
type Client struct {
	*cfg

	client NetmapServiceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	callOpts []grpc.CallOption
}

// ErrNilNetmapServiceClient is returned by functions that expect
// a non-nil ContainerServiceClient, but received nil.
var ErrNilNetmapServiceClient = errors.New("netmap gRPC client is nil")

func defaultCfg() *cfg {
	return new(cfg)
}

// NewClient creates, initializes and returns a new Client instance.
//
// Options are applied one by one in order.
func NewClient(c NetmapServiceClient, opts ...Option) (*Client, error) {
	if c == nil {
		return nil, ErrNilNetmapServiceClient
	}

	cfg := defaultCfg()
	for i := range opts {
		opts[i](cfg)
	}

	return &Client{
		cfg:    cfg,
		client: c,
	}, nil
}

func (c *Client) LocalNodeInfo(ctx context.Context, req *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error) {
	return c.client.LocalNodeInfo(ctx, req, c.callOpts...)
}

func (c *Client) NetworkInfo(ctx context.Context, req *NetworkInfoRequest) (*NetworkInfoResponse, error) {
	return c.client.NetworkInfo(ctx, req, c.callOpts...)
}

// WithCallOptions returns Option that configures
// Client to attach call options to each rpc call.
func WithCallOptions(opts []grpc.CallOption) Option {
	return func(c *cfg) {
		c.callOpts = opts
	}
}
