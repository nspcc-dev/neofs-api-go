package storagegroup

import (
	"encoding/binary"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	// SizeField order number from storage group proto definition.
	SizeField = 1
	// HashField order number from storage group proto definition.
	HashField = 2
	// ExpirationField order number from storage group proto definition.
	ExpirationField = 3
	// ObjectIDsField order number from storage group proto definition.
	ObjectIDsField = 4
)

// StableMarshal marshals unified storage group structure in a protobuf
// compatible way without field order shuffle.
func (s *StorageGroup) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	n, err = proto.UInt64Marshal(SizeField, buf, s.size)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(HashField, buf[offset:], s.hash)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(ExpirationField, buf[offset:], s.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	prefix, _ = proto.NestedStructurePrefix(ObjectIDsField)

	for i := range s.members {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = s.members[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = s.members[i].StableMarshal(buf[offset:])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of storage group structure marshalled by StableMarshal function.
func (s *StorageGroup) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.UInt64Size(SizeField, s.size)
	size += proto.BytesSize(HashField, s.hash)
	size += proto.UInt64Size(ExpirationField, s.exp)

	_, ln := proto.NestedStructurePrefix(ObjectIDsField)

	for i := range s.members {
		n := s.members[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}
