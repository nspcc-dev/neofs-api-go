package service

import (
	"github.com/nspcc-dev/neofs-api-go/refs"
)

// VerbContainer is an interface of the container of a token verb value.
type VerbContainer interface {
	GetVerb() Token_Info_Verb
	SetVerb(Token_Info_Verb)
}

// TokenIDContainer is an interface of the container of a token ID value.
type TokenIDContainer interface {
	GetID() TokenID
	SetID(TokenID)
}

// CreationEpochContainer is an interface of the container of a creation epoch number.
type CreationEpochContainer interface {
	CreationEpoch() uint64
	SetCreationEpoch(uint64)
}

// ExpirationEpochContainer is an interface of the container of an expiration epoch number.
type ExpirationEpochContainer interface {
	ExpirationEpoch() uint64
	SetExpirationEpoch(uint64)
}

// SessionKeyContainer is an interface of the container of session key bytes.
type SessionKeyContainer interface {
	GetSessionKey() []byte
	SetSessionKey([]byte)
}

// SignatureContainer is an interface of the container of signature bytes.
type SignatureContainer interface {
	GetSignature() []byte
	SetSignature([]byte)
}

// SessionTokenInfo is an interface that determines the information scope of session token.
type SessionTokenInfo interface {
	TokenIDContainer
	refs.OwnerIDContainer
	VerbContainer
	refs.AddressContainer
	CreationEpochContainer
	ExpirationEpochContainer
	SessionKeyContainer
}

// SessionToken is an interface of token information and signature pair.
type SessionToken interface {
	SessionTokenInfo
	SignatureContainer
}

var _ SessionToken = (*Token)(nil)

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
