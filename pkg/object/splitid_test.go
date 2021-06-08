package object_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/stretchr/testify/require"
)

func TestSplitID(t *testing.T) {
	id := object.NewSplitID()

	t.Run("toV2/fromV2", func(t *testing.T) {
		data := id.ToV2()

		newID := object.NewSplitIDFromV2(data)
		require.NotNil(t, newID)

		require.Equal(t, id, newID)
	})

	t.Run("string/parse", func(t *testing.T) {
		idStr := id.String()

		newID := object.NewSplitID()
		require.NoError(t, newID.Parse(idStr))

		require.Equal(t, id, newID)
	})

	t.Run("set UUID", func(t *testing.T) {
		newUUID := uuid.New()
		id.SetUUID(newUUID)

		require.Equal(t, newUUID.String(), id.String())
	})

	t.Run("nil value", func(t *testing.T) {
		var newID *object.SplitID

		require.NotPanics(t, func() {
			require.Nil(t, newID.ToV2())
			require.Equal(t, "", newID.String())
		})
	})
}

