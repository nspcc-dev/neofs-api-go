package object
// todo: all extensions must be transferred to the separate util library

import "github.com/nspcc-dev/neofs-proto/storagegroup"

// IsLinking checks if object has children links to another objects.
// We have to check payload size because zero-object must have zero
// payload and non-zero payload length field in system header.
func (m Object) IsLinking() bool {
	for i := range m.Headers {
		switch v := m.Headers[i].Value.(type) {
		case *Header_Link:
			if v.Link.GetType() == Link_Child {
				return m.SystemHeader.PayloadLength > 0 && len(m.Payload) == 0
			}
		}
	}
	return false
}

// VerificationHeader returns verification header if it is presented in extended headers.
func (m Object) VerificationHeader() (*VerificationHeader, error) {
	_, vh := m.LastHeader(HeaderType(VerifyHdr))
	if vh == nil {
		return nil, ErrHeaderNotFound
	}
	return vh.Value.(*Header_Verify).Verify, nil
}

// SetVerificationHeader sets verification header in the object.
// It will replace existing verification header or add a new one.
func (m *Object) SetVerificationHeader(header *VerificationHeader) {
	m.SetHeader(&Header{Value: &Header_Verify{Verify: header}})
}

// Links returns slice of ids of specified link type
func (m *Object) Links(t Link_Type) []ID {
	var res []ID
	for i := range m.Headers {
		switch v := m.Headers[i].Value.(type) {
		case *Header_Link:
			if v.Link.GetType() == t {
				res = append(res, v.Link.ID)
			}
		}
	}
	return res
}

// Tombstone returns tombstone header if it is presented in extended headers.
func (m Object) Tombstone() *Tombstone {
	_, h := m.LastHeader(HeaderType(TombstoneHdr))
	if h != nil {
		return h.Value.(*Header_Tombstone).Tombstone
	}
	return nil
}

// IsTombstone checks if object has tombstone header.
func (m Object) IsTombstone() bool {
	n, _ := m.LastHeader(HeaderType(TombstoneHdr))
	return n != -1
}

// StorageGroup returns storage group structure if it is presented in extended headers.
func (m Object) StorageGroup() (*storagegroup.StorageGroup, error) {
	_, sgHdr := m.LastHeader(HeaderType(StorageGroupHdr))
	if sgHdr == nil {
		return nil, ErrHeaderNotFound
	}
	return sgHdr.Value.(*Header_StorageGroup).StorageGroup, nil
}

// SetStorageGroup sets storage group header in the object.
// It will replace existing storage group header or add a new one.
func (m *Object) SetStorageGroup(group *storagegroup.StorageGroup) {
	m.SetHeader(&Header{Value: &Header_StorageGroup{StorageGroup: group}})
}
