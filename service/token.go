package service

import (
	"crypto/ecdsa"
	"encoding/binary"
	"io"

	"github.com/nspcc-dev/neofs-api-go/refs"
)

const verbSize = 4

const fixedTokenDataSize = 0 +
	refs.UUIDSize +
	refs.OwnerIDSize +
	verbSize +
	refs.UUIDSize +
	refs.CIDSize +
	8 +
	8

var tokenEndianness = binary.BigEndian

// GetID is an ID field getter.
func (m Token_Info) GetID() TokenID {
	return m.ID
}

// SetID is an ID field setter.
func (m *Token_Info) SetID(id TokenID) {
	m.ID = id
}

// GetOwnerID is an OwnerID field getter.
func (m Token_Info) GetOwnerID() OwnerID {
	return m.OwnerID
}

// SetOwnerID is an OwnerID field setter.
func (m *Token_Info) SetOwnerID(id OwnerID) {
	m.OwnerID = id
}

// SetVerb is a Verb field setter.
func (m *Token_Info) SetVerb(verb Token_Info_Verb) {
	m.Verb = verb
}

// GetAddress is an Address field getter.
func (m Token_Info) GetAddress() Address {
	return m.Address
}

// SetAddress is an Address field setter.
func (m *Token_Info) SetAddress(addr Address) {
	m.Address = addr
}

// CreationEpoch is a Created field getter.
func (m Token_Info) CreationEpoch() uint64 {
	return m.Created
}

// SetCreationEpoch is a Created field setter.
func (m *Token_Info) SetCreationEpoch(e uint64) {
	m.Created = e
}

// ExpirationEpoch is a ValidUntil field getter.
func (m Token_Info) ExpirationEpoch() uint64 {
	return m.ValidUntil
}

// SetExpirationEpoch is a ValidUntil field setter.
func (m *Token_Info) SetExpirationEpoch(e uint64) {
	m.ValidUntil = e
}

// SetSessionKey is a SessionKey field setter.
func (m *Token_Info) SetSessionKey(key []byte) {
	m.SessionKey = key
}

// SetSignature is a Signature field setter.
func (m *Token) SetSignature(sig []byte) {
	m.Signature = sig
}

// Size returns the size of a binary representation of the verb.
func (x Token_Info_Verb) Size() int {
	return verbSize
}

// Bytes returns a binary representation of the verb.
func (x Token_Info_Verb) Bytes() []byte {
	data := make([]byte, verbSize)
	tokenEndianness.PutUint32(data, uint32(x))
	return data
}

// AddSignKey calls a Signature field setter with passed signature.
func (m *Token) AddSignKey(sig []byte, _ *ecdsa.PublicKey) {
	m.SetSignature(sig)
}

// SignedData returns token information in a binary representation.
func (m *Token) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	copyTokenSignedData(data, m)

	return data, nil
}

// ReadSignedData copies a binary representation of the token information to passed buffer.
//
// If buffer length is less than required, io.ErrUnexpectedEOF returns.
func (m *Token_Info) ReadSignedData(p []byte) (int, error) {
	sz := m.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	copyTokenSignedData(p, m)

	return sz, nil
}

// SignedDataSize returns the length of signed token information slice.
func (m Token_Info) SignedDataSize() int {
	return fixedTokenDataSize + len(m.GetSessionKey())
}

// Fills passed buffer with signing token information bytes.
// Does not check buffer length, it is understood that enough space is allocated in it.
func copyTokenSignedData(buf []byte, token SessionTokenInfo) {
	var off int

	off += copy(buf[off:], token.GetID().Bytes())

	off += copy(buf[off:], token.GetOwnerID().Bytes())

	off += copy(buf[off:], token.GetVerb().Bytes())

	addr := token.GetAddress()
	off += copy(buf[off:], addr.CID.Bytes())
	off += copy(buf[off:], addr.ObjectID.Bytes())

	tokenEndianness.PutUint64(buf[off:], token.CreationEpoch())
	off += 8

	tokenEndianness.PutUint64(buf[off:], token.ExpirationEpoch())
	off += 8

	copy(buf[off:], token.GetSessionKey())
}
