package object

import (
	"bytes"
	"sort"
)

// Here are defined getter functions for objects that contain storage group
// information.

type (
	// IDList is a slice of object ids, that can be sorted.
	IDList []ID

	// ZoneInfo provides validation info of storage group.
	ZoneInfo struct {
		Hash
		Size uint64
	}

	// IdentificationInfo provides meta information about storage group.
	IdentificationInfo struct {
		SGID
		CID
		OwnerID
	}
)

// Len returns amount of object ids in IDList.
func (s IDList) Len() int { return len(s) }

// Less returns byte comparision between IDList[i] and IDList[j].
func (s IDList) Less(i, j int) bool { return bytes.Compare(s[i].Bytes(), s[j].Bytes()) == -1 }

// Swap swaps element with i and j index in IDList.
func (s IDList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Group returns slice of object ids that are part of a storage group.
func (m *Object) Group() []ID {
	sgLinks := m.Links(Link_StorageGroup)
	sort.Sort(IDList(sgLinks))
	return sgLinks
}

// Zones returns validation zones of storage group.
func (m *Object) Zones() []ZoneInfo {
	sgInfo, err := m.StorageGroup()
	if err != nil {
		return nil
	}
	return []ZoneInfo{
		{
			Hash: sgInfo.ValidationHash,
			Size: sgInfo.ValidationDataSize,
		},
	}
}

// IDInfo returns meta information about storage group.
func (m *Object) IDInfo() *IdentificationInfo {
	return &IdentificationInfo{
		SGID:    m.SystemHeader.ID,
		CID:     m.SystemHeader.CID,
		OwnerID: m.SystemHeader.OwnerID,
	}
}
