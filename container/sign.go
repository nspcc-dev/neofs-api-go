package container

import (
	"encoding/binary"
	"io"
)

var requestEndianness = binary.BigEndian

// SignedData returns payload bytes of the request.
func (m PutRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m PutRequest) SignedDataSize() (sz int) {
	sz += m.GetMessageID().Size()

	sz += 8

	sz += m.GetOwnerID().Size()

	rules := m.GetRules()
	sz += rules.Size()

	sz += 4

	return
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the Request size is insufficient, io.ErrUnexpectedEOF returns.
func (m PutRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetMessageID().Bytes())

	requestEndianness.PutUint64(p[off:], m.GetCapacity())
	off += 8

	off += copy(p[off:], m.GetOwnerID().Bytes())

	rules := m.GetRules()
	// FIXME: implement and use stable functions
	n, err := rules.MarshalTo(p[off:])
	off += n
	if err != nil {
		return off, err
	}

	requestEndianness.PutUint32(p[off:], m.GetBasicACL())
	off += 4

	return off, nil
}

// SignedData returns payload bytes of the request.
func (m DeleteRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m DeleteRequest) SignedDataSize() (sz int) {
	return m.GetCID().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the Request size is insufficient, io.ErrUnexpectedEOF returns.
func (m DeleteRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetCID().Bytes())

	return off, nil
}

// SignedData returns payload bytes of the request.
func (m GetRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m GetRequest) SignedDataSize() (sz int) {
	return m.GetCID().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the Request size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetCID().Bytes())

	return off, nil
}
