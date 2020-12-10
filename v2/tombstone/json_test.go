package tombstone_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/tombstone"
	"github.com/stretchr/testify/require"
)

func TestTombstoneJSON(t *testing.T) {
	from := generateTombstone()

	data, err := from.MarshalJSON()
	require.NoError(t, err)

	to := new(tombstone.Tombstone)
	require.NoError(t, to.UnmarshalJSON(data))

	require.Equal(t, from, to)
}
