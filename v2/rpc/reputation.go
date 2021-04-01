package rpc

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/rpc/common"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
)

const serviceReputation = serviceNamePrefix + "reputation.ReputationService"

const (
	rpcReputationSendLocalTrust         = "SendLocalTrust"
	rpcReputationSendIntermediateResult = "SendIntermediateResult"
)

// SendLocalTrust executes ReputationService.SendLocalTrust RPC.
func SendLocalTrust(
	cli *client.Client,
	req *reputation.SendLocalTrustRequest,
	opts ...client.CallOption,
) (*reputation.SendLocalTrustResponse, error) {
	resp := new(reputation.SendLocalTrustResponse)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceReputation, rpcReputationSendLocalTrust), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SendIntermediateResult executes ReputationService.SendIntermediateResult RPC.
func SendIntermediateResult(
	cli *client.Client,
	req *reputation.SendIntermediateResultRequest,
	opts ...client.CallOption,
) (*reputation.SendIntermediateResultRequest, error) {
	resp := new(reputation.SendIntermediateResultRequest)

	err := client.SendUnary(cli, common.CallMethodInfoUnary(serviceReputation, rpcReputationSendIntermediateResult), req, resp, opts...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
