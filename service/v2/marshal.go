package v2

func (m *RequestMetaHeader) StableMarshal(buf []byte) ([]byte, error) {
	// fixme: implement stable
	return m.Marshal()
}

func (m *RequestMetaHeader) StableSize() int {
	// fixme: implement stable
	return m.Size()
}

func (m *RequestVerificationHeader) StableMarshal(buf []byte) ([]byte, error) {
	// fixme: implement stable
	return m.Marshal()
}

func (m *RequestVerificationHeader) StableSize() int {
	// fixme: implement stable
	return m.Size()
}
