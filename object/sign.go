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

	return append(
		addr.CID.Bytes(),
		addr.ObjectID.Bytes()...,
	), nil
}

// ReadSignedData copies marshaled Address field to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRequest) ReadSignedData(p []byte) error {
	addr := m.GetAddress()

	if len(p) < addr.CID.Size()+addr.ObjectID.Size() {
		return io.ErrUnexpectedEOF
	}

	off := copy(p, addr.CID.Bytes())

	copy(p[off:], addr.ObjectID.Bytes())

	return nil
}

// SignedDataSize returns the size of object address.
func (m GetRequest) SignedDataSize() int {
	addr := m.GetAddress()

	return addr.CID.Size() + addr.ObjectID.Size()
}
