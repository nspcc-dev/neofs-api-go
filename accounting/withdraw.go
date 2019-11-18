package accounting

import (
	"encoding/binary"

	"github.com/nspcc-dev/neofs-proto/refs"
)

type (
	// MessageID type alias.
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
	var offset int
	// MessageID-len + OwnerID-len + Amount + Height
	buf := make([]byte, refs.UUIDSize+refs.OwnerIDSize+binary.MaxVarintLen64+binary.MaxVarintLen64)
	offset += copy(buf[offset:], m.MessageID.Bytes())
	offset += copy(buf[offset:], m.OwnerID.Bytes())
	offset += binary.PutVarint(buf[offset:], m.Amount.Value)
	binary.PutUvarint(buf[offset:], m.Height)
	return buf, nil
}

// PrepareData prepares bytes representation of DeleteRequest to satisfy SignedRequest interface.
func (m *DeleteRequest) PrepareData() ([]byte, error) {
	var offset int
	// ID-len + OwnerID-len + MessageID-len
	buf := make([]byte, refs.UUIDSize+refs.OwnerIDSize+refs.UUIDSize)
	offset += copy(buf[offset:], m.ID.Bytes())
	offset += copy(buf[offset:], m.OwnerID.Bytes())
	copy(buf[offset:], m.MessageID.Bytes())
	return buf, nil
}
