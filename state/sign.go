package state

import (
	"io"
)

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m NetmapRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m MetricsRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m HealthRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m DumpRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
//
// Always returns an empty slice.
func (m DumpVarsRequest) SignedData() ([]byte, error) {
	return make([]byte, 0), nil
}

// SignedData returns payload bytes of the request.
func (m ChangeStateRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m ChangeStateRequest) SignedDataSize() int {
	return m.GetState().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the Request size is insufficient, io.ErrUnexpectedEOF returns.
func (m ChangeStateRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetState().Bytes())

	return off, nil
}
