package session

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/stretchr/testify/require"
)

func TestMapTokenStore(t *testing.T) {
	// create new private token
	pToken, err := NewPrivateToken(0)
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

func TestMapTokenStore_RemoveExpired(t *testing.T) {
	// create some epoch number
	e1 := uint64(1)

	// create private token that expires after e1
	tok1, err := NewPrivateToken(e1)
	require.NoError(t, err)

	// create some greater than e1 epoch number
	e2 := e1 + 1

	// create private token that expires after e2
	tok2, err := NewPrivateToken(e2)
	require.NoError(t, err)

	// create token store instance
	s := NewMapTokenStore()

	// create storage keys for tokens
	id1, err := refs.NewUUID()
	require.NoError(t, err)
	id2, err := refs.NewUUID()
	require.NoError(t, err)

	assertPresence := func(ids ...TokenID) {
		for i := range ids {
			_, err = s.Fetch(ids[i])
			require.NoError(t, err)
		}
	}

	assertAbsence := func(ids ...TokenID) {
		for i := range ids {
			_, err = s.Fetch(ids[i])
			require.EqualError(t, err, ErrPrivateTokenNotFound.Error())
		}
	}

	// store both tokens
	require.NoError(t, s.Store(id1, tok1))
	require.NoError(t, s.Store(id2, tok2))

	// ascertain that both tokens are available
	assertPresence(id1, id2)

	// perform cleaning for epoch in which both tokens are not expired
	require.NoError(t, s.RemoveExpired(e1))

	// ascertain that both tokens are still available
	assertPresence(id1, id2)

	// perform cleaning for epoch greater than e1 and not greater than e2
	require.NoError(t, s.RemoveExpired(e1+1))

	// ascertain that tok1 was removed
	assertAbsence(id1)

	// ascertain that tok2 was not removed
	assertPresence(id2)
}
