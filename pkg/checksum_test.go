package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/stretchr/testify/require"
)

func TestChecksum(t *testing.T) {
	c := NewChecksum()

	cSHA256 := [sha256.Size]byte{}
	_, _ = rand.Read(cSHA256[:])

	c.SetSHA256(cSHA256)

	require.Equal(t, ChecksumSHA256, c.GetType())
	require.Equal(t, cSHA256[:], c.GetSum())

	cV2 := c.ToV2()

	require.Equal(t, refs.SHA256, cV2.GetType())
	require.Equal(t, cSHA256[:], cV2.GetSum())

	cTZ := [64]byte{}
	_, _ = rand.Read(cSHA256[:])

	c.SetTillichZemor(cTZ)

	require.Equal(t, ChecksumTZ, c.GetType())
	require.Equal(t, cTZ[:], c.GetSum())

	cV2 = c.ToV2()

	require.Equal(t, refs.TillichZemor, cV2.GetType())
	require.Equal(t, cTZ[:], cV2.GetSum())
}
