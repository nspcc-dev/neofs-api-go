package session

import (
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetKey sets key to the X-Header.
func (m *XHeader) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetValue sets value of the X-Header.
func (m *XHeader) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetExp sets epoch number of the token expiration.
func (m *SessionToken_Body_TokenLifetime) SetExp(v uint64) {
	if m != nil {
		m.Exp = v
	}
}

// SetNbf sets starting epoch number of the token.
func (m *SessionToken_Body_TokenLifetime) SetNbf(v uint64) {
	if m != nil {
		m.Nbf = v
	}
}

// SetIat sets the number of the epoch in which the token was issued.
func (m *SessionToken_Body_TokenLifetime) SetIat(v uint64) {
	if m != nil {
		m.Iat = v
	}
}

// SetId sets identifier of the session token.
func (m *SessionToken_Body) SetId(v []byte) {
	if m != nil {
		m.Id = v
	}
}

// SetOwnerId sets identifier of the session token owner.
func (m *SessionToken_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetLifetime sets lifetime of the session token.
func (m *SessionToken_Body) SetLifetime(v *SessionToken_Body_TokenLifetime) {
	if m != nil {
		m.Lifetime = v
	}
}

// SetSessionKey sets public session key in a binary format.
func (m *SessionToken_Body) SetSessionKey(v []byte) {
	if m != nil {
		m.SessionKey = v
	}
}

// SetObjectAddressContext sets object context of the session token.
func (m *SessionToken_Body) SetObjectSessionContext(v *ObjectSessionContext) {
	if m != nil {
		m.Context = &SessionToken_Body_Object{
			Object: v,
		}
	}
}

// SetContainerSessionContext sets container context of the session token.
func (m *SessionToken_Body) SetContainerSessionContext(v *ContainerSessionContext) {
	if m != nil {
		m.Context = &SessionToken_Body_Container{
			Container: v,
		}
	}
}

// SetAddress sets address of the object related to the session.
func (m *ObjectSessionContext) SetAddress(v *refs.Address) {
	if m != nil {
		m.Address = v
	}
}

// SetVerb sets type of request for which the token is issued.
func (m *ObjectSessionContext) SetVerb(v ObjectSessionContext_Verb) {
	if m != nil {
		m.Verb = v
	}
}

// SetVerb sets type of request for which the token is issued.
func (x *ContainerSessionContext) SetVerb(v ContainerSessionContext_Verb) {
	if x != nil {
		x.Verb = v
	}
}

// SetWildcard sets wildcard flag of the container session.
func (x *ContainerSessionContext) SetWildcard(v bool) {
	if x != nil {
		x.Wildcard = v
	}
}

// SetContainerId sets identifier of the container related to the session.
func (x *ContainerSessionContext) SetContainerId(v *refs.ContainerID) {
	if x != nil {
		x.ContainerId = v
	}
}

// SetBody sets session token body.
func (m *SessionToken) SetBody(v *SessionToken_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetSignature sets session token signature.
func (m *SessionToken) SetSignature(v *refs.Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetVersion sets client protocol version.
func (m *RequestMetaHeader) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

// SetEpoch sets client local epoch.
func (m *RequestMetaHeader) SetEpoch(v uint64) {
	if m != nil {
		m.Epoch = v
	}
}

// SetTtl sets request TTL.
func (m *RequestMetaHeader) SetTtl(v uint32) {
	if m != nil {
		m.Ttl = v
	}
}

// SetXHeaders sets request X-Headers.
func (m *RequestMetaHeader) SetXHeaders(v []*XHeader) {
	if m != nil {
		m.XHeaders = v
	}
}

// SetSessionToken sets session token of the request.
func (m *RequestMetaHeader) SetSessionToken(v *SessionToken) {
	if m != nil {
		m.SessionToken = v
	}
}

// SetBearerToken sets bearer token of the request.
func (m *RequestMetaHeader) SetBearerToken(v *acl.BearerToken) {
	if m != nil {
		m.BearerToken = v
	}
}

// SetOrigin sets origin request meta header.
func (m *RequestMetaHeader) SetOrigin(v *RequestMetaHeader) {
	if m != nil {
		m.Origin = v
	}
}

// SetVersion sets server protocol version.
func (m *ResponseMetaHeader) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

// SetEpoch sets server local epoch.
func (m *ResponseMetaHeader) SetEpoch(v uint64) {
	if m != nil {
		m.Epoch = v
	}
}

// SetTtl sets response TTL.
func (m *ResponseMetaHeader) SetTtl(v uint32) {
	if m != nil {
		m.Ttl = v
	}
}

// SetXHeaders sets response X-Headers.
func (m *ResponseMetaHeader) SetXHeaders(v []*XHeader) {
	if m != nil {
		m.XHeaders = v
	}
}

// SetOrigin sets origin response meta header.
func (m *ResponseMetaHeader) SetOrigin(v *ResponseMetaHeader) {
	if m != nil {
		m.Origin = v
	}
}

// SetBodySignature sets signature of the request body.
func (m *RequestVerificationHeader) SetBodySignature(v *refs.Signature) {
	if m != nil {
		m.BodySignature = v
	}
}

// SetMetaSignature sets signature of the request meta.
func (m *RequestVerificationHeader) SetMetaSignature(v *refs.Signature) {
	if m != nil {
		m.MetaSignature = v
	}
}

// SetOriginSignature sets signature of the origin verification header of the request.
func (m *RequestVerificationHeader) SetOriginSignature(v *refs.Signature) {
	if m != nil {
		m.OriginSignature = v
	}
}

// SetOrigin sets origin verification header of the request.
func (m *RequestVerificationHeader) SetOrigin(v *RequestVerificationHeader) {
	if m != nil {
		m.Origin = v
	}
}

// SetBodySignature sets signature of the response body.
func (m *ResponseVerificationHeader) SetBodySignature(v *refs.Signature) {
	if m != nil {
		m.BodySignature = v
	}
}

// SetMetaSignature sets signature of the response meta.
func (m *ResponseVerificationHeader) SetMetaSignature(v *refs.Signature) {
	if m != nil {
		m.MetaSignature = v
	}
}

// SetOriginSignature sets signature of the origin verification header of the response.
func (m *ResponseVerificationHeader) SetOriginSignature(v *refs.Signature) {
	if m != nil {
		m.OriginSignature = v
	}
}

// SetOrigin sets origin verification header of the response.
func (m *ResponseVerificationHeader) SetOrigin(v *ResponseVerificationHeader) {
	if m != nil {
		m.Origin = v
	}
}
