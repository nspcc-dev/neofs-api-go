package container

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/client"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client represents universal container
// transport client.
type Client struct {
	cPut *putClient

	cGet *getClient

	cDel *delClient

	cList *listClient

	cSetEACL *setEACLClient

	cGetEACL *getEACLClient

	cAnnounce *announceUsedSpaceClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	proto client.Protocol

	globalOpts []client.Option

	gRPC cfgGRPC
}

type cfgGRPC struct {
	serviceClient container.ContainerServiceClient

	grpcCallOpts []grpc.CallOption

	callOpts []container.Option

	client *container.Client
}

type putClient struct {
	requestConverter func(*PutRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *PutResponse
}

type getClient struct {
	requestConverter func(*GetRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *GetResponse
}

type delClient struct {
	requestConverter func(*DeleteRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *DeleteResponse
}

type listClient struct {
	requestConverter func(*ListRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *ListResponse
}

type setEACLClient struct {
	requestConverter func(*SetExtendedACLRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *SetExtendedACLResponse
}

type getEACLClient struct {
	requestConverter func(*GetExtendedACLRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *GetExtendedACLResponse
}

type announceUsedSpaceClient struct {
	requestConverter func(request *AnnounceUsedSpaceRequest) interface{}

	caller func(context.Context, interface{}) (interface{}, error)

	responseConverter func(interface{}) *AnnounceUsedSpaceResponse
}

// Put sends PutRequest over the network and returns PutResponse.
//
// It returns any error encountered during the call.
func (c *Client) Put(ctx context.Context, req *PutRequest) (*PutResponse, error) {
	resp, err := c.cPut.caller(ctx, c.cPut.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container put request")
	}

	return c.cPut.responseConverter(resp), nil
}

// Get sends GetRequest over the network and returns GetResponse.
//
// It returns any error encountered during the call.
func (c *Client) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	resp, err := c.cGet.caller(ctx, c.cGet.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container get request")
	}

	return c.cGet.responseConverter(resp), nil
}

// Delete sends GetRequest over the network and returns GetResponse.
//
// It returns any error encountered during the call.
func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	resp, err := c.cDel.caller(ctx, c.cDel.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container delete request")
	}

	return c.cDel.responseConverter(resp), nil
}

// List sends ListRequest over the network and returns ListResponse.
//
// It returns any error encountered during the call.
func (c *Client) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	resp, err := c.cList.caller(ctx, c.cList.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container list request")
	}

	return c.cList.responseConverter(resp), nil
}

// SetExtendedACL sends SetExtendedACLRequest over the network and returns SetExtendedACLResponse.
//
// It returns any error encountered during the call.
func (c *Client) SetExtendedACL(ctx context.Context, req *SetExtendedACLRequest) (*SetExtendedACLResponse, error) {
	resp, err := c.cSetEACL.caller(ctx, c.cSetEACL.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container set EACL request")
	}

	return c.cSetEACL.responseConverter(resp), nil
}

// GetExtendedACL sends GetExtendedACLRequest over the network and returns GetExtendedACLResponse.
//
// It returns any error encountered during the call.
func (c *Client) GetExtendedACL(ctx context.Context, req *GetExtendedACLRequest) (*GetExtendedACLResponse, error) {
	resp, err := c.cGetEACL.caller(ctx, c.cGetEACL.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send container get EACL request")
	}

	return c.cGetEACL.responseConverter(resp), nil
}

// AnnounceUsedSpace sends AnnounceUsedSpaceRequest over the network and returns
// AnnounceUsedSpaceResponse.
//
// It returns any error encountered during the call.
func (c *Client) AnnounceUsedSpace(ctx context.Context, req *AnnounceUsedSpaceRequest) (*AnnounceUsedSpaceResponse, error) {
	resp, err := c.cAnnounce.caller(ctx, c.cAnnounce.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send announce used space request")
	}

	return c.cAnnounce.responseConverter(resp), nil
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
		var c *container.Client
		if c, err = newGRPCClient(cfg); err != nil {
			break
		}

		return &Client{
			cPut: &putClient{
				requestConverter: func(req *PutRequest) interface{} {
					return PutRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.Put(ctx, req.(*container.PutRequest))
				},
				responseConverter: func(resp interface{}) *PutResponse {
					return PutResponseFromGRPCMessage(resp.(*container.PutResponse))
				},
			},
			cGet: &getClient{
				requestConverter: func(req *GetRequest) interface{} {
					return GetRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.Get(ctx, req.(*container.GetRequest))
				},
				responseConverter: func(resp interface{}) *GetResponse {
					return GetResponseFromGRPCMessage(resp.(*container.GetResponse))
				},
			},
			cDel: &delClient{
				requestConverter: func(req *DeleteRequest) interface{} {
					return DeleteRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.Delete(ctx, req.(*container.DeleteRequest))
				},
				responseConverter: func(resp interface{}) *DeleteResponse {
					return DeleteResponseFromGRPCMessage(resp.(*container.DeleteResponse))
				},
			},
			cList: &listClient{
				requestConverter: func(req *ListRequest) interface{} {
					return ListRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.List(ctx, req.(*container.ListRequest))
				},
				responseConverter: func(resp interface{}) *ListResponse {
					return ListResponseFromGRPCMessage(resp.(*container.ListResponse))
				},
			},
			cSetEACL: &setEACLClient{
				requestConverter: func(req *SetExtendedACLRequest) interface{} {
					return SetExtendedACLRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.SetExtendedACL(ctx, req.(*container.SetExtendedACLRequest))
				},
				responseConverter: func(resp interface{}) *SetExtendedACLResponse {
					return SetExtendedACLResponseFromGRPCMessage(resp.(*container.SetExtendedACLResponse))
				},
			},
			cGetEACL: &getEACLClient{
				requestConverter: func(req *GetExtendedACLRequest) interface{} {
					return GetExtendedACLRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.GetExtendedACL(ctx, req.(*container.GetExtendedACLRequest))
				},
				responseConverter: func(resp interface{}) *GetExtendedACLResponse {
					return GetExtendedACLResponseFromGRPCMessage(resp.(*container.GetExtendedACLResponse))
				},
			},
			cAnnounce: &announceUsedSpaceClient{
				requestConverter: func(req *AnnounceUsedSpaceRequest) interface{} {
					return AnnounceUsedSpaceRequestToGRPCMessage(req)
				},
				caller: func(ctx context.Context, req interface{}) (interface{}, error) {
					return c.AnnounceUsedSpace(ctx, req.(*container.AnnounceUsedSpaceRequest))
				},
				responseConverter: func(resp interface{}) *AnnounceUsedSpaceResponse {
					return AnnounceUsedSpaceResponseFromGRPCMessage(resp.(*container.AnnounceUsedSpaceResponse))
				},
			},
		}, nil
	default:
		err = client.ErrProtoUnsupported
	}

	return nil, errors.Wrapf(err, "could not create %s Session client", cfg.proto)
}

func newGRPCClient(cfg *cfg) (*container.Client, error) {
	var err error

	if cfg.gRPC.client == nil {
		if cfg.gRPC.serviceClient == nil {
			conn, err := client.NewGRPCClientConn(cfg.globalOpts...)
			if err != nil {
				return nil, errors.Wrap(err, "could not open gRPC client connection")
			}

			cfg.gRPC.serviceClient = container.NewContainerServiceClient(conn)
		}

		cfg.gRPC.client, err = container.NewClient(
			cfg.gRPC.serviceClient,
			append(
				cfg.gRPC.callOpts,
				container.WithCallOptions(cfg.gRPC.grpcCallOpts),
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

func WithGRPCServiceClient(v container.ContainerServiceClient) Option {
	return func(c *cfg) {
		c.gRPC.serviceClient = v
	}
}

func WithGRPCCallOpts(v []grpc.CallOption) Option {
	return func(c *cfg) {
		c.gRPC.grpcCallOpts = v
	}
}

func WithGRPCClientOpts(v []container.Option) Option {
	return func(c *cfg) {
		c.gRPC.callOpts = v
	}
}

func WithGRPCClient(v *container.Client) Option {
	return func(c *cfg) {
		c.gRPC.client = v
	}
}
