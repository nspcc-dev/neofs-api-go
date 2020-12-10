package tombstone_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/tombstone"
	"github.com/stretchr/testify/require"
)

func TestTombstone_StableMarshal(t *testing.T) {
	from := generateTombstone()

	t.Run("non empty", func(t *testing.T) {
		wire, err := from.StableMarshal(nil)
		require.NoError(t, err)

		to := new(tombstone.Tombstone)
		require.NoError(t, to.Unmarshal(wire))

		require.Equal(t, from, to)
	})
}

func generateTombstone() *tombstone.Tombstone {
	t := new(tombstone.Tombstone)

	oid1 := new(refs.ObjectID)
	oid1.SetValue([]byte("Object ID 1"))

	oid2 := new(refs.ObjectID)
	oid2.SetValue([]byte("Object ID 2"))

	t.SetExpirationEpoch(100)
	t.SetSplitID([]byte("split ID"))
	t.SetMembers([]*refs.ObjectID{oid1, oid2})

	return t
}
