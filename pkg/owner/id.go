package owner

import (
	"github.com/mr-tron/base58"
	"github.com/nspcc-dev/neo-go/pkg/encoding/address"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// ID represents v2-compatible owner identifier.
type ID refs.OwnerID

// NewIDFromV2 wraps v2 OwnerID message to ID.
func NewIDFromV2(idV2 *refs.OwnerID) *ID {
	return (*ID)(idV2)
}

// NewID creates and initializes blank ID.
//
// Works similar as NewIDFromV2(new(OwnerID)).
func NewID() *ID {
	return NewIDFromV2(new(refs.OwnerID))
}

// SetNeo3Wallet sets owner identifier value to NEO3 wallet address.
func (id *ID) SetNeo3Wallet(v *NEO3Wallet) {
	(*refs.OwnerID)(id).SetValue(v.Bytes())
}

// ToV2 returns the v2 owner ID message.
func (id *ID) ToV2() *refs.OwnerID {
	return (*refs.OwnerID)(id)
}

func (id *ID) String() string {
	return base58.Encode((*refs.OwnerID)(id).GetValue())
}

func ScriptHashBE(id *ID) ([]byte, error) {
	addr, err := address.StringToUint160(id.String())
	if err != nil {
		return nil, err
	}

	return addr.BytesBE(), nil
}

// NewIDFromNeo3Wallet creates new owner identity from 25-byte neo wallet.
func NewIDFromNeo3Wallet(v *NEO3Wallet) *ID {
	id := NewID()
	id.SetNeo3Wallet(v)

	return id
}
