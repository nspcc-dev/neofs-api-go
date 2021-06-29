package storagegroup

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	sg "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
)

func (s *StorageGroup) ToGRPCMessage() neofsgrpc.Message {
	m := new(sg.StorageGroup)

	if s != nil {
		m = new(sg.StorageGroup)

		m.SetMembers(refs.ObjectIDListToGRPCMessage(s.members))
		m.SetExpirationEpoch(s.exp)
		m.SetValidationDataSize(s.size)
		m.SetValidationHash(s.hash.ToGRPCMessage().(*refsGRPC.Checksum))
	}

	return m
}

func (s *StorageGroup) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*sg.StorageGroup)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	hash := v.GetValidationHash()
	if hash == nil {
		s.hash = nil
	} else {
		if s.hash == nil {
			s.hash = new(refs.Checksum)
		}

		err = s.hash.FromGRPCMessage(hash)
		if err != nil {
			return err
		}
	}

	s.members, err = refs.ObjectIDListFromGRPCMessage(v.GetMembers())
	if err != nil {
		return err
	}

	s.exp = v.GetExpirationEpoch()
	s.size = v.GetValidationDataSize()

	return nil
}
