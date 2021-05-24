package cidtest

import (
	"crypto/sha256"
	"math/rand"

	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
)

// Generate returns random cid.ID.
func Generate() *cid.ID {
	checksum := [sha256.Size]byte{}

	rand.Read(checksum[:])

	return GenerateWithChecksum(checksum)
}

// GenerateWithChecksum returns cid.ID initialized
// with specified checksum.
func GenerateWithChecksum(cs [sha256.Size]byte) *cid.ID {
	id := cid.New()
	id.SetSHA256(cs)

	return id
}
