package object

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// Address represents v2-compatible object address.
type Address refs.Address

// NewAddressFromV2 converts v2 Address message to Address.
func NewAddressFromV2(aV2 *refs.Address) *Address {
	return (*Address)(aV2)
}

// NewAddress creates and initializes blank Address.
//
// Works similar as NewAddressFromV2(new(Address)).
func NewAddress() *Address {
	return NewAddressFromV2(new(refs.Address))
}

// ToV2 converts Address to v2 Address message.
func (a *Address) ToV2() *refs.Address {
	return (*refs.Address)(a)
}

// AddressFromBytes restores Address from a binary representation.
func AddressFromBytes(data []byte) (*Address, error) {
	addrV2 := new(refs.Address)
	if err := addrV2.StableUnmarshal(data); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal object address")
	}

	return NewAddressFromV2(addrV2), nil
}

// GetContainerID returns container identifier.
func (a *Address) GetContainerID() *container.ID {
	return container.NewIDFromV2(
		(*refs.Address)(a).GetContainerID(),
	)
}

// SetContainerID sets container identifier.
func (a *Address) SetContainerID(id *container.ID) {
	(*refs.Address)(a).SetContainerID(id.ToV2())
}

// GetObjectID returns object identifier.
func (a *Address) GetObjectID() *ID {
	return NewIDFromV2(
		(*refs.Address)(a).GetObjectID(),
	)
}

// SetObjectID sets object identifier.
func (a *Address) SetObjectID(id *ID) {
	(*refs.Address)(a).SetObjectID(id.ToV2())
}
