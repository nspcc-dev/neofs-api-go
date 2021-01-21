package container

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client wraps ContainerServiceClient
// with pre-defined configurations.
type Client struct {
	*cfg

	client ContainerServiceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	callOpts []grpc.CallOption
}

// ErrNilContainerServiceClient is returned by functions that expect
// a non-nil ContainerServiceClient, but received nil.
var ErrNilContainerServiceClient = errors.New("container gRPC client is nil")

func defaultCfg() *cfg {
	return new(cfg)
}

// NewClient creates, initializes and returns a new Client instance.
//
// Options are applied one by one in order.
func NewClient(c ContainerServiceClient, opts ...Option) (*Client, error) {
	if c == nil {
		return nil, ErrNilContainerServiceClient
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

func (c *Client) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	return c.client.Put(ctx, req, c.callOpts...)
}

func (c *Client) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return c.client.Get(ctx, req, c.callOpts...)
}

func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return c.client.Delete(ctx, req, c.callOpts...)
}

func (c *Client) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return c.client.List(ctx, req, c.callOpts...)
}

func (c *Client) SetExtendedACL(ctx context.Context, req *SetExtendedACLRequest) (*SetExtendedACLResponse, error) {
	return c.client.SetExtendedACL(ctx, req, c.callOpts...)
}

func (c *Client) GetExtendedACL(ctx context.Context, req *GetExtendedACLRequest) (*GetExtendedACLResponse, error) {
	return c.client.GetExtendedACL(ctx, req, c.callOpts...)
}

func (c *Client) AnnounceUsedSpace(ctx context.Context, req *AnnounceUsedSpaceRequest) (*AnnounceUsedSpaceResponse, error) {
	return c.client.AnnounceUsedSpace(ctx, req, c.callOpts...)
}

// WithCallOptions returns Option that configures
// Client to attach call options to each rpc call.
func WithCallOptions(opts []grpc.CallOption) Option {
	return func(c *cfg) {
		c.callOpts = opts
	}
}
