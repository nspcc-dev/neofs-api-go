package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func randSHA256(t *testing.T) [sha256.Size]byte {
	cSHA256 := [sha256.Size]byte{}
	_, err := rand.Read(cSHA256[:])
	require.NoError(t, err)

	return cSHA256
}

func TestChecksum(t *testing.T) {
	c := NewChecksum()

	cSHA256 := [sha256.Size]byte{}
	_, _ = rand.Read(cSHA256[:])

	c.SetSHA256(cSHA256)

	require.Equal(t, ChecksumSHA256, c.Type())
	require.Equal(t, cSHA256[:], c.Sum())

	cV2 := c.ToV2()

	require.Equal(t, refs.SHA256, cV2.GetType())
	require.Equal(t, cSHA256[:], cV2.GetSum())

	cTZ := [64]byte{}
	_, _ = rand.Read(cSHA256[:])

	c.SetTillichZemor(cTZ)

	require.Equal(t, ChecksumTZ, c.Type())
	require.Equal(t, cTZ[:], c.Sum())

	cV2 = c.ToV2()

	require.Equal(t, refs.TillichZemor, cV2.GetType())
	require.Equal(t, cTZ[:], cV2.GetSum())
}

func TestEqualChecksums(t *testing.T) {
	require.True(t, EqualChecksums(nil, nil))

	csSHA := [sha256.Size]byte{}
	_, _ = rand.Read(csSHA[:])

	cs1 := NewChecksum()
	cs1.SetSHA256(csSHA)

	cs2 := NewChecksum()
	cs2.SetSHA256(csSHA)

	require.True(t, EqualChecksums(cs1, cs2))

	csSHA[0]++
	cs2.SetSHA256(csSHA)

	require.False(t, EqualChecksums(cs1, cs2))
}

func TestChecksumEncoding(t *testing.T) {
	cs := NewChecksum()
	cs.SetSHA256(randSHA256(t))

	t.Run("binary", func(t *testing.T) {
		data, err := cs.Marshal()
		require.NoError(t, err)

		c2 := NewChecksum()
		require.NoError(t, c2.Unmarshal(data))

		require.Equal(t, cs, c2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := cs.MarshalJSON()
		require.NoError(t, err)

		cs2 := NewChecksum()
		require.NoError(t, cs2.UnmarshalJSON(data))

		require.Equal(t, cs, cs2)
	})

	t.Run("string", func(t *testing.T) {
		cs2 := NewChecksum()

		require.NoError(t, cs2.Parse(cs.String()))

		require.Equal(t, cs, cs2)
	})
}

func TestNewChecksumFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *refs.Checksum

		require.Nil(t, NewChecksumFromV2(x))
	})
}

func TestChecksum_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *Checksum

		require.Nil(t, x.ToV2())
	})
}

func TestNewChecksum(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		chs := NewChecksum()

		// check initial values
		require.Equal(t, ChecksumUnknown, chs.Type())
		require.Nil(t, chs.Sum())

		// convert to v2 message
		chsV2 := chs.ToV2()

		require.Equal(t, refs.UnknownChecksum, chsV2.GetType())
		require.Nil(t, chsV2.GetSum())
	})
}
