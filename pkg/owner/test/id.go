package ownertest

import (
	"math/rand"

	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Generate returns owner.ID calculated
// from a random owner.NEO3Wallet.
func Generate() *owner.ID {
	data := make([]byte, owner.NEO3WalletSize)

	rand.Read(data)

	return GenerateFromBytes(data)
}

// GenerateFromBytes returns owner.ID generated
// from a passed byte slice.
func GenerateFromBytes(val []byte) *owner.ID {
	idV2 := new(refs.OwnerID)
	idV2.SetValue(val)

	return owner.NewIDFromV2(idV2)
}
