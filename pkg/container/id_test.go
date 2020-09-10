package container

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDV2_0(t *testing.T) {
	cid := NewID()

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	cid.SetSHA256(checksum)

	cidV2 := cid.ToV2()

	require.Equal(t, checksum[:], cidV2.GetValue())
}
