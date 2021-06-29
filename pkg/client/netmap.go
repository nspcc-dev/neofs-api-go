package client

import (
	"context"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
	v2netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

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
func (x Client) EndpointInfo(ctx context.Context, opts ...CallOption) (*EndpointInfo, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2netmap.LocalNodeInfoRequestBody)

	var req v2netmap.LocalNodeInfoRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.LocalNodeInfoPrm

	prm.SetRequest(req)

	var res rpcapi.LocalNodeInfoRes

	err = rpcapi.LocalNodeInfo(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	body := resp.GetBody()

	return &EndpointInfo{
		version: pkg.NewVersionFromV2(body.GetVersion()),
		ni:      netmap.NewNodeInfoFromV2(body.GetNodeInfo()),
	}, nil
}

// NetworkInfo returns information about the NeoFS network of which the remote server is a part.
func (x Client) NetworkInfo(ctx context.Context, opts ...CallOption) (*netmap.NetworkInfo, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2netmap.NetworkInfoRequestBody)

	var req v2netmap.NetworkInfoRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var prm rpcapi.NetworkInfoPrm

	prm.SetRequest(req)

	var res rpcapi.NetworkInfoRes

	err = rpcapi.NetworkInfo(ctx, x.c, prm, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("response message verification failed: %w", err)
	}

	return netmap.NewNetworkInfoFromV2(resp.GetBody().GetNetworkInfo()), nil
}
