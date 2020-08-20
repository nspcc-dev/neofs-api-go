package session

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

func CreateRequestBodyToGRPCMessage(c *CreateRequestBody) *session.CreateRequest_Body {
	if c == nil {
		return nil
	}

	m := new(session.CreateRequest_Body)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(c.GetOwnerID()),
	)

	m.SetExpiration(c.GetExpiration())

	return m
}

func CreateRequestBodyFromGRPCMessage(m *session.CreateRequest_Body) *CreateRequestBody {
	if m == nil {
		return nil
	}

	c := new(CreateRequestBody)

	c.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	c.SetExpiration(m.GetExpiration())

	return c
}

func CreateRequestToGRPCMessage(c *CreateRequest) *session.CreateRequest {
	if c == nil {
		return nil
	}

	m := new(session.CreateRequest)

	m.SetBody(
		CreateRequestBodyToGRPCMessage(c.GetBody()),
	)

	RequestHeadersToGRPC(c, m)

	return m
}

func CreateRequestFromGRPCMessage(m *session.CreateRequest) *CreateRequest {
	if m == nil {
		return nil
	}

	c := new(CreateRequest)

	c.SetBody(
		CreateRequestBodyFromGRPCMessage(m.GetBody()),
	)

	RequestHeadersFromGRPC(m, c)

	return c
}

func CreateResponseBodyToGRPCMessage(c *CreateResponseBody) *session.CreateResponse_Body {
	if c == nil {
		return nil
	}

	m := new(session.CreateResponse_Body)

	m.SetId(c.GetID())
	m.SetSessionKey(c.GetSessionKey())

	return m
}

func CreateResponseBodyFromGRPCMessage(m *session.CreateResponse_Body) *CreateResponseBody {
	if m == nil {
		return nil
	}

	c := new(CreateResponseBody)

	c.SetID(m.GetId())
	c.SetSessionKey(m.GetSessionKey())

	return c
}

func CreateResponseToGRPCMessage(c *CreateResponse) *session.CreateResponse {
	if c == nil {
		return nil
	}

	m := new(session.CreateResponse)

	m.SetBody(
		CreateResponseBodyToGRPCMessage(c.GetBody()),
	)

	ResponseHeadersToGRPC(c, m)

	return m
}

func CreateResponseFromGRPCMessage(m *session.CreateResponse) *CreateResponse {
	if m == nil {
		return nil
	}

	c := new(CreateResponse)

	c.SetBody(
		CreateResponseBodyFromGRPCMessage(m.GetBody()),
	)

	ResponseHeadersFromGRPC(m, c)

	return c
}

func TokenLifetimeToGRPCMessage(tl *TokenLifetime) *session.SessionToken_Body_TokenLifetime {
	if tl == nil {
		return nil
	}

	m := new(session.SessionToken_Body_TokenLifetime)

	m.SetExp(tl.GetExp())
	m.SetNbf(tl.GetNbf())
	m.SetIat(tl.GetIat())

	return m
}

func TokenLifetimeFromGRPCMessage(m *session.SessionToken_Body_TokenLifetime) *TokenLifetime {
	if m == nil {
		return nil
	}

	tl := new(TokenLifetime)

	tl.SetExp(m.GetExp())
	tl.SetNbf(m.GetNbf())
	tl.SetIat(m.GetIat())

	return tl
}

func XHeaderToGRPCMessage(x *XHeader) *session.XHeader {
	if x == nil {
		return nil
	}

	m := new(session.XHeader)

	m.SetKey(x.GetKey())
	m.SetValue(x.GetValue())

	return m
}

func XHeaderFromGRPCMessage(m *session.XHeader) *XHeader {
	if m == nil {
		return nil
	}

	x := new(XHeader)

	x.SetKey(m.GetKey())
	x.SetValue(m.GetValue())

	return x
}

func SessionTokenToGRPCMessage(t *SessionToken) *session.SessionToken {
	if t == nil {
		return nil
	}

	m := new(session.SessionToken)

	m.SetBody(
		SessionTokenBodyToGRPCMessage(t.GetBody()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(t.GetSignature()),
	)

	return m
}

func SessionTokenFromGRPCMessage(m *session.SessionToken) *SessionToken {
	if m == nil {
		return nil
	}

	t := new(SessionToken)

	t.SetBody(
		SessionTokenBodyFromGRPCMessage(m.GetBody()),
	)

	t.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return t
}

func RequestVerificationHeaderToGRPCMessage(r *RequestVerificationHeader) *session.RequestVerificationHeader {
	if r == nil {
		return nil
	}

	m := new(session.RequestVerificationHeader)

	m.SetBodySignature(
		refs.SignatureToGRPCMessage(r.GetBodySignature()),
	)

	m.SetMetaSignature(
		refs.SignatureToGRPCMessage(r.GetMetaSignature()),
	)

	m.SetOriginSignature(
		refs.SignatureToGRPCMessage(r.GetOriginSignature()),
	)

	m.SetOrigin(
		RequestVerificationHeaderToGRPCMessage(r.GetOrigin()),
	)

	return m
}

func RequestVerificationHeaderFromGRPCMessage(m *session.RequestVerificationHeader) *RequestVerificationHeader {
	if m == nil {
		return nil
	}

	r := new(RequestVerificationHeader)

	r.SetBodySignature(
		refs.SignatureFromGRPCMessage(m.GetBodySignature()),
	)

	r.SetMetaSignature(
		refs.SignatureFromGRPCMessage(m.GetMetaSignature()),
	)

	r.SetOriginSignature(
		refs.SignatureFromGRPCMessage(m.GetOriginSignature()),
	)

	r.SetOrigin(
		RequestVerificationHeaderFromGRPCMessage(m.GetOrigin()),
	)

	return r
}

func RequestMetaHeaderToGRPCMessage(r *RequestMetaHeader) *session.RequestMetaHeader {
	if r == nil {
		return nil
	}

	m := new(session.RequestMetaHeader)

	m.SetTtl(r.GetTTL())
	m.SetEpoch(r.GetEpoch())

	m.SetVersion(
		refs.VersionToGRPCMessage(r.GetVersion()),
	)

	m.SetSessionToken(
		SessionTokenToGRPCMessage(r.GetSessionToken()),
	)

	m.SetBearerToken(
		acl.BearerTokenToGRPCMessage(r.GetBearerToken()),
	)

	m.SetOrigin(
		RequestMetaHeaderToGRPCMessage(r.GetOrigin()),
	)

	xHeaders := r.GetXHeaders()
	xHdrMsg := make([]*session.XHeader, 0, len(xHeaders))

	for i := range xHeaders {
		xHdrMsg = append(xHdrMsg, XHeaderToGRPCMessage(xHeaders[i]))
	}

	m.SetXHeaders(xHdrMsg)

	return m
}

func RequestMetaHeaderFromGRPCMessage(m *session.RequestMetaHeader) *RequestMetaHeader {
	if m == nil {
		return nil
	}

	r := new(RequestMetaHeader)

	r.SetTTL(m.GetTtl())
	r.SetEpoch(m.GetEpoch())

	r.SetVersion(
		refs.VersionFromGRPCMessage(m.GetVersion()),
	)

	r.SetSessionToken(
		SessionTokenFromGRPCMessage(m.GetSessionToken()),
	)

	r.SetBearerToken(
		acl.BearerTokenFromGRPCMessage(m.GetBearerToken()),
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

func RequestHeadersToGRPC(
	src interface {
		GetMetaHeader() *RequestMetaHeader
		GetVerificationHeader() *RequestVerificationHeader
	},
	dst interface {
		SetMetaHeader(*session.RequestMetaHeader)
		SetVerifyHeader(*session.RequestVerificationHeader)
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
		GetMetaHeader() *session.RequestMetaHeader
		GetVerifyHeader() *session.RequestVerificationHeader
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

func ResponseVerificationHeaderToGRPCMessage(r *ResponseVerificationHeader) *session.ResponseVerificationHeader {
	if r == nil {
		return nil
	}

	m := new(session.ResponseVerificationHeader)

	m.SetBodySignature(
		refs.SignatureToGRPCMessage(r.GetBodySignature()),
	)

	m.SetMetaSignature(
		refs.SignatureToGRPCMessage(r.GetMetaSignature()),
	)

	m.SetOriginSignature(
		refs.SignatureToGRPCMessage(r.GetOriginSignature()),
	)

	m.SetOrigin(
		ResponseVerificationHeaderToGRPCMessage(r.GetOrigin()),
	)

	return m
}

func ResponseVerificationHeaderFromGRPCMessage(m *session.ResponseVerificationHeader) *ResponseVerificationHeader {
	if m == nil {
		return nil
	}

	r := new(ResponseVerificationHeader)

	r.SetBodySignature(
		refs.SignatureFromGRPCMessage(m.GetBodySignature()),
	)

	r.SetMetaSignature(
		refs.SignatureFromGRPCMessage(m.GetMetaSignature()),
	)

	r.SetOriginSignature(
		refs.SignatureFromGRPCMessage(m.GetOriginSignature()),
	)

	r.SetOrigin(
		ResponseVerificationHeaderFromGRPCMessage(m.GetOrigin()),
	)

	return r
}

func ResponseMetaHeaderToGRPCMessage(r *ResponseMetaHeader) *session.ResponseMetaHeader {
	if r == nil {
		return nil
	}

	m := new(session.ResponseMetaHeader)

	m.SetTtl(r.GetTTL())
	m.SetEpoch(r.GetEpoch())

	m.SetVersion(
		refs.VersionToGRPCMessage(r.GetVersion()),
	)

	m.SetOrigin(
		ResponseMetaHeaderToGRPCMessage(r.GetOrigin()),
	)

	xHeaders := r.GetXHeaders()
	xHdrMsg := make([]*session.XHeader, 0, len(xHeaders))

	for i := range xHeaders {
		xHdrMsg = append(xHdrMsg, XHeaderToGRPCMessage(xHeaders[i]))
	}

	m.SetXHeaders(xHdrMsg)

	return m
}

func ResponseMetaHeaderFromGRPCMessage(m *session.ResponseMetaHeader) *ResponseMetaHeader {
	if m == nil {
		return nil
	}

	r := new(ResponseMetaHeader)

	r.SetTTL(m.GetTtl())
	r.SetEpoch(m.GetEpoch())

	r.SetVersion(
		refs.VersionFromGRPCMessage(m.GetVersion()),
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
		SetMetaHeader(*session.ResponseMetaHeader)
		SetVerifyHeader(*session.ResponseVerificationHeader)
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
		GetMetaHeader() *session.ResponseMetaHeader
		GetVerifyHeader() *session.ResponseVerificationHeader
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

func ObjectSessionVerbToGRPCField(v ObjectSessionVerb) session.ObjectSessionContext_Verb {
	switch v {
	case ObjectVerbPut:
		return session.ObjectSessionContext_PUT
	case ObjectVerbGet:
		return session.ObjectSessionContext_GET
	case ObjectVerbHead:
		return session.ObjectSessionContext_HEAD
	case ObjectVerbSearch:
		return session.ObjectSessionContext_SEARCH
	case ObjectVerbDelete:
		return session.ObjectSessionContext_DELETE
	case ObjectVerbRange:
		return session.ObjectSessionContext_RANGE
	case ObjectVerbRangeHash:
		return session.ObjectSessionContext_RANGEHASH
	default:
		return session.ObjectSessionContext_VERB_UNSPECIFIED
	}
}

func ObjectSessionVerbFromGRPCField(v session.ObjectSessionContext_Verb) ObjectSessionVerb {
	switch v {
	case session.ObjectSessionContext_PUT:
		return ObjectVerbPut
	case session.ObjectSessionContext_GET:
		return ObjectVerbGet
	case session.ObjectSessionContext_HEAD:
		return ObjectVerbHead
	case session.ObjectSessionContext_SEARCH:
		return ObjectVerbSearch
	case session.ObjectSessionContext_DELETE:
		return ObjectVerbDelete
	case session.ObjectSessionContext_RANGE:
		return ObjectVerbRange
	case session.ObjectSessionContext_RANGEHASH:
		return ObjectVerbRangeHash
	default:
		return ObjectVerbUnknown
	}
}

func ObjectSessionContextToGRPCMessage(c *ObjectSessionContext) *session.ObjectSessionContext {
	if c == nil {
		return nil
	}

	m := new(session.ObjectSessionContext)

	m.SetVerb(
		ObjectSessionVerbToGRPCField(c.GetVerb()),
	)

	m.SetAddress(
		refs.AddressToGRPCMessage(c.GetAddress()),
	)

	return m
}

func ObjectSessionContextFromGRPCMessage(m *session.ObjectSessionContext) *ObjectSessionContext {
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

func SessionTokenBodyToGRPCMessage(t *SessionTokenBody) *session.SessionToken_Body {
	if t == nil {
		return nil
	}

	m := new(session.SessionToken_Body)

	switch v := t.GetContext(); t := v.(type) {
	case nil:
	case *ObjectSessionContext:
		m.SetObjectSessionContext(
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

func SessionTokenBodyFromGRPCMessage(m *session.SessionToken_Body) *SessionTokenBody {
	if m == nil {
		return nil
	}

	t := new(SessionTokenBody)

	switch v := m.GetContext().(type) {
	case nil:
	case *session.SessionToken_Body_Object:
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
