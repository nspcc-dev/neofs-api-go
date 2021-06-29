package tombstone

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	tombstone "github.com/nspcc-dev/neofs-api-go/v2/tombstone/grpc"
)

func (s *Tombstone) ToGRPCMessage() neofsgrpc.Message {
	var m *tombstone.Tombstone

	if s != nil {
		m = new(tombstone.Tombstone)

		m.SetMembers(refs.ObjectIDListToGRPCMessage(s.members))
		m.SetExpirationEpoch(s.exp)
		m.SetSplitId(s.splitID)
	}

	return m
}

func (s *Tombstone) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*tombstone.Tombstone)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	s.members, err = refs.ObjectIDListFromGRPCMessage(v.GetMembers())
	if err != nil {
		return err
	}

	s.exp = v.GetExpirationEpoch()
	s.splitID = v.GetSplitId()

	return nil
}
