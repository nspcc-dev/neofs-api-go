package object

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/client"
	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client represents universal object
// transport client.
type Client struct {
	getClient *getObjectClient

	putClient *putObjectClient

	headClient *headObjectClient

	searchClient *searchObjectClient

	deleteClient *deleteObjectClient

	getRangeClient *getRangeObjectClient

	getRangeHashClient *getRangeHashObjectClient
}

// Option represents Client option.
type Option func(*cfg)

type cfg struct {
	proto client.Protocol

	globalOpts []client.Option

	gRPC cfgGRPC
}

type cfgGRPC struct {
	serviceClient object.ObjectServiceClient

	grpcCallOpts []grpc.CallOption

	callOpts []object.Option

	client *object.Client
}

// types of upper level sub-clients, accessed directly from object.Client
type (
	getObjectClient struct {
		streamClientGetter func(context.Context, *GetRequest) (interface{}, error)

		streamerConstructor func(interface{}) (GetObjectStreamer, error)
	}

	putObjectClient struct {
		streamClientGetter func(context.Context) (interface{}, error)

		streamerConstructor func(interface{}) (PutObjectStreamer, error)
	}

	headObjectClient struct {
		requestConverter func(request *HeadRequest) interface{}

		caller func(context.Context, interface{}) (interface{}, error)

		responseConverter func(interface{}) *HeadResponse
	}

	deleteObjectClient struct {
		requestConverter func(request *DeleteRequest) interface{}

		caller func(context.Context, interface{}) (interface{}, error)

		responseConverter func(interface{}) *DeleteResponse
	}

	searchObjectClient struct {
		streamClientGetter func(context.Context, *SearchRequest) (interface{}, error)

		streamerConstructor func(interface{}) (SearchObjectStreamer, error)
	}

	getRangeObjectClient struct {
		streamClientGetter func(context.Context, *GetRangeRequest) (interface{}, error)

		streamerConstructor func(interface{}) (GetRangeObjectStreamer, error)
	}

	getRangeHashObjectClient struct {
		requestConverter func(request *GetRangeHashRequest) interface{}

		caller func(context.Context, interface{}) (interface{}, error)

		responseConverter func(interface{}) *GetRangeHashResponse
	}
)

func (c *Client) Get(ctx context.Context, req *GetRequest) (GetObjectStreamer, error) {
	cli, err := c.getClient.streamClientGetter(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "could not send get object request")
	}

	return c.getClient.streamerConstructor(cli)
}

func (c *Client) Put(ctx context.Context) (PutObjectStreamer, error) {
	cli, err := c.putClient.streamClientGetter(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not prepare put object streamer")
	}

	return c.putClient.streamerConstructor(cli)
}

func (c *Client) Head(ctx context.Context, req *HeadRequest) (*HeadResponse, error) {
	resp, err := c.headClient.caller(ctx, c.headClient.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send head object request")
	}

	return c.headClient.responseConverter(resp), nil
}

func (c *Client) Search(ctx context.Context, req *SearchRequest) (SearchObjectStreamer, error) {
	cli, err := c.searchClient.streamClientGetter(ctx, req)
	if err != nil {
		return nil, err
	}

	return c.searchClient.streamerConstructor(cli)
}

func (c *Client) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	resp, err := c.deleteClient.caller(ctx, c.deleteClient.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send delete object request")
	}

	return c.deleteClient.responseConverter(resp), nil
}

func (c *Client) GetRange(ctx context.Context, req *GetRangeRequest) (GetRangeObjectStreamer, error) {
	cli, err := c.getRangeClient.streamClientGetter(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "could not send get object range request")
	}

	return c.getRangeClient.streamerConstructor(cli)
}

func (c *Client) GetRangeHash(ctx context.Context, req *GetRangeHashRequest) (*GetRangeHashResponse, error) {
	resp, err := c.getRangeHashClient.caller(ctx, c.getRangeHashClient.requestConverter(req))
	if err != nil {
		return nil, errors.Wrap(err, "could not send get object range hash request")
	}

	return c.getRangeHashClient.responseConverter(resp), nil
}

func defaultCfg() *cfg {
	return &cfg{
		proto: client.ProtoGRPC,
	}
}

func New(opts ...Option) (*Client, error) {
	cfg := defaultCfg()

	for i := range opts {
		opts[i](cfg)
	}

	var err error

	switch cfg.proto {
	case client.ProtoGRPC:
		var c *object.Client
		if c, err = newGRPCClient(cfg); err != nil {
			break
		}

		return &Client{
			getClient:          newGRPCGetClient(c),
			putClient:          newGRPCPutClient(c),
			headClient:         newGRPCHeadClient(c),
			searchClient:       newGRPCSearchClient(c),
			deleteClient:       newGRPCDeleteClient(c),
			getRangeClient:     newGRPCGetRangeClient(c),
			getRangeHashClient: newGRPCGetRangeHashClient(c),
		}, nil
	default:
		err = client.ErrProtoUnsupported
	}

	return nil, errors.Wrapf(err, "could not create %s object client", cfg.proto)
}

func newGRPCClient(cfg *cfg) (*object.Client, error) {
	var err error

	if cfg.gRPC.client == nil {
		if cfg.gRPC.serviceClient == nil {
			conn, err := client.NewGRPCClientConn(cfg.globalOpts...)
			if err != nil {
				return nil, errors.Wrap(err, "could not open gRPC getClient connection")
			}

			cfg.gRPC.serviceClient = object.NewObjectServiceClient(conn)
		}

		cfg.gRPC.client, err = object.NewClient(
			cfg.gRPC.serviceClient,
			append(
				cfg.gRPC.callOpts,
				object.WithCallOptions(cfg.gRPC.grpcCallOpts),
			)...,
		)
	}

	return cfg.gRPC.client, err
}

func newGRPCGetClient(c *object.Client) *getObjectClient {
	cli := &getObjectClient{
		streamClientGetter: func(ctx context.Context, request *GetRequest) (interface{}, error) {
			return c.Get(ctx, GetRequestToGRPCMessage(request))
		},
		streamerConstructor: func(i interface{}) (GetObjectStreamer, error) {
			cli, ok := i.(object.ObjectService_GetClient)
			if !ok {
				return nil, errors.New("can't convert interface to grpc get getClient")
			}
			return &getObjectStream{
				recv: func() (*GetResponse, error) {
					resp, err := cli.Recv()
					if err != nil {
						return nil, err
					}

					return GetResponseFromGRPCMessage(resp), nil
				},
			}, nil
		},
	}

	return cli
}

func newGRPCPutClient(c *object.Client) *putObjectClient {
	cli := &putObjectClient{
		streamClientGetter: func(ctx context.Context) (interface{}, error) {
			return c.Put(ctx)
		},
		streamerConstructor: func(i interface{}) (PutObjectStreamer, error) {
			cli, ok := i.(object.ObjectService_PutClient)
			if !ok {
				return nil, errors.New("can't convert interface to grpc get getClient")
			}

			return &putObjectStream{
				send: func(request *PutRequest) error {
					return cli.Send(PutRequestToGRPCMessage(request))
				},
				closeAndRecv: func() (*PutResponse, error) {
					resp, err := cli.CloseAndRecv()
					if err != nil {
						return nil, err
					}

					return PutResponseFromGRPCMessage(resp), nil
				},
			}, nil
		},
	}

	return cli
}

func newGRPCHeadClient(c *object.Client) *headObjectClient {
	return &headObjectClient{
		requestConverter: func(req *HeadRequest) interface{} {
			return HeadRequestToGRPCMessage(req)
		},
		caller: func(ctx context.Context, req interface{}) (interface{}, error) {
			return c.Head(ctx, req.(*object.HeadRequest))
		},
		responseConverter: func(resp interface{}) *HeadResponse {
			return HeadResponseFromGRPCMessage(resp.(*object.HeadResponse))
		},
	}
}

func newGRPCSearchClient(c *object.Client) *searchObjectClient {
	cli := &searchObjectClient{
		streamClientGetter: func(ctx context.Context, request *SearchRequest) (interface{}, error) {
			return c.Search(ctx, SearchRequestToGRPCMessage(request))
		},
		streamerConstructor: func(i interface{}) (SearchObjectStreamer, error) {
			cli, ok := i.(object.ObjectService_SearchClient)
			if !ok {
				return nil, errors.New("can't convert interface to grpc get getClient")
			}
			return &searchObjectStream{
				recv: func() (*SearchResponse, error) {
					resp, err := cli.Recv()
					if err != nil {
						return nil, err
					}

					return SearchResponseFromGRPCMessage(resp), nil
				},
			}, nil
		},
	}

	return cli
}

func newGRPCDeleteClient(c *object.Client) *deleteObjectClient {
	return &deleteObjectClient{
		requestConverter: func(req *DeleteRequest) interface{} {
			return DeleteRequestToGRPCMessage(req)
		},
		caller: func(ctx context.Context, req interface{}) (interface{}, error) {
			return c.Delete(ctx, req.(*object.DeleteRequest))
		},
		responseConverter: func(resp interface{}) *DeleteResponse {
			return DeleteResponseFromGRPCMessage(resp.(*object.DeleteResponse))
		},
	}
}

func newGRPCGetRangeClient(c *object.Client) *getRangeObjectClient {
	cli := &getRangeObjectClient{
		streamClientGetter: func(ctx context.Context, request *GetRangeRequest) (interface{}, error) {
			return c.GetRange(ctx, GetRangeRequestToGRPCMessage(request))
		},
		streamerConstructor: func(i interface{}) (GetRangeObjectStreamer, error) {
			cli, ok := i.(object.ObjectService_GetRangeClient)
			if !ok {
				return nil, errors.New("can't convert interface to grpc get getClient")
			}
			return &getRangeObjectStream{
				recv: func() (*GetRangeResponse, error) {
					resp, err := cli.Recv()
					if err != nil {
						return nil, err
					}

					return GetRangeResponseFromGRPCMessage(resp), nil
				},
			}, nil
		},
	}

	return cli
}

func newGRPCGetRangeHashClient(c *object.Client) *getRangeHashObjectClient {
	return &getRangeHashObjectClient{
		requestConverter: func(req *GetRangeHashRequest) interface{} {
			return GetRangeHashRequestToGRPCMessage(req)
		},
		caller: func(ctx context.Context, req interface{}) (interface{}, error) {
			return c.GetRangeHash(ctx, req.(*object.GetRangeHashRequest))
		},
		responseConverter: func(resp interface{}) *GetRangeHashResponse {
			return GetRangeHashResponseFromGRPCMessage(resp.(*object.GetRangeHashResponse))
		},
	}
}

func WithGlobalOpts(v ...client.Option) Option {
	return func(c *cfg) {
		if len(v) > 0 {
			c.globalOpts = v
		}
	}
}

func WithGRPCServiceClient(v object.ObjectServiceClient) Option {
	return func(c *cfg) {
		c.gRPC.serviceClient = v
	}
}

func WithGRPCCallOpts(v []grpc.CallOption) Option {
	return func(c *cfg) {
		c.gRPC.grpcCallOpts = v
	}
}

func WithGRPCClientOpts(v []object.Option) Option {
	return func(c *cfg) {
		c.gRPC.callOpts = v
	}
}

func WithGRPCClient(v *object.Client) Option {
	return func(c *cfg) {
		c.gRPC.client = v
	}
}
