package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/object"
)

// Attribute represents v2-compatible object attribute.
type Attribute object.Attribute

// NewAttributeFromV2 wraps v2 Attribute message to Attribute.
func NewAttributeFromV2(aV2 *object.Attribute) *Attribute {
	return (*Attribute)(aV2)
}

// NewAttribute creates and initializes blank Attribute.
//
// Works similar as NewAttributeFromV2(new(Attribute)).
func NewAttribute() *Attribute {
	return NewAttributeFromV2(new(object.Attribute))
}

// GetKey returns key to the object attribute.
func (a *Attribute) GetKey() string {
	return (*object.Attribute)(a).GetKey()
}

// SetKey sets key to the object attribute.
func (a *Attribute) SetKey(v string) {
	(*object.Attribute)(a).SetKey(v)
}

// GetValue return value of the object attribute.
func (a *Attribute) GetValue() string {
	return (*object.Attribute)(a).GetValue()
}

// SetValue sets value of the object attribute.
func (a *Attribute) SetValue(v string) {
	(*object.Attribute)(a).SetValue(v)
}

// ToV2 converts Attribute to v2 Attribute message.
func (a *Attribute) ToV2() *object.Attribute {
	return (*object.Attribute)(a)
}

// Marshal marshals Attribute into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (d *Attribute) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*object.Attribute)(d).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Attribute.
func (d *Attribute) Unmarshal(data []byte) error {
	return (*object.Attribute)(d).
		Unmarshal(data)
}

// MarshalJSON encodes Attribute to protobuf JSON format.
func (d *Attribute) MarshalJSON() ([]byte, error) {
	return (*object.Attribute)(d).
		MarshalJSON()
}

// UnmarshalJSON decodes Attribute from protobuf JSON format.
func (d *Attribute) UnmarshalJSON(data []byte) error {
	return (*object.Attribute)(d).
		UnmarshalJSON(data)
}
