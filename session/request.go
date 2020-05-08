package session

import (
	"encoding/binary"
	"io"

	"github.com/nspcc-dev/neofs-api-go/refs"
)

const signedRequestDataSize = 0 +
	refs.OwnerIDSize +
	8 +
	8

var requestEndianness = binary.BigEndian

// GetOwnerID is an OwnerID field getter.
func (m CreateRequest) GetOwnerID() OwnerID {
	return m.OwnerID
}

// SetOwnerID is an OwnerID field setter.
func (m *CreateRequest) SetOwnerID(id OwnerID) {
	m.OwnerID = id
}

func (m CreateRequest) SignedData() ([]byte, error) {
	data := make([]byte, m.SignedDataSize())

	_, err := m.ReadSignedData(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m CreateRequest) SignedDataSize() int {
	return signedRequestDataSize
}

func (m CreateRequest) ReadSignedData(p []byte) (int, error) {
	sz := m.SignedDataSize()
	if len(p) < sz {
		return 0, io.ErrUnexpectedEOF
	}

	var off int

	off += copy(p[off:], m.GetOwnerID().Bytes())

	requestEndianness.PutUint64(p[off:], m.CreationEpoch())
	off += 8

	requestEndianness.PutUint64(p[off:], m.ExpirationEpoch())

	return sz, nil
}
