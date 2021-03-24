package client

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/pkg/reputation"
	"github.com/nspcc-dev/neofs-api-go/rpc/client"
	v2reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation"
	rpcapi "github.com/nspcc-dev/neofs-api-go/v2/rpc"
	v2signature "github.com/nspcc-dev/neofs-api-go/v2/signature"
	"github.com/pkg/errors"
)

// Reputation contains methods for working with Reputation system values.
type Reputation interface {
	// SendLocalTrust sends local trust values of local peer.
	SendLocalTrust(context.Context, SendLocalTrustPrm, ...CallOption) (*SendLocalTrustRes, error)
}

// SendLocalTrustPrm groups parameters of SendLocalTrust operation.
type SendLocalTrustPrm struct {
	epoch uint64

	trusts []*reputation.Trust
}

// Epoch returns epoch in which the trust was assessed.
func (x SendLocalTrustPrm) Epoch() uint64 {
	return x.epoch
}

// SetEpoch sets epoch in which the trust was assessed.
func (x *SendLocalTrustPrm) SetEpoch(epoch uint64) {
	x.epoch = epoch
}

// Trusts returns list of local trust values.
func (x SendLocalTrustPrm) Trusts() []*reputation.Trust {
	return x.trusts
}

// SetTrusts sets list of local trust values.
func (x *SendLocalTrustPrm) SetTrusts(trusts []*reputation.Trust) {
	x.trusts = trusts
}

// SendLocalTrustPrm groups results of SendLocalTrust operation.
type SendLocalTrustRes struct{}

func (c *clientImpl) SendLocalTrust(ctx context.Context, prm SendLocalTrustPrm, opts ...CallOption) (*SendLocalTrustRes, error) {
	// apply all available options
	callOptions := c.defaultCallOptions()

	for i := range opts {
		opts[i](callOptions)
	}

	reqBody := new(v2reputation.SendLocalTrustRequestBody)
	reqBody.SetEpoch(prm.Epoch())
	reqBody.SetTrusts(reputation.TrustsToV2(prm.Trusts()))

	req := new(v2reputation.SendLocalTrustRequest)
	req.SetBody(reqBody)
	req.SetMetaHeader(v2MetaHeaderFromOpts(callOptions))

	err := v2signature.SignServiceMessage(callOptions.key, req)
	if err != nil {
		return nil, err
	}

	resp, err := rpcapi.SendLocalTrust(c.Raw(), req, client.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	err = v2signature.VerifyServiceMessage(resp)
	if err != nil {
		return nil, errors.Wrap(err, "can't verify response message")
	}

	return new(SendLocalTrustRes), nil
}
