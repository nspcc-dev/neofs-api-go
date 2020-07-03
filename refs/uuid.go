package refs

import (
	"bytes"
	"encoding/hex"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

// Size returns size in bytes of UUID (UUIDSize).
func (UUID) Size() int { return UUIDSize }

// Empty checks that current UUID is empty value.
func (u UUID) Empty() bool { return bytes.Equal(u.Bytes(), emptyUUID) }

// Reset sets current UUID to empty value.
func (u *UUID) Reset() { *u = [UUIDSize]byte{} }

// ProtoMessage method to satisfy proto.Message.
func (UUID) ProtoMessage() {}

// Marshal returns UUID bytes representation.
func (u UUID) Marshal() ([]byte, error) { return u.Bytes(), nil }

// MarshalTo returns UUID bytes representation.
func (u UUID) MarshalTo(data []byte) (int, error) { return copy(data, u[:]), nil }

// Bytes returns UUID bytes representation.
func (u UUID) Bytes() []byte {
	buf := make([]byte, UUIDSize)
	copy(buf, u[:])
	return buf
}

// Equal checks that current UUID is equal to passed UUID.
func (u UUID) Equal(u2 UUID) bool { return bytes.Equal(u.Bytes(), u2.Bytes()) }

func (u UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], u)
	return string(buf[:])
}

// Unmarshal tries to parse UUID bytes representation.
func (u *UUID) Unmarshal(data []byte) error {
	if ln := len(data); ln != UUIDSize {
		return errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", UUIDSize, ln)
	}

	copy((*u)[:], data)
	return nil
}

// Parse tries to parse UUID string representation.
func (u *UUID) Parse(id string) error {
	tmp, err := uuid.Parse(id)
	if err != nil {
		return errors.Wrapf(err, "could not parse `%s`", id)
	}

	copy((*u)[:], tmp[:])
	return nil
}

// Merge used by proto.Clone
func (u *UUID) Merge(src proto.Message) {
	if tmp, ok := src.(*UUID); ok {
		*u = *tmp
	}
}
