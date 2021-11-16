package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	storagegroup "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

const (
	sizeField       = 1
	hashField       = 2
	expirationField = 3
	objectIDsField  = 4
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
		err       error
	)

	n, err = proto.UInt64Marshal(sizeField, buf[offset:], s.size)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(hashField, buf[offset:], s.hash)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(expirationField, buf[offset:], s.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = refs.ObjectIDNestedListMarshal(objectIDsField, buf[offset:], s.members)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize of storage group structure marshalled by StableMarshal function.
func (s *StorageGroup) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.UInt64Size(sizeField, s.size)
	size += proto.NestedStructureSize(hashField, s.hash)
	size += proto.UInt64Size(expirationField, s.exp)
	size += refs.ObjectIDNestedListSize(objectIDsField, s.members)

	return size
}

func (s *StorageGroup) Unmarshal(data []byte) error {
	return message.Unmarshal(s, data, new(storagegroup.StorageGroup))
}
