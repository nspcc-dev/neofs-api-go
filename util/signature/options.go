package signature

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/sha512"

	crypto "github.com/nspcc-dev/neofs-crypto"
)

type Options struct {
	SignFunc        func(key *ecdsa.PrivateKey, stream DataSource) ([]byte, error)
	VerifyFunc      func(key *ecdsa.PublicKey, stream DataSource, sig []byte) error
	UnmarshalPublic func([]byte) *ecdsa.PublicKey
}

// DefaultOptions represents default set of options.
func DefaultOptions() *Options {
	return &Options{
		SignFunc:        cryptoSign,
		VerifyFunc:      cryptoVerify,
		UnmarshalPublic: crypto.UnmarshalPublicKey,
	}
}

func SignWithRFC6979() SignOption {
	return func(c *Options) {
		c.SignFunc = rfc6979Sign
		c.VerifyFunc = rfc6979Verify
	}
}

// WithUnmarshalPublicKey sets f as a function for unmarshaling public keys.
func WithUnmarshalPublicKey(f func([]byte) *ecdsa.PublicKey) SignOption {
	return func(c *Options) {
		c.UnmarshalPublic = f
	}
}

func hash512(m DataSource) []byte {
	w := sha512.New()
	_, err := m.WriteSignedDataTo(w)
	if err != nil {
		panic(err)
	}
	return w.Sum(nil)
}

func hash256(m DataSource) []byte {
	w := sha256.New()
	_, err := m.WriteSignedDataTo(w)
	if err != nil {
		panic(err)
	}
	return w.Sum(nil)
}

func cryptoSign(key *ecdsa.PrivateKey, m DataSource) ([]byte, error) {
	return crypto.SignHash(key, hash512(m))
}

func cryptoVerify(pub *ecdsa.PublicKey, msg DataSource, sig []byte) error {
	return crypto.VerifyHash(pub, hash512(msg), sig)
}

func rfc6979Sign(key *ecdsa.PrivateKey, m DataSource) ([]byte, error) {
	return crypto.SignRFC6979Hash(key, hash256(m))
}

func rfc6979Verify(pub *ecdsa.PublicKey, m DataSource, sig []byte) error {
	return crypto.VerifyRFC6979Hash(pub, hash256(m), sig)
}
