package object

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents object identifier that
// supports different type of values.
type ID struct {
	val []byte
}

// SetSHA256 sets object identifier value to SHA256 checksum.
func (id *ID) SetSHA256(v [sha256.Size]byte) {
	if id != nil {
		id.val = v[:]
	}
}

// ToV2 converts ID to v2 ObjectID message.
func (id *ID) ToV2() *refs.ObjectID {
	if id != nil {
		idV2 := new(refs.ObjectID)
		idV2.SetValue(id.val)

		return idV2
	}

	return nil
}

// IDFromV2 converts v2 ObjectID message to ID.
//
// Returns an error if the format of the identifier
// in the message is broken.
func IDFromV2(idV2 *refs.ObjectID) (*ID, error) {
	val := idV2.GetValue()
	if ln := len(val); ln != sha256.Size {
		return nil, errors.Errorf("could not convert %T to %T: invalid length %d",
			idV2, (*ID)(nil), ln,
		)
	}

	return &ID{
		val: val,
	}, nil
}
