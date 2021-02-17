package netmap

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/client"
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client represents universal netmap transport client.
type Client struct {
	cLocalNodeInfo *localNodeInfoClient

	cNetworkInfo *networkInfoClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	proto client.Protocol

	globalOpts []client.Option

	gRPC cfgGRPC
}

type cfgGRPC struct {
	serviceClient netmap.NetmapServiceClient

	grpcCallOpts []grpc.CallOption

	callOpts []netmap.Option

	client *netmap.Client
}

type localNodeInfoClient struct {
	requestConverter func(*LocalNodeInfoRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *LocalNodeInfoResponse
}

type networkInfoClient struct {
	requestConverter func(*NetworkInfoRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *NetworkInfoResponse
}

// LocalNodeInfo sends LocalNodeInfoRequest over the network.
func (c *Client) LocalNodeInfo(ctx context.Context, req *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error) {
	resp, err := c.cLocalNodeInfo.caller(ctx, c.cLocalNodeInfo.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send local node info request")
	}

	return c.cLocalNodeInfo.responseConverter(resp), nil
}

// NetworkInfo sends NetworkInfoRequest over the network.
func (c *Client) NetworkInfo(ctx context.Context, req *NetworkInfoRequest) (*NetworkInfoResponse, error) {
	resp, err := c.cNetworkInfo.caller(ctx, c.cNetworkInfo.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send network info request")
	}

	return c.cNetworkInfo.responseConverter(resp), nil
}

func defaultCfg() *cfg {
	return &cfg{
		proto: client.ProtoGRPC,
	}
}

// NewClient is a constructor for netmap transport client.
func NewClient(opts ...Option) (*Client, error) {
	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	var err error

	switch cfg.proto {
	case client.ProtoGRPC:
		var c *netmap.Client
		if c, err = newGRPCClient(cfg); err != nil {
			break
		}

		return &Client{
			cLocalNodeInfo: &localNodeInfoClient{
				requestConverter: func(req *LocalNodeInfoRequest) interface{} {
					return LocalNodeInfoRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.LocalNodeInfo(ctx, req.(*netmap.LocalNodeInfoRequest))
				},
				responseConverter: func(resp interface{}) *LocalNodeInfoResponse {
					return LocalNodeInfoResponseFromGRPCMessage(resp.(*netmap.LocalNodeInfoResponse))
				},
			},
			cNetworkInfo: &networkInfoClient{
				requestConverter: func(req *NetworkInfoRequest) interface{} {
					return NetworkInfoRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.NetworkInfo(ctx, req.(*netmap.NetworkInfoRequest))
				},
				responseConverter: func(resp interface{}) *NetworkInfoResponse {
					return NetworkInfoResponseFromGRPCMessage(resp.(*netmap.NetworkInfoResponse))
				},
			},
		}, nil
	default:
		err = client.ErrProtoUnsupported
	}

	return nil, errors.Wrapf(err, "could not create %s Netmap client", cfg.proto)
}

func newGRPCClient(cfg *cfg) (*netmap.Client, error) {
	var err error

	if cfg.gRPC.client == nil {
		if cfg.gRPC.serviceClient == nil {
			conn, err := client.NewGRPCClientConn(cfg.globalOpts...)
			if err != nil {
				return nil, errors.Wrap(err, "could not open gRPC client connection")
			}

			cfg.gRPC.serviceClient = netmap.NewNetmapServiceClient(conn)
		}

		cfg.gRPC.client, err = netmap.NewClient(
			cfg.gRPC.serviceClient,
			append(
				cfg.gRPC.callOpts,
				netmap.WithCallOptions(cfg.gRPC.grpcCallOpts),
			)...,
		)
	}

	return cfg.gRPC.client, err
}

// WithGlobalOpts sets global client options to client.
func WithGlobalOpts(v ...client.Option) Option {
	return func(c *cfg) {
		if len(v) > 0 {
			c.globalOpts = v
		}
	}
}

// WithGRPCServiceClient sets existing service client.
func WithGRPCServiceClient(v netmap.NetmapServiceClient) Option {
	return func(c *cfg) {
		c.gRPC.serviceClient = v
	}
}

// WithGRPCServiceClient sets GRPC specific call options.
func WithGRPCCallOpts(v []grpc.CallOption) Option {
	return func(c *cfg) {
		c.gRPC.grpcCallOpts = v
	}
}

// WithGRPCServiceClient sets GRPC specific client options.
func WithGRPCClientOpts(v []netmap.Option) Option {
	return func(c *cfg) {
		c.gRPC.callOpts = v
	}
}

// WithGRPCServiceClient sets existing GRPC client.
func WithGRPCClient(v *netmap.Client) Option {
	return func(c *cfg) {
		c.gRPC.client = v
	}
}
