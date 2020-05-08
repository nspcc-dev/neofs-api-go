package session

// GetOwnerID is an OwnerID field getter.
func (m CreateRequest) GetOwnerID() OwnerID {
	return m.OwnerID
}

// SetOwnerID is an OwnerID field setter.
func (m *CreateRequest) SetOwnerID(id OwnerID) {
	m.OwnerID = id
}
