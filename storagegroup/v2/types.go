package v2

import (
	refs "github.com/nspcc-dev/neofs-api-go/refs/v2"
)

// SetValidationDataSize sets the total size of the payloads of the storage group.
func (m *StorageGroup) SetValidationDataSize(v uint64) {
	if m != nil {
		m.ValidationDataSize = v
	}
}

// SetValidationHash sets total homomorphic hash of the storage group payloads.
func (m *StorageGroup) SetValidationHash(v []byte) {
	if m != nil {
		m.ValidationHash = v
	}
}

// SetExpirationEpoch sets number of the last epoch of the storage group lifetime.
func (m *StorageGroup) SetExpirationEpoch(v uint64) {
	if m != nil {
		m.ExpirationEpoch = v
	}
}

// SetMembers sets list of the identifiers of the storage group members.
func (m *StorageGroup) SetMembers(v []*refs.ObjectID) {
	if m != nil {
		m.Members = v
	}
}
