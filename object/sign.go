package object

import (
	"encoding/binary"
	"io"

	"github.com/nspcc-dev/neofs-api-go/service"
)

// SignedData returns payload bytes of the request.
//
// If payload is nil, ErrHeaderNotFound returns.
func (m PutRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m PutRequest) ReadSignedData(p []byte) (int, error) {
	r := m.GetR()
	if r == nil {
		return 0, ErrHeaderNotFound
	}

	return r.MarshalTo(p)
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
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRequest) ReadSignedData(p []byte) (int, error) {
	addr := m.GetAddress()

	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], addr.CID.Bytes())

	off += copy(p[off:], addr.ObjectID.Bytes())

	return off, nil
}

// SignedDataSize returns payload size of the request.
func (m GetRequest) SignedDataSize() int {
	return addressSize(m.GetAddress())
}

// SignedData returns payload bytes of the request.
func (m HeadRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m HeadRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	if m.GetFullHeaders() {
		p[0] = 1
	} else {
		p[0] = 0
	}

	off := 1

	off += copy(p[off:], m.Address.CID.Bytes())

	off += copy(p[off:], m.Address.ObjectID.Bytes())

	return off, nil
}

// SignedDataSize returns payload size of the request.
func (m HeadRequest) SignedDataSize() int {
	return addressSize(m.Address) + 1
}

// SignedData returns payload bytes of the request.
func (m DeleteRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m DeleteRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.OwnerID.Bytes())

	off += copy(p[off:], addressBytes(m.Address))

	return off, nil
}

// SignedDataSize returns payload size of the request.
func (m DeleteRequest) SignedDataSize() int {
	return m.OwnerID.Size() + addressSize(m.Address)
}

// SignedData returns payload bytes of the request.
func (m GetRangeRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRangeRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	n, err := (&m.Range).MarshalTo(p)
	if err != nil {
		return 0, err
	}

	n += copy(p[n:], addressBytes(m.GetAddress()))

	return n, nil
}

// SignedDataSize returns payload size of the request.
func (m GetRangeRequest) SignedDataSize() int {
	return (&m.Range).Size() + addressSize(m.GetAddress())
}

// SignedData returns payload bytes of the request.
func (m GetRangeHashRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m GetRangeHashRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], addressBytes(m.GetAddress()))

	off += copy(p[off:], rangeSetBytes(m.GetRanges()))

	off += copy(p[off:], m.GetSalt())

	return off, nil
}

// SignedDataSize returns payload size of the request.
func (m GetRangeHashRequest) SignedDataSize() int {
	var sz int

	sz += addressSize(m.GetAddress())

	sz += rangeSetSize(m.GetRanges())

	sz += len(m.GetSalt())

	return sz
}

// SignedData returns payload bytes of the request.
func (m SearchRequest) SignedData() ([]byte, error) {
	return service.SignedDataFromReader(m)
}

// ReadSignedData copies payload bytes to passed buffer.
//
// If the buffer size is insufficient, io.ErrUnexpectedEOF returns.
func (m SearchRequest) ReadSignedData(p []byte) (int, error) {
	if len(p) < m.SignedDataSize() {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.CID().Bytes())

	binary.BigEndian.PutUint32(p[off:], m.GetQueryVersion())
	off += 4

	off += copy(p[off:], m.GetQuery())

	return off, nil
}

// SignedDataSize returns payload size of the request.
func (m SearchRequest) SignedDataSize() int {
	var sz int

	sz += m.CID().Size()

	sz += 4 // uint32 Version

	sz += len(m.GetQuery())

	return sz
}

func rangeSetSize(rs []Range) int {
	return 4 + len(rs)*16 // two uint64 fields
}

func rangeSetBytes(rs []Range) []byte {
	data := make([]byte, rangeSetSize(rs))

	binary.BigEndian.PutUint32(data, uint32(len(rs)))

	off := 4

	for i := range rs {
		binary.BigEndian.PutUint64(data[off:], rs[i].Offset)
		off += 8

		binary.BigEndian.PutUint64(data[off:], rs[i].Length)
		off += 8
	}

	return data
}

func addressSize(addr Address) int {
	return addr.CID.Size() + addr.ObjectID.Size()
}

func addressBytes(addr Address) []byte {
	return append(addr.CID.Bytes(), addr.ObjectID.Bytes()...)
}
