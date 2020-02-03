package hash

import (
	"bytes"

	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neofs-api/internal"
	"github.com/nspcc-dev/tzhash/tz"
	"github.com/pkg/errors"
)

// HomomorphicHashSize contains size of HH.
const HomomorphicHashSize = 64

// Hash is implementation of HomomorphicHash.
type Hash [HomomorphicHashSize]byte

// ErrWrongDataSize raised when wrong size of bytes is passed to unmarshal HH.
const ErrWrongDataSize = internal.Error("wrong data size")

var (
	_ internal.Custom = (*Hash)(nil)

	emptyHH [HomomorphicHashSize]byte
)

// Size returns size of Hash (HomomorphicHashSize).
func (h Hash) Size() int { return HomomorphicHashSize }

// Empty checks that Hash is empty.
func (h Hash) Empty() bool { return bytes.Equal(h.Bytes(), emptyHH[:]) }

// Reset sets current Hash to empty value.
func (h *Hash) Reset() { *h = Hash{} }

// ProtoMessage method to satisfy proto.Message interface.
func (h Hash) ProtoMessage() {}

// Bytes represents Hash as bytes.
func (h Hash) Bytes() []byte {
	buf := make([]byte, HomomorphicHashSize)
	copy(buf, h[:])
	return h[:]
}

// Marshal returns bytes representation of Hash.
func (h Hash) Marshal() ([]byte, error) { return h.Bytes(), nil }

// MarshalTo tries to marshal Hash into passed bytes and returns count of copied bytes.
func (h *Hash) MarshalTo(data []byte) (int, error) { return copy(data, h.Bytes()), nil }

// Unmarshal tries to parse bytes into valid Hash.
func (h *Hash) Unmarshal(data []byte) error {
	if ln := len(data); ln != HomomorphicHashSize {
		return errors.Wrapf(ErrWrongDataSize, "expect=%d, actual=%d", HomomorphicHashSize, ln)
	}

	copy((*h)[:], data)
	return nil
}

// String returns string representation of Hash.
func (h Hash) String() string { return base58.Encode(h[:]) }

// Equal checks that current Hash is equal to passed Hash.
func (h Hash) Equal(hash Hash) bool { return h == hash }

// Verify validates if current hash generated from passed data.
func (h Hash) Verify(data []byte) bool { return h.Equal(Sum(data)) }

// Validate checks if combined hashes are equal to current Hash.
func (h Hash) Validate(hashes []Hash) bool {
	hashBytes := make([][]byte, 0, len(hashes))
	for i := range hashes {
		hashBytes = append(hashBytes, hashes[i].Bytes())
	}
	ok, err := tz.Validate(h.Bytes(), hashBytes)
	return err == nil && ok
}

// Sum returns Tillich-Zémor checksum of data.
func Sum(data []byte) Hash { return tz.Sum(data) }

// Concat combines hashes based on homomorphic property.
func Concat(hashes []Hash) (Hash, error) {
	var (
		hash Hash
		h    = make([][]byte, 0, len(hashes))
	)
	for i := range hashes {
		h = append(h, hashes[i].Bytes())
	}
	cat, err := tz.Concat(h)
	if err != nil {
		return hash, err
	}
	return hash, hash.Unmarshal(cat)
}
