package container

import (
	"bytes"
	"encoding/binary"

	"github.com/nspcc-dev/neofs-proto/refs"
	"github.com/pkg/errors"
)

type (
	// CID type alias.
	CID = refs.CID
	// UUID type alias.
	UUID = refs.UUID
	// OwnerID type alias.
	OwnerID = refs.OwnerID
	// OwnerID type alias.
	MessageID = refs.MessageID
)

// SetTTL sets ttl to GetRequest to satisfy TTLRequest interface.
func (m *GetRequest) SetTTL(v uint32) { m.TTL = v }

// SetTTL sets ttl to PutRequest to satisfy TTLRequest interface.
func (m *PutRequest) SetTTL(v uint32) { m.TTL = v }

// SetTTL sets ttl to ListRequest to satisfy TTLRequest interface.
func (m *ListRequest) SetTTL(v uint32) { m.TTL = v }

// SetTTL sets ttl to DeleteRequest to satisfy TTLRequest interface.
func (m *DeleteRequest) SetTTL(v uint32) { m.TTL = v }

// SetSignature sets signature to PutRequest to satisfy SignedRequest interface.
func (m *PutRequest) SetSignature(v []byte) { m.Signature = v }

// SetSignature sets signature to DeleteRequest to satisfy SignedRequest interface.
func (m *DeleteRequest) SetSignature(v []byte) { m.Signature = v }

// PrepareData prepares bytes representation of PutRequest to satisfy SignedRequest interface.
func (m *PutRequest) PrepareData() ([]byte, error) {
	var (
		err      error
		buf      = new(bytes.Buffer)
		capBytes = make([]byte, 8)
	)

	binary.BigEndian.PutUint64(capBytes, m.Capacity)

	if _, err = buf.Write(m.MessageID.Bytes()); err != nil {
		return nil, errors.Wrap(err, "could not write message id")
	} else if _, err = buf.Write(capBytes); err != nil {
		return nil, errors.Wrap(err, "could not write capacity")
	} else if _, err = buf.Write(m.OwnerID.Bytes()); err != nil {
		return nil, errors.Wrap(err, "could not write pub")
	} else if data, err := m.Rules.Marshal(); err != nil {
		return nil, errors.Wrap(err, "could not marshal placement")
	} else if _, err = buf.Write(data); err != nil {
		return nil, errors.Wrap(err, "could not write placement")
	}

	return buf.Bytes(), nil
}

// PrepareData prepares bytes representation of DeleteRequest to satisfy SignedRequest interface.
func (m *DeleteRequest) PrepareData() ([]byte, error) {
	return m.CID.Bytes(), nil
}
