package object

import (
	"io"
)

// SignedData returns marshaled payload of the Put request.
//
// If payload is nil, ErrHeaderNotFound returns.
func (m PutRequest) SignedData() ([]byte, error) {
	r := m.GetR()
	if r == nil {
		return nil, ErrHeaderNotFound
	}

	data := make([]byte, r.Size())

	if _, err := r.MarshalTo(data); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadSignedData copies marshaled payload of the Put request to passed buffer.
//
// If payload is nil, ErrHeaderNotFound returns.
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

// SignedData returns marshaled Address field.
//
// Resulting error is always nil.
func (m GetRequest) SignedData() ([]byte, error) {
	addr := m.GetAddress()

	return addressBytes(addr), nil
}

// ReadSignedData copies marshaled Address field to passed buffer.
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

// SignedDataSize returns the size of object address.
func (m GetRequest) SignedDataSize() int {
	return addressSize(m.GetAddress())
}

// SignedData returns marshaled Address field.
//
// Resulting error is always nil.
func (m HeadRequest) SignedData() ([]byte, error) {
	sz := addressSize(m.Address)

	data := make([]byte, sz+1)

	if m.GetFullHeaders() {
		data[0] = 1
	}

	copy(data[1:], addressBytes(m.Address))

	return data, nil
}

// ReadSignedData copies marshaled Address field to passed buffer.
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

// SignedDataSize returns the size of object address.
func (m HeadRequest) SignedDataSize() int {
	return addressSize(m.Address) + 1
}

func addressSize(addr Address) int {
	return addr.CID.Size() + addr.ObjectID.Size()
}

func addressBytes(addr Address) []byte {
	return append(addr.CID.Bytes(), addr.ObjectID.Bytes()...)
}
