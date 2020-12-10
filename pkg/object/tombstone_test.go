package object

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func generateIDList(sz int) []*ID {
	res := make([]*ID, sz)
	cs := [sha256.Size]byte{}

	for i := 0; i < sz; i++ {
		res[i] = NewID()
		rand.Read(cs[:])
		res[i].SetSHA256(cs)
	}

	return res
}

func TestTombstone(t *testing.T) {
	ts := NewTombstone()

	exp := uint64(13)
	ts.SetExpirationEpoch(exp)
	require.Equal(t, exp, ts.ExpirationEpoch())

	splitID := NewSplitID()
	ts.SetSplitID(splitID)
	require.Equal(t, splitID, ts.SplitID())

	members := generateIDList(3)
	ts.SetMembers(members)
	require.Equal(t, members, ts.Members())
}

func TestTombstoneEncoding(t *testing.T) {
	ts := NewTombstone()
	ts.SetExpirationEpoch(13)
	ts.SetSplitID(NewSplitID())
	ts.SetMembers(generateIDList(5))

	t.Run("binary", func(t *testing.T) {
		data, err := ts.Marshal()
		require.NoError(t, err)

		ts2 := NewTombstone()
		require.NoError(t, ts2.Unmarshal(data))

		require.Equal(t, ts, ts2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := ts.MarshalJSON()
		require.NoError(t, err)

		ts2 := NewTombstone()
		require.NoError(t, ts2.UnmarshalJSON(data))

		require.Equal(t, ts, ts2)
	})
}
