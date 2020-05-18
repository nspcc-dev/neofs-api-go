package session

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	crypto "github.com/nspcc-dev/neofs-crypto"
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

// PublicSessionToken returns a binary representation of session public key.
//
// If passed PrivateToken is nil, ErrNilPrivateToken returns.
// If passed PrivateToken carries nil private key, crypto.ErrEmptyPrivateKey returns.
func PublicSessionToken(pToken PrivateToken) ([]byte, error) {
	if pToken == nil {
		return nil, ErrNilPrivateToken
	}

	sk := pToken.PrivateKey()
	if sk == nil {
		return nil, crypto.ErrEmptyPrivateKey
	}

	return crypto.MarshalPublicKey(&sk.PublicKey), nil
}

// PrivateKey is a session private key getter.
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
