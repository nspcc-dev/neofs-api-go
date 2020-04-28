package session

import (
	"crypto/ecdsa"
	"sync"

	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/nspcc-dev/neofs-api-go/service"
	crypto "github.com/nspcc-dev/neofs-crypto"
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

	// PToken is a wrapper around Token that allows to sign data
	// and to do thread-safe manipulations.
	PToken struct {
		Token

		mtx        *sync.Mutex
		PrivateKey *ecdsa.PrivateKey
	}
)

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

// SignData signs data with session private key.
func (t *PToken) SignData(data []byte) ([]byte, error) {
	return crypto.Sign(t.PrivateKey, data)
}
