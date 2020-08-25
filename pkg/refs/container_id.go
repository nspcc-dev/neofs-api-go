package refs

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
)

type (
	ContainerID [sha256.Size]byte
)

func (c ContainerID) String() string {
	return base58.Encode(c[:])
}
