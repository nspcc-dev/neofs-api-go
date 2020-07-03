package refs

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api-go/chain"
	"github.com/pkg/errors"
)

// NewOwnerID returns generated OwnerID from passed public keys.
func NewOwnerID(keys ...*ecdsa.PublicKey) (owner OwnerID, err error) {
	if len(keys) == 0 {
		return
	}
	var d []byte
	d, err = base58.Decode(chain.KeysToAddress(keys...))
	if err != nil {
		return
	}
	copy(owner[:], d)
	return owner, nil
}

// Size returns OwnerID size in bytes (OwnerIDSize).
func (OwnerID) Size() int { return OwnerIDSize }

// Empty checks that current OwnerID is empty value.
func (o OwnerID) Empty() bool { return bytes.Equal(o.Bytes(), emptyOwner) }

// Equal checks that current OwnerID is equal to passed OwnerID.
func (o OwnerID) Equal(id OwnerID) bool { return bytes.Equal(o.Bytes(), id.Bytes()) }

// Reset sets current OwnerID to empty value.
func (o *OwnerID) Reset() { *o = OwnerID{} }

// ProtoMessage method to satisfy proto.Message interface.
func (OwnerID) ProtoMessage() {}

// Marshal returns OwnerID bytes representation.
func (o OwnerID) Marshal() ([]byte, error) { return o.Bytes(), nil }

// MarshalTo copies OwnerID bytes representation into passed slice of bytes.
func (o OwnerID) MarshalTo(data []byte) (int, error) { return copy(data, o.Bytes()), nil }

// String returns string representation of OwnerID.
func (o OwnerID) String() string { return base58.Encode(o[:]) }

// Bytes returns OwnerID bytes representation.
func (o OwnerID) Bytes() []byte {
	buf := make([]byte, OwnerIDSize)
	copy(buf, o[:])
	return buf
}

// Unmarshal tries to parse OwnerID bytes representation into current OwnerID.
func (o *OwnerID) Unmarshal(data []byte) error {
	if ln := len(data); ln != OwnerIDSize {
		return errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", OwnerIDSize, ln)
	}

	copy((*o)[:], data)
	return nil
}

// Merge used by proto.Clone
func (o *OwnerID) Merge(src proto.Message) {
	if uid, ok := src.(*OwnerID); ok {
		*o = *uid
	}
}
