package container

import (
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

func AttributeToGRPCMessage(a *Attribute) *container.Container_Attribute {
	if a == nil {
		return nil
	}

	m := new(container.Container_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())

	return m
}

func AttributeFromGRPCMessage(m *container.Container_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	a := new(Attribute)

	a.SetKey(m.GetKey())
	a.SetValue(m.GetValue())

	return a
}

func ContainerToGRPCMessage(c *Container) *container.Container {
	if c == nil {
		return nil
	}

	m := new(container.Container)

	m.SetVersion(
		service.VersionToGRPCMessage(c.GetVersion()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(c.GetOwnerID()),
	)

	m.SetNonce(c.GetNonce())

	m.SetBasicAcl(c.GetBasicACL())

	m.SetPlacementPolicy(
		netmap.PlacementPolicyToGRPCMessage(c.GetPlacementPolicy()),
	)

	attr := c.GetAttributes()
	attrMsg := make([]*container.Container_Attribute, 0, len(attr))

	for i := range attr {
		attrMsg = append(attrMsg, AttributeToGRPCMessage(attr[i]))
	}

	m.SetAttributes(attrMsg)

	return m
}

func ContainerFromGRPCMessage(m *container.Container) *Container {
	if m == nil {
		return nil
	}

	c := new(Container)

	c.SetVersion(
		service.VersionFromGRPCMessage(m.GetVersion()),
	)

	c.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	c.SetNonce(m.GetNonce())

	c.SetBasicACL(m.GetBasicAcl())

	c.SetPlacementPolicy(
		netmap.PlacementPolicyFromGRPCMessage(m.GetPlacementPolicy()),
	)

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	c.SetAttributes(attr)

	return c
}
