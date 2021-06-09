package cid_test

import (
	"crypto/sha256"
	"math/rand"
	"testing"

	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	cidtest "github.com/nspcc-dev/neofs-api-go/pkg/container/id/test"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func randSHA256Checksum() (cs [sha256.Size]byte) {
	rand.Read(cs[:])
	return
}

func TestID_ToV2(t *testing.T) {
	t.Run("non-nil", func(t *testing.T) {
		checksum := randSHA256Checksum()

		id := cidtest.GenerateWithChecksum(checksum)

		idV2 := id.ToV2()

		require.Equal(t, id, cid.NewFromV2(idV2))
		require.Equal(t, checksum[:], idV2.GetValue())
	})

	t.Run("nil", func(t *testing.T) {
		var x *cid.ID

		require.Nil(t, x.ToV2())
	})
}

func TestID_Equal(t *testing.T) {
	cs := randSHA256Checksum()

	id1 := cidtest.GenerateWithChecksum(cs)
	id2 := cidtest.GenerateWithChecksum(cs)

	require.True(t, id1.Equal(id2))

	id3 := cidtest.Generate()

	require.False(t, id1.Equal(id3))
}

func TestID_String(t *testing.T) {
	id := cidtest.Generate()
	id2 := cid.New()

	require.NoError(t, id2.Parse(id.String()))
	require.Equal(t, id, id2)
}

func TestContainerIDEncoding(t *testing.T) {
	id := cidtest.Generate()

	t.Run("binary", func(t *testing.T) {
		data, err := id.Marshal()
		require.NoError(t, err)

		id2 := cid.New()
		require.NoError(t, id2.Unmarshal(data))

		require.Equal(t, id, id2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := id.MarshalJSON()
		require.NoError(t, err)

		a2 := cid.New()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, id, a2)
	})
}

func TestNewFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *refs.ContainerID

		require.Nil(t, cid.NewFromV2(x))
	})
}
