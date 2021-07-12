package neofsecdsa_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	neofsecdsa "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa"
	neofscryptotest "github.com/nspcc-dev/neofs-api-go/crypto/test"
	"github.com/stretchr/testify/require"
)

func TestEcdsaSigner(t *testing.T) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	s := neofsecdsa.Signer(*key)

	neofscryptotest.TestSigner(t, s)

	neofscryptotest.TestPublicKeyMarshal(t, s.Public(), neofsecdsa.PublicBlank)
}
