package object

import (
	"sort"

	"github.com/nspcc-dev/neofs-api/refs"
	"github.com/nspcc-dev/neofs-api/storagegroup"
)

// Here are defined getter functions for objects that contain storage group
// information.

var _ storagegroup.Provider = (*Object)(nil)

// Group returns slice of object ids that are part of a storage group.
func (m *Object) Group() []refs.ObjectID {
	sgLinks := m.Links(Link_StorageGroup)
	sort.Sort(storagegroup.IDList(sgLinks))
	return sgLinks
}

// Zones returns validation zones of storage group.
func (m *Object) Zones() []storagegroup.ZoneInfo {
	sgInfo, err := m.StorageGroup()
	if err != nil {
		return nil
	}
	return []storagegroup.ZoneInfo{
		{
			Hash: sgInfo.ValidationHash,
			Size: sgInfo.ValidationDataSize,
		},
	}
}

// IDInfo returns meta information about storage group.
func (m *Object) IDInfo() *storagegroup.IdentificationInfo {
	return &storagegroup.IdentificationInfo{
		SGID:    m.SystemHeader.ID,
		CID:     m.SystemHeader.CID,
		OwnerID: m.SystemHeader.OwnerID,
	}
}
