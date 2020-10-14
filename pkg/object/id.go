package object

import (
	"bytes"
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/internal"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents v2-compatible object identifier.
type ID refs.ObjectID

// ErrBadID should be returned when bytes slice hasn't sha256.Size
// Notice: if byte slice changed, please, replace error message.
const ErrBadID = internal.Error("object.ID should be 32 bytes length")

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

// Parse converts base58 string representation into ID.
func (id *ID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return errors.Wrap(err, "could not parse object.ID from string")
	} else if len(data) != sha256.Size {
		return ErrBadID
	}

	(*refs.ObjectID)(id).SetValue(data)

	return nil
}

// String returns base58 string representation of ID.
func (id *ID) String() string {
	return base58.Encode((*refs.ObjectID)(id).GetValue())
}
