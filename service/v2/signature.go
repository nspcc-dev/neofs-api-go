package v2

func (m *RequestMetaHeader) ReadSignedData(buf []byte) ([]byte, error) {
	return m.StableMarshal(buf)
}

func (m *RequestMetaHeader) SignedDataSize() int {
	return m.StableSize()
}

func (m *RequestVerificationHeader) ReadSignedData(buf []byte) ([]byte, error) {
	return m.StableMarshal(buf)
}

func (m *RequestVerificationHeader) SignedDataSize() int {
	return m.StableSize()
}
