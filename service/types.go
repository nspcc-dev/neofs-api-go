package service

// NodeRole to identify in Bootstrap service.
type NodeRole int32

// TTLCondition is a function type that used to verify that TTL values match a specific criterion.
// Nil error indicates compliance with the criterion.
type TTLCondition func(uint32) error

// RawSource is an interface of the container of a boolean Raw value with read access.
type RawSource interface {
	GetRaw() bool
}

// RawContainer is an interface of the container of a boolean Raw value.
type RawContainer interface {
	RawSource
	SetRaw(bool)
}

// VersionContainer is an interface of the container of a numerical Version value with read access.
type VersionSource interface {
	GetVersion() uint32
}

// VersionContainer is an interface of the container of a numerical Version value.
type VersionContainer interface {
	VersionSource
	SetVersion(uint32)
}

// EpochSource is an interface of the container of a NeoFS epoch number with read access.
type EpochSource interface {
	GetEpoch() uint64
}

// EpochContainer is an interface of the container of a NeoFS epoch number.
type EpochContainer interface {
	EpochSource
	SetEpoch(uint64)
}

// TTLSource is an interface of the container of a numerical TTL value with read access.
type TTLSource interface {
	GetTTL() uint32
}

// TTLContainer is an interface of the container of a numerical TTL value.
type TTLContainer interface {
	TTLSource
	SetTTL(uint32)
}

// RequestMetaContainer is an interface of a fixed set of request meta value containers.
// Contains:
// - TTL value;
// - NeoFS epoch number;
// - Protocol version;
// - Raw toggle option.
type RequestMetaContainer interface {
	TTLContainer
	EpochContainer
	VersionContainer
	RawContainer
}

// VerbSource is an interface of the container of a token verb value with read access.
type VerbSource interface {
	GetVerb() Token_Info_Verb
}

// VerbContainer is an interface of the container of a token verb value.
type VerbContainer interface {
	VerbSource
	SetVerb(Token_Info_Verb)
}

// TokenIDSource is an interface of the container of a token ID value with read access.
type TokenIDSource interface {
	GetID() TokenID
}

// TokenIDContainer is an interface of the container of a token ID value.
type TokenIDContainer interface {
	TokenIDSource
	SetID(TokenID)
}

// CreationEpochSource is an interface of the container of a creation epoch number with read access.
type CreationEpochSource interface {
	CreationEpoch() uint64
}

// CreationEpochContainer is an interface of the container of a creation epoch number.
type CreationEpochContainer interface {
	CreationEpochSource
	SetCreationEpoch(uint64)
}

// ExpirationEpochSource is an interface of the container of an expiration epoch number with read access.
type ExpirationEpochSource interface {
	ExpirationEpoch() uint64
}

// ExpirationEpochContainer is an interface of the container of an expiration epoch number.
type ExpirationEpochContainer interface {
	ExpirationEpochSource
	SetExpirationEpoch(uint64)
}

// SessionKeySource is an interface of the container of session key bytes with read access.
type SessionKeySource interface {
	GetSessionKey() []byte
}

// SessionKeyContainer is an interface of the container of public session key bytes.
type SessionKeyContainer interface {
	SessionKeySource
	SetSessionKey([]byte)
}

// SignatureSource is an interface of the container of signature bytes with read access.
type SignatureSource interface {
	GetSignature() []byte
}

// SignatureContainer is an interface of the container of signature bytes.
type SignatureContainer interface {
	SignatureSource
	SetSignature([]byte)
}

// SessionTokenSource is an interface of the container of a SessionToken with read access.
type SessionTokenSource interface {
	GetSessionToken() SessionToken
}

// SessionTokenInfo is an interface of a fixed set of token information value containers.
// Contains:
// - ID of the token;
// - ID of the token's owner;
// - verb of the session;
// - address of the session object;
// - creation epoch number of the token;
// - expiration epoch number of the token;
// - public session key bytes.
type SessionTokenInfo interface {
	TokenIDContainer
	OwnerIDContainer
	VerbContainer
	AddressContainer
	CreationEpochContainer
	ExpirationEpochContainer
	SessionKeyContainer
}

// SessionToken is an interface of token information and signature pair.
type SessionToken interface {
	SessionTokenInfo
	SignatureContainer
}
