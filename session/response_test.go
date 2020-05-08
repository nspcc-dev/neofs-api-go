package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateResponseGettersSetters(t *testing.T) {
	t.Run("id", func(t *testing.T) {
		id := TokenID{1, 2, 3}
		m := new(CreateResponse)

		m.SetID(id)

		require.Equal(t, id, m.GetID())
	})

	t.Run("session key", func(t *testing.T) {
		key := []byte{1, 2, 3}
		m := new(CreateResponse)

		m.SetSessionKey(key)

		require.Equal(t, key, m.GetSessionKey())
	})
}
