package reputation

import (
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

// SetEpoch sets epoch in which the trust was assessed.
func (x *AnnounceLocalTrustRequest_Body) SetEpoch(v uint64) {
	if x != nil {
		x.Epoch = v
	}
}

// SetTrusts sets list of normalized trust values.
func (x *AnnounceLocalTrustRequest_Body) SetTrusts(v []*Trust) {
	if x != nil {
		x.Trusts = v
	}
}

// SetBody sets body of the request.
func (x *AnnounceLocalTrustRequest) SetBody(v *AnnounceLocalTrustRequest_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (x *AnnounceLocalTrustRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (x *AnnounceLocalTrustRequest) SetVerifyHeader(v *session.RequestVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (x *AnnounceLocalTrustResponse) SetBody(v *AnnounceLocalTrustResponse_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (x *AnnounceLocalTrustResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (x *AnnounceLocalTrustResponse) SetVerifyHeader(v *session.ResponseVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetEpoch sets epoch in which the intermediate trust was assessed.
func (x *AnnounceIntermediateResultRequest_Body) SetEpoch(v uint64) {
	if x != nil {
		x.Epoch = v
	}
}

// SetIteration sets sequence number of the iteration.
func (x *AnnounceIntermediateResultRequest_Body) SetIteration(v uint32) {
	if x != nil {
		x.Iteration = v
	}
}

// SetTrust sets current global trust value.
func (x *AnnounceIntermediateResultRequest_Body) SetTrust(v *PeerToPeerTrust) {
	if x != nil {
		x.Trust = v
	}
}

// SetBody sets body of the request.
func (x *AnnounceIntermediateResultRequest) SetBody(v *AnnounceIntermediateResultRequest_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (x *AnnounceIntermediateResultRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (x *AnnounceIntermediateResultRequest) SetVerifyHeader(v *session.RequestVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}

// SetBody sets body of the response.
func (x *AnnounceIntermediateResultResponse) SetBody(v *AnnounceIntermediateResultResponse_Body) {
	if x != nil {
		x.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (x *AnnounceIntermediateResultResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if x != nil {
		x.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (x *AnnounceIntermediateResultResponse) SetVerifyHeader(v *session.ResponseVerificationHeader) {
	if x != nil {
		x.VerifyHeader = v
	}
}
