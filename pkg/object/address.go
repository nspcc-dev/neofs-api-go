package object

import (
	"errors"
	"strings"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// Address represents v2-compatible object address.
type Address refs.Address

var errInvalidAddressString = errors.New("incorrect format of the string object address")

const (
	addressParts     = 2
	addressSeparator = "/"
)

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

// ContainerID returns container identifier.
func (a *Address) ContainerID() *container.ID {
	return container.NewIDFromV2(
		(*refs.Address)(a).GetContainerID(),
	)
}

// SetContainerID sets container identifier.
func (a *Address) SetContainerID(id *container.ID) {
	(*refs.Address)(a).SetContainerID(id.ToV2())
}

// ObjectID returns object identifier.
func (a *Address) ObjectID() *ID {
	return NewIDFromV2(
		(*refs.Address)(a).GetObjectID(),
	)
}

// SetObjectID sets object identifier.
func (a *Address) SetObjectID(id *ID) {
	(*refs.Address)(a).SetObjectID(id.ToV2())
}

// Parse converts base58 string representation into Address.
func (a *Address) Parse(s string) error {
	var (
		err   error
		oid   = NewID()
		cid   = container.NewID()
		parts = strings.Split(s, addressSeparator)
	)

	if len(parts) != addressParts {
		return errInvalidAddressString
	} else if err = cid.Parse(parts[0]); err != nil {
		return err
	} else if err = oid.Parse(parts[1]); err != nil {
		return err
	}

	a.SetObjectID(oid)
	a.SetContainerID(cid)

	return nil
}

// String returns string representation of Object.Address.
func (a *Address) String() string {
	return strings.Join([]string{
		a.ContainerID().String(),
		a.ObjectID().String(),
	}, addressSeparator)
}

// Marshal marshals Address into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (a *Address) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*refs.Address)(a).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Address.
func (a *Address) Unmarshal(data []byte) error {
	return (*refs.Address)(a).
		Unmarshal(data)
}

// MarshalJSON encodes Address to protobuf JSON format.
func (a *Address) MarshalJSON() ([]byte, error) {
	return (*refs.Address)(a).
		MarshalJSON()
}

// UnmarshalJSON decodes Address from protobuf JSON format.
func (a *Address) UnmarshalJSON(data []byte) error {
	return (*refs.Address)(a).
		UnmarshalJSON(data)
}
