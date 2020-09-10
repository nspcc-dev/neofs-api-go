package container

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// ID represents v2-compatible container identifier.
type ID refs.ContainerID

// NewIDFromV2 wraps v2 ContainerID message to ID.
func NewIDFromV2(idV2 *refs.ContainerID) *ID {
	return (*ID)(idV2)
}

// NewID creates and initializes blank ID.
//
// Works similar to NewIDFromV2(new(ContainerID)).
func NewID() *ID {
	return NewIDFromV2(new(refs.ContainerID))
}

// SetSHA256 sets container identifier value to SHA256 checksum.
func (id *ID) SetSHA256(v [sha256.Size]byte) {
	(*refs.ContainerID)(id).SetValue(v[:])
}

// ToV2 returns the v2 container ID message.
func (id *ID) ToV2() *refs.ContainerID {
	return (*refs.ContainerID)(id)
}
