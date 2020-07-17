package refs

import (
	"bytes"
	"crypto/sha256"

	"github.com/gogo/protobuf/proto"
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
)

// CIDForBytes creates CID for passed bytes.
func CIDForBytes(data []byte) CID { return sha256.Sum256(data) }

// CIDFromBytes parses CID from passed bytes.
func CIDFromBytes(data []byte) (cid CID, err error) {
	if ln := len(data); ln != CIDSize {
		return CID{}, errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", CIDSize, ln)
	}

	copy(cid[:], data)
	return
}

// CIDFromString parses CID from string representation of CID.
func CIDFromString(c string) (CID, error) {
	var cid CID
	decoded, err := base58.Decode(c)
	if err != nil {
		return cid, err
	}
	return CIDFromBytes(decoded)
}

// Size returns size of CID (CIDSize).
func (c CID) Size() int { return CIDSize }

// Parse tries to parse CID from string representation.
func (c *CID) Parse(cid string) error {
	var err error
	if *c, err = CIDFromString(cid); err != nil {
		return err
	}
	return nil
}

// Empty checks that current CID is empty.
func (c CID) Empty() bool { return bytes.Equal(c.Bytes(), emptyCID) }

// Equal checks that current CID is equal to passed CID.
func (c CID) Equal(cid CID) bool { return bytes.Equal(c.Bytes(), cid.Bytes()) }

// Marshal returns CID bytes representation.
func (c CID) Marshal() ([]byte, error) { return c.Bytes(), nil }

// MarshalBinary returns CID bytes representation.
func (c CID) MarshalBinary() ([]byte, error) { return c.Bytes(), nil }

// MarshalTo marshal CID to bytes representation into passed bytes.
func (c *CID) MarshalTo(data []byte) (int, error) { return copy(data, c.Bytes()), nil }

// ProtoMessage method to satisfy proto.Message interface.
func (c CID) ProtoMessage() {}

// String returns string representation of CID.
func (c CID) String() string { return base58.Encode(c[:]) }

// Reset resets current CID to zero value.
func (c *CID) Reset() { *c = CID{} }

// Bytes returns CID bytes representation.
func (c CID) Bytes() []byte {
	buf := make([]byte, CIDSize)
	copy(buf, c[:])
	return buf
}

// UnmarshalBinary tries to parse bytes representation of CID.
func (c *CID) UnmarshalBinary(data []byte) error { return c.Unmarshal(data) }

// Unmarshal tries to parse bytes representation of CID.
func (c *CID) Unmarshal(data []byte) error {
	if ln := len(data); ln != CIDSize {
		return errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", CIDSize, ln)
	}

	copy((*c)[:], data)
	return nil
}

// Verify validates that current CID is generated for passed bytes data.
func (c CID) Verify(data []byte) error {
	if id := CIDForBytes(data); !bytes.Equal(c[:], id[:]) {
		return errors.New("wrong hash for data")
	}
	return nil
}

// Merge used by proto.Clone
func (c *CID) Merge(src proto.Message) {
	if cid, ok := src.(*CID); ok {
		*c = *cid
	}
}
