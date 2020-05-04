package service

import (
	"crypto/ecdsa"
	"encoding/binary"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

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

// Returns byte slice that is used for creation/verification of the token signature.
func verificationTokenData(token SessionToken) []byte {
	var sz int

	id := token.GetID()
	sz += id.Size()

	ownerID := token.GetOwnerID()
	sz += ownerID.Size()

	verb := uint32(token.GetVerb())
	sz += 4

	addr := token.GetAddress()
	sz += addr.CID.Size() + addr.ObjectID.Size()

	cEpoch := token.CreationEpoch()
	sz += 8

	fEpoch := token.ExpirationEpoch()
	sz += 8

	key := token.GetSessionKey()
	sz += len(key)

	data := make([]byte, sz)

	var off int

	tokenEndianness.PutUint32(data, verb)
	off += 4

	tokenEndianness.PutUint64(data[off:], cEpoch)
	off += 8

	tokenEndianness.PutUint64(data[off:], fEpoch)
	off += 8

	off += copy(data[off:], id.Bytes())
	off += copy(data[off:], ownerID.Bytes())
	off += copy(data[off:], addr.CID.Bytes())
	off += copy(data[off:], addr.ObjectID.Bytes())
	off += copy(data[off:], key)

	return data
}

// SignToken calculates and stores the signature of token information.
//
// If passed token is nil, ErrNilToken returns.
// If passed private key is nil, crypto.ErrEmptyPrivateKey returns.
func SignToken(token SessionToken, key *ecdsa.PrivateKey) error {
	if token == nil {
		return ErrNilToken
	} else if key == nil {
		return crypto.ErrEmptyPrivateKey
	}

	sig, err := crypto.Sign(key, verificationTokenData(token))
	if err != nil {
		return err
	}

	token.SetSignature(sig)

	return nil
}

// VerifyTokenSignature checks if token was signed correctly.
//
// If passed token is nil, ErrNilToken returns.
// If passed public key is nil, crypto.ErrEmptyPublicKey returns.
func VerifyTokenSignature(token SessionToken, key *ecdsa.PublicKey) error {
	if token == nil {
		return ErrNilToken
	} else if key == nil {
		return crypto.ErrEmptyPublicKey
	}

	return crypto.Verify(
		key,
		verificationTokenData(token),
		token.GetSignature(),
	)
}
