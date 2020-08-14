package service

import (
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
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

// SetMajor sets major version number.
func (m *Version) SetMajor(v uint32) {
	if m != nil {
		m.Major = v
	}
}

// SetMinor sets minor version number.
func (m *Version) SetMinor(v uint32) {
	if m != nil {
		m.Minor = v
	}
}

// SetExp sets epoch number of the token expiration.
func (m *TokenLifetime) SetExp(v uint64) {
	if m != nil {
		m.Exp = v
	}
}

// SetNbf sets starting epoch number of the token.
func (m *TokenLifetime) SetNbf(v uint64) {
	if m != nil {
		m.Nbf = v
	}
}

// SetIat sets the number of the epoch in which the token was issued.
func (m *TokenLifetime) SetIat(v uint64) {
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

// SetVerb sets verb of the session token.
func (m *SessionToken_Body) SetVerb(v SessionToken_Body_Verb) {
	if m != nil {
		m.Verb = v
	}
}

// SetLifetime sets lifetime of the session token.
func (m *SessionToken_Body) SetLifetime(v *TokenLifetime) {
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
func (m *SessionToken_Body) SetObjectAddressContext(v *refs.Address) {
	if m != nil {
		m.Context = &SessionToken_Body_ObjectAddress{
			ObjectAddress: v,
		}
	}
}

// SetBody sets session token body.
func (m *SessionToken) SetBody(v *SessionToken_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetSignature sets session token signature.
func (m *SessionToken) SetSignature(v *Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetEaclTable sets eACL table of the bearer token.
func (m *BearerToken_Body) SetEaclTable(v *acl.EACLTable) {
	if m != nil {
		m.EaclTable = v
	}
}

// SetOwnerId sets identifier of the bearer token owner.
func (m *BearerToken_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetLifetime sets lifetime of the bearer token.
func (m *BearerToken_Body) SetLifetime(v *TokenLifetime) {
	if m != nil {
		m.Lifetime = v
	}
}

// SetBody sets bearer token body.
func (m *BearerToken) SetBody(v *BearerToken_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetSignature sets bearer token signature.
func (m *BearerToken) SetSignature(v *Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetVersion sets client protocol version.
func (m *RequestMetaHeader) SetVersion(v *Version) {
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
func (m *RequestMetaHeader) SetBearerToken(v *BearerToken) {
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
func (m *ResponseMetaHeader) SetVersion(v *Version) {
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
