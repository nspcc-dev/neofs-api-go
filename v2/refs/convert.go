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
