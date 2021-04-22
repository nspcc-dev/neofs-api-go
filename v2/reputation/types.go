package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

// PeerID represents reputation.PeerID message
// from NeoFS API v2.
type PeerID struct {
	val []byte
}

// GetValue returns peer's binary ID.
func (x *PeerID) GetValue() []byte {
	if x != nil {
		return x.val
	}

	return nil
}

// SetValue sets peer's binary ID.
func (x *PeerID) SetValue(v []byte) {
	if x != nil {
		x.val = v
	}
}

// Trust represents reputation.Trust message
// from NeoFS API v2.
type Trust struct {
	val float64

	peer *PeerID
}

// GetPeer returns trusted peer's ID.
func (x *Trust) GetPeer() *PeerID {
	if x != nil {
		return x.peer
	}

	return nil
}

// SetPeer sets trusted peer's ID.
func (x *Trust) SetPeer(v *PeerID) {
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

// PeerToPeerTrust represents reputation.PeerToPeerTrust message
// from NeoFS API v2.
type PeerToPeerTrust struct {
	trusting *PeerID

	trust *Trust
}

// GetTrustingPeer returns trusting peer ID.
func (x *PeerToPeerTrust) GetTrustingPeer() *PeerID {
	if x != nil {
		return x.trusting
	}

	return nil
}

// SetTrustingPeer sets trusting peer ID.
func (x *PeerToPeerTrust) SetTrustingPeer(v *PeerID) {
	if x != nil {
		x.trusting = v
	}
}

// GetTrust returns trust value of trusting peer to the trusted one.
func (x *PeerToPeerTrust) GetTrust() *Trust {
	if x != nil {
		return x.trust
	}

	return nil
}

// SetTrust sets trust value of trusting peer to the trusted one.
func (x *PeerToPeerTrust) SetTrust(v *Trust) {
	if x != nil {
		x.trust = v
	}
}

// GlobalTrustBody represents reputation.GlobalTrust.Body message
// from NeoFS API v2.
type GlobalTrustBody struct {
	manager *PeerID

	trust *Trust
}

// GetManager returns node manager ID.
func (x *GlobalTrustBody) GetManager() *PeerID {
	if x != nil {
		return x.manager
	}

	return nil
}

// SetManager sets node manager ID.
func (x *GlobalTrustBody) SetManager(v *PeerID) {
	if x != nil {
		x.manager = v
	}
}

// GetTrust returns global trust value.
func (x *GlobalTrustBody) GetTrust() *Trust {
	if x != nil {
		return x.trust
	}

	return nil
}

// SetTrust sets global trust value.
func (x *GlobalTrustBody) SetTrust(v *Trust) {
	if x != nil {
		x.trust = v
	}
}

// GlobalTrust represents reputation.GlobalTrust message
// from NeoFS API v2.
type GlobalTrust struct {
	version *refs.Version

	body *GlobalTrustBody

	sig *refs.Signature
}

// GetVersion returns message format version.
func (x *GlobalTrust) GetVersion() *refs.Version {
	if x != nil {
		return x.version
	}

	return nil
}

// SetVersion sets message format version.
func (x *GlobalTrust) SetVersion(v *refs.Version) {
	if x != nil {
		x.version = v
	}
}

// GetBody returns message body.
func (x *GlobalTrust) GetBody() *GlobalTrustBody {
	if x != nil {
		return x.body
	}

	return nil
}

// SetBody sets message body.
func (x *GlobalTrust) SetBody(v *GlobalTrustBody) {
	if x != nil {
		x.body = v
	}
}

// GetSignature returns body signature.
func (x *GlobalTrust) GetSignature() *refs.Signature {
	if x != nil {
		return x.sig
	}

	return nil
}

// SetSignature sets body signature.
func (x *GlobalTrust) SetSignature(v *refs.Signature) {
	if x != nil {
		x.sig = v
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

// SendIntermediateResultRequestBody is a structure of SendIntermediateResult request body.
type SendIntermediateResultRequestBody struct {
	epoch uint64

	iter uint32

	trust *PeerToPeerTrust
}

// GetEpoch returns epoch number in which the intermediate trust was assessed.
func (x *SendIntermediateResultRequestBody) GetEpoch() uint64 {
	if x != nil {
		return x.epoch
	}

	return 0
}

// SetEpoch sets epoch number in which the intermediate trust was assessed.
func (x *SendIntermediateResultRequestBody) SetEpoch(v uint64) {
	if x != nil {
		x.epoch = v
	}
}

// GetIteration returns sequence number of the iteration.
func (x *SendIntermediateResultRequestBody) GetIteration() uint32 {
	if x != nil {
		return x.iter
	}

	return 0
}

// SetIteration sets sequence number of the iteration.
func (x *SendIntermediateResultRequestBody) SetIteration(v uint32) {
	if x != nil {
		x.iter = v
	}
}

// GetTrust returns current global trust value.
func (x *SendIntermediateResultRequestBody) GetTrust() *PeerToPeerTrust {
	if x != nil {
		return x.trust
	}

	return nil
}

// SetTrust sets current global trust value.
func (x *SendIntermediateResultRequestBody) SetTrust(v *PeerToPeerTrust) {
	if x != nil {
		x.trust = v
	}
}

// SendLocalTrustResponseBody is a structure of SendIntermediateResult response body.
type SendIntermediateResultResponseBody struct{}

// SendIntermediateResultRequest represents reputation.SendIntermediateResult
// message from NeoFS API v2.
type SendIntermediateResultRequest struct {
	body *SendIntermediateResultRequestBody

	session.RequestHeaders
}

// GetBody returns request body.
func (x *SendIntermediateResultRequest) GetBody() *SendIntermediateResultRequestBody {
	if x != nil {
		return x.body
	}

	return nil
}

// SetBody sets request body.
func (x *SendIntermediateResultRequest) SetBody(v *SendIntermediateResultRequestBody) {
	if x != nil {
		x.body = v
	}
}

// SendIntermediateResultResponse represents reputation.SendIntermediateResultResponse
// message from NeoFS API v2.
type SendIntermediateResultResponse struct {
	body *SendIntermediateResultResponseBody

	session.ResponseHeaders
}

// GetBody returns response body.
func (x *SendIntermediateResultResponse) GetBody() *SendIntermediateResultResponseBody {
	if x != nil {
		return x.body
	}

	return nil
}

// SetBody sets response body.
func (x *SendIntermediateResultResponse) SetBody(v *SendIntermediateResultResponseBody) {
	if x != nil {
		x.body = v
	}
}
