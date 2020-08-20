package session

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client wraps SessionServiceClient
// with pre-defined configurations.
type Client struct {
	*cfg

	client SessionServiceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	callOpts []grpc.CallOption
}

// ErrNilSessionServiceClient is returned by functions that expect
// a non-nil SessionServiceClient, but received nil.
var ErrNilSessionServiceClient = errors.New("session gRPC client is nil")

func defaultCfg() *cfg {
	return new(cfg)
}

// NewClient creates, initializes and returns a new Client instance.
//
// Options are applied one by one in order.
func NewClient(c SessionServiceClient, opts ...Option) (*Client, error) {
	if c == nil {
		return nil, ErrNilSessionServiceClient
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

func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return c.client.Create(ctx, req, c.callOpts...)
}

// WithCallOptions returns Option that configures
// Client to attach call options to each rpc call.
func WithCallOptions(opts []grpc.CallOption) Option {
	return func(c *cfg) {
		c.callOpts = opts
	}
}
