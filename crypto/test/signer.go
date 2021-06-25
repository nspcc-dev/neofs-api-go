package neofscryptotest

import (
	"testing"

	neofscrypto "github.com/nspcc-dev/neofs-api-go/crypto"
	"github.com/stretchr/testify/require"
)

var data = []byte("test data")

// TestSigner asserts Sign method of the neofscrypto.Signer implementation
// and Verify method of the corresponding neofscrypto.PublicKey.
//
// s should not be nil.
func TestSigner(t testing.TB, s neofscrypto.Signer) {
	sig, err := s.Sign(data)
	require.NoError(t, err)

	valid := s.Public().Verify(data, sig)
	require.True(t, valid)
}

// TestPublicKeyEncoding asserts binary methods of the neofscrypto.PublicKey implementation.
//
// key and constructor result should be both non-nil pointers.
func TestPublicKeyMarshal(t testing.TB, key neofscrypto.PublicKey, blankCons func() neofscrypto.PublicKey) {
	data, err := key.MarshalBinary()
	require.NoError(t, err)

	key2 := blankCons()

	err = key2.UnmarshalBinary(data)
	require.NoError(t, err)

	require.Equal(t, key, key2)
}
