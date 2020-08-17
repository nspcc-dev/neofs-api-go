package refs

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

func OwnerIDToGRPCMessage(o *OwnerID) *refs.OwnerID {
	if o == nil {
		return nil
	}

	m := new(refs.OwnerID)

	m.SetValue(o.GetValue())

	return m
}

func OwnerIDFromGRPCMessage(m *refs.OwnerID) *OwnerID {
	if m == nil {
		return nil
	}

	o := new(OwnerID)

	o.SetValue(m.GetValue())

	return o
}

func ContainerIDToGRPCMessage(c *ContainerID) *refs.ContainerID {
	if c == nil {
		return nil
	}

	m := new(refs.ContainerID)

	m.SetValue(c.GetValue())

	return m
}

func ContainerIDFromGRPCMessage(m *refs.ContainerID) *ContainerID {
	if m == nil {
		return nil
	}

	c := new(ContainerID)

	c.SetValue(m.GetValue())

	return c
}

func ObjectIDToGRPCMessage(o *ObjectID) *refs.ObjectID {
	if o == nil {
		return nil
	}

	m := new(refs.ObjectID)

	m.SetValue(o.GetValue())

	return m
}

func ObjectIDFromGRPCMessage(m *refs.ObjectID) *ObjectID {
	if m == nil {
		return nil
	}

	o := new(ObjectID)

	o.SetValue(m.GetValue())

	return o
}

func AddressToGRPCMessage(a *Address) *refs.Address {
	if a == nil {
		return nil
	}

	m := new(refs.Address)

	m.SetContainerId(
		ContainerIDToGRPCMessage(a.GetContainerID()),
	)

	m.SetObjectId(
		ObjectIDToGRPCMessage(a.GetObjectID()),
	)

	return m
}

func AddressFromGRPCMessage(m *refs.Address) *Address {
	if m == nil {
		return nil
	}

	a := new(Address)

	a.SetContainerID(
		ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	a.SetObjectID(
		ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	return a
}
