package container

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents container identifier
// that supports different type of values.
type ID struct {
	val []byte
}

// SetSHA256 sets container identifier value to SHA256 checksum.
func (id *ID) SetSHA256(v [sha256.Size]byte) {
	if id != nil {
		id.val = v[:]
	}
}

// ToV2 returns the v2 container ID message.
func (id *ID) ToV2() *refs.ContainerID {
	if id != nil {
		idV2 := new(refs.ContainerID)
		idV2.SetValue(id.val)

		return idV2
	}

	return nil
}

func IDFromV2(idV2 *refs.ContainerID) (*ID, error) {
	val := idV2.GetValue()
	if ln := len(val); ln != sha256.Size {
		return nil, errors.Errorf(
			"could not convert %T to %T: expected length %d, received %d",
			idV2, (*ID)(nil), sha256.Size, ln,
		)
	}

	return &ID{
		val: val,
	}, nil
}
