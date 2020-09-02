package object

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/pkg/errors"
)

// Address represents address of NeoFS object.
type Address struct {
	cid *container.ID

	oid *ID
}

// ToV2 converts Address to v2 Address message.
func (a *Address) ToV2() *refs.Address {
	if a != nil {
		aV2 := new(refs.Address)
		aV2.SetContainerID(a.cid.ToV2())
		aV2.SetObjectID(a.oid.ToV2())

		return aV2
	}

	return nil
}

// GetObjectIDV2 converts object identifier to v2 ObjectID message.
func (a *Address) GetObjectIDV2() *refs.ObjectID {
	if a != nil {
		return a.oid.ToV2()
	}

	return nil
}

// AddressFromV2 converts v2 Address message to Address.
func AddressFromV2(aV2 *refs.Address) (*Address, error) {
	if aV2 == nil {
		return nil, nil
	}

	oid, err := IDFromV2(aV2.GetObjectID())
	if err != nil {
		return nil, errors.Wrap(err, "could not convert object identifier")
	}

	cid, err := container.IDFromV2(aV2.GetContainerID())
	if err != nil {
		return nil, errors.Wrap(err, "could not convert container identifier")
	}

	return &Address{
		cid: cid,
		oid: oid,
	}, nil
}

// AddressFromBytes restores Address from a binary representation.
func AddressFromBytes(data []byte) (*Address, error) {
	addrV2 := new(refs.Address)
	if err := addrV2.StableUnmarshal(data); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal object address")
	}

	return AddressFromV2(addrV2)
}
