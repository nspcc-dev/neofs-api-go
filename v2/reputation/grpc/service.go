package reputation

import (
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

// SetEpoch sets epoch in which the trust was assessed.
func (x *SendLocalTrustRequest_Body) SetEpoch(v uint64) {
	if x != nil {
		x.Epoch = v
	}
}

// SetTrusts sets list of normalized trust values.
func (x *SendLocalTrustRequest_Body) SetTrusts(v []*Trust) {
	if x != nil {
		x.Trusts = v
	}
}

// SetBody sets body of the request.
func (x *SendLocalTrustRequest) SetBody(v *SendLocalTrustRequest_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (x *SendLocalTrustRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (x *SendLocalTrustRequest) SetVerifyHeader(v *session.RequestVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (x *SendLocalTrustResponse) SetBody(v *SendLocalTrustResponse_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (x *SendLocalTrustResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (x *SendLocalTrustResponse) SetVerifyHeader(v *session.ResponseVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetIteration sets sequence number of the iteration.
func (x *SendIntermediateResultRequest_Body) SetIteration(v uint32) {
	if x != nil {
		x.Iteration = v
	}
}

// SetTrust sets current global trust value.
func (x *SendIntermediateResultRequest_Body) SetTrust(v *PeerToPeerTrust) {
	if x != nil {
		x.Trust = v
	}
}

// SetBody sets body of the request.
func (x *SendIntermediateResultRequest) SetBody(v *SendIntermediateResultRequest_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (x *SendIntermediateResultRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (x *SendIntermediateResultRequest) SetVerifyHeader(v *session.RequestVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (x *SendIntermediateResultResponse) SetBody(v *SendIntermediateResultResponse_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (x *SendIntermediateResultResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (x *SendIntermediateResultResponse) SetVerifyHeader(v *session.ResponseVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}
