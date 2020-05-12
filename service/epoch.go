package service

// SetEpoch is an Epoch field setter.
func (m *ResponseMetaHeader) SetEpoch(v uint64) {
	m.Epoch = v
}

// SetEpoch is an Epoch field setter.
func (m *RequestMetaHeader) SetEpoch(v uint64) {
	m.Epoch = v
}
