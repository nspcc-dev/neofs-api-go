package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

type StorageGroup struct {
	size uint64

	hash []byte

	exp uint64

	members []*refs.ObjectID
}

func (s *StorageGroup) GetValidationDataSize() uint64 {
	if s != nil {
		return s.size
	}

	return 0
}

func (s *StorageGroup) SetValidationDataSize(v uint64) {
	if s != nil {
		s.size = v
	}
}

func (s *StorageGroup) GetValidationHash() []byte {
	if s != nil {
		return s.hash
	}

	return nil
}

func (s *StorageGroup) SetValidationHash(v []byte) {
	if s != nil {
		s.hash = v
	}
}

func (s *StorageGroup) GetExpirationEpoch() uint64 {
	if s != nil {
		return s.exp
	}

	return 0
}

func (s *StorageGroup) SetExpirationEpoch(v uint64) {
	if s != nil {
		s.exp = v
	}
}

func (s *StorageGroup) GetMembers() []*refs.ObjectID {
	if s != nil {
		return s.members
	}

	return nil
}

func (s *StorageGroup) SetMembers(v []*refs.ObjectID) {
	if s != nil {
		s.members = v
	}
}
