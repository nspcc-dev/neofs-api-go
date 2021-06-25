package client

import (
	"context"
	"fmt"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	v2reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
)

// Reputation contains methods for working with Reputation system values.
type Reputation interface {
	// AnnounceLocalTrust announces local trust values of local peer.
	AnnounceLocalTrust(context.Context, AnnounceLocalTrustPrm, ...CallOption) (*AnnounceLocalTrustRes, error)

	// AnnounceIntermediateTrust announces the intermediate result of the iterative algorithm for calculating
	// the global reputation of the node.
	AnnounceIntermediateTrust(context.Context, AnnounceIntermediateTrustPrm, ...CallOption) (*AnnounceIntermediateTrustRes, error)
}

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

func (c *clientImpl) AnnounceLocalTrust(ctx context.Context, prm AnnounceLocalTrustPrm, opts ...CallOption) (*AnnounceLocalTrustRes, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2reputation.AnnounceLocalTrustRequestBody)
	reqBody.SetEpoch(prm.Epoch())
	reqBody.SetTrusts(reputation.TrustsToV2(prm.Trusts()))

	req := new(v2reputation.AnnounceLocalTrustRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.AnnounceLocalTrust(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	err = v2signature.VerifyServiceMessage(resp)
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

func (c *clientImpl) AnnounceIntermediateTrust(ctx context.Context, prm AnnounceIntermediateTrustPrm, opts ...CallOption) (*AnnounceIntermediateTrustRes, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2reputation.AnnounceIntermediateResultRequestBody)
	reqBody.SetEpoch(prm.Epoch())
	reqBody.SetIteration(prm.Iteration())
	reqBody.SetTrust(prm.Trust().ToV2())

	req := new(v2reputation.AnnounceIntermediateResultRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(neofsecdsa.Signer(callOptions.key), req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.AnnounceIntermediateResult(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, fmt.Errorf("can't verify response message: %w", err)
	}

	return new(AnnounceIntermediateTrustRes), nil
}
