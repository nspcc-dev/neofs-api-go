package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	v2netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

// Netmap contains methods related to netmap.
type Netmap interface {
	// EndpointInfo returns attributes, address and public key of the node, specified
	// in client constructor via address or open connection. This can be used as a
	// health check to see if node is alive and responses to requests.
	EndpointInfo(context.Context, ...CallOption) (*EndpointInfo, error)

	// NetworkInfo returns information about the NeoFS network of which the remote server is a part.
	NetworkInfo(context.Context, ...CallOption) (*netmap.NetworkInfo, error)
}

// EACLWithSignature represents eACL table/signature pair.
type EndpointInfo struct {
	version *pkg.Version

	ni *netmap.NodeInfo
}

// LatestVersion returns latest NeoFS API version in use.
func (e *EndpointInfo) LatestVersion() *pkg.Version {
	return e.version
}

// NodeInfo returns returns information about the NeoFS node.
func (e *EndpointInfo) NodeInfo() *netmap.NodeInfo {
	return e.ni
}

// EndpointInfo returns attributes, address and public key of the node, specified
// in client constructor via address or open connection. This can be used as a
// health check to see if node is alive and responses to requests.
func (c *clientImpl) EndpointInfo(ctx context.Context, opts ...CallOption) (*EndpointInfo, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2netmap.LocalNodeInfoRequestBody)

	req := new(v2netmap.LocalNodeInfoRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.LocalNodeInfo(c.Raw(), req)
	if err != nil {
		return nil, errors.Wrap(err, "transport error")
	}

	err = v2signature.VerifyServiceMessage(resp, callOptions.signOpts()...)
	if err != nil {
		return nil, errors.Wrap(err, "can't verify response message")
	}

	body := resp.GetBody()

	return &EndpointInfo{
		version: pkg.NewVersionFromV2(body.GetVersion()),
		ni:      netmap.NewNodeInfoFromV2(body.GetNodeInfo()),
	}, nil
}

// NetworkInfo returns information about the NeoFS network of which the remote server is a part.
func (c *clientImpl) NetworkInfo(ctx context.Context, opts ...CallOption) (*netmap.NetworkInfo, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2netmap.NetworkInfoRequestBody)

	req := new(v2netmap.NetworkInfoRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.NetworkInfo(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, errors.Wrap(err, "v2 NetworkInfo RPC failure")
	}

	err = v2signature.VerifyServiceMessage(resp, callOptions.signOpts()...)
	if err != nil {
		return nil, errors.Wrap(err, "response message verification failed")
	}

	return netmap.NewNetworkInfoFromV2(resp.GetBody().GetNetworkInfo()), nil
}
