package session

import (
	"context"
	"crypto/ecdsa"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-proto/refs"
)

type (
	// KeyStore is an interface that describes storage,
	// that allows to fetch public keys by OwnerID.
	KeyStore interface {
		Get(ctx context.Context, id refs.OwnerID) ([]*ecdsa.PublicKey, error)
	}

	// TokenStore is a PToken storage manipulation interface.
	TokenStore interface {
		// New returns new token with specified parameters.
		New(p TokenParams) *PToken

		// Fetch tries to fetch a token with specified id.
		Fetch(id TokenID) *PToken

		// Remove removes token with id from store.
		Remove(id TokenID)
	}

	// TokenParams contains params to create new PToken.
	TokenParams struct {
		FirstEpoch uint64
		LastEpoch  uint64
		ObjectID   []ObjectID
		OwnerID    OwnerID
		PublicKeys [][]byte
	}
)

// NewInitRequest returns new initialization CreateRequest from passed Token.
func NewInitRequest(t *Token) *CreateRequest {
	return &CreateRequest{Message: &CreateRequest_Init{Init: t}}
}

// NewSignedRequest returns new signed CreateRequest from passed Token.
func NewSignedRequest(t *Token) *CreateRequest {
	return &CreateRequest{Message: &CreateRequest_Signed{Signed: t}}
}

// Sign signs contents of the header with the private key.
func (m *VerificationHeader) Sign(key *ecdsa.PrivateKey) error {
	s, err := crypto.Sign(key, m.PublicKey)
	if err != nil {
		return err
	}
	m.KeySignature = s
	return nil
}
