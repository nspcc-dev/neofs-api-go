package service

import "encoding/binary"

const (
	_ NodeRole = iota
	// InnerRingNode that work like IR node.
	InnerRingNode
	// StorageNode that work like a storage node.
	StorageNode
)

// String is method, that represent NodeRole as string.
func (nt NodeRole) String() string {
	switch nt {
	case InnerRingNode:
		return "InnerRingNode"
	case StorageNode:
		return "StorageNode"
	default:
		return "Unknown"
	}
}

func (nt NodeRole) Size() int {
	return 4
}

func (nt NodeRole) Bytes() []byte {
	data := make([]byte, nt.Size())

	binary.BigEndian.PutUint32(data, uint32(nt))

	return data
}
