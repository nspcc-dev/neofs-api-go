package container

import (
	"crypto/sha256"

	"github.com/nspcc-dev/neofs-api-go/v2/container"
)

type Container struct {
	container.Container
}

func New(opts ...NewOption) *Container {
	cnrOptions := defaultContainerOptions()
	for i := range opts {
		opts[i].apply(&cnrOptions)
	}

	cnr := new(Container)
	cnr.SetNonce(cnrOptions.nonce[:])
	cnr.SetBasicACL(cnrOptions.acl)

	if cnrOptions.owner != nil {
		cnr.SetOwnerID(cnrOptions.owner.ToV2())
	}

	if cnrOptions.policy != nil {
		cnr.SetPlacementPolicy(cnrOptions.policy)
	}

	attributes := make([]*container.Attribute, len(cnrOptions.attributes))
	for i := range cnrOptions.attributes {
		attribute := new(container.Attribute)
		attribute.SetKey(cnrOptions.attributes[i].key)
		attribute.SetValue(cnrOptions.attributes[i].value)
		attributes[i] = attribute
	}
	if len(attributes) > 0 {
		cnr.SetAttributes(attributes)
	}

	return cnr
}

func (c Container) ToV2() *container.Container {
	return &c.Container
}

func NewContainerFromV2(c *container.Container) *Container {
	cnr := new(Container)

	if c != nil {
		cnr.Container = *c
	}

	return cnr
}

// CalculateID calculates container identifier
// based on its structure.
func CalculateID(c *Container) *ID {
	data, err := c.ToV2().StableMarshal(nil)
	if err != nil {
		panic(err)
	}

	id := NewID()
	id.SetSHA256(sha256.Sum256(data))

	return id
}
