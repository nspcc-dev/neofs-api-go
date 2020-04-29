package session

import (
	"context"
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-api-go/service"
)

type (
	// ObjectID type alias.
	ObjectID = refs.ObjectID
	// OwnerID type alias.
	OwnerID = refs.OwnerID
	// TokenID type alias.
	TokenID = refs.UUID
	// Token type alias
	Token = service.Token
	// Address type alias
	Address = refs.Address
	// Verb is Token_Info_Verb type alias
	Verb = service.Token_Info_Verb
)

// PrivateToken is an interface of session private part.
type PrivateToken interface {
	// PublicKey must return a binary representation of session public key.
	PublicKey() []byte

	// Sign must return the signature of passed data.
	//
	// Resulting signature must be verified by crypto.Verify function
	// with the session public key.
	Sign([]byte) ([]byte, error)
}

// PrivateTokenSource is an interface of private token storage with read access.
type PrivateTokenSource interface {
	// Fetch must return the storage record corresponding to the passed key.
	//
	// Resulting error must be ErrPrivateTokenNotFound if there is no corresponding record.
	Fetch(TokenID) (PrivateToken, error)
}

// PrivateTokenStore is an interface of the storage of private tokens addressable by TokenID.
type PrivateTokenStore interface {
	PrivateTokenSource

	// Store must save passed private token in the storage under the given key.
	//
	// Resulting error must be nil if private token was stored successfully.
	Store(TokenID, PrivateToken) error
}

// KeyStore is an interface of the storage of public keys addressable by OwnerID,
type KeyStore interface {
	// Get must return the storage record corresponding to the passed key.
	//
	// Resulting error must be ErrKeyNotFound if there is no corresponding record.
	Get(context.Context, OwnerID) ([]*ecdsa.PublicKey, error)
}

const ErrPrivateTokenNotFound = internal.Error("private token not found")

const (
	// ErrWrongFirstEpoch is raised when passed Token contains wrong first epoch.
	// First epoch is an epoch since token is valid
	ErrWrongFirstEpoch = internal.Error("wrong first epoch")

	// ErrWrongLastEpoch is raised when passed Token contains wrong last epoch.
	// Last epoch is an epoch until token is valid
	ErrWrongLastEpoch = internal.Error("wrong last epoch")

	// ErrWrongOwner is raised when passed Token contains wrong OwnerID.
	ErrWrongOwner = internal.Error("wrong owner")

	// ErrEmptyPublicKey is raised when passed Token contains wrong public key.
	ErrEmptyPublicKey = internal.Error("empty public key")

	// ErrWrongObjectsCount is raised when passed Token contains wrong objects count.
	ErrWrongObjectsCount = internal.Error("wrong objects count")

	// ErrWrongObjects is raised when passed Token contains wrong object ids.
	ErrWrongObjects = internal.Error("wrong objects")

	// ErrInvalidSignature is raised when wrong signature is passed to VerificationHeader.VerifyData().
	ErrInvalidSignature = internal.Error("invalid signature")
)
