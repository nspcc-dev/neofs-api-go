package session

import (
	"context"
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/refs"
)

type (
	// KeyStore is an interface that describes storage,
	// that allows to fetch public keys by OwnerID.
	KeyStore interface {
		Get(ctx context.Context, id refs.OwnerID) ([]*ecdsa.PublicKey, error)
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
