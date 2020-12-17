package container

import (
	"bytes"
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents v2-compatible container identifier.
type ID refs.ContainerID

// ErrIDMismatch is returned when container structure does not match
// a specific identifier.
var ErrIDMismatch = errors.New("container structure does not match the identifier")

var errInvalidIDString = errors.New("incorrect format of the string container ID")

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

// Equal returns true if identifiers are identical.
func (id *ID) Equal(id2 *ID) bool {
	return bytes.Equal(
		(*refs.ContainerID)(id).GetValue(),
		(*refs.ContainerID)(id2).GetValue(),
	)
}

// Parse converts base58 string representation into ID.
func (id *ID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return errors.Wrap(err, "could not parse container.ID from string")
	} else if len(data) != sha256.Size {
		return errInvalidIDString
	}

	(*refs.ContainerID)(id).SetValue(data)

	return nil
}

// String returns base58 string representation of ID.
func (id *ID) String() string {
	return base58.Encode((*refs.ContainerID)(id).GetValue())
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

	return (*refs.ContainerID)(id).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of ID.
func (id *ID) Unmarshal(data []byte) error {
	return (*refs.ContainerID)(id).
		Unmarshal(data)
}

// MarshalJSON encodes ID to protobuf JSON format.
func (id *ID) MarshalJSON() ([]byte, error) {
	return (*refs.ContainerID)(id).
		MarshalJSON()
}

// UnmarshalJSON decodes ID from protobuf JSON format.
func (id *ID) UnmarshalJSON(data []byte) error {
	return (*refs.ContainerID)(id).
		UnmarshalJSON(data)
}
