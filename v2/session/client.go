package session

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/client"
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client represents universal session
// transport client.
type Client struct {
	client *createClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	proto client.Protocol

	globalOpts []client.Option

	gRPC cfgGRPC
}

type cfgGRPC struct {
	serviceClient session.SessionServiceClient

	grpcCallOpts []grpc.CallOption

	callOpts []session.Option

	client *session.Client
}

type createClient struct {
	requestConverter func(*CreateRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *CreateResponse
}

// Create sends CreateRequest over the network and returns CreateResponse.
//
// It returns any error encountered during the call.
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	resp, err := c.client.caller(ctx, c.client.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send session init request")
	}

	return c.client.responseConverter(resp), nil
}

func defaultCfg() *cfg {
	return &cfg{
		proto: client.ProtoGRPC,
	}
}

func NewClient(opts ...Option) (*Client, error) {
	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	var err error

	switch cfg.proto {
	case client.ProtoGRPC:
		var c *session.Client
		if c, err = newGRPCClient(cfg); err != nil {
			break
		}

		return &Client{
			client: &createClient{
				requestConverter: func(req *CreateRequest) interface{} {
					return CreateRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.Create(ctx, req.(*session.CreateRequest))
				},
				responseConverter: func(resp interface{}) *CreateResponse {
					return CreateResponseFromGRPCMessage(resp.(*session.CreateResponse))
				},
			},
		}, nil
	default:
		err = client.ErrProtoUnsupported
	}

	return nil, errors.Wrapf(err, "could not create %s Session client", cfg.proto)
}

func newGRPCClient(cfg *cfg) (*session.Client, error) {
	var err error

	if cfg.gRPC.client == nil {
		if cfg.gRPC.serviceClient == nil {
			conn, err := client.NewGRPCClientConn(cfg.globalOpts...)
			if err != nil {
				return nil, errors.Wrap(err, "could not open gRPC client connection")
			}

			cfg.gRPC.serviceClient = session.NewSessionServiceClient(conn)
		}

		cfg.gRPC.client, err = session.NewClient(
			cfg.gRPC.serviceClient,
			append(
				cfg.gRPC.callOpts,
				session.WithCallOptions(cfg.gRPC.grpcCallOpts),
			)...,
		)
	}

	return cfg.gRPC.client, err
}

func WithGlobalOpts(v ...client.Option) Option {
	return func(c *cfg) {
		if len(v) > 0 {
			c.globalOpts = v
		}
	}
}

func WithGRPCServiceClient(v session.SessionServiceClient) Option {
	return func(c *cfg) {
		c.gRPC.serviceClient = v
	}
}

func WithGRPCCallOpts(v []grpc.CallOption) Option {
	return func(c *cfg) {
		c.gRPC.grpcCallOpts = v
	}
}

func WithGRPCClientOpts(v []session.Option) Option {
	return func(c *cfg) {
		c.gRPC.callOpts = v
	}
}

func WithGRPCClient(v *session.Client) Option {
	return func(c *cfg) {
		c.gRPC.client = v
	}
}
