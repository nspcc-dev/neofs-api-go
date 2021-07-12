package neofsecdsatest

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"sync"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
)

var (
	k ecdsa.PrivateKey
	o sync.Once
)

func initKey() {
	const str = "307702010104203ee1fd84dd7199925f8d32f897aaa7f2d6484aa3738e5e0abd03f8240d7c6d8ca00a06082a8648ce3d030107a1440342000475099c302b77664a2508bec1cae47903857b762c62713f190e8d99912ef76737f36191e4c0ea50e47b0e0edbae24fd6529df84f9bd63f87219df3a086efe9195"

	var (
		buf  []byte
		err  error
		kptr *ecdsa.PrivateKey
	)

	if buf, err = hex.DecodeString(str); err == nil {
		kptr, err = x509.ParseECPrivateKey(buf)
	}

	if err != nil {
		panic(err)
	}

	k = *kptr
}

// Key returns ecdsa.PrivateKey.
//
// Result is always the same.
func Key() ecdsa.PrivateKey {
	o.Do(initKey)
	return k
}

// Signer returns ECDSA neofscrypto.Signer.
//
// Result is always the same.
func Signer() neofscrypto.Signer {
	return neofsecdsa.Signer(Key())
}

// Public returns ECDSA neofscrypto.PublicKey.
//
// Result is always the same.
func Public() neofscrypto.PublicKey {
	return Signer().Public()
}

// PublicKey returns ecdsa.PublicKey.
//
// Result is always the same.
func PublicKey() ecdsa.PublicKey {
	return Key().PublicKey
}

// PublicBytes returns ECDSA public key in a binary format.
//
// Result is always the same.
func PublicBytes() []byte {
	d, err := Public().MarshalBinary()
	if err != nil {
		panic(err)
	}

	return d
}
