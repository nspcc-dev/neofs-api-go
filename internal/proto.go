package internal

import "github.com/gogo/protobuf/proto"

// Custom contains methods to satisfy proto.Message
// including custom methods to satisfy protobuf for
// non-proto defined types.
type Custom interface {
	Size() int
	Empty() bool
	Bytes() []byte
	Marshal() ([]byte, error)
	MarshalTo(data []byte) (int, error)
	Unmarshal(data []byte) error
	proto.Message

	// Should contains for proto.Clone
	proto.Merger
}
