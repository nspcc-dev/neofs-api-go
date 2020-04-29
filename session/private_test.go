package session

import (
	"crypto/rand"
	"testing"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/stretchr/testify/require"
)

func TestPrivateToken(t *testing.T) {
	// create new private token
	pToken, err := NewPrivateToken()
	require.NoError(t, err)

	// generate data to sign
	data := make([]byte, 10)
	_, err = rand.Read(data)
	require.NoError(t, err)

	// sign data via private token
	sig, err := pToken.Sign(data)
	require.NoError(t, err)

	// check signature
	require.NoError(t,
		crypto.Verify(
			crypto.UnmarshalPublicKey(pToken.PublicKey()),
			data,
			sig,
		),
	)
}
