package accounting

import (
	"encoding/binary"
	"io"
)

// SignedData returns payload bytes of the request.
func (m BalanceRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	if _, err := m.ReadSignedData(data); err != nil {
		return nil, err
	}

	return data, nil
}

// SignedDataSize returns payload size of the request.
func (m BalanceRequest) SignedDataSize() int {
	return m.GetOwnerID().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m BalanceRequest) ReadSignedData(p []byte) (int, error) {
	sz := m.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	copy(p, m.GetOwnerID().Bytes())

	return sz, nil
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
func (m GetRequest) SignedDataSize() int {
	return m.GetID().Size() + m.GetOwnerID().Size()
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRequest) ReadSignedData(p []byte) (int, error) {
	sz := m.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetID().Bytes())

	copy(p[off:], m.GetOwnerID().Bytes())

	return sz, nil
}

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
	sz += m.GetOwnerID().Size()

	sz += m.GetMessageID().Size()

	sz += 8

	if amount := m.GetAmount(); amount != nil {
		sz += amount.Size()
	}

	return
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m PutRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetOwnerID().Bytes())

	off += copy(p[off:], m.GetMessageID().Bytes())

	binary.BigEndian.PutUint64(p[off:], m.GetHeight())
	off += 8

	if amount := m.GetAmount(); amount != nil {
		n, err := amount.MarshalTo(p[off:])
		off += n
		if err != nil {
			return off + n, err
		}
	}

	return off, nil
}
