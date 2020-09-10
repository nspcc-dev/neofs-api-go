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
