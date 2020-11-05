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

// NodeState is an enumeration of various states of the NeoFS node.
type NodeState uint32

// NodeAttribute represents v2 compatible attribute of the NeoFS Storage Node.
type NodeAttribute netmap.Attribute

// NodeInfo represents v2 compatible descriptor of the NeoFS node.
type NodeInfo netmap.NodeInfo

const (
	_ NodeState = iota

	// NodeStateOffline is network unavailable state.
	NodeStateOffline

	// NodeStateOnline is an active state in the network.
	NodeStateOnline
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

// PublicKey returns public key of the node in bytes.
func (n Node) PublicKey() []byte {
	return n.InfoV2.GetPublicKey()
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

// NodeStateFromV2 converts v2 NodeState to NodeState.
func NodeStateFromV2(s netmap.NodeState) NodeState {
	switch s {
	default:
		return 0
	case netmap.Online:
		return NodeStateOnline
	case netmap.Offline:
		return NodeStateOffline
	}
}

// ToV2 converts NodeState to v2 NodeState.
func (s NodeState) ToV2() netmap.NodeState {
	switch s {
	default:
		return netmap.UnspecifiedState
	case NodeStateOffline:
		return netmap.Offline
	case NodeStateOnline:
		return netmap.Online
	}
}

func (s NodeState) String() string {
	switch s {
	default:
		return "UNSPECIFIED"
	case NodeStateOffline:
		return "OFFLINE"
	case NodeStateOnline:
		return "ONLINE"
	}
}

// NewNodeAttribute creates and returns new NodeAttribute instance.
func NewNodeAttribute() *NodeAttribute {
	return NewNodeAttributeFromV2(new(netmap.Attribute))
}

// NodeAttributeFromV2 converts v2 node Attribute to NodeAttribute.
func NewNodeAttributeFromV2(a *netmap.Attribute) *NodeAttribute {
	return (*NodeAttribute)(a)
}

// ToV2 converts NodeAttribute to v2 node Attribute.
func (a *NodeAttribute) ToV2() *netmap.Attribute {
	return (*netmap.Attribute)(a)
}

// Key returns key to the node attribute.
func (a *NodeAttribute) Key() string {
	return (*netmap.Attribute)(a).
		GetKey()
}

// SetKey sets key to the node attribute.
func (a *NodeAttribute) SetKey(key string) {
	(*netmap.Attribute)(a).
		SetKey(key)
}

// Value returns value of the node attribute.
func (a *NodeAttribute) Value() string {
	return (*netmap.Attribute)(a).
		GetValue()
}

// SetValue sets value of the node attribute.
func (a *NodeAttribute) SetValue(val string) {
	(*netmap.Attribute)(a).
		SetValue(val)
}

// ParentKeys returns list of parent keys.
func (a *NodeAttribute) ParentKeys() []string {
	return (*netmap.Attribute)(a).
		GetParents()
}

// SetParentKeys sets list of parent keys.
func (a *NodeAttribute) SetParentKeys(keys ...string) {
	(*netmap.Attribute)(a).
		SetParents(keys)
}

// NewNodeInfo creates and returns new NodeInfo instance.
func NewNodeInfo() *NodeInfo {
	return NewNodeInfoFromV2(new(netmap.NodeInfo))
}

// NewNodeInfoFromV2 converts v2 NodeInfo to NodeInfo.
func NewNodeInfoFromV2(i *netmap.NodeInfo) *NodeInfo {
	return (*NodeInfo)(i)
}

// NodeInfoToJSON encodes NodeInfo to JSON format.
func NodeInfoToJSON(i *NodeInfo) ([]byte, error) {
	return netmap.NodeInfoToJSON(i.ToV2())
}

// NodeInfoFromJSON decodes NodeInfo from JSON-encoded data.
func NodeInfoFromJSON(data []byte) (*NodeInfo, error) {
	i, err := netmap.NodeInfoFromJSON(data)
	if err != nil {
		return nil, err
	}

	return NewNodeInfoFromV2(i), nil
}

// ToV2 converts NodeInfo to v2 NodeInfo.
func (i *NodeInfo) ToV2() *netmap.NodeInfo {
	return (*netmap.NodeInfo)(i)
}

// PublicKey returns public key of the node in a binary format.
func (i *NodeInfo) PublicKey() []byte {
	return (*netmap.NodeInfo)(i).
		GetPublicKey()
}

// SetPublicKey sets public key of the node in a binary format.
func (i *NodeInfo) SetPublicKey(key []byte) {
	(*netmap.NodeInfo)(i).
		SetPublicKey(key)
}

// Address returns network endpoint address of the node.
func (i *NodeInfo) Address() string {
	return (*netmap.NodeInfo)(i).
		GetAddress()
}

// SetAddress sets network endpoint address of the node.
func (i *NodeInfo) SetAddress(addr string) {
	(*netmap.NodeInfo)(i).
		SetAddress(addr)
}

// Attributes returns list of the node attributes.
func (i *NodeInfo) Attributes() []*NodeAttribute {
	as := (*netmap.NodeInfo)(i).
		GetAttributes()

	res := make([]*NodeAttribute, 0, len(as))

	for i := range as {
		res = append(res, NewNodeAttributeFromV2(as[i]))
	}

	return res
}

// SetAttributes sets list of the node attributes.
func (i *NodeInfo) SetAttributes(as ...*NodeAttribute) {
	asV2 := make([]*netmap.Attribute, 0, len(as))

	for i := range as {
		asV2 = append(asV2, as[i].ToV2())
	}

	(*netmap.NodeInfo)(i).
		SetAttributes(asV2)
}

// State returns node state.
func (i *NodeInfo) State() NodeState {
	return NodeStateFromV2(
		(*netmap.NodeInfo)(i).
			GetState(),
	)
}

// SetState sets node state.
func (i *NodeInfo) SetState(s NodeState) {
	(*netmap.NodeInfo)(i).
		SetState(s.ToV2())
}
