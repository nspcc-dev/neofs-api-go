package cid

import (
	"bytes"
	"crypto/sha256"
	"errors"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// ID represents v2-compatible container identifier.
type ID refs.ContainerID

// NewFromV2 wraps v2 ContainerID message to ID.
//
// Nil refs.ContainerID converts to nil.
func NewFromV2(idV2 *refs.ContainerID) *ID {
	return (*ID)(idV2)
}

// New creates and initializes blank ID.
func New() *ID {
	return NewFromV2(new(refs.ContainerID))
}

// SetSHA256 sets container identifier value to SHA256 checksum.
func (id *ID) SetSHA256(v [sha256.Size]byte) {
	(*refs.ContainerID)(id).SetValue(v[:])
}

// ToV2 returns the v2 container ID message.
//
// Nil Result converts to nil.
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

// Parse parses string representation of ID.
//
// Returns error if s is not a base58 encoded
// ID data.
func (id *ID) Parse(s string) error {
	data, err := base58.Decode(s)
	if err != nil {
		return err
	} else if len(data) != sha256.Size {
		return errors.New("incorrect format of the string container ID")
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
