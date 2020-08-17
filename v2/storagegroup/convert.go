package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	sg "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
)

// StorageGroupToGRPCMessage converts unified proto structure into grpc structure.
func StorageGroupToGRPCMessage(s *StorageGroup) *sg.StorageGroup {
	if s == nil {
		return nil
	}

	m := new(sg.StorageGroup)

	m.SetValidationDataSize(s.GetValidationDataSize())
	m.SetValidationHash(s.GetValidationHash())
	m.SetExpirationEpoch(s.GetExpirationEpoch())

	members := s.GetMembers()
	memberMsg := make([]*refsGRPC.ObjectID, 0, len(members))

	for i := range members {
		memberMsg = append(memberMsg, refs.ObjectIDToGRPCMessage(members[i]))
	}

	m.SetMembers(memberMsg)

	return m
}

// StorageGroupFromGRPCMessage converts grpc structure into unified proto structure.
func StorageGroupFromGRPCMessage(m *sg.StorageGroup) *StorageGroup {
	if m == nil {
		return nil
	}

	s := new(StorageGroup)

	s.SetValidationDataSize(m.GetValidationDataSize())
	s.SetValidationHash(m.GetValidationHash())
	s.SetExpirationEpoch(m.GetExpirationEpoch())

	memberMsg := m.GetMembers()
	members := make([]*refs.ObjectID, 0, len(memberMsg))

	for i := range memberMsg {
		members = append(members, refs.ObjectIDFromGRPCMessage(memberMsg[i]))
	}

	s.SetMembers(members)

	return s
}
