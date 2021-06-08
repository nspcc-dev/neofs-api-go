package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Checksum represents v2-compatible checksum.
type Checksum refs.Checksum

// ChecksumType represents the enumeration
// of checksum types.
type ChecksumType uint8

const (
	// ChecksumUnknown is an undefined checksum type.
	ChecksumUnknown ChecksumType = iota

	// ChecksumSHA256 is a SHA256 checksum type.
	ChecksumSHA256

	// ChecksumTZ is a Tillich-Zemor checksum type.
	ChecksumTZ
)

// NewChecksumFromV2 wraps v2 Checksum message to Checksum.
//
// Nil refs.Checksum converts to nil.
func NewChecksumFromV2(cV2 *refs.Checksum) *Checksum {
	return (*Checksum)(cV2)
}

// NewVersion creates and initializes blank Version.
//
// Works similar as NewVersionFromV2(new(Version)).
func NewChecksum() *Checksum {
	return NewChecksumFromV2(new(refs.Checksum))
}

// Type returns checksum type.
func (c *Checksum) Type() ChecksumType {
	switch (*refs.Checksum)(c).GetType() {
	case refs.SHA256:
		return ChecksumSHA256
	case refs.TillichZemor:
		return ChecksumTZ
	default:
		return ChecksumUnknown
	}
}

// Sum returns checksum bytes.
func (c *Checksum) Sum() []byte {
	return (*refs.Checksum)(c).GetSum()
}

// SetSHA256 sets checksum to SHA256 hash.
func (c *Checksum) SetSHA256(v [sha256.Size]byte) {
	checksum := (*refs.Checksum)(c)

	checksum.SetType(refs.SHA256)
	checksum.SetSum(v[:])
}

// SetTillichZemor sets checksum to Tillich-Zemor hash.
func (c *Checksum) SetTillichZemor(v [64]byte) {
	checksum := (*refs.Checksum)(c)

	checksum.SetType(refs.TillichZemor)
	checksum.SetSum(v[:])
}

// ToV2 converts Checksum to v2 Checksum message.
//
// Nil Checksum converts to nil.
func (c *Checksum) ToV2() *refs.Checksum {
	return (*refs.Checksum)(c)
}

func EqualChecksums(cs1, cs2 *Checksum) bool {
	return cs1.Type() == cs2.Type() && bytes.Equal(cs1.Sum(), cs2.Sum())
}

// Marshal marshals Checksum into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (c *Checksum) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*refs.Checksum)(c).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Checksum.
func (c *Checksum) Unmarshal(data []byte) error {
	return (*refs.Checksum)(c).
		Unmarshal(data)
}

// MarshalJSON encodes Checksum to protobuf JSON format.
func (c *Checksum) MarshalJSON() ([]byte, error) {
	return (*refs.Checksum)(c).
		MarshalJSON()
}

// UnmarshalJSON decodes Checksum from protobuf JSON format.
func (c *Checksum) UnmarshalJSON(data []byte) error {
	return (*refs.Checksum)(c).
		UnmarshalJSON(data)
}

func (c *Checksum) String() string {
	return hex.EncodeToString(
		(*refs.Checksum)(c).
			GetSum(),
	)
}

// Parse parses Checksum from its string representation.
func (c *Checksum) Parse(s string) error {
	data, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	var typ refs.ChecksumType

	switch ln := len(data); ln {
	default:
		return fmt.Errorf("unsupported checksum length %d", ln)
	case sha256.Size:
		typ = refs.SHA256
	case 64:
		typ = refs.TillichZemor
	}

	cV2 := (*refs.Checksum)(c)
	cV2.SetType(typ)
	cV2.SetSum(data)

	return nil
}
