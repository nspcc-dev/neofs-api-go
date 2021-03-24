package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

// Trust represents reputation.Trust message
// from NeoFS API v2.
type Trust struct {
	val float64

	peer []byte
}

// GetPeer returns trusted peer's ID.
func (x *Trust) GetPeer() []byte {
	if x != nil {
		return x.peer
	}

	return nil
}

// SetPeer sets trusted peer's ID.
func (x *Trust) SetPeer(v []byte) {
	if x != nil {
		x.peer = v
	}
}

// GetValue returns trust value.
func (x *Trust) GetValue() float64 {
	if x != nil {
		return x.val
	}

	return 0
}

// SetValue sets trust value.
func (x *Trust) SetValue(v float64) {
	if x != nil {
		x.val = v
	}
}

// SendLocalTrustRequestBody is a structure of SendLocalTrust request body.
type SendLocalTrustRequestBody struct {
	epoch uint64

	trusts []*Trust
}

// GetEpoch returns epoch in which the trust was assessed.
func (x *SendLocalTrustRequestBody) GetEpoch() uint64 {
	if x != nil {
		return x.epoch
	}

	return 0
}

// SetEpoch sets epoch in which the trust was assessed.
func (x *SendLocalTrustRequestBody) SetEpoch(v uint64) {
	if x != nil {
		x.epoch = v
	}
}

// GetTrusts returns list of normalized trust values.
func (x *SendLocalTrustRequestBody) GetTrusts() []*Trust {
	if x != nil {
		return x.trusts
	}

	return nil
}

// SetTrusts sets list of normalized trust values.
func (x *SendLocalTrustRequestBody) SetTrusts(v []*Trust) {
	if x != nil {
		x.trusts = v
	}
}

// SendLocalTrustResponseBody is a structure of SendLocalTrust response body.
type SendLocalTrustResponseBody struct{}

// SendLocalTrustRequest represents reputation.SendLocalTrustRequest
// message from NeoFS API v2.
type SendLocalTrustRequest struct {
	body *SendLocalTrustRequestBody

	session.RequestHeaders
}

// GetBody returns request body.
func (x *SendLocalTrustRequest) GetBody() *SendLocalTrustRequestBody {
	if x != nil {
		return x.body
	}

	return nil
}

// SetBody sets request body.
func (x *SendLocalTrustRequest) SetBody(v *SendLocalTrustRequestBody) {
	if x != nil {
		x.body = v
	}
}

// SendLocalTrustResponse represents reputation.SendLocalTrustResponse
// message from NeoFS API v2.
type SendLocalTrustResponse struct {
	body *SendLocalTrustResponseBody

	session.ResponseHeaders
}

// GetBody returns response body.
func (x *SendLocalTrustResponse) GetBody() *SendLocalTrustResponseBody {
	if x != nil {
		return x.body
	}

	return nil
}

// SetBody sets response body.
func (x *SendLocalTrustResponse) SetBody(v *SendLocalTrustResponseBody) {
	if x != nil {
		x.body = v
	}
}
