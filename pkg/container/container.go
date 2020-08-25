package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/container"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

type Container struct {
	container.Container
}

func New(opts ...NewOption) (*Container, error) {
	cnrOptions := defaultContainerOptions()
	for i := range opts {
		opts[i].apply(&cnrOptions)
	}

	cnr := new(Container)
	cnr.SetNonce(cnrOptions.nonce[:])
	cnr.SetBasicACL(cnrOptions.acl)

	if cnrOptions.policy != "" {
		// todo: set placement policy
	}

	if cnrOptions.owner != nil {
		owner := new(refs.OwnerID)
		owner.SetValue(cnrOptions.owner[:])

		cnr.SetOwnerID(owner)
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

	return cnr, nil
}
