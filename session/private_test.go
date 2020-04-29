package session

import (
	"crypto/rand"
	"testing"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/stretchr/testify/require"
)

func TestPrivateToken(t *testing.T) {
	// create new private token
	pToken, err := NewPrivateToken(0)
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

func TestPToken_Expired(t *testing.T) {
	e := uint64(10)

	var token PrivateToken = &pToken{
		validUntil: e,
	}

	// must not be expired in the epoch before last
	require.False(t, token.Expired(e-1))

	// must not be expired in the last epoch
	require.False(t, token.Expired(e))

	// must be expired in the epoch after last
	require.True(t, token.Expired(e+1))
}
