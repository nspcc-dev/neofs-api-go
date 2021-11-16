package rpc

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/common"
)

const serviceNetmap = serviceNamePrefix + "netmap.NetmapService"

const (
	rpcNetmapNodeInfo = "LocalNodeInfo"
	rpcNetmapNetInfo  = "NetworkInfo"
)

// LocalNodeInfo executes NetmapService.LocalNodeInfo RPC.
func LocalNodeInfo(
	cli *client.Client,
	req *netmap.LocalNodeInfoRequest,
	opts ...client.CallOption,
) (*netmap.LocalNodeInfoResponse, error) {
	resp := new(netmap.LocalNodeInfoResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceNetmap, rpcNetmapNodeInfo), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// NetworkInfo executes NetmapService.NetworkInfo RPC.
func NetworkInfo(
	cli *client.Client,
	req *netmap.NetworkInfoRequest,
	opts ...client.CallOption,
) (*netmap.NetworkInfoResponse, error) {
	resp := new(netmap.NetworkInfoResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceNetmap, rpcNetmapNetInfo), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
