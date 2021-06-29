package session

import (
	"fmt"

	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	aclGRPC "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	session "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
)

func (c *CreateRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *session.CreateRequest_Body

	if c != nil {
		m = new(session.CreateRequest_Body)

		m.SetOwnerId(c.ownerID.ToGRPCMessage().(*refsGRPC.OwnerID))
		m.SetExpiration(c.expiration)
	}

	return m
}

func (c *CreateRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.CreateRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	ownerID := v.GetOwnerId()
	if ownerID == nil {
		c.ownerID = nil
	} else {
		if c.ownerID == nil {
			c.ownerID = new(refs.OwnerID)
		}

		err = c.ownerID.FromGRPCMessage(ownerID)
		if err != nil {
			return err
		}
	}

	c.expiration = v.GetExpiration()

	return nil
}

func (c *CreateRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *session.CreateRequest

	if c != nil {
		m = new(session.CreateRequest)

		m.SetBody(c.body.ToGRPCMessage().(*session.CreateRequest_Body))
		c.RequestHeaders.ToMessage(m)
	}

	return m
}

func (c *CreateRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.CreateRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		c.body = nil
	} else {
		if c.body == nil {
			c.body = new(CreateRequestBody)
		}

		err = c.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return c.RequestHeaders.FromMessage(v)
}

func (c *CreateResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *session.CreateResponse_Body

	if c != nil {
		m = new(session.CreateResponse_Body)

		m.SetSessionKey(c.sessionKey)
		m.SetId(c.id)
	}

	return m
}

func (c *CreateResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.CreateResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	c.sessionKey = v.GetSessionKey()
	c.id = v.GetId()

	return nil
}

func (c *CreateResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *session.CreateResponse

	if c != nil {
		m = new(session.CreateResponse)

		m.SetBody(c.body.ToGRPCMessage().(*session.CreateResponse_Body))
		c.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (c *CreateResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.CreateResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		c.body = nil
	} else {
		if c.body == nil {
			c.body = new(CreateResponseBody)
		}

		err = c.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return c.ResponseHeaders.FromMessage(v)
}

func (l *TokenLifetime) ToGRPCMessage() neofsgrpc.Message {
	var m *session.SessionToken_Body_TokenLifetime

	if l != nil {
		m = new(session.SessionToken_Body_TokenLifetime)

		m.SetExp(l.exp)
		m.SetIat(l.iat)
		m.SetNbf(l.nbf)
	}

	return m
}

func (l *TokenLifetime) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.SessionToken_Body_TokenLifetime)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	l.exp = v.GetExp()
	l.iat = v.GetIat()
	l.nbf = v.GetNbf()

	return nil
}

func (x *XHeader) ToGRPCMessage() neofsgrpc.Message {
	var m *session.XHeader

	if x != nil {
		m = new(session.XHeader)

		m.SetKey(x.key)
		m.SetValue(x.val)
	}

	return m
}

func (x *XHeader) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.XHeader)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	x.key = v.GetKey()
	x.val = v.GetValue()

	return nil
}

func XHeadersToGRPC(xs []*XHeader) (res []*session.XHeader) {
	if xs != nil {
		res = make([]*session.XHeader, 0, len(xs))

		for i := range xs {
			res = append(res, xs[i].ToGRPCMessage().(*session.XHeader))
		}
	}

	return
}

func XHeadersFromGRPC(xs []*session.XHeader) (res []*XHeader, err error) {
	if xs != nil {
		res = make([]*XHeader, 0, len(xs))

		for i := range xs {
			var x *XHeader

			if xs[i] != nil {
				x = new(XHeader)

				err = x.FromGRPCMessage(xs[i])
				if err != nil {
					return
				}
			}

			res = append(res, x)
		}
	}

	return
}

func (t *SessionToken) ToGRPCMessage() neofsgrpc.Message {
	var m *session.SessionToken

	if t != nil {
		m = new(session.SessionToken)

		m.SetBody(t.body.ToGRPCMessage().(*session.SessionToken_Body))
		m.SetSignature(t.sig.ToGRPCMessage().(*refsGRPC.Signature))
	}

	return m
}

func (t *SessionToken) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.SessionToken)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		t.body = nil
	} else {
		if t.body == nil {
			t.body = new(SessionTokenBody)
		}

		err = t.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	sig := v.GetSignature()
	if sig == nil {
		t.sig = nil
	} else {
		if t.sig == nil {
			t.sig = new(refs.Signature)
		}

		err = t.sig.FromGRPCMessage(sig)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RequestVerificationHeader) ToGRPCMessage() neofsgrpc.Message {
	var m *session.RequestVerificationHeader

	if r != nil {
		m = new(session.RequestVerificationHeader)

		m.SetBodySignature(r.bodySig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetMetaSignature(r.metaSig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetOriginSignature(r.originSig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetOrigin(r.origin.ToGRPCMessage().(*session.RequestVerificationHeader))
	}

	return m
}

func (r *RequestVerificationHeader) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.RequestVerificationHeader)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	originSig := v.GetOriginSignature()
	if originSig == nil {
		r.originSig = nil
	} else {
		if r.originSig == nil {
			r.originSig = new(refs.Signature)
		}

		err = r.originSig.FromGRPCMessage(originSig)
		if err != nil {
			return err
		}
	}

	metaSig := v.GetMetaSignature()
	if metaSig == nil {
		r.metaSig = nil
	} else {
		if r.metaSig == nil {
			r.metaSig = new(refs.Signature)
		}

		err = r.metaSig.FromGRPCMessage(metaSig)
		if err != nil {
			return err
		}
	}

	bodySig := v.GetBodySignature()
	if bodySig == nil {
		r.bodySig = nil
	} else {
		if r.bodySig == nil {
			r.bodySig = new(refs.Signature)
		}

		err = r.bodySig.FromGRPCMessage(bodySig)
		if err != nil {
			return err
		}
	}

	origin := v.GetOrigin()
	if origin == nil {
		r.origin = nil
	} else {
		if r.origin == nil {
			r.origin = new(RequestVerificationHeader)
		}

		err = r.origin.FromGRPCMessage(origin)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RequestMetaHeader) ToGRPCMessage() neofsgrpc.Message {
	var m *session.RequestMetaHeader

	if r != nil {
		m = new(session.RequestMetaHeader)

		m.SetVersion(r.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetSessionToken(r.sessionToken.ToGRPCMessage().(*session.SessionToken))
		m.SetBearerToken(r.bearerToken.ToGRPCMessage().(*aclGRPC.BearerToken))
		m.SetXHeaders(XHeadersToGRPC(r.xHeaders))
		m.SetEpoch(r.epoch)
		m.SetTtl(r.ttl)
		m.SetOrigin(r.origin.ToGRPCMessage().(*session.RequestMetaHeader))
	}

	return m
}

func (r *RequestMetaHeader) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.RequestMetaHeader)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	version := v.GetVersion()
	if version == nil {
		r.version = nil
	} else {
		if r.version == nil {
			r.version = new(refs.Version)
		}

		err = r.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	sessionToken := v.GetSessionToken()
	if sessionToken == nil {
		r.sessionToken = nil
	} else {
		if r.sessionToken == nil {
			r.sessionToken = new(SessionToken)
		}

		err = r.sessionToken.FromGRPCMessage(sessionToken)
		if err != nil {
			return err
		}
	}

	bearerToken := v.GetBearerToken()
	if bearerToken == nil {
		r.bearerToken = nil
	} else {
		if r.bearerToken == nil {
			r.bearerToken = new(acl.BearerToken)
		}

		err = r.bearerToken.FromGRPCMessage(bearerToken)
		if err != nil {
			return err
		}
	}

	origin := v.GetOrigin()
	if origin == nil {
		r.origin = nil
	} else {
		if r.origin == nil {
			r.origin = new(RequestMetaHeader)
		}

		err = r.origin.FromGRPCMessage(origin)
		if err != nil {
			return err
		}
	}

	r.xHeaders, err = XHeadersFromGRPC(v.GetXHeaders())
	if err != nil {
		return err
	}

	r.epoch = v.GetEpoch()
	r.ttl = v.GetTtl()

	return nil
}

func (r *ResponseVerificationHeader) ToGRPCMessage() neofsgrpc.Message {
	var m *session.ResponseVerificationHeader

	if r != nil {
		m = new(session.ResponseVerificationHeader)

		m.SetBodySignature(r.bodySig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetMetaSignature(r.metaSig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetOriginSignature(r.originSig.ToGRPCMessage().(*refsGRPC.Signature))
		m.SetOrigin(r.origin.ToGRPCMessage().(*session.ResponseVerificationHeader))
	}

	return m
}

func (r *ResponseVerificationHeader) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.ResponseVerificationHeader)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	originSig := v.GetOriginSignature()
	if originSig == nil {
		r.originSig = nil
	} else {
		if r.originSig == nil {
			r.originSig = new(refs.Signature)
		}

		err = r.originSig.FromGRPCMessage(originSig)
		if err != nil {
			return err
		}
	}

	metaSig := v.GetMetaSignature()
	if metaSig == nil {
		r.metaSig = nil
	} else {
		if r.metaSig == nil {
			r.metaSig = new(refs.Signature)
		}

		err = r.metaSig.FromGRPCMessage(metaSig)
		if err != nil {
			return err
		}
	}

	bodySig := v.GetBodySignature()
	if bodySig == nil {
		r.bodySig = nil
	} else {
		if r.bodySig == nil {
			r.bodySig = new(refs.Signature)
		}

		err = r.bodySig.FromGRPCMessage(bodySig)
		if err != nil {
			return err
		}
	}

	origin := v.GetOrigin()
	if origin == nil {
		r.origin = nil
	} else {
		if r.origin == nil {
			r.origin = new(ResponseVerificationHeader)
		}

		err = r.origin.FromGRPCMessage(origin)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ResponseMetaHeader) ToGRPCMessage() neofsgrpc.Message {
	var m *session.ResponseMetaHeader

	if r != nil {
		m = new(session.ResponseMetaHeader)

		m.SetVersion(r.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetXHeaders(XHeadersToGRPC(r.xHeaders))
		m.SetEpoch(r.epoch)
		m.SetTtl(r.ttl)
		m.SetOrigin(r.origin.ToGRPCMessage().(*session.ResponseMetaHeader))
	}

	return m
}

func (r *ResponseMetaHeader) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.ResponseMetaHeader)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	version := v.GetVersion()
	if version == nil {
		r.version = nil
	} else {
		if r.version == nil {
			r.version = new(refs.Version)
		}

		err = r.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	origin := v.GetOrigin()
	if origin == nil {
		r.origin = nil
	} else {
		if r.origin == nil {
			r.origin = new(ResponseMetaHeader)
		}

		err = r.origin.FromGRPCMessage(origin)
		if err != nil {
			return err
		}
	}

	r.xHeaders, err = XHeadersFromGRPC(v.GetXHeaders())
	if err != nil {
		return err
	}

	r.epoch = v.GetEpoch()
	r.ttl = v.GetTtl()

	return nil
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

func (c *ObjectSessionContext) ToGRPCMessage() neofsgrpc.Message {
	var m *session.ObjectSessionContext

	if c != nil {
		m = new(session.ObjectSessionContext)

		m.SetVerb(ObjectSessionVerbToGRPCField(c.verb))
		m.SetAddress(c.addr.ToGRPCMessage().(*refsGRPC.Address))
	}

	return m
}

func (c *ObjectSessionContext) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.ObjectSessionContext)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	addr := v.GetAddress()
	if addr == nil {
		c.addr = nil
	} else {
		if c.addr == nil {
			c.addr = new(refs.Address)
		}

		err = c.addr.FromGRPCMessage(addr)
		if err != nil {
			return err
		}
	}

	c.verb = ObjectSessionVerbFromGRPCField(v.GetVerb())

	return nil
}

func (t *SessionTokenBody) ToGRPCMessage() neofsgrpc.Message {
	var m *session.SessionToken_Body

	if t != nil {
		m = new(session.SessionToken_Body)

		switch typ := t.ctx.(type) {
		default:
			panic(fmt.Sprintf("unknown session context %T", typ))
		case nil:
			m.Context = nil
		case *ObjectSessionContext:
			m.SetObjectSessionContext(typ.ToGRPCMessage().(*session.ObjectSessionContext))
		case *ContainerSessionContext:
			m.SetContainerSessionContext(typ.ToGRPCMessage().(*session.ContainerSessionContext))
		}

		m.SetOwnerId(t.ownerID.ToGRPCMessage().(*refsGRPC.OwnerID))
		m.SetId(t.id)
		m.SetSessionKey(t.sessionKey)
		m.SetLifetime(t.lifetime.ToGRPCMessage().(*session.SessionToken_Body_TokenLifetime))
	}

	return m
}

func (t *SessionTokenBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.SessionToken_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	t.ctx = nil

	switch val := v.GetContext().(type) {
	default:
		err = fmt.Errorf("unknown session context %T", val)
	case nil:
	case *session.SessionToken_Body_Object:
		ctx, ok := t.ctx.(*ObjectSessionContext)
		if !ok {
			ctx = new(ObjectSessionContext)
			t.ctx = ctx
		}

		err = ctx.FromGRPCMessage(val.Object)
	case *session.SessionToken_Body_Container:
		ctx, ok := t.ctx.(*ContainerSessionContext)
		if !ok {
			ctx = new(ContainerSessionContext)
			t.ctx = ctx
		}

		err = ctx.FromGRPCMessage(val.Container)
	}

	if err != nil {
		return err
	}

	ownerID := v.GetOwnerId()
	if ownerID == nil {
		t.ownerID = nil
	} else {
		if t.ownerID == nil {
			t.ownerID = new(refs.OwnerID)
		}

		err = t.ownerID.FromGRPCMessage(ownerID)
		if err != nil {
			return err
		}
	}

	lifetime := v.GetLifetime()
	if lifetime == nil {
		t.lifetime = nil
	} else {
		if t.lifetime == nil {
			t.lifetime = new(TokenLifetime)
		}

		err = t.lifetime.FromGRPCMessage(lifetime)
		if err != nil {
			return err
		}
	}

	t.id = v.GetId()
	t.sessionKey = v.GetSessionKey()

	return nil
}

// ContainerSessionVerbToGRPCField converts ContainerSessionVerb
// to gRPC-generated session.ContainerSessionContext_Verb.
//
// If v is outside of the ContainerSessionVerb enum,
// session.ContainerSessionContext_VERB_UNSPECIFIED is returned.
func ContainerSessionVerbToGRPCField(v ContainerSessionVerb) session.ContainerSessionContext_Verb {
	switch v {
	default:
		return session.ContainerSessionContext_VERB_UNSPECIFIED
	case ContainerVerbPut:
		return session.ContainerSessionContext_PUT
	case ContainerVerbDelete:
		return session.ContainerSessionContext_DELETE
	case ContainerVerbSetEACL:
		return session.ContainerSessionContext_SETEACL
	}
}

// ContainerSessionVerbFromGRPCField converts gRPC-generated
// session.ContainerSessionContext_Verb to ContainerSessionVerb.
//
// If v is outside of the session.ContainerSessionContext_Verb enum,
// ContainerVerbUnknown is returned.
func ContainerSessionVerbFromGRPCField(v session.ContainerSessionContext_Verb) ContainerSessionVerb {
	switch v {
	default:
		return ContainerVerbUnknown
	case session.ContainerSessionContext_PUT:
		return ContainerVerbPut
	case session.ContainerSessionContext_DELETE:
		return ContainerVerbDelete
	case session.ContainerSessionContext_SETEACL:
		return ContainerVerbSetEACL
	}
}

// ToGRPCMessage converts ContainerSessionContext to gRPC-generated
// session.ContainerSessionContext message.
func (x *ContainerSessionContext) ToGRPCMessage() neofsgrpc.Message {
	var m *session.ContainerSessionContext

	if x != nil {
		m = new(session.ContainerSessionContext)

		m.SetVerb(ContainerSessionVerbToGRPCField(x.verb))
		m.SetWildcard(x.wildcard)
		m.SetContainerId(x.cid.ToGRPCMessage().(*refsGRPC.ContainerID))
	}

	return m
}

// FromGRPCMessage tries to restore ContainerSessionContext from grpc.Message.
//
// Returns message.ErrUnexpectedMessageType if m is not
// a gRPC-generated session.ContainerSessionContext message.
func (x *ContainerSessionContext) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*session.ContainerSessionContext)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	cid := v.GetContainerId()
	if cid == nil {
		x.cid = nil
	} else {
		if x.cid == nil {
			x.cid = new(refs.ContainerID)
		}

		err = x.cid.FromGRPCMessage(cid)
		if err != nil {
			return err
		}
	}

	x.verb = ContainerSessionVerbFromGRPCField(v.GetVerb())
	x.wildcard = v.GetWildcard()

	return nil
}
