package accounting

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client wraps AccountingServiceClient
// with pre-defined configurations.
type Client struct {
	*cfg

	client AccountingServiceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	callOpts []grpc.CallOption
}

// ErrNilAccountingServiceClient is returned by functions that expect
// a non-nil AccountingServiceClient, but received nil.
var ErrNilAccountingServiceClient = errors.New("accounting gRPC client is nil")

func defaultCfg() *cfg {
	return new(cfg)
}

// NewClient creates, initializes and returns a new Client instance.
//
// Options are applied one by one in order.
func NewClient(c AccountingServiceClient, opts ...Option) (*Client, error) {
	if c == nil {
		return nil, ErrNilAccountingServiceClient
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

func (c *Client) Balance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	return c.client.Balance(ctx, req, c.callOpts...)
}

// WithCallOptions returns Option that configures
// Client to attach call options to each rpc call.
func WithCallOptions(opts []grpc.CallOption) Option {
	return func(c *cfg) {
		c.callOpts = opts
	}
}
