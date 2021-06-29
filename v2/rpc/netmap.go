package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

const serviceNetmap = serviceNamePrefix + "netmap.NetmapService"

const (
	rpcNetmapNodeInfo = "LocalNodeInfo"
	rpcNetmapNetInfo  = "NetworkInfo"
)

// LocalNodeInfoPrm groups the parameters of LocalNodeInfo call.
type LocalNodeInfoPrm struct {
	req netmap.LocalNodeInfoRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *LocalNodeInfoPrm) SetRequest(req netmap.LocalNodeInfoRequest) {
	x.req = req
}

// LocalNodeInfoRes groups the results of LocalNodeInfo call.
type LocalNodeInfoRes struct {
	resp netmap.LocalNodeInfoResponse
}

// Response returns the server response.
func (x *LocalNodeInfoRes) Response() netmap.LocalNodeInfoResponse {
	return x.resp
}

// LocalNodeInfo executes NetmapService.LocalNodeInfo RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func LocalNodeInfo(ctx context.Context, cli protoclient.Client, prm LocalNodeInfoPrm, res *LocalNodeInfoRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceNetmap, rpcNetmapNodeInfo)
}

// NetworkInfoPrm groups the parameters of NetworkInfo call.
type NetworkInfoPrm struct {
	req netmap.NetworkInfoRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *NetworkInfoPrm) SetRequest(req netmap.NetworkInfoRequest) {
	x.req = req
}

// NetworkInfoRes groups the results of NetworkInfo call.
type NetworkInfoRes struct {
	resp netmap.NetworkInfoResponse
}

// Response returns the server response.
func (x *NetworkInfoRes) Response() netmap.NetworkInfoResponse {
	return x.resp
}

// NetworkInfo executes NetmapService.NetworkInfo RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func NetworkInfo(ctx context.Context, cli protoclient.Client, prm NetworkInfoPrm, res *NetworkInfoRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceNetmap, rpcNetmapNetInfo)
}
