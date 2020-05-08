package session

// GetID is an ID field getter.
func (m CreateResponse) GetID() TokenID {
	return m.ID
}

// SetID is an ID field setter.
func (m *CreateResponse) SetID(id TokenID) {
	m.ID = id
}

// SetSessionKey is a SessionKey field setter.
func (m *CreateResponse) SetSessionKey(key []byte) {
	m.SessionKey = key
}
