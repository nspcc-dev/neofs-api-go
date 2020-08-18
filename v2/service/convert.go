package service

import (
	"fmt"

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
	if t == nil {
		return nil
	}

	m := new(service.SessionToken)

	m.SetBody(
		SessionTokenBodyToGRPCMessage(t.GetBody()),
	)

	m.SetSignature(
		SignatureToGRPCMessage(t.GetSignature()),
	)

	return m
}

func SessionTokenFromGRPCMessage(m *service.SessionToken) *SessionToken {
	if m == nil {
		return nil
	}

	t := new(SessionToken)

	t.SetBody(
		SessionTokenBodyFromGRPCMessage(m.GetBody()),
	)

	t.SetSignature(
		SignatureFromGRPCMessage(m.GetSignature()),
	)

	return t
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

	m.SetXHeaders(xHdrMsg)

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

	r.SetXHeaders(xHeaders)

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
		GetMetaHeader() *RequestMetaHeader
		GetVerificationHeader() *RequestVerificationHeader
	},
	dst interface {
		SetMetaHeader(*service.RequestMetaHeader)
		SetVerifyHeader(*service.RequestVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		RequestMetaHeaderToGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetVerifyHeader(
		RequestVerificationHeaderToGRPCMessage(src.GetVerificationHeader()),
	)
}

func RequestHeadersFromGRPC(
	src interface {
		GetMetaHeader() *service.RequestMetaHeader
		GetVerifyHeader() *service.RequestVerificationHeader
	},
	dst interface {
		SetMetaHeader(*RequestMetaHeader)
		SetVerificationHeader(*RequestVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		RequestMetaHeaderFromGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetVerificationHeader(
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

	m.SetXHeaders(xHdrMsg)

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

	r.SetXHeaders(xHeaders)

	return r
}

func ResponseHeadersToGRPC(
	src interface {
		GetMetaHeader() *ResponseMetaHeader
		GetVerificationHeader() *ResponseVerificationHeader
	},
	dst interface {
		SetMetaHeader(*service.ResponseMetaHeader)
		SetVerifyHeader(*service.ResponseVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		ResponseMetaHeaderToGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetVerifyHeader(
		ResponseVerificationHeaderToGRPCMessage(src.GetVerificationHeader()),
	)
}

func ResponseHeadersFromGRPC(
	src interface {
		GetMetaHeader() *service.ResponseMetaHeader
		GetVerifyHeader() *service.ResponseVerificationHeader
	},
	dst interface {
		SetMetaHeader(*ResponseMetaHeader)
		SetVerificationHeader(*ResponseVerificationHeader)
	},
) {
	dst.SetMetaHeader(
		ResponseMetaHeaderFromGRPCMessage(src.GetMetaHeader()),
	)

	dst.SetVerificationHeader(
		ResponseVerificationHeaderFromGRPCMessage(src.GetVerifyHeader()),
	)
}

func ObjectSessionVerbToGRPCField(v ObjectSessionVerb) service.ObjectSessionContext_Verb {
	switch v {
	case ObjectVerbPut:
		return service.ObjectSessionContext_PUT
	case ObjectVerbGet:
		return service.ObjectSessionContext_GET
	case ObjectVerbHead:
		return service.ObjectSessionContext_HEAD
	case ObjectVerbSearch:
		return service.ObjectSessionContext_SEARCH
	case ObjectVerbDelete:
		return service.ObjectSessionContext_DELETE
	case ObjectVerbRange:
		return service.ObjectSessionContext_RANGE
	case ObjectVerbRangeHash:
		return service.ObjectSessionContext_RANGEHASH
	default:
		return service.ObjectSessionContext_VERB_UNSPECIFIED
	}
}

func ObjectSessionVerbFromGRPCField(v service.ObjectSessionContext_Verb) ObjectSessionVerb {
	switch v {
	case service.ObjectSessionContext_PUT:
		return ObjectVerbPut
	case service.ObjectSessionContext_GET:
		return ObjectVerbGet
	case service.ObjectSessionContext_HEAD:
		return ObjectVerbHead
	case service.ObjectSessionContext_SEARCH:
		return ObjectVerbSearch
	case service.ObjectSessionContext_DELETE:
		return ObjectVerbDelete
	case service.ObjectSessionContext_RANGE:
		return ObjectVerbRange
	case service.ObjectSessionContext_RANGEHASH:
		return ObjectVerbRangeHash
	default:
		return ObjectVerbUnknown
	}
}

func ObjectSessionContextToGRPCMessage(c *ObjectSessionContext) *service.ObjectSessionContext {
	if c == nil {
		return nil
	}

	m := new(service.ObjectSessionContext)

	m.SetVerb(
		ObjectSessionVerbToGRPCField(c.GetVerb()),
	)

	m.SetAddress(
		refs.AddressToGRPCMessage(c.GetAddress()),
	)

	return m
}

func ObjectSessionContextFromGRPCMessage(m *service.ObjectSessionContext) *ObjectSessionContext {
	if m == nil {
		return nil
	}

	c := new(ObjectSessionContext)

	c.SetVerb(
		ObjectSessionVerbFromGRPCField(m.GetVerb()),
	)

	c.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	return c
}

func SessionTokenBodyToGRPCMessage(t *SessionTokenBody) *service.SessionToken_Body {
	if t == nil {
		return nil
	}

	m := new(service.SessionToken_Body)

	switch v := t.GetContext(); t := v.(type) {
	case nil:
	case *ObjectSessionContext:
		m.SetObjectServiceContext(
			ObjectSessionContextToGRPCMessage(t),
		)
	default:
		panic(fmt.Sprintf("unknown session context %T", t))
	}

	m.SetId(t.GetID())

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(t.GetOwnerID()),
	)

	m.SetLifetime(
		TokenLifetimeToGRPCMessage(t.GetLifetime()),
	)

	m.SetSessionKey(t.GetSessionKey())

	return m
}

func SessionTokenBodyFromGRPCMessage(m *service.SessionToken_Body) *SessionTokenBody {
	if m == nil {
		return nil
	}

	t := new(SessionTokenBody)

	switch v := m.GetContext().(type) {
	case nil:
	case *service.SessionToken_Body_Object:
		t.SetContext(
			ObjectSessionContextFromGRPCMessage(v.Object),
		)
	default:
		panic(fmt.Sprintf("unknown session context %T", v))
	}

	t.SetID(m.GetId())

	t.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	t.SetLifetime(
		TokenLifetimeFromGRPCMessage(m.GetLifetime()),
	)

	t.SetSessionKey(m.GetSessionKey())

	return t
}
