package object

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// ID represents v2-compatible object identifier.
type ID refs.ObjectID

var errInvalidIDString = errors.New("incorrect format of the string object ID")

// NewIDFromV2 wraps v2 ObjectID message to ID.
//
// Nil refs.ObjectID converts to nil.
func NewIDFromV2(idV2 *refs.ObjectID) *ID {
	return (*ID)(idV2)
}

// NewID creates and initializes blank ID.
//
// Works similar as NewIDFromV2(new(ObjectID)).
//
// Defaults:
// 	- value: nil.
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
//
// Nil ID converts to nil.
func (id *ID) ToV2() *refs.ObjectID {
	return (*refs.ObjectID)(id)
}

// Parse converts base58 string representation into ID.
func (id *ID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return fmt.Errorf("could not parse object.ID from string: %w", err)
	} else if len(data) != sha256.Size {
		return errInvalidIDString
	}

	(*refs.ObjectID)(id).SetValue(data)

	return nil
}

// String returns base58 string representation of ID.
func (id *ID) String() string {
	return base58.Encode((*refs.ObjectID)(id).GetValue())
}

// Marshal marshals ID into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (id *ID) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*refs.ObjectID)(id).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of ID.
func (id *ID) Unmarshal(data []byte) error {
	return (*refs.ObjectID)(id).
		Unmarshal(data)
}

// MarshalJSON encodes ID to protobuf JSON format.
func (id *ID) MarshalJSON() ([]byte, error) {
	return (*refs.ObjectID)(id).
		MarshalJSON()
}

// UnmarshalJSON decodes ID from protobuf JSON format.
func (id *ID) UnmarshalJSON(data []byte) error {
	return (*refs.ObjectID)(id).
		UnmarshalJSON(data)
}
