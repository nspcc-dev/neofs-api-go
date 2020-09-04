package object

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// RawObject represents NeoFS object that provides
// a convenient interface to fill in the fields of
// an object in isolation from its internal structure.
type RawObject struct {
	rwObject
}

func (o *RawObject) set(setter func()) {
	o.fin = false
	setter()
}

// SetContainerID sets object's container identifier.
func (o *RawObject) SetContainerID(v *container.ID) {
	if o != nil {
		o.set(func() {
			o.cid = v
		})
	}
}

// SetOwnerID sets identifier of the object's owner.
func (o *RawObject) SetOwnerID(v *owner.ID) {
	if o != nil {
		o.set(func() {
			o.ownerID = v
		})
	}
}

// Release returns read-only Object instance.
func (o *RawObject) Release() *Object {
	if o != nil {
		return &Object{
			rwObject: o.rwObject,
		}
	}

	return nil
}

// SetPayloadChecksumSHA256 sets payload checksum as a SHA256 checksum.
func (o *RawObject) SetPayloadChecksumSHA256(v [sha256.Size]byte) {
	if o != nil {
		o.set(func() {
			if o.payloadChecksum == nil {
				o.payloadChecksum = new(refs.Checksum)
			}

			o.payloadChecksum.SetType(refs.SHA256)
			o.payloadChecksum.SetSum(v[:])
		})
	}
}
