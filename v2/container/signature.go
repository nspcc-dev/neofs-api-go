package container

func (m *Container) ReadSignedData(buf []byte) ([]byte, error) {
	return m.StableMarshal(buf)
}

func (m *Container) SignedDataSize() int {
	return m.StableSize()
}

func (m *PutRequest_Body) ReadSignedData(buf []byte) ([]byte, error) {
	return m.StableMarshal(buf)
}

func (m *PutRequest_Body) SignedDataSize() int {
	return m.StableSize()
}
