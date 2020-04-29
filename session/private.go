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

// Sign signs data with session private key.
func (t *pToken) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(t.sessionKey, data)
}

// PublicKey returns a binary representation of the session public key.
func (t *pToken) PublicKey() []byte {
	return crypto.MarshalPublicKey(&t.sessionKey.PublicKey)
}

func (t *pToken) Expired(epoch uint64) bool {
	return t.validUntil < epoch
}
