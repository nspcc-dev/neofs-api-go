package service

// SetVersion is a Version field setter.
func (m *ResponseMetaHeader) SetVersion(v uint32) {
	m.Version = v
}

// SetVersion is a Version field setter.
func (m *RequestMetaHeader) SetVersion(v uint32) {
	m.Version = v
}
