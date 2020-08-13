package refs

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOwnerID_StableMarshal(t *testing.T) {
	owner := make([]byte, 25)
	_, err := rand.Read(owner)
	require.NoError(t, err)

	expectedOwner := new(OwnerID)
	expectedOwner.Value = owner

	gotOwner := new(OwnerID)

	t.Run("small buffer", func(t *testing.T) {
		_, err = expectedOwner.StableMarshal(make([]byte, 1))
		require.Error(t, err)
	})

	t.Run("empty owner", func(t *testing.T) {
		data, err := new(OwnerID).StableMarshal(nil)
		require.NoError(t, err)

		err = gotOwner.Unmarshal(data)
		require.NoError(t, err)

		require.Len(t, gotOwner.Value, 0)
	})

	t.Run("non empty owner", func(t *testing.T) {
		data, err := expectedOwner.StableMarshal(nil)
		require.NoError(t, err)

		err = gotOwner.Unmarshal(data)
		require.NoError(t, err)

		require.Equal(t, expectedOwner, gotOwner)
	})
}
