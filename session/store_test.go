package session

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/stretchr/testify/require"
)

func TestMapTokenStore(t *testing.T) {
	// create new private token
	pToken, err := NewPrivateToken()
	require.NoError(t, err)

	// create map token store
	s := NewMapTokenStore()

	// create new storage key
	id, err := refs.NewUUID()
	require.NoError(t, err)

	// ascertain that there is no record for the key
	_, err = s.Fetch(id)
	require.EqualError(t, err, ErrPrivateTokenNotFound.Error())

	// save private token record
	require.NoError(t, s.Store(id, pToken))

	// fetch private token by the key
	res, err := s.Fetch(id)
	require.NoError(t, err)

	// ascertain that returned token equals to initial
	require.Equal(t, pToken, res)
}
