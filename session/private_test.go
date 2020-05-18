package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPToken_PrivateKey(t *testing.T) {
	// create new private token
	pToken, err := NewPrivateToken(0)
	require.NoError(t, err)
	require.NotNil(t, pToken.PrivateKey())
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

func TestPrivateTokenKey_SetOwnerID(t *testing.T) {
	ownerID := OwnerID{1, 2, 3}

	s := new(PrivateTokenKey)

	s.SetOwnerID(ownerID)

	require.Equal(t, ownerID, s.owner)
}

func TestPrivateTokenKey_SetTokenID(t *testing.T) {
	tokenID := TokenID{1, 2, 3}

	s := new(PrivateTokenKey)

	s.SetTokenID(tokenID)

	require.Equal(t, tokenID, s.token)
}
