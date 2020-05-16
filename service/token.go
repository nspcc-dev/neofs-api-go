package service

import (
	"crypto/ecdsa"
	"encoding/binary"
	"io"

	"github.com/nspcc-dev/neofs-api-go/refs"
)

type signAccumWithToken struct {
	SignedDataSource
	SignKeyPairAccumulator
	SignKeyPairSource

	token SessionToken
}

type signDataReaderWithToken struct {
	SignedDataSource
	SignKeyPairAccumulator
	SignKeyPairSource

	rdr SignedDataReader

	token SessionToken
}

type signedSessionToken struct {
	SessionToken
}

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
func (m TokenLifetime) CreationEpoch() uint64 {
	return m.Created
}

// SetCreationEpoch is a Created field setter.
func (m *TokenLifetime) SetCreationEpoch(e uint64) {
	m.Created = e
}

// ExpirationEpoch is a ValidUntil field getter.
func (m TokenLifetime) ExpirationEpoch() uint64 {
	return m.ValidUntil
}

// SetExpirationEpoch is a ValidUntil field setter.
func (m *TokenLifetime) SetExpirationEpoch(e uint64) {
	m.ValidUntil = e
}

// SetSessionKey is a SessionKey field setter.
func (m *Token_Info) SetSessionKey(key []byte) {
	m.SessionKey = key
}

// SetOwnerKey is an OwnerKey field setter.
func (m *Token_Info) SetOwnerKey(key []byte) {
	m.OwnerKey = key
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

// AddSignKey calls a Signature field setter of token with passed signature.
func (s signedSessionToken) AddSignKey(sig []byte, _ *ecdsa.PublicKey) {
	if s.SessionToken != nil {
		s.SessionToken.SetSignature(sig)
	}
}

// SignedData returns token information in a binary representation.
func (s signedSessionToken) SignedData() ([]byte, error) {
	return SignedDataFromReader(s)
}

// SignedDataSize returns the length of signed token information slice.
func (s signedSessionToken) SignedDataSize() int {
	return tokenInfoSize(s.SessionToken)
}

// ReadSignedData copies a binary representation of the token information to passed buffer.
//
// If buffer length is less than required, io.ErrUnexpectedEOF returns.
func (s signedSessionToken) ReadSignedData(p []byte) (int, error) {
	sz := s.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	copyTokenSignedData(p, s.SessionToken)

	return sz, nil
}

// NewSignedSessionToken wraps passed SessionToken in a component suitable for signing.
//
// Result can be used in AddSignatureWithKey function.
func NewSignedSessionToken(token SessionToken) DataWithSignKeyAccumulator {
	return &signedSessionToken{
		SessionToken: token,
	}
}

// NewVerifiedSessionToken wraps passed SessionToken in a component suitable for signature verification.
//
// Result can be used in VerifySignatureWithKey function.
func NewVerifiedSessionToken(token SessionToken) DataWithSignature {
	return &signedSessionToken{
		SessionToken: token,
	}
}

func tokenInfoSize(v SessionKeySource) int {
	if v == nil {
		return 0
	}
	return fixedTokenDataSize + len(v.GetSessionKey())
}

// Fills passed buffer with signing token information bytes.
// Does not check buffer length, it is understood that enough space is allocated in it.
//
// If passed SessionTokenInfo, buffer remains unchanged.
func copyTokenSignedData(buf []byte, token SessionTokenInfo) {
	if token == nil {
		return
	}

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

// SignedData concatenates signed data with session token information. Returns concatenation result.
//
// Token bytes are added if and only if token is not nil.
func (s signAccumWithToken) SignedData() ([]byte, error) {
	data, err := s.SignedDataSource.SignedData()
	if err != nil {
		return nil, err
	}

	tokenData := make([]byte, tokenInfoSize(s.token))

	copyTokenSignedData(tokenData, s.token)

	return append(data, tokenData...), nil
}

func (s signDataReaderWithToken) SignedDataSize() int {
	sz := s.rdr.SignedDataSize()
	if sz < 0 {
		return -1
	}

	sz += tokenInfoSize(s.token)

	return sz
}

func (s signDataReaderWithToken) ReadSignedData(p []byte) (int, error) {
	dataSize := s.rdr.SignedDataSize()
	if dataSize < 0 {
		return 0, ErrNegativeLength
	}

	sumSize := dataSize + tokenInfoSize(s.token)

	if len(p) < sumSize {
		return 0, io.ErrUnexpectedEOF
	}

	if n, err := s.rdr.ReadSignedData(p); err != nil {
		return n, err
	}

	copyTokenSignedData(p[dataSize:], s.token)

	return sumSize, nil
}
