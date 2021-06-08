package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

type (
	Attribute  container.Attribute
	Attributes []*Attribute
)

func NewAttribute() *Attribute {
	return NewAttributeFromV2(new(container.Attribute))
}

func (a *Attribute) SetKey(v string) {
	(*container.Attribute)(a).SetKey(v)
}

func (a *Attribute) SetValue(v string) {
	(*container.Attribute)(a).SetValue(v)
}

func (a *Attribute) Key() string {
	return (*container.Attribute)(a).GetKey()
}

func (a *Attribute) Value() string {
	return (*container.Attribute)(a).GetValue()
}

// NewAttributeFromV2 wraps protocol dependent version of
// Attribute message.
//
// Nil container.Attribute converts to nil.
func NewAttributeFromV2(v *container.Attribute) *Attribute {
	return (*Attribute)(v)
}

// ToV2 converts Attribute to v2 Attribute message.
//
// Nil Attribute converts to nil.
func (a *Attribute) ToV2() *container.Attribute {
	return (*container.Attribute)(a)
}

func NewAttributesFromV2(v []*container.Attribute) Attributes {
	attrs := make(Attributes, 0, len(v))
	for i := range v {
		attrs = append(attrs, NewAttributeFromV2(v[i]))
	}

	return attrs
}

func (a Attributes) ToV2() []*container.Attribute {
	attrs := make([]*container.Attribute, 0, len(a))
	for i := range a {
		attrs = append(attrs, a[i].ToV2())
	}

	return attrs
}
