package object

import (
	"bytes"
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// ID represents v2-compatible object identifier.
type ID refs.ObjectID

// NewIDFromV2 wraps v2 ObjectID message to ID.
func NewIDFromV2(idV2 *refs.ObjectID) *ID {
	return (*ID)(idV2)
}

// NewID creates and initializes blank ID.
//
// Works similar as NewIDFromV2(new(ObjectID)).
func NewID() *ID {
	return NewIDFromV2(new(refs.ObjectID))
}

// SetSHA256 sets object identifier value to SHA256 checksum.
func (id *ID) SetSHA256(v [sha256.Size]byte) {
	(*refs.ObjectID)(id).SetValue(v[:])
}

// Equal returns true if identifiers are identical.
func (id *ID) Equal(id2 *ID) bool {
	return bytes.Equal(
		(*refs.ObjectID)(id).GetValue(),
		(*refs.ObjectID)(id2).GetValue(),
	)
}

// ToV2 converts ID to v2 ObjectID message.
func (id *ID) ToV2() *refs.ObjectID {
	return (*refs.ObjectID)(id)
}
