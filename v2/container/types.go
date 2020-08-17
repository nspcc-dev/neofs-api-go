package container

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type Attribute struct {
	key, val string
}

type Container struct {
	version *service.Version

	ownerID *refs.OwnerID

	nonce []byte

	basicACL uint32

	attr []*Attribute

	policy *netmap.PlacementPolicy
}

func (a *Attribute) GetKey() string {
	if a != nil {
		return a.key
	}

	return ""
}

func (a *Attribute) SetKey(v string) {
	if a != nil {
		a.key = v
	}
}

func (a *Attribute) GetValue() string {
	if a != nil {
		return a.val
	}

	return ""
}

func (a *Attribute) SetValue(v string) {
	if a != nil {
		a.val = v
	}
}

func (c *Container) GetVersion() *service.Version {
	if c != nil {
		return c.version
	}

	return nil
}

func (c *Container) SetVersion(v *service.Version) {
	if c != nil {
		c.version = v
	}
}

func (c *Container) GetOwnerID() *refs.OwnerID {
	if c != nil {
		return c.ownerID
	}

	return nil
}

func (c *Container) SetOwnerID(v *refs.OwnerID) {
	if c != nil {
		c.ownerID = v
	}
}

func (c *Container) GetNonce() []byte {
	if c != nil {
		return c.nonce
	}

	return nil
}

func (c *Container) SetNonce(v []byte) {
	if c != nil {
		c.nonce = v
	}
}

func (c *Container) GetBasicACL() uint32 {
	if c != nil {
		return c.basicACL
	}

	return 0
}

func (c *Container) SetBasicACL(v uint32) {
	if c != nil {
		c.basicACL = v
	}
}

func (c *Container) GetAttributes() []*Attribute {
	if c != nil {
		return c.attr
	}

	return nil
}

func (c *Container) SetAttributes(v []*Attribute) {
	if c != nil {
		c.attr = v
	}
}

func (c *Container) GetPlacementPolicy() *netmap.PlacementPolicy {
	if c != nil {
		return c.policy
	}

	return nil
}

func (c *Container) SetPlacementPolicy(v *netmap.PlacementPolicy) {
	if c != nil {
		c.policy = v
	}
}
