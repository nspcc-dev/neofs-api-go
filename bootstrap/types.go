package bootstrap

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/nspcc-dev/neofs-proto/object"
)

type (
	// NodeStatus is a bitwise status field of the node.
	NodeStatus uint64
)

const (
	storageFullMask = 0x1

	optionCapacity = "/Capacity:"
	optionPrice    = "/Price:"
)

var (
	_ proto.Message = (*NodeInfo)(nil)
	_ proto.Message = (*SpreadMap)(nil)
)

// Equals checks whether two NodeInfo has same address.
func (m NodeInfo) Equals(n1 NodeInfo) bool {
	return m.Address == n1.Address && bytes.Equal(m.PubKey, n1.PubKey)
}

// Full checks if node has enough space for storing users objects.
func (n NodeStatus) Full() bool {
	return n&storageFullMask > 0
}

// SetFull changes state of node to indicate if node has enough space for storing users objects.
// If value is true - there's not enough space.
func (n *NodeStatus) SetFull(value bool) {
	switch value {
	case true:
		*n |= NodeStatus(storageFullMask)
	case false:
		*n &= NodeStatus(^uint64(storageFullMask))
	}
}

// Price returns price in 1e-8*GAS/Megabyte per month.
// User set price in GAS/Terabyte per month.
func (m NodeInfo) Price() uint64 {
	for i := range m.Options {
		if strings.HasPrefix(m.Options[i], optionPrice) {
			n, err := strconv.ParseFloat(m.Options[i][len(optionPrice):], 64)
			if err != nil {
				return 0
			}
			return uint64(n*1e8) / uint64(object.UnitsMB) // UnitsMB == megabytes in 1 terabyte
		}
	}
	return 0
}

// Capacity returns node's capacity as reported by user.
func (m NodeInfo) Capacity() uint64 {
	for i := range m.Options {
		if strings.HasPrefix(m.Options[i], optionCapacity) {
			n, err := strconv.ParseUint(m.Options[i][len(optionCapacity):], 10, 64)
			if err != nil {
				return 0
			}
			return n
		}
	}
	return 0
}

// String returns string representation of NodeInfo.
func (m NodeInfo) String() string {
	return "(NodeInfo)<" +
		"Address:" + m.Address +
		", " +
		"PublicKey:" + hex.EncodeToString(m.PubKey) +
		", " +
		"Options: [" + strings.Join(m.Options, ",") + "]>"
}

// String returns string representation of SpreadMap.
func (m SpreadMap) String() string {
	result := make([]string, 0, len(m.NetMap))
	for i := range m.NetMap {
		result = append(result, m.NetMap[i].String())
	}
	return "(SpreadMap)<" +
		"Epoch: " + strconv.FormatUint(m.Epoch, 10) +
		", " +
		"Netmap: [" + strings.Join(result, ",") + "]>"
}
