package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

// Object represents v2-compatible NeoFS object that provides
// a convenient interface for working in isolation
// from the internal structure of an object.
//
// Object allows to work with the object in read-only
// mode as a reflection of the immutability of objects
// in the system.
type Object struct {
	*rwObject
}

// NewFromV2 wraps v2 Object message to Object.
func NewFromV2(oV2 *object.Object) *Object {
	return &Object{
		rwObject: (*rwObject)(oV2),
	}
}

// New creates and initializes blank Object.
//
// Works similar as NewFromV2(new(Object)).
func New() *Object {
	return NewFromV2(new(object.Object))
}

// ToV2 converts Object to v2 Object message.
func (o *Object) ToV2() *object.Object {
	if o != nil {
		return (*object.Object)(o.rwObject)
	}

	return nil
}

// FromBytes restores Object instance from a binary representation.
func FromBytes(data []byte) (*Object, error) {
	oV2 := new(object.Object)

	if err := oV2.StableUnmarshal(data); err != nil {
		return nil, err
	}

	return NewFromV2(oV2), nil
}
