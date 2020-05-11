package bootstrap

import "io"

// SignedData returns payload bytes of the request.
func (m Request) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m Request) SignedDataSize() (sz int) {
	sz += m.GetType().Size()

	sz += m.GetState().Size()

	info := m.GetInfo()
	sz += info.Size()

	return
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the Request size is insufficient, io.ErrUnexpectedEOF returns.
func (m Request) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetType().Bytes())

	off += copy(p[off:], m.GetState().Bytes())

	info := m.GetInfo()
	// FIXME: implement and use stable functions
	n, err := info.MarshalTo(p[off:])
	off += n

	return off, err
}
