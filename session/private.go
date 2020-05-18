package session

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type pToken struct {
	// private session token
	sessionKey *ecdsa.PrivateKey
	// last epoch of the lifetime
	validUntil uint64
}

// NewPrivateToken creates PrivateToken instance that expires after passed epoch.
//
// Returns non-nil error on key generation error.
func NewPrivateToken(validUntil uint64) (PrivateToken, error) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &pToken{
		sessionKey: sk,
		validUntil: validUntil,
	}, nil
}

// PrivateKey returns a binary representation of the session public key.
func (t *pToken) PrivateKey() *ecdsa.PrivateKey {
	return t.sessionKey
}

func (t *pToken) Expired(epoch uint64) bool {
	return t.validUntil < epoch
}

// SetOwnerID is an owner ID field setter.
func (s *PrivateTokenKey) SetOwnerID(id OwnerID) {
	s.owner = id
}

// SetTokenID is a token ID field setter.
func (s *PrivateTokenKey) SetTokenID(id TokenID) {
	s.token = id
}
