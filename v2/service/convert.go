package service

import (
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	service "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

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
	if t == nil {
		return nil
	}

	m := new(service.BearerToken)

	m.SetBody(
		BearerTokenBodyToGRPCMessage(t.GetBody()),
	)

	m.SetSignature(
		SignatureToGRPCMessage(t.GetSignature()),
	)

	return m
}

func BearerTokenFromGRPCMessage(m *service.BearerToken) *BearerToken {
	if m == nil {
		return nil
	}

	bt := new(BearerToken)

	bt.SetBody(
		BearerTokenBodyFromGRPCMessage(m.GetBody()),
	)

	bt.SetSignature(
		SignatureFromGRPCMessage(m.GetSignature()),
	)

	return bt
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

func TokenLifetimeToGRPCMessage(tl *TokenLifetime) *service.TokenLifetime {
	if tl == nil {
		return nil
	}

	m := new(service.TokenLifetime)

	m.SetExp(tl.GetExp())
	m.SetNbf(tl.GetNbf())
	m.SetIat(tl.GetIat())

	return m
}

func TokenLifetimeFromGRPCMessage(m *service.TokenLifetime) *TokenLifetime {
	if m == nil {
		return nil
	}

	tl := new(TokenLifetime)

	tl.SetExp(m.GetExp())
	tl.SetNbf(m.GetNbf())
	tl.SetIat(m.GetIat())

	return tl
}

func BearerTokenBodyToGRPCMessage(v *BearerTokenBody) *service.BearerToken_Body {
	if v == nil {
		return nil
	}

	m := new(service.BearerToken_Body)

	m.SetEaclTable(
		acl.TableToGRPCMessage(v.GetEACL()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(v.GetOwnerID()),
	)

	m.SetLifetime(
		TokenLifetimeToGRPCMessage(v.GetLifetime()),
	)

	return m
}

func BearerTokenBodyFromGRPCMessage(m *service.BearerToken_Body) *BearerTokenBody {
	if m == nil {
		return nil
	}

	bt := new(BearerTokenBody)

	bt.SetEACL(
		acl.TableFromGRPCMessage(m.GetEaclTable()),
	)

	bt.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	bt.SetLifetime(
		TokenLifetimeFromGRPCMessage(m.GetLifetime()),
	)

	return bt
}

func RequestHeadersToGRPC(
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

func RequestHeadersFromGRPC(
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

func ResponseVerificationHeaderToGRPCMessage(r *ResponseVerificationHeader) *service.ResponseVerificationHeader {
	if r == nil {
		return nil
	}

	m := new(service.ResponseVerificationHeader)

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
		ResponseVerificationHeaderToGRPCMessage(r.GetOrigin()),
	)

	return m
}

func ResponseVerificationHeaderFromGRPCMessage(m *service.ResponseVerificationHeader) *ResponseVerificationHeader {
	if m == nil {
		return nil
	}

	r := new(ResponseVerificationHeader)

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
		ResponseVerificationHeaderFromGRPCMessage(m.GetOrigin()),
	)

	return r
}

func ResponseMetaHeaderToGRPCMessage(r *ResponseMetaHeader) *service.ResponseMetaHeader {
	if r == nil {
		return nil
	}

	m := new(service.ResponseMetaHeader)

	m.SetTtl(r.GetTTL())
	m.SetEpoch(r.GetEpoch())

	m.SetVersion(
		VersionToGRPCMessage(r.GetVersion()),
	)

	m.SetOrigin(
		ResponseMetaHeaderToGRPCMessage(r.GetOrigin()),
	)

	xHeaders := r.GetXHeaders()
	xHdrMsg := make([]*service.XHeader, 0, len(xHeaders))

	for i := range xHeaders {
		xHdrMsg = append(xHdrMsg, XHeaderToGRPCMessage(xHeaders[i]))
	}

	return m
}

func ResponseMetaHeaderFromGRPCMessage(m *service.ResponseMetaHeader) *ResponseMetaHeader {
	if m == nil {
		return nil
	}

	r := new(ResponseMetaHeader)

	r.SetTTL(m.GetTtl())
	r.SetEpoch(m.GetEpoch())

	r.SetVersion(
		VersionFromGRPCMessage(m.GetVersion()),
	)

	r.SetOrigin(
		ResponseMetaHeaderFromGRPCMessage(m.GetOrigin()),
	)

	xHdrMsg := m.GetXHeaders()
	xHeaders := make([]*XHeader, 0, len(xHdrMsg))

	for i := range xHdrMsg {
		xHeaders = append(xHeaders, XHeaderFromGRPCMessage(xHdrMsg[i]))
	}

	return r
}

func ResponseHeadersToGRPC(
	src interface {
		GetResponseMetaHeader() *ResponseMetaHeader
		GetResponseVerificationHeader() *ResponseVerificationHeader
	},
	dst interface {
		SetMetaHeader(*service.ResponseMetaHeader)
		SetVerifyHeader(*service.ResponseVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		ResponseMetaHeaderToGRPCMessage(src.GetResponseMetaHeader()),
	)

	dst.SetVerifyHeader(
		ResponseVerificationHeaderToGRPCMessage(src.GetResponseVerificationHeader()),
	)
}

func ResponseHeadersFromGRPC(
	src interface {
		GetMetaHeader() *service.ResponseMetaHeader
		GetVerifyHeader() *service.ResponseVerificationHeader
	},
	dst interface {
		SetResponseMetaHeader(*ResponseMetaHeader)
		SetResponseVerificationHeader(*ResponseVerificationHeader)
	},
) {
	dst.SetResponseMetaHeader(
		ResponseMetaHeaderFromGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetResponseVerificationHeader(
		ResponseVerificationHeaderFromGRPCMessage(src.GetVerifyHeader()),
	)
}
