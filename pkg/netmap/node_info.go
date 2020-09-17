package netmap

import (
	"strconv"

	"github.com/nspcc-dev/hrw"
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

type (
	// Node is a wrapper over NodeInfo.
	Node struct {
		ID       uint64
		Index    int
		Capacity uint64
		Price    uint64
		AttrMap  map[string]string

		InfoV2 *netmap.NodeInfo
	}

	// Nodes represents slice of graph leafs.
	Nodes []*Node
)

// Enumeration of well-known attributes.
const (
	CapacityAttr = "Capacity"
	PriceAttr    = "Price"
)

var _ hrw.Hasher = (*Node)(nil)

// Hash is a function from hrw.Hasher interface. It is implemented
// to support weighted hrw therefore sort function sorts nodes
// based on their `N` value.
func (n Node) Hash() uint64 {
	return n.ID
}

// NetworkAddress returns network address
// of the node in a string format.
func (n Node) NetworkAddress() string {
	return n.InfoV2.GetAddress()
}

// NodesFromV2 converts slice of v2 netmap.NodeInfo to a generic node slice.
func NodesFromV2(infos []netmap.NodeInfo) Nodes {
	nodes := make(Nodes, len(infos))
	for i := range infos {
		nodes[i] = newNodeV2(i, &infos[i])
	}
	return nodes
}

func newNodeV2(index int, ni *netmap.NodeInfo) *Node {
	n := &Node{
		ID:      hrw.Hash(ni.GetPublicKey()),
		Index:   index,
		AttrMap: make(map[string]string, len(ni.GetAttributes())),
		InfoV2:  ni,
	}
	for _, attr := range ni.GetAttributes() {
		switch attr.GetKey() {
		case CapacityAttr:
			n.Capacity, _ = strconv.ParseUint(attr.GetValue(), 10, 64)
		case PriceAttr:
			n.Price, _ = strconv.ParseUint(attr.GetValue(), 10, 64)
		}
		n.AttrMap[attr.GetKey()] = attr.GetValue()
	}
	return n
}

// Weights returns slice of nodes weights W.
func (n Nodes) Weights(wf weightFunc) []float64 {
	w := make([]float64, 0, len(n))
	for i := range n {
		w = append(w, wf(n[i]))
	}
	return w
}

// Attribute returns value of attribute k.
func (n *Node) Attribute(k string) string {
	return n.AttrMap[k]
}

// GetBucketWeight computes weight for a Bucket.
func GetBucketWeight(ns Nodes, a aggregator, wf weightFunc) float64 {
	for i := range ns {
		a.Add(wf(ns[i]))
	}
	return a.Compute()
}
