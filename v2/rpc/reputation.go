package rpc

import (
	"context"

	protoclient "github.com/nspcc-dev/neofs-api-go/rpc/client"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
)

const serviceReputation = serviceNamePrefix + "reputation.ReputationService"

const (
	rpcReputationAnnounceLocalTrust         = "AnnounceLocalTrust"
	rpcReputationAnnounceIntermediateResult = "AnnounceIntermediateResult"
)

// AnnounceLocalTrustPrm groups the parameters of AnnounceLocalTrust call.
type AnnounceLocalTrustPrm struct {
	req reputation.AnnounceLocalTrustRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *AnnounceLocalTrustPrm) SetRequest(req reputation.AnnounceLocalTrustRequest) {
	x.req = req
}

// AnnounceLocalTrustRes groups the results of AnnounceLocalTrust call.
type AnnounceLocalTrustRes struct {
	resp reputation.AnnounceLocalTrustResponse
}

// Response returns the server response.
func (x *AnnounceLocalTrustRes) Response() reputation.AnnounceLocalTrustResponse {
	return x.resp
}

// AnnounceLocalTrust executes ReputationService.NetworkInfo RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func AnnounceLocalTrust(ctx context.Context, cli protoclient.Client, prm AnnounceLocalTrustPrm, res *AnnounceLocalTrustRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceReputation, rpcReputationAnnounceLocalTrust)
}

// AnnounceIntermediateResultPrm groups the parameters of AnnounceIntermediateResult call.
type AnnounceIntermediateResultPrm struct {
	req reputation.AnnounceIntermediateResultRequest
}

// SetRequest sets request to send to the server.
//
// Required parameter.
func (x *AnnounceIntermediateResultPrm) SetRequest(req reputation.AnnounceIntermediateResultRequest) {
	x.req = req
}

// AnnounceIntermediateResultRes groups the results of AnnounceIntermediateResult call.
type AnnounceIntermediateResultRes struct {
	resp reputation.AnnounceIntermediateResultResponse
}

// Response returns the server response.
func (x *AnnounceIntermediateResultRes) Response() reputation.AnnounceIntermediateResultResponse {
	return x.resp
}

// AnnounceIntermediateResult executes ReputationService.AnnounceIntermediateResult RPC.
//
// Client connection should be established.
//
// If there is no error, res.Response() contains the server response.
//
// Context and res must not be nil.
func AnnounceIntermediateResult(ctx context.Context, cli protoclient.Client, prm AnnounceIntermediateResultPrm, res *AnnounceIntermediateResultRes) error {
	return sendUnaryRPC(ctx, cli, &prm.req, &res.resp, serviceReputation, rpcReputationAnnounceIntermediateResult)
}
