package pkg

import (
	"crypto/sha256"

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
func NewChecksumFromV2(cV2 *refs.Checksum) *Checksum {
	return (*Checksum)(cV2)
}

// NewVersion creates and initializes blank Version.
//
// Works similar as NewVersionFromV2(new(Version)).
func NewChecksum() *Checksum {
	return NewChecksumFromV2(new(refs.Checksum))
}

// GetType returns checksum type.
func (c *Checksum) GetType() ChecksumType {
	switch (*refs.Checksum)(c).GetType() {
	case refs.SHA256:
		return ChecksumSHA256
	case refs.TillichZemor:
		return ChecksumTZ
	default:
		return ChecksumUnknown
	}
}

// GetSum returns checksum bytes.
func (c *Checksum) GetSum() []byte {
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

func (c *Checksum) ToV2() *refs.Checksum {
	return (*refs.Checksum)(c)
}
