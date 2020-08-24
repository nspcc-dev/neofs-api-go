package accounting

import (
	"context"

	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/client"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client represents universal accounting
// transport client.
type Client struct {
	client *getBalanceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	proto client.Protocol

	globalOpts []client.Option

	gRPC cfgGRPC
}

type cfgGRPC struct {
	serviceClient accounting.AccountingServiceClient

	grpcCallOpts []grpc.CallOption

	callOpts []accounting.Option

	client *accounting.Client
}

type getBalanceClient struct {
	requestConverter func(*BalanceRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *BalanceResponse
}

// Balance sends BalanceRequest over the network and returns BalanceResponse.
//
// It returns any error encountered during the call.
func (c *Client) Balance(ctx context.Context, req *BalanceRequest) (*BalanceResponse, error) {
	resp, err := c.client.caller(ctx, c.client.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send balance request")
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
		var c *accounting.Client
		if c, err = newGRPCClient(cfg); err != nil {
			break
		}

		return &Client{
			client: &getBalanceClient{
				requestConverter: func(req *BalanceRequest) interface{} {
					return BalanceRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.Balance(ctx, req.(*accounting.BalanceRequest))
				},
				responseConverter: func(resp interface{}) *BalanceResponse {
					return BalanceResponseFromGRPCMessage(resp.(*accounting.BalanceResponse))
				},
			},
		}, nil
	default:
		err = client.ErrProtoUnsupported
	}

	return nil, errors.Wrapf(err, "could not create %s Accounting client", cfg.proto)
}

func newGRPCClient(cfg *cfg) (*accounting.Client, error) {
	var err error

	if cfg.gRPC.client == nil {
		if cfg.gRPC.serviceClient == nil {
			conn, err := client.NewGRPCClientConn(cfg.globalOpts...)
			if err != nil {
				return nil, errors.Wrap(err, "could not open gRPC client connection")
			}

			cfg.gRPC.serviceClient = accounting.NewAccountingServiceClient(conn)
		}

		cfg.gRPC.client, err = accounting.NewClient(
			cfg.gRPC.serviceClient,
			append(
				cfg.gRPC.callOpts,
				accounting.WithCallOptions(cfg.gRPC.grpcCallOpts),
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

func WithGRPCServiceClient(v accounting.AccountingServiceClient) Option {
	return func(c *cfg) {
		c.gRPC.serviceClient = v
	}
}

func WithGRPCCallOpts(v []grpc.CallOption) Option {
	return func(c *cfg) {
		c.gRPC.grpcCallOpts = v
	}
}

func WithGRPCClientOpts(v []accounting.Option) Option {
	return func(c *cfg) {
		c.gRPC.callOpts = v
	}
}

func WithGRPCClient(v *accounting.Client) Option {
	return func(c *cfg) {
		c.gRPC.client = v
	}
}
