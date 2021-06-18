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

		*NodeInfo
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
	// AttrPrice is a key to the node attribute that indicates the
	// price in GAS tokens for storing one GB of data during one Epoch.
	AttrPrice = "Price"

	// AttrCapacity is a key to the node attribute that indicates the
	// total available disk space in Gigabytes.
	AttrCapacity = "Capacity"

	// AttrSubnet is a key to the node attribute that indicates the
	// string ID of node's storage subnet.
	AttrSubnet = "Subnet"

	// AttrUNLOCODE is a key to the node attribute that indicates the
	// node's geographic location in UN/LOCODE format.
	AttrUNLOCODE = "UN-LOCODE"

	// AttrCountryCode is a key to the node attribute that indicates the
	// Country code in ISO 3166-1_alpha-2 format.
	AttrCountryCode = "CountryCode"

	// AttrCountry is a key to the node attribute that indicates the
	// country short name in English, as defined in ISO-3166.
	AttrCountry = "Country"

	// AttrLocation is a key to the node attribute that indicates the
	// place name of the node location.
	AttrLocation = "Location"

	// AttrSubDivCode is a key to the node attribute that indicates the
	// country's administrative subdivision where node is located
	// in ISO 3166-2 format.
	AttrSubDivCode = "SubDivCode"

	// AttrSubDiv is a key to the node attribute that indicates the
	// country's administrative subdivision name, as defined in
	// ISO 3166-2.
	AttrSubDiv = "SubDiv"

	// AttrContinent is a key to the node attribute that indicates the
	// node's continent name according to the Seven-Continent model.
	AttrContinent = "Continent"
)

var _ hrw.Hasher = (*Node)(nil)

// Hash is a function from hrw.Hasher interface. It is implemented
// to support weighted hrw therefore sort function sorts nodes
// based on their `N` value.
func (n Node) Hash() uint64 {
	return n.ID
}

// NodesFromInfo converts slice of NodeInfo to a generic node slice.
func NodesFromInfo(infos []NodeInfo) Nodes {
	nodes := make(Nodes, len(infos))
	for i := range infos {
		nodes[i] = newNodeV2(i, &infos[i])
	}

	return nodes
}

func newNodeV2(index int, ni *NodeInfo) *Node {
	n := &Node{
		ID:       hrw.Hash(ni.PublicKey()),
		Index:    index,
		AttrMap:  make(map[string]string, len(ni.Attributes())),
		NodeInfo: ni,
	}

	for _, attr := range ni.Attributes() {
		switch attr.Key() {
		case AttrCapacity:
			n.Capacity, _ = strconv.ParseUint(attr.Value(), 10, 64)
		case AttrPrice:
			n.Price, _ = strconv.ParseUint(attr.Value(), 10, 64)
		}

		n.AttrMap[attr.Key()] = attr.Value()
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

// String returns string representation of NodeState.
//
// String mapping:
//  * NodeStateOnline: ONLINE;
//  * NodeStateOffline: OFFLINE;
//  * default: UNSPECIFIED.
func (s NodeState) String() string {
	return s.ToV2().String()
}

// FromString parses NodeState from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (s *NodeState) FromString(str string) bool {
	var g netmap.NodeState

	ok := g.FromString(str)

	if ok {
		*s = NodeStateFromV2(g)
	}

	return ok
}

// NewNodeAttribute creates and returns new NodeAttribute instance.
//
// Defaults:
//  - key: "";
//  - value: "";
//  - parents: nil.
func NewNodeAttribute() *NodeAttribute {
	return NewNodeAttributeFromV2(new(netmap.Attribute))
}

// NodeAttributeFromV2 converts v2 node Attribute to NodeAttribute.
//
// Nil netmap.Attribute converts to nil.
func NewNodeAttributeFromV2(a *netmap.Attribute) *NodeAttribute {
	return (*NodeAttribute)(a)
}

// ToV2 converts NodeAttribute to v2 node Attribute.
//
// Nil NodeAttribute converts to nil.
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

// Marshal marshals NodeAttribute into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (a *NodeAttribute) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.Attribute)(a).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of NodeAttribute.
func (a *NodeAttribute) Unmarshal(data []byte) error {
	return (*netmap.Attribute)(a).
		Unmarshal(data)
}

// MarshalJSON encodes NodeAttribute to protobuf JSON format.
func (a *NodeAttribute) MarshalJSON() ([]byte, error) {
	return (*netmap.Attribute)(a).
		MarshalJSON()
}

// UnmarshalJSON decodes NodeAttribute from protobuf JSON format.
func (a *NodeAttribute) UnmarshalJSON(data []byte) error {
	return (*netmap.Attribute)(a).
		UnmarshalJSON(data)
}

// NewNodeInfo creates and returns new NodeInfo instance.
//
// Defaults:
//  - publicKey: nil;
//  - address: "";
//  - attributes nil;
//  - state: 0.
func NewNodeInfo() *NodeInfo {
	return NewNodeInfoFromV2(new(netmap.NodeInfo))
}

// NewNodeInfoFromV2 converts v2 NodeInfo to NodeInfo.
//
// Nil netmap.NodeInfo converts to nil.
func NewNodeInfoFromV2(i *netmap.NodeInfo) *NodeInfo {
	return (*NodeInfo)(i)
}

// ToV2 converts NodeInfo to v2 NodeInfo.
//
// Nil NodeInfo converts to nil.
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
//
// Deprecated: use IterateAddresses method.
func (i *NodeInfo) Address() (addr string) {
	i.IterateAddresses(func(s string) bool {
		addr = s
		return true
	})

	return
}

// SetAddress sets network endpoint address of the node.
//
// Deprecated: use SetAddresses method.
func (i *NodeInfo) SetAddress(addr string) {
	i.SetAddresses(addr)
}

// NumberOfAddresses returns number of network addresses of the node.
func (i *NodeInfo) NumberOfAddresses() int {
	return (*netmap.NodeInfo)(i).
		NumberOfAddresses()
}

// IterateAddresses iterates over network addresses of the node.
// Breaks iteration on f's true return.
//
// Handler should not be nil.
func (i *NodeInfo) IterateAddresses(f func(string) bool) {
	(*netmap.NodeInfo)(i).
		IterateAddresses(f)
}

// IterateAllAddresses is a helper function to unconditionally
// iterate over all node addresses.
func IterateAllAddresses(i *NodeInfo, f func(string)) {
	i.IterateAddresses(func(addr string) bool {
		f(addr)
		return false
	})
}

// SetAddresses sets list of network addresses of the node.
func (i *NodeInfo) SetAddresses(v ...string) {
	(*netmap.NodeInfo)(i).
		SetAddresses(v...)
}

// Attributes returns list of the node attributes.
func (i *NodeInfo) Attributes() []*NodeAttribute {
	if i == nil {
		return nil
	}

	as := (*netmap.NodeInfo)(i).
		GetAttributes()

	if as == nil {
		return nil
	}

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

// Marshal marshals NodeInfo into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (i *NodeInfo) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return (*netmap.NodeInfo)(i).
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of NodeInfo.
func (i *NodeInfo) Unmarshal(data []byte) error {
	return (*netmap.NodeInfo)(i).
		Unmarshal(data)
}

// MarshalJSON encodes NodeInfo to protobuf JSON format.
func (i *NodeInfo) MarshalJSON() ([]byte, error) {
	return (*netmap.NodeInfo)(i).
		MarshalJSON()
}

// UnmarshalJSON decodes NodeInfo from protobuf JSON format.
func (i *NodeInfo) UnmarshalJSON(data []byte) error {
	return (*netmap.NodeInfo)(i).
		UnmarshalJSON(data)
}
