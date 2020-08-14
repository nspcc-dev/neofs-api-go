package v2

type Signature struct {
	key, sign []byte
}

type Version struct {
	major, minor uint32
}

type XHeader struct {
	key, val string
}

type SessionToken struct {
	// TODO: fill me
}

type BearerToken struct {
	// TODO: fill me
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

type OwnerID struct {
	val []byte
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

func (o *OwnerID) GetValue() []byte {
	if o != nil {
		return o.val
	}

	return nil
}

func (o *OwnerID) SetValue(v []byte) {
	if o != nil {
		o.val = v
	}
}
