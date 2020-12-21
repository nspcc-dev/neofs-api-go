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

func ObjectIDListToGRPCMessage(ids []*ObjectID) []*refs.ObjectID {
	if ids == nil {
		return nil
	}

	idsV2 := make([]*refs.ObjectID, 0, len(ids))

	for i := range ids {
		idsV2 = append(idsV2, ObjectIDToGRPCMessage(ids[i]))
	}

	return idsV2
}

func ObjectIDListFromGRPCMessage(idsV2 []*refs.ObjectID) []*ObjectID {
	if idsV2 == nil {
		return nil
	}

	ids := make([]*ObjectID, 0, len(idsV2))

	for i := range idsV2 {
		ids = append(ids, ObjectIDFromGRPCMessage(idsV2[i]))
	}

	return ids
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

func ChecksumToGRPCMessage(c *Checksum) *refs.Checksum {
	if c == nil {
		return nil
	}

	m := new(refs.Checksum)

	m.SetChecksumType(refs.ChecksumType(c.GetType()))

	m.SetSum(c.GetSum())

	return m
}

func ChecksumFromGRPCMessage(m *refs.Checksum) *Checksum {
	if m == nil {
		return nil
	}

	c := new(Checksum)

	c.SetType(ChecksumType(m.GetType()))

	c.SetSum(m.GetSum())

	return c
}

func VersionToGRPCMessage(v *Version) *refs.Version {
	if v == nil {
		return nil
	}

	msg := new(refs.Version)

	msg.SetMajor(v.GetMajor())
	msg.SetMinor(v.GetMinor())

	return msg
}

func VersionFromGRPCMessage(m *refs.Version) *Version {
	if m == nil {
		return nil
	}

	v := new(Version)

	v.SetMajor(m.GetMajor())
	v.SetMinor(m.GetMinor())

	return v
}

func SignatureToGRPCMessage(s *Signature) *refs.Signature {
	if s == nil {
		return nil
	}

	m := new(refs.Signature)

	m.SetKey(s.GetKey())
	m.SetSign(s.GetSign())

	return m
}

func SignatureFromGRPCMessage(m *refs.Signature) *Signature {
	if m == nil {
		return nil
	}

	s := new(Signature)

	s.SetKey(m.GetKey())
	s.SetSign(m.GetSign())

	return s
}
