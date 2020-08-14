package v2

import (
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

func SignatureToGRPCMessage(s *Signature) *service.Signature {
	if s == nil {
		return nil
	}

	m := new(service.Signature)

	m.SetKey(s.GetKey())
	m.SetSign(s.GetSign())

	return m
}

func SignatureFromGRPCMessage(m *service.Signature) *Signature {
	if m == nil {
		return nil
	}

	s := new(Signature)

	s.SetKey(m.GetKey())
	s.SetSign(m.GetSign())

	return s
}

func VersionToGRPCMessage(v *Version) *service.Version {
	if v == nil {
		return nil
	}

	msg := new(service.Version)

	msg.SetMajor(v.GetMajor())
	msg.SetMinor(v.GetMinor())

	return msg
}

func VersionFromGRPCMessage(m *service.Version) *Version {
	if m == nil {
		return nil
	}

	v := new(Version)

	v.SetMajor(m.GetMajor())
	v.SetMinor(m.GetMinor())

	return v
}

func XHeaderToGRPCMessage(x *XHeader) *service.XHeader {
	if x == nil {
		return nil
	}

	m := new(service.XHeader)

	m.SetKey(x.GetKey())
	m.SetValue(x.GetValue())

	return m
}

func XHeaderFromGRPCMessage(m *service.XHeader) *XHeader {
	if m == nil {
		return nil
	}

	x := new(XHeader)

	x.SetKey(m.GetKey())
	x.SetValue(m.GetValue())

	return x
}

func SessionTokenToGRPCMessage(t *SessionToken) *service.SessionToken {
	// TODO: fill me
	return nil
}

func SessionTokenFromGRPCMessage(m *service.SessionToken) *SessionToken {
	// TODO: fill me
	return nil
}

func BearerTokenToGRPCMessage(t *BearerToken) *service.BearerToken {
	// TODO: fill me
	return nil
}

func BearerTokenFromGRPCMessage(m *service.BearerToken) *BearerToken {
	// TODO: fill me
	return nil
}

func RequestVerificationHeaderToGRPCMessage(r *RequestVerificationHeader) *service.RequestVerificationHeader {
	if r == nil {
		return nil
	}

	m := new(service.RequestVerificationHeader)

	m.SetBodySignature(
		SignatureToGRPCMessage(r.GetBodySignature()),
	)

	m.SetMetaSignature(
		SignatureToGRPCMessage(r.GetMetaSignature()),
	)

	m.SetOriginSignature(
		SignatureToGRPCMessage(r.GetOriginSignature()),
	)

	m.SetOrigin(
		RequestVerificationHeaderToGRPCMessage(r.GetOrigin()),
	)

	return m
}

func RequestVerificationHeaderFromGRPCMessage(m *service.RequestVerificationHeader) *RequestVerificationHeader {
	if m == nil {
		return nil
	}

	r := new(RequestVerificationHeader)

	r.SetBodySignature(
		SignatureFromGRPCMessage(m.GetBodySignature()),
	)

	r.SetMetaSignature(
		SignatureFromGRPCMessage(m.GetMetaSignature()),
	)

	r.SetOriginSignature(
		SignatureFromGRPCMessage(m.GetOriginSignature()),
	)

	r.SetOrigin(
		RequestVerificationHeaderFromGRPCMessage(m.GetOrigin()),
	)

	return r
}

func RequestMetaHeaderToGRPCMessage(r *RequestMetaHeader) *service.RequestMetaHeader {
	if r == nil {
		return nil
	}

	m := new(service.RequestMetaHeader)

	m.SetTtl(r.GetTTL())
	m.SetEpoch(r.GetEpoch())

	m.SetVersion(
		VersionToGRPCMessage(r.GetVersion()),
	)

	m.SetSessionToken(
		SessionTokenToGRPCMessage(r.GetSessionToken()),
	)

	m.SetBearerToken(
		BearerTokenToGRPCMessage(r.GetBearerToken()),
	)

	m.SetOrigin(
		RequestMetaHeaderToGRPCMessage(r.GetOrigin()),
	)

	xHeaders := r.GetXHeaders()
	xHdrMsg := make([]*service.XHeader, 0, len(xHeaders))

	for i := range xHeaders {
		xHdrMsg = append(xHdrMsg, XHeaderToGRPCMessage(xHeaders[i]))
	}

	return m
}

func RequestMetaHeaderFromGRPCMessage(m *service.RequestMetaHeader) *RequestMetaHeader {
	if m == nil {
		return nil
	}

	r := new(RequestMetaHeader)

	r.SetTTL(m.GetTtl())
	r.SetEpoch(m.GetEpoch())

	r.SetVersion(
		VersionFromGRPCMessage(m.GetVersion()),
	)

	r.SetSessionToken(
		SessionTokenFromGRPCMessage(m.GetSessionToken()),
	)

	r.SetBearerToken(
		BearerTokenFromGRPCMessage(m.GetBearerToken()),
	)

	r.SetOrigin(
		RequestMetaHeaderFromGRPCMessage(m.GetOrigin()),
	)

	xHdrMsg := m.GetXHeaders()
	xHeaders := make([]*XHeader, 0, len(xHdrMsg))

	for i := range xHdrMsg {
		xHeaders = append(xHeaders, XHeaderFromGRPCMessage(xHdrMsg[i]))
	}

	return r
}

func OwnerIDToGRPCMessage(o *OwnerID) *refs.OwnerID {
	if o == nil {
		return nil
	}

	m := new(refs.OwnerID)

	m.SetValue(o.GetValue())

	return m
}

func OwnerIDFromGRPCMessage(m *refs.OwnerID) *OwnerID {
	if m == nil {
		return nil
	}

	o := new(OwnerID)

	o.SetValue(m.GetValue())

	return o
}

func BalanceRequestBodyToGRPCMessage(b *BalanceRequestBody) *accounting.BalanceRequest_Body {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest_Body)

	m.SetOwnerId(
		OwnerIDToGRPCMessage(b.GetOwnerID()),
	)

	return m
}

func BalanceRequestBodyFromGRPCMessage(m *accounting.BalanceRequest_Body) *BalanceRequestBody {
	if m == nil {
		return nil
	}

	b := new(BalanceRequestBody)

	b.SetOwnerID(
		OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	return b
}

func headersToGRPC(
	src interface {
		GetRequestMetaHeader() *RequestMetaHeader
		GetRequestVerificationHeader() *RequestVerificationHeader
	},
	dst interface {
		SetMetaHeader(*service.RequestMetaHeader)
		SetVerifyHeader(*service.RequestVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		RequestMetaHeaderToGRPCMessage(src.GetRequestMetaHeader()),
	)

	dst.SetVerifyHeader(
		RequestVerificationHeaderToGRPCMessage(src.GetRequestVerificationHeader()),
	)
}

func headersFromGRPC(
	src interface {
		GetMetaHeader() *service.RequestMetaHeader
		GetVerifyHeader() *service.RequestVerificationHeader
	},
	dst interface {
		SetRequestMetaHeader(*RequestMetaHeader)
		SetRequestVerificationHeader(*RequestVerificationHeader)
	},
) {
	dst.SetRequestMetaHeader(
		RequestMetaHeaderFromGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetRequestVerificationHeader(
		RequestVerificationHeaderFromGRPCMessage(src.GetVerifyHeader()),
	)
}

func BalanceRequestToGRPCMessage(b *BalanceRequest) *accounting.BalanceRequest {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest)

	m.SetBody(
		BalanceRequestBodyToGRPCMessage(b.GetBody()),
	)

	headersToGRPC(b, m)

	return m
}

func BalanceRequestFromGRPCMessage(m *accounting.BalanceRequest) *BalanceRequest {
	if m == nil {
		return nil
	}

	b := new(BalanceRequest)

	b.SetBody(
		BalanceRequestBodyFromGRPCMessage(m.GetBody()),
	)

	headersFromGRPC(m, b)

	return b
}
