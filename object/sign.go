package object

import (
	"io"
)

// SignedData returns payload bytes of the request.
//
// If payload is nil, ErrHeaderNotFound returns.
func (m PutRequest) SignedData() ([]byte, error) {
	sz := m.SignedDataSize()
	if sz < 0 {
		return nil, ErrHeaderNotFound
	}

	data := make([]byte, sz)

	return data, m.ReadSignedData(data)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m PutRequest) ReadSignedData(p []byte) error {
	r := m.GetR()
	if r == nil {
		return ErrHeaderNotFound
	}

	_, err := r.MarshalTo(p)

	return err
}

// SignedDataSize returns the size of payload of the Put request.
//
// If payload is nil, -1 returns.
func (m PutRequest) SignedDataSize() int {
	r := m.GetR()
	if r == nil {
		return -1
	}

	return r.Size()
}

// SignedData returns payload bytes of the request.
func (m GetRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	return data, m.ReadSignedData(data)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRequest) ReadSignedData(p []byte) error {
	addr := m.GetAddress()

	if len(p) < m.SignedDataSize() {
		return io.ErrUnexpectedEOF
	}

	off := copy(p, addr.CID.Bytes())

	copy(p[off:], addr.ObjectID.Bytes())

	return nil
}

// SignedDataSize returns payload size of the request.
func (m GetRequest) SignedDataSize() int {
	return addressSize(m.GetAddress())
}

// SignedData returns payload bytes of the request.
func (m HeadRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	return data, m.ReadSignedData(data)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m HeadRequest) ReadSignedData(p []byte) error {
	if len(p) < m.SignedDataSize() {
		return io.ErrUnexpectedEOF
	}

	if m.GetFullHeaders() {
		p[0] = 1
	}

	off := 1 + copy(p[1:], m.Address.CID.Bytes())

	copy(p[off:], m.Address.ObjectID.Bytes())

	return nil
}

// SignedDataSize returns payload size of the request.
func (m HeadRequest) SignedDataSize() int {
	return addressSize(m.Address) + 1
}

// SignedData returns payload bytes of the request.
func (m DeleteRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	return data, m.ReadSignedData(data)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m DeleteRequest) ReadSignedData(p []byte) error {
	if len(p) < m.SignedDataSize() {
		return io.ErrUnexpectedEOF
	}

	off := copy(p, m.OwnerID.Bytes())

	copy(p[off:], addressBytes(m.Address))

	return nil
}

// SignedDataSize returns payload size of the request.
func (m DeleteRequest) SignedDataSize() int {
	return m.OwnerID.Size() + addressSize(m.Address)
}

func addressSize(addr Address) int {
	return addr.CID.Size() + addr.ObjectID.Size()
}

func addressBytes(addr Address) []byte {
	return append(addr.CID.Bytes(), addr.ObjectID.Bytes()...)
}
