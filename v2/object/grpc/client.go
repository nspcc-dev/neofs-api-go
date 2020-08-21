package object

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client wraps ObjectServiceClient
// with pre-defined configurations.
type Client struct {
	*cfg

	client ObjectServiceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	callOpts []grpc.CallOption
}

// ErrNilObjectServiceClient is returned by functions that expect
// a non-nil ObjectServiceClient, but received nil.
var ErrNilObjectServiceClient = errors.New("object gRPC client is nil")

func defaultCfg() *cfg {
	return new(cfg)
}

// NewClient creates, initializes and returns a new Client instance.
//
// Options are applied one by one in order.
func NewClient(c ObjectServiceClient, opts ...Option) (*Client, error) {
	if c == nil {
		return nil, ErrNilObjectServiceClient
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

func (c *Client) Get(ctx context.Context, req *GetRequest) (ObjectService_GetClient, error) {
	return c.client.Get(ctx, req, c.callOpts...)
}

func (c *Client) Put(ctx context.Context) (ObjectService_PutClient, error) {
	return c.client.Put(ctx, c.callOpts...)
}

func (c *Client) Head(ctx context.Context, req *HeadRequest) (*HeadResponse, error) {
	return c.client.Head(ctx, req, c.callOpts...)
}

func (c *Client) Search(ctx context.Context, req *SearchRequest) (ObjectService_SearchClient, error) {
	return c.client.Search(ctx, req, c.callOpts...)
}

func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return c.client.Delete(ctx, req, c.callOpts...)
}

func (c *Client) GetRange(ctx context.Context, req *GetRangeRequest) (ObjectService_GetRangeClient, error) {
	return c.client.GetRange(ctx, req, c.callOpts...)
}

func (c *Client) GetRangeHash(ctx context.Context, req *GetRangeHashRequest) (*GetRangeHashResponse, error) {
	return c.client.GetRangeHash(ctx, req, c.callOpts...)
}

// WithCallOptions returns Option that configures
// Client to attach call options to each rpc call.
func WithCallOptions(opts []grpc.CallOption) Option {
	return func(c *cfg) {
		c.callOpts = opts
	}
}
