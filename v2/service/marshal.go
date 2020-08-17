package service

import (
	"encoding/binary"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	signatureKeyField   = 1
	signatureValueField = 2

	versionMajorField = 1
	versionMinorField = 2

	xheaderKeyField   = 1
	xheaderValueField = 2

	lifetimeExpirationField     = 1
	lifetimeNotValidBeforeField = 2
	lifetimeIssuedAtField       = 3

	objectCtxVerbField    = 1
	objectCtxAddressField = 2

	sessionTokenBodyIDField        = 1
	sessionTokenBodyOwnerField     = 2
	sessionTokenBodyLifetimeField  = 3
	sessionTokenBodyKeyField       = 4
	sessionTokenBodyObjectCtxField = 5

	sessionTokenBodyField      = 1
	sessionTokenSignatureField = 2

	bearerTokenBodyACLField      = 1
	bearerTokenBodyOwnerField    = 2
	bearerTokenBodyLifetimeField = 3

	bearerTokenBodyField      = 1
	bearerTokenSignatureField = 2

	reqMetaHeaderVersionField      = 1
	reqMetaHeaderEpochField        = 2
	reqMetaHeaderTTLField          = 3
	reqMetaHeaderXHeadersField     = 4
	reqMetaHeaderSessionTokenField = 5
	reqMetaHeaderBearerTokenField  = 6
	reqMetaHeaderOriginField       = 7

	reqVerifHeaderBodySignatureField   = 1
	reqVerifHeaderMetaSignatureField   = 2
	reqVerifHeaderOriginSignatureField = 3
	reqVerifHeaderOriginField          = 4

	respMetaHeaderVersionField  = 1
	respMetaHeaderEpochField    = 2
	respMetaHeaderTTLField      = 3
	respMetaHeaderXHeadersField = 4
	respMetaHeaderOriginField   = 5

	respVerifHeaderBodySignatureField   = 1
	respVerifHeaderMetaSignatureField   = 2
	respVerifHeaderOriginSignatureField = 3
	respVerifHeaderOriginField          = 4
)

func (s *Signature) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.BytesMarshal(signatureKeyField, buf, s.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.BytesMarshal(signatureValueField, buf[offset:], s.sign)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Signature) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.BytesSize(signatureKeyField, s.key)
	size += proto.BytesSize(signatureValueField, s.sign)

	return size
}

func (v *Version) StableMarshal(buf []byte) ([]byte, error) {
	if v == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, v.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt32Marshal(versionMajorField, buf, v.major)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt32Marshal(versionMinorField, buf[offset:], v.minor)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (v *Version) StableSize() (size int) {
	if v == nil {
		return 0
	}

	size += proto.UInt32Size(versionMajorField, v.major)
	size += proto.UInt32Size(versionMinorField, v.minor)

	return size
}

func (x *XHeader) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.StringMarshal(xheaderKeyField, buf, x.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.StringMarshal(xheaderValueField, buf[offset:], x.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *XHeader) StableSize() (size int) {
	if x == nil {
		return 0
	}

	size += proto.StringSize(xheaderKeyField, x.key)
	size += proto.StringSize(xheaderValueField, x.val)

	return size
}

func (l *TokenLifetime) StableMarshal(buf []byte) ([]byte, error) {
	if l == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, l.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt64Marshal(lifetimeExpirationField, buf, l.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(lifetimeNotValidBeforeField, buf[offset:], l.nbf)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt64Marshal(lifetimeIssuedAtField, buf[offset:], l.iat)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (l *TokenLifetime) StableSize() (size int) {
	if l == nil {
		return 0
	}

	size += proto.UInt64Size(lifetimeExpirationField, l.exp)
	size += proto.UInt64Size(lifetimeNotValidBeforeField, l.nbf)
	size += proto.UInt64Size(lifetimeIssuedAtField, l.iat)

	return size
}

func (c *ObjectSessionContext) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	n, err = proto.EnumMarshal(objectCtxVerbField, buf, int32(c.verb))
	if err != nil {
		return nil, err
	}

	offset += n

	if c.addr != nil {
		prefix, _ = proto.NestedStructurePrefix(objectCtxAddressField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = c.addr.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = c.addr.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (c *ObjectSessionContext) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.EnumSize(objectCtxVerbField, int32(c.verb))

	if c.addr != nil {
		_, ln := proto.NestedStructurePrefix(objectCtxAddressField)
		n := c.addr.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (t *SessionTokenBody) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	n, err = proto.BytesMarshal(sessionTokenBodyIDField, buf, t.id)
	if err != nil {
		return nil, err
	}

	offset += n

	if t.ownerID != nil {
		prefix, _ = proto.NestedStructurePrefix(sessionTokenBodyOwnerField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.ownerID.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.ownerID.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if t.lifetime != nil {
		prefix, _ = proto.NestedStructurePrefix(sessionTokenBodyLifetimeField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.lifetime.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.lifetime.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	n, err = proto.BytesMarshal(sessionTokenBodyKeyField, buf[offset:], t.sessionKey)
	if err != nil {
		return nil, err
	}

	offset += n

	if t.ctx != nil {
		switch v := t.ctx.(type) {
		case *ObjectSessionContext:
			prefix, _ = proto.NestedStructurePrefix(sessionTokenBodyObjectCtxField)
			offset += binary.PutUvarint(buf[offset:], prefix)

			n = v.StableSize()
			offset += binary.PutUvarint(buf[offset:], uint64(n))

			_, err = v.StableMarshal(buf[offset:])
			if err != nil {
				return nil, err

			}
		default:
			panic("cannot marshal unknown session token context")
		}
	}

	return buf, nil
}

func (t *SessionTokenBody) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.BytesSize(sessionTokenBodyIDField, t.id)

	if t.ownerID != nil {
		_, ln := proto.NestedStructurePrefix(sessionTokenBodyOwnerField)
		n := t.ownerID.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if t.lifetime != nil {
		_, ln := proto.NestedStructurePrefix(sessionTokenBodyLifetimeField)
		n := t.lifetime.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	size += proto.BytesSize(sessionTokenBodyKeyField, t.sessionKey)

	if t.ctx != nil {
		switch v := t.ctx.(type) {
		case *ObjectSessionContext:
			_, ln := proto.NestedStructurePrefix(sessionTokenBodyObjectCtxField)
			n := v.StableSize()
			size += ln + proto.VarUIntSize(uint64(n)) + n
		default:
			panic("cannot marshal unknown session token context")
		}
	}

	return size
}

func (t *SessionToken) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if t.body != nil {
		prefix, _ = proto.NestedStructurePrefix(sessionTokenBodyField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.body.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.body.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if t.sig != nil {
		prefix, _ = proto.NestedStructurePrefix(sessionTokenSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.sig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.sig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (t *SessionToken) StableSize() (size int) {
	if t == nil {
		return 0
	}

	if t.body != nil {
		_, ln := proto.NestedStructurePrefix(sessionTokenBodyField)
		n := t.body.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if t.sig != nil {
		_, ln := proto.NestedStructurePrefix(sessionTokenSignatureField)
		n := t.sig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (bt *BearerTokenBody) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if bt.eacl != nil {
		prefix, _ = proto.NestedStructurePrefix(bearerTokenBodyACLField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = bt.eacl.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = bt.eacl.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if bt.ownerID != nil {
		prefix, _ = proto.NestedStructurePrefix(bearerTokenBodyOwnerField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = bt.ownerID.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = bt.ownerID.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if bt.lifetime != nil {
		prefix, _ = proto.NestedStructurePrefix(bearerTokenBodyLifetimeField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = bt.lifetime.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = bt.lifetime.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (bt *BearerTokenBody) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	if bt.eacl != nil {
		_, ln := proto.NestedStructurePrefix(bearerTokenBodyACLField)
		n := bt.eacl.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if bt.ownerID != nil {
		_, ln := proto.NestedStructurePrefix(bearerTokenBodyOwnerField)
		n := bt.ownerID.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if bt.lifetime != nil {
		_, ln := proto.NestedStructurePrefix(bearerTokenBodyLifetimeField)
		n := bt.lifetime.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (bt *BearerToken) StableMarshal(buf []byte) ([]byte, error) {
	if bt == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, bt.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if bt.body != nil {
		prefix, _ = proto.NestedStructurePrefix(bearerTokenBodyField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = bt.body.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = bt.body.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if bt.sig != nil {
		prefix, _ = proto.NestedStructurePrefix(bearerTokenSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = bt.sig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = bt.sig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (bt *BearerToken) StableSize() (size int) {
	if bt == nil {
		return 0
	}

	if bt.body != nil {
		_, ln := proto.NestedStructurePrefix(bearerTokenBodyField)
		n := bt.body.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if bt.sig != nil {
		_, ln := proto.NestedStructurePrefix(bearerTokenSignatureField)
		n := bt.sig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (r *RequestMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if r.version != nil {
		prefix, _ = proto.NestedStructurePrefix(reqMetaHeaderVersionField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.version.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.version.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	n, err = proto.UInt64Marshal(reqMetaHeaderEpochField, buf[offset:], r.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(reqMetaHeaderTTLField, buf[offset:], r.ttl)
	if err != nil {
		return nil, err
	}

	offset += n

	prefix, _ = proto.NestedStructurePrefix(reqMetaHeaderXHeadersField)

	for i := range r.xHeaders {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.xHeaders[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.xHeaders[i].StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.sessionToken != nil {
		prefix, _ = proto.NestedStructurePrefix(reqMetaHeaderSessionTokenField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.sessionToken.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.sessionToken.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.bearerToken != nil {
		prefix, _ = proto.NestedStructurePrefix(reqMetaHeaderBearerTokenField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.bearerToken.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.bearerToken.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.origin != nil {
		prefix, _ = proto.NestedStructurePrefix(reqMetaHeaderOriginField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.origin.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.origin.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (r *RequestMetaHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.version != nil {
		_, ln := proto.NestedStructurePrefix(reqMetaHeaderVersionField)
		n := r.version.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	size += proto.UInt64Size(reqMetaHeaderEpochField, r.epoch)
	size += proto.UInt32Size(reqMetaHeaderTTLField, r.ttl)

	_, ln := proto.NestedStructurePrefix(reqMetaHeaderXHeadersField)

	for i := range r.xHeaders {
		n := r.xHeaders[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.sessionToken != nil {
		_, ln := proto.NestedStructurePrefix(reqMetaHeaderSessionTokenField)
		n := r.sessionToken.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.bearerToken != nil {
		_, ln := proto.NestedStructurePrefix(reqMetaHeaderBearerTokenField)
		n := r.bearerToken.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.origin != nil {
		_, ln := proto.NestedStructurePrefix(reqMetaHeaderOriginField)
		n := r.origin.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (r *RequestVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if r.bodySig != nil {
		prefix, _ = proto.NestedStructurePrefix(reqVerifHeaderBodySignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.bodySig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.bodySig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.metaSig != nil {
		prefix, _ = proto.NestedStructurePrefix(reqVerifHeaderMetaSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.metaSig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.metaSig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.originSig != nil {
		prefix, _ = proto.NestedStructurePrefix(reqVerifHeaderOriginSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.originSig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.originSig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.origin != nil {
		prefix, _ = proto.NestedStructurePrefix(reqVerifHeaderOriginField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.origin.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.origin.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (r *RequestVerificationHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.bodySig != nil {
		_, ln := proto.NestedStructurePrefix(reqVerifHeaderBodySignatureField)
		n := r.bodySig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.metaSig != nil {
		_, ln := proto.NestedStructurePrefix(reqVerifHeaderMetaSignatureField)
		n := r.metaSig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.originSig != nil {
		_, ln := proto.NestedStructurePrefix(reqVerifHeaderOriginSignatureField)
		n := r.originSig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.origin != nil {
		_, ln := proto.NestedStructurePrefix(reqVerifHeaderOriginField)
		n := r.origin.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (r *ResponseMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if r.version != nil {
		prefix, _ = proto.NestedStructurePrefix(respMetaHeaderVersionField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.version.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.version.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	n, err = proto.UInt64Marshal(respMetaHeaderEpochField, buf[offset:], r.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(respMetaHeaderTTLField, buf[offset:], r.ttl)
	if err != nil {
		return nil, err
	}

	offset += n

	prefix, _ = proto.NestedStructurePrefix(respMetaHeaderXHeadersField)

	for i := range r.xHeaders {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.xHeaders[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.xHeaders[i].StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.origin != nil {
		prefix, _ = proto.NestedStructurePrefix(respMetaHeaderOriginField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.origin.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.origin.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (r *ResponseMetaHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.version != nil {
		_, ln := proto.NestedStructurePrefix(respMetaHeaderVersionField)
		n := r.version.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	size += proto.UInt64Size(respMetaHeaderEpochField, r.epoch)
	size += proto.UInt32Size(respMetaHeaderTTLField, r.ttl)

	_, ln := proto.NestedStructurePrefix(respMetaHeaderXHeadersField)

	for i := range r.xHeaders {
		n := r.xHeaders[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.origin != nil {
		_, ln := proto.NestedStructurePrefix(respMetaHeaderOriginField)
		n := r.origin.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

func (r *ResponseVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if r.bodySig != nil {
		prefix, _ = proto.NestedStructurePrefix(respVerifHeaderBodySignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.bodySig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.bodySig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.metaSig != nil {
		prefix, _ = proto.NestedStructurePrefix(respVerifHeaderMetaSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.metaSig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.metaSig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.originSig != nil {
		prefix, _ = proto.NestedStructurePrefix(respVerifHeaderOriginSignatureField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.originSig.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.originSig.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}

		offset += n
	}

	if r.origin != nil {
		prefix, _ = proto.NestedStructurePrefix(respVerifHeaderOriginField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.origin.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.origin.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err

		}
	}

	return buf, nil
}

func (r *ResponseVerificationHeader) StableSize() (size int) {
	if r == nil {
		return 0
	}

	if r.bodySig != nil {
		_, ln := proto.NestedStructurePrefix(respVerifHeaderBodySignatureField)
		n := r.bodySig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.metaSig != nil {
		_, ln := proto.NestedStructurePrefix(respVerifHeaderMetaSignatureField)
		n := r.metaSig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.originSig != nil {
		_, ln := proto.NestedStructurePrefix(respVerifHeaderOriginSignatureField)
		n := r.originSig.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if r.origin != nil {
		_, ln := proto.NestedStructurePrefix(respVerifHeaderOriginField)
		n := r.origin.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}
