package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	sg "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
)

// StorageGroupToGRPCMessage converts unified proto structure into grpc structure.
func StorageGroupToGRPCMessage(s *StorageGroup) *sg.StorageGroup {
	if s == nil {
		return nil
	}

	m := new(sg.StorageGroup)

	m.SetValidationDataSize(s.GetValidationDataSize())
	m.SetValidationHash(
		refs.ChecksumToGRPCMessage(s.GetValidationHash()),
	)
	m.SetExpirationEpoch(s.GetExpirationEpoch())

	m.SetMembers(
		refs.ObjectIDListToGRPCMessage(s.GetMembers()),
	)

	return m
}

// StorageGroupFromGRPCMessage converts grpc structure into unified proto structure.
func StorageGroupFromGRPCMessage(m *sg.StorageGroup) *StorageGroup {
	if m == nil {
		return nil
	}

	s := new(StorageGroup)

	s.SetValidationDataSize(m.GetValidationDataSize())
	s.SetValidationHash(
		refs.ChecksumFromGRPCMessage(m.GetValidationHash()),
	)
	s.SetExpirationEpoch(m.GetExpirationEpoch())

	s.SetMembers(
		refs.ObjectIDListFromGRPCMessage(m.GetMembers()),
	)

	return s
}
