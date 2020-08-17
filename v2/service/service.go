package service

import (
	"github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

type Signature struct {
	key, sign []byte
}

type Version struct {
	major, minor uint32
}

type XHeader struct {
	key, val string
}

type TokenLifetime struct {
	exp, nbf, iat uint64
}

type SessionToken struct {
	// TODO: fill me
}

type BearerTokenBody struct {
	eacl *acl.Table

	ownerID *refs.OwnerID

	lifetime *TokenLifetime
}

type BearerToken struct {
	body *BearerTokenBody

	sig *Signature
}

type RequestVerificationHeader struct {
	bodySig, metaSig, originSig *Signature

	origin *RequestVerificationHeader
}

type RequestMetaHeader struct {
	version *Version

	ttl uint32

	epoch uint64

	xHeaders []*XHeader

	sessionToken *SessionToken

	bearerToken *BearerToken

	origin *RequestMetaHeader
}

type ResponseVerificationHeader struct {
	bodySig, metaSig, originSig *Signature

	origin *ResponseVerificationHeader
}

type ResponseMetaHeader struct {
	version *Version

	ttl uint32

	epoch uint64

	xHeaders []*XHeader

	origin *ResponseMetaHeader
}

func (s *Signature) GetKey() []byte {
	if s != nil {
		return s.key
	}

	return nil
}

func (s *Signature) SetKey(v []byte) {
	if s != nil {
		s.key = v
	}
}

func (s *Signature) GetSign() []byte {
	if s != nil {
		return s.sign
	}

	return nil
}

func (s *Signature) SetSign(v []byte) {
	if s != nil {
		s.sign = v
	}
}

func (v *Version) GetMajor() uint32 {
	if v != nil {
		return v.major
	}

	return 0
}

func (v *Version) SetMajor(val uint32) {
	if v != nil {
		v.major = val
	}
}

func (v *Version) GetMinor() uint32 {
	if v != nil {
		return v.minor
	}

	return 0
}

func (v *Version) SetMinor(val uint32) {
	if v != nil {
		v.minor = val
	}
}

func (x *XHeader) GetKey() string {
	if x != nil {
		return x.key
	}

	return ""
}

func (x *XHeader) SetKey(v string) {
	if x != nil {
		x.key = v
	}
}

func (x *XHeader) GetValue() string {
	if x != nil {
		return x.val
	}

	return ""
}

func (x *XHeader) SetValue(v string) {
	if x != nil {
		x.val = v
	}
}

func (r *RequestVerificationHeader) GetBodySignature() *Signature {
	if r != nil {
		return r.bodySig
	}

	return nil
}

func (r *RequestVerificationHeader) SetBodySignature(v *Signature) {
	if r != nil {
		r.bodySig = v
	}
}

func (r *RequestVerificationHeader) GetMetaSignature() *Signature {
	if r != nil {
		return r.metaSig
	}

	return nil
}

func (r *RequestVerificationHeader) SetMetaSignature(v *Signature) {
	if r != nil {
		r.metaSig = v
	}
}

func (r *RequestVerificationHeader) GetOriginSignature() *Signature {
	if r != nil {
		return r.originSig
	}

	return nil
}

func (r *RequestVerificationHeader) SetOriginSignature(v *Signature) {
	if r != nil {
		r.originSig = v
	}
}

func (r *RequestVerificationHeader) GetOrigin() *RequestVerificationHeader {
	if r != nil {
		return r.origin
	}

	return nil
}

func (r *RequestVerificationHeader) SetOrigin(v *RequestVerificationHeader) {
	if r != nil {
		r.origin = v
	}
}

func (r *RequestVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	// TODO: do not use hack
	_, err := RequestVerificationHeaderToGRPCMessage(r).MarshalTo(buf)

	return buf, err
}

func (r *RequestVerificationHeader) StableSize() int {
	// TODO: do not use hack
	return RequestVerificationHeaderToGRPCMessage(r).Size()
}

func (r *RequestMetaHeader) GetVersion() *Version {
	if r != nil {
		return r.version
	}

	return nil
}

func (r *RequestMetaHeader) SetVersion(v *Version) {
	if r != nil {
		r.version = v
	}
}

func (r *RequestMetaHeader) GetTTL() uint32 {
	if r != nil {
		return r.ttl
	}

	return 0
}

func (r *RequestMetaHeader) SetTTL(v uint32) {
	if r != nil {
		r.ttl = v
	}
}

func (r *RequestMetaHeader) GetEpoch() uint64 {
	if r != nil {
		return r.epoch
	}

	return 0
}

func (r *RequestMetaHeader) SetEpoch(v uint64) {
	if r != nil {
		r.epoch = v
	}
}

func (r *RequestMetaHeader) GetXHeaders() []*XHeader {
	if r != nil {
		return r.xHeaders
	}

	return nil
}

func (r *RequestMetaHeader) SetXHeaders(v []*XHeader) {
	if r != nil {
		r.xHeaders = v
	}
}

func (r *RequestMetaHeader) GetSessionToken() *SessionToken {
	if r != nil {
		return r.sessionToken
	}

	return nil
}

func (r *RequestMetaHeader) SetSessionToken(v *SessionToken) {
	if r != nil {
		r.sessionToken = v
	}
}

func (r *RequestMetaHeader) GetBearerToken() *BearerToken {
	if r != nil {
		return r.bearerToken
	}

	return nil
}

func (r *RequestMetaHeader) SetBearerToken(v *BearerToken) {
	if r != nil {
		r.bearerToken = v
	}
}

func (r *RequestMetaHeader) GetOrigin() *RequestMetaHeader {
	if r != nil {
		return r.origin
	}

	return nil
}

func (r *RequestMetaHeader) SetOrigin(v *RequestMetaHeader) {
	if r != nil {
		r.origin = v
	}
}

func (r *RequestMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	// TODO: do not use hack
	_, err := RequestMetaHeaderToGRPCMessage(r).MarshalTo(buf)

	return buf, err
}

func (r *RequestMetaHeader) StableSize() int {
	// TODO: do not use hack
	return RequestMetaHeaderToGRPCMessage(r).Size()
}

func (tl *TokenLifetime) GetExp() uint64 {
	if tl != nil {
		return tl.exp
	}

	return 0
}

func (tl *TokenLifetime) SetExp(v uint64) {
	if tl != nil {
		tl.exp = v
	}
}

func (tl *TokenLifetime) GetNbf() uint64 {
	if tl != nil {
		return tl.nbf
	}

	return 0
}

func (tl *TokenLifetime) SetNbf(v uint64) {
	if tl != nil {
		tl.nbf = v
	}
}

func (tl *TokenLifetime) GetIat() uint64 {
	if tl != nil {
		return tl.iat
	}

	return 0
}

func (tl *TokenLifetime) SetIat(v uint64) {
	if tl != nil {
		tl.iat = v
	}
}

func (bt *BearerTokenBody) GetEACL() *acl.Table {
	if bt != nil {
		return bt.eacl
	}

	return nil
}

func (bt *BearerTokenBody) SetEACL(v *acl.Table) {
	if bt != nil {
		bt.eacl = v
	}
}

func (bt *BearerTokenBody) GetOwnerID() *refs.OwnerID {
	if bt != nil {
		return bt.ownerID
	}

	return nil
}

func (bt *BearerTokenBody) SetOwnerID(v *refs.OwnerID) {
	if bt != nil {
		bt.ownerID = v
	}
}

func (bt *BearerTokenBody) GetLifetime() *TokenLifetime {
	if bt != nil {
		return bt.lifetime
	}

	return nil
}

func (bt *BearerTokenBody) SetLifetime(v *TokenLifetime) {
	if bt != nil {
		bt.lifetime = v
	}
}

func (bt *BearerToken) GetBody() *BearerTokenBody {
	if bt != nil {
		return bt.body
	}

	return nil
}

func (bt *BearerToken) SetBody(v *BearerTokenBody) {
	if bt != nil {
		bt.body = v
	}
}

func (bt *BearerToken) GetSignature() *Signature {
	if bt != nil {
		return bt.sig
	}

	return nil
}

func (bt *BearerToken) SetSignature(v *Signature) {
	if bt != nil {
		bt.sig = v
	}
}

func (r *ResponseVerificationHeader) GetBodySignature() *Signature {
	if r != nil {
		return r.bodySig
	}

	return nil
}

func (r *ResponseVerificationHeader) SetBodySignature(v *Signature) {
	if r != nil {
		r.bodySig = v
	}
}

func (r *ResponseVerificationHeader) GetMetaSignature() *Signature {
	if r != nil {
		return r.metaSig
	}

	return nil
}

func (r *ResponseVerificationHeader) SetMetaSignature(v *Signature) {
	if r != nil {
		r.metaSig = v
	}
}

func (r *ResponseVerificationHeader) GetOriginSignature() *Signature {
	if r != nil {
		return r.originSig
	}

	return nil
}

func (r *ResponseVerificationHeader) SetOriginSignature(v *Signature) {
	if r != nil {
		r.originSig = v
	}
}

func (r *ResponseVerificationHeader) GetOrigin() *ResponseVerificationHeader {
	if r != nil {
		return r.origin
	}

	return nil
}

func (r *ResponseVerificationHeader) SetOrigin(v *ResponseVerificationHeader) {
	if r != nil {
		r.origin = v
	}
}

func (r *ResponseMetaHeader) GetVersion() *Version {
	if r != nil {
		return r.version
	}

	return nil
}

func (r *ResponseMetaHeader) SetVersion(v *Version) {
	if r != nil {
		r.version = v
	}
}

func (r *ResponseMetaHeader) GetTTL() uint32 {
	if r != nil {
		return r.ttl
	}

	return 0
}

func (r *ResponseMetaHeader) SetTTL(v uint32) {
	if r != nil {
		r.ttl = v
	}
}

func (r *ResponseMetaHeader) GetEpoch() uint64 {
	if r != nil {
		return r.epoch
	}

	return 0
}

func (r *ResponseMetaHeader) SetEpoch(v uint64) {
	if r != nil {
		r.epoch = v
	}
}

func (r *ResponseMetaHeader) GetXHeaders() []*XHeader {
	if r != nil {
		return r.xHeaders
	}

	return nil
}

func (r *ResponseMetaHeader) SetXHeaders(v []*XHeader) {
	if r != nil {
		r.xHeaders = v
	}
}

func (r *ResponseMetaHeader) GetOrigin() *ResponseMetaHeader {
	if r != nil {
		return r.origin
	}

	return nil
}

func (r *ResponseMetaHeader) SetOrigin(v *ResponseMetaHeader) {
	if r != nil {
		r.origin = v
	}
}
