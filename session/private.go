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
}

// NewSessionPrivateToken creates PrivateToken instance.
//
// Returns non-nil error on key generation error.
func NewPrivateToken() (PrivateToken, error) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &pToken{
		sessionKey: sk,
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
