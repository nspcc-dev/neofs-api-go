package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateRequestGettersSetters(t *testing.T) {
	t.Run("owner ID", func(t *testing.T) {
		id := OwnerID{1, 2, 3}
		m := new(CreateRequest)

		m.SetOwnerID(id)

		require.Equal(t, id, m.GetOwnerID())
	})

	t.Run("lifetime", func(t *testing.T) {
		e1, e2 := uint64(3), uint64(4)
		m := new(CreateRequest)

		m.SetCreationEpoch(e1)
		m.SetExpirationEpoch(e2)

		require.Equal(t, e1, m.CreationEpoch())
		require.Equal(t, e2, m.ExpirationEpoch())
	})
}
