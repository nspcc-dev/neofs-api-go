package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/client"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

// EndpointInfo returns attributes, address and public key of the node, specified
// in client constructor via address or open connection. This can be used as a
// health check to see if node is alive and responses to requests.
func (c Client) EndpointInfo(ctx context.Context, opts ...CallOption) (*netmap.NodeInfo, error) {
	switch c.remoteNode.Version.GetMajor() {
	case 2:
		return c.endpointInfoV2(ctx, opts...)
	default:
		return nil, unsupportedProtocolErr
	}
}

func (c Client) endpointInfoV2(ctx context.Context, opts ...CallOption) (*netmap.NodeInfo, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()
	for i := range opts {
		opts[i].apply(&callOptions)
	}

	reqBody := new(netmap.LocalNodeInfoRequestBody)

	req := new(netmap.LocalNodeInfoRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(c.key, req)
	if err != nil {
		return nil, err
	}

	switch c.remoteNode.Protocol {
	case GRPC:
		cli, err := v2NetmapClientFromOptions(c.opts)
		if err != nil {
			return nil, errors.Wrap(err, "can't create grpc client")
		}

		resp, err := cli.LocalNodeInfo(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "transport error")
		}

		err = v2signature.VerifyServiceMessage(resp)
		if err != nil {
			return nil, errors.Wrap(err, "can't verify response message")
		}

		return resp.GetBody().GetNodeInfo(), nil
	default:
		return nil, unsupportedProtocolErr
	}
}

func v2NetmapClientFromOptions(opts *clientOptions) (cli *netmap.Client, err error) {
	switch {
	case opts.grpcOpts.v2NetmapClient != nil:
		// return value from client cache
		return opts.grpcOpts.v2NetmapClient, nil

	case opts.grpcOpts.conn != nil:
		cli, err = netmap.NewClient(netmap.WithGlobalOpts(
			client.WithGRPCConn(opts.grpcOpts.conn)),
		)

	case opts.addr != "":
		cli, err = netmap.NewClient(netmap.WithGlobalOpts(
			client.WithNetworkAddress(opts.addr)),
		)

	default:
		return nil, errors.New("lack of sdk client options to create accounting client")
	}

	// check if client correct and save in cache
	if err != nil {
		return nil, err
	}

	opts.grpcOpts.v2NetmapClient = cli

	return cli, nil
}
