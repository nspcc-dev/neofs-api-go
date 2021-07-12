package neofsrfc6979test

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"sync"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	neofsrfc6979 "github.com/nspcc-dev/neofs-api-go/crypto/rfc6979"
)

var (
	s neofscrypto.Signer
	o sync.Once
)

// Signer returns RFC6979 neofscrypto.Signer.
//
// Result is always the same.
func Signer() neofscrypto.Signer {
	o.Do(func() {
		const str = "307702010104203ee1fd84dd7199925f8d32f897aaa7f2d6484aa3738e5e0abd03f8240d7c6d8ca00a06082a8648ce3d030107a1440342000475099c302b77664a2508bec1cae47903857b762c62713f190e8d99912ef76737f36191e4c0ea50e47b0e0edbae24fd6529df84f9bd63f87219df3a086efe9195"

		var (
			key *ecdsa.PrivateKey
			buf []byte
			err error
		)

		if buf, err = hex.DecodeString(str); err == nil {
			key, err = x509.ParseECPrivateKey(buf)
		}

		if err != nil {
			panic(err)
		}

		s = neofsrfc6979.Signer(*key)
	})

	return s
}
