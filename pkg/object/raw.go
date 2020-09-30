package object

import (
	"github.com/nspcc-dev/neofs-api-go/pkg"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

// RawObject represents v2-compatible NeoFS object that provides
// a convenient interface to fill in the fields of
// an object in isolation from its internal structure.
type RawObject struct {
	*rwObject
}

// NewRawFromV2 wraps v2 Object message to RawObject.
func NewRawFromV2(oV2 *object.Object) *RawObject {
	return &RawObject{
		rwObject: (*rwObject)(oV2),
	}
}

// NewRawFrom wraps Object instance to RawObject.
func NewRawFrom(obj *Object) *RawObject {
	return NewRawFromV2(obj.ToV2())
}

// NewRaw creates and initializes blank RawObject.
//
// Works similar as NewRawFromV2(new(Object)).
func NewRaw() *RawObject {
	return NewRawFromV2(new(object.Object))
}

// Object returns read-only object instance.
func (o *RawObject) Object() *Object {
	if o != nil {
		return &Object{
			rwObject: o.rwObject,
		}
	}

	return nil
}

// SetID sets object identifier.
func (o *RawObject) SetID(v *ID) {
	o.setID(v)
}

// SetSignature sets signature of the object identifier.
func (o *RawObject) SetSignature(v *pkg.Signature) {
	o.setSignature(v)
}

// SetPayload sets payload bytes.
func (o *RawObject) SetPayload(v []byte) {
	o.setPayload(v)
}

// SetVersion sets version of the object.
func (o *RawObject) SetVersion(v *pkg.Version) {
	o.setVersion(v)
}

// SetPayloadSize sets payload length of the object.
func (o *RawObject) SetPayloadSize(v uint64) {
	o.setPayloadSize(v)
}

// SetContainerID sets identifier of the related container.
func (o *RawObject) SetContainerID(v *container.ID) {
	o.setContainerID(v)
}

// SetOwnerID sets identifier of the object owner.
func (o *RawObject) SetOwnerID(v *owner.ID) {
	o.setOwnerID(v)
}

// SetCreationEpoch sets epoch number in which object was created.
func (o *RawObject) SetCreationEpoch(v uint64) {
	o.setCreationEpoch(v)
}

// SetPayloadChecksum sets checksum of the object payload.
func (o *RawObject) SetPayloadChecksum(v *pkg.Checksum) {
	o.setPayloadChecksum(v)
}

// SetPayloadHomomorphicHash sets homomorphic hash of the object payload.
func (o *RawObject) SetPayloadHomomorphicHash(v *pkg.Checksum) {
	o.setPayloadHomomorphicHash(v)
}

// SetAttributes sets object attributes.
func (o *RawObject) SetAttributes(v ...*Attribute) {
	o.setAttributes(v...)
}

// SetPreviousID sets identifier of the previous sibling object.
func (o *RawObject) SetPreviousID(v *ID) {
	o.setPreviousID(v)
}

// SetChildren sets list of the identifiers of the child objects.
func (o *RawObject) SetChildren(v ...*ID) {
	o.setChildren(v...)
}

// SetParentID sets identifier of the parent object.
func (o *RawObject) SetParentID(v *ID) {
	o.setParentID(v)
}

// SetParent sets parent object w/o payload.
func (o *RawObject) SetParent(v *Object) {
	o.setParent(v)
}

// SetSessionToken sets token of the session
// within which object was created.
func (o *RawObject) SetSessionToken(v *token.SessionToken) {
	o.setSessionToken(v)
}

// SetType sets type of the object.
func (o *RawObject) SetType(v Type) {
	o.setType(v)
}

// CutPayload returns RawObject w/ empty payload.
//
// Changes of non-payload fields affect source object.
func (o *RawObject) CutPayload() *RawObject {
	if o != nil {
		return &RawObject{
			rwObject: o.rwObject.cutPayload(),
		}
	}

	return nil
}
