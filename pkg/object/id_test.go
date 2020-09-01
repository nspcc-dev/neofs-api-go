package object

import (
	"crypto/rand"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := new(ID)

	checksum := [sha256.Size]byte{}

	_, err := rand.Read(checksum[:])
	require.NoError(t, err)

	id.SetSHA256(checksum)

	idV2 := id.ToV2()

	id2, err := IDFromV2(idV2)
	require.NoError(t, err)

	require.Equal(t, id, id2)
}
