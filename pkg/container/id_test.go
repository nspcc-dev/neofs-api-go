package container

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDV2_0(t *testing.T) {
	cid := new(ID)

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	cid.SetSHA256(checksum)

	cidV2 := cid.ToV2()

	cid2, err := IDFromV2(cidV2)
	require.NoError(t, err)

	require.Equal(t, cid, cid2)
}
