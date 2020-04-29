package session

import (
	"context"
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/internal"
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

	// Expired must return true if and only if private token is expired in the given epoch number.
	Expired(uint64) bool
}

// PrivateTokenSource is an interface of private token storage with read access.
type PrivateTokenSource interface {
	// Fetch must return the storage record corresponding to the passed key.
	//
	// Resulting error must be ErrPrivateTokenNotFound if there is no corresponding record.
	Fetch(TokenID) (PrivateToken, error)
}

// EpochLifetimeStore is an interface of the storage of elements that lifetime is limited by NeoFS epoch.
type EpochLifetimeStore interface {
	// RemoveExpired must remove all elements that are expired in the given epoch.
	RemoveExpired(uint64) error
}

// PrivateTokenStore is an interface of the storage of private tokens addressable by TokenID.
type PrivateTokenStore interface {
	PrivateTokenSource
	EpochLifetimeStore

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

// ErrPrivateTokenNotFound is raised when addressed private token was not found in storage.
const ErrPrivateTokenNotFound = internal.Error("private token not found")
