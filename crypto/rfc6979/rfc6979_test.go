package neofsrfc6979_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	neofsrfc6979 "github.com/nspcc-dev/neofs-api-go/crypto/rfc6979"
	neofscryptotest "github.com/nspcc-dev/neofs-api-go/crypto/test"
	"github.com/stretchr/testify/require"
)

func TestRFC6979Signer(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	s := neofsrfc6979.Signer(*key)

	neofscryptotest.TestSigner(t, s)

	neofscryptotest.TestPublicKeyMarshal(t, s.Public(), neofsrfc6979.PublicBlank)
}
