package tombstone

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
	tombstone "github.com/nspcc-dev/neofs-api-go/v2/tombstone/grpc"
	goproto "google.golang.org/protobuf/proto"
)

const (
	expFNum     = 1
	splitIDFNum = 2
	membersFNum = 3
)

// StableMarshal marshals unified tombstone message in a protobuf
// compatible way without field order shuffle.
func (s *Tombstone) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.UInt64Marshal(expFNum, buf[offset:], s.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(splitIDFNum, buf[offset:], s.splitID)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range s.members {
		n, err = proto.NestedStructureMarshal(membersFNum, buf[offset:], s.members[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize returns size of tombstone message marshalled by StableMarshal function.
func (s *Tombstone) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.UInt64Size(expFNum, s.exp)
	size += proto.BytesSize(splitIDFNum, s.splitID)

	for i := range s.members {
		size += proto.NestedStructureSize(membersFNum, s.members[i])
	}

	return size
}

// Unmarshal unmarshal tombstone message from its binary representation.
func (s *Tombstone) Unmarshal(data []byte) error {
	m := new(tombstone.Tombstone)
	if err := goproto.Unmarshal(data, m); err != nil {
		return err
	}

	*s = *TombstoneFromGRPCMessage(m)

	return nil
}
