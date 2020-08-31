package owner

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// ID represents owner identifier that
// supports different type of values.
type ID struct {
	val []byte
}

// SetNeo3Wallet sets owner identifier value to NEO3 wallet address.
func (id *ID) SetNeo3Wallet(v *NEO3Wallet) {
	if id != nil {
		id.val = v.Bytes()
	}
}

// ToV2 returns the v2 owner ID message.
func (id *ID) ToV2() *refs.OwnerID {
	if id != nil {
		idV2 := new(refs.OwnerID)
		idV2.SetValue(id.val)

		return idV2
	}

	return nil
}

// IDFromV2 converts owner ID v2 structure to ID.
func IDFromV2(idV2 *refs.OwnerID) (*ID, error) {
	if idV2 == nil {
		return nil, nil
	}

	val := idV2.GetValue()
	if ln := len(val); ln != 25 {
		return nil, errors.Errorf(
			"could not convert %T to %T: expected length %d, received %d",
			idV2, (*ID)(nil), sha256.Size, ln,
		)
	}

	return &ID{
		val: val,
	}, nil
}
