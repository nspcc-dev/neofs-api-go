package object

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := NewID()

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	id.SetSHA256(checksum)

	idV2 := id.ToV2()

	require.Equal(t, checksum[:], idV2.GetValue())
}

func TestID_Equal(t *testing.T) {
	cs := randSHA256Checksum(t)

	id1 := NewID()
	id1.SetSHA256(cs)

	id2 := NewID()
	id2.SetSHA256(cs)

	id3 := NewID()
	id3.SetSHA256(randSHA256Checksum(t))

	require.True(t, id1.Equal(id2))
	require.False(t, id1.Equal(id3))
}
