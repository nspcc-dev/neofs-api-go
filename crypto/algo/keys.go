package cryptoalgo

import (
	"errors"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	neofsrfc6979 "github.com/nspcc-dev/neofs-api-go/crypto/rfc6979"
)

type kc func() neofscrypto.PublicKey

type kcs struct {
	m map[SignatureAlgorithm]kc
}

var v = &kcs{
	m: make(map[SignatureAlgorithm]kc),
}

func (v *kcs) register(a SignatureAlgorithm, c kc) {
	v.m[a] = c
}

func init() {
	v.register(ECDSA, neofsecdsa.PublicBlank)
	v.register(RFC6979, neofsrfc6979.PublicBlank)
}

// UnmarshalKey unmarshals neofscrypto.PublicKey from a binary representation
// according to SignatureAlgorithm.
//
// If algo is not supported, an error returns.
func UnmarshalKey(algo SignatureAlgorithm, binKey []byte) (neofscrypto.PublicKey, error) {
	c, ok := v.m[algo]
	if !ok {
		return nil, errors.New("unsupported signature algorithm")
	}

	key := c()

	return key, key.UnmarshalBinary(binKey)
}
