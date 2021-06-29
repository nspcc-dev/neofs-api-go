package client

import (
	"context"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	v2reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// AnnounceLocalTrustPrm groups parameters of AnnounceLocalTrust operation.
type AnnounceLocalTrustPrm struct {
	epoch uint64

	trusts []*reputation.Trust
}

// Epoch returns epoch in which the trust was assessed.
func (x AnnounceLocalTrustPrm) Epoch() uint64 {
	return x.epoch
}

// SetEpoch sets epoch in which the trust was assessed.
func (x *AnnounceLocalTrustPrm) SetEpoch(epoch uint64) {
	x.epoch = epoch
}

// Trusts returns list of local trust values.
func (x AnnounceLocalTrustPrm) Trusts() []*reputation.Trust {
	return x.trusts
}

// SetTrusts sets list of local trust values.
func (x *AnnounceLocalTrustPrm) SetTrusts(trusts []*reputation.Trust) {
	x.trusts = trusts
}

// AnnounceLocalTrustRes groups results of AnnounceLocalTrust operation.
type AnnounceLocalTrustRes struct{}

func (x Client) AnnounceLocalTrust(ctx context.Context, prm AnnounceLocalTrustPrm, opts ...CallOption) (*AnnounceLocalTrustRes, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2reputation.AnnounceLocalTrustRequestBody)
	reqBody.SetEpoch(prm.Epoch())
	reqBody.SetTrusts(reputation.TrustsToV2(prm.Trusts()))

	var req v2reputation.AnnounceLocalTrustRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var p rpcapi.AnnounceLocalTrustPrm

	p.SetRequest(req)

	var res rpcapi.AnnounceLocalTrustRes

	err = rpcapi.AnnounceLocalTrust(ctx, x.c, p, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return new(AnnounceLocalTrustRes), nil
}

// AnnounceIntermediateTrustPrm groups parameters of AnnounceIntermediateTrust operation.
type AnnounceIntermediateTrustPrm struct {
	epoch uint64

	iter uint32

	trust *reputation.PeerToPeerTrust
}

func (x *AnnounceIntermediateTrustPrm) Epoch() uint64 {
	return x.epoch
}

func (x *AnnounceIntermediateTrustPrm) SetEpoch(epoch uint64) {
	x.epoch = epoch
}

// Iteration returns sequence number of the iteration.
func (x AnnounceIntermediateTrustPrm) Iteration() uint32 {
	return x.iter
}

// SetIteration sets sequence number of the iteration.
func (x *AnnounceIntermediateTrustPrm) SetIteration(iter uint32) {
	x.iter = iter
}

// Trust returns current global trust value computed at the specified iteration.
func (x AnnounceIntermediateTrustPrm) Trust() *reputation.PeerToPeerTrust {
	return x.trust
}

// SetTrust sets current global trust value computed at the specified iteration.
func (x *AnnounceIntermediateTrustPrm) SetTrust(trust *reputation.PeerToPeerTrust) {
	x.trust = trust
}

// AnnounceIntermediateTrustRes groups results of AnnounceIntermediateTrust operation.
type AnnounceIntermediateTrustRes struct{}

func (x Client) AnnounceIntermediateTrust(ctx context.Context, prm AnnounceIntermediateTrustPrm, opts ...CallOption) (*AnnounceIntermediateTrustRes, error) {
	// apply all available options
	callOptions := defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2reputation.AnnounceIntermediateResultRequestBody)
	reqBody.SetEpoch(prm.Epoch())
	reqBody.SetIteration(prm.Iteration())
	reqBody.SetTrust(prm.Trust().ToV2())

	var req v2reputation.AnnounceIntermediateResultRequest
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), &req)
	if err != nil {
		return nil, err
	}

	var p rpcapi.AnnounceIntermediateResultPrm

	p.SetRequest(req)

	var res rpcapi.AnnounceIntermediateResultRes

	err = rpcapi.AnnounceIntermediateResult(ctx, x.c, p, &res)
	if err != nil {
		return nil, fmt.Errorf("transport error: %w", err)
	}

	resp := res.Response()

	err = v2signature.VerifyServiceMessage(&resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return new(AnnounceIntermediateTrustRes), nil
}
