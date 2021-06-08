package netmaptest

import (
	"github.com/nspcc-dev/neofs-api-go/pkg/netmap"
)

func filter(withInner bool) *netmap.Filter {
	x := netmap.NewFilter()

	x.SetName("name")
	x.SetKey("key")
	x.SetValue("value")
	x.SetOperation(netmap.OpAND)

	if withInner {
		x.SetInnerFilters(filter(false), filter(false))
	}

	return x
}

// Filter returns random netmap.Filter.
func Filter() *netmap.Filter {
	return filter(true)
}

// Replica returns random netmap.Replica.
func Replica() *netmap.Replica {
	x := netmap.NewReplica()

	x.SetCount(666)
	x.SetSelector("selector")

	return x
}

// Selector returns random netmap.Selector.
func Selector() *netmap.Selector {
	x := netmap.NewSelector()

	x.SetCount(11)
	x.SetName("name")
	x.SetFilter("filter")
	x.SetAttribute("attribute")
	x.SetClause(netmap.ClauseDistinct)

	return x
}

// PlacementPolicy returns random netmap.PlacementPolicy.
func PlacementPolicy() *netmap.PlacementPolicy {
	x := netmap.NewPlacementPolicy()

	x.SetContainerBackupFactor(9)
	x.SetFilters(Filter(), Filter())
	x.SetReplicas(Replica(), Replica())
	x.SetSelectors(Selector(), Selector())

	return x
}

// NetworkInfo returns random netmap.NetworkInfo.
func NetworkInfo() *netmap.NetworkInfo {
	x := netmap.NewNetworkInfo()

	x.SetCurrentEpoch(21)
	x.SetMagicNumber(32)

	return x
}

// NodeAttribute returns random netmap.NodeAttribute.
func NodeAttribute() *netmap.NodeAttribute {
	x := netmap.NewNodeAttribute()

	x.SetKey("key")
	x.SetValue("value")
	x.SetParentKeys("parent1", "parent2")

	return x
}

// NodeInfo returns random netmap.NodeInfo.
func NodeInfo() *netmap.NodeInfo {
	x := netmap.NewNodeInfo()

	x.SetAddress("address")
	x.SetPublicKey([]byte("public key"))
	x.SetState(netmap.NodeStateOnline)
	x.SetAttributes(NodeAttribute(), NodeAttribute())

	return x
}

// Node returns random netmap.Node.
func Node() *netmap.Node {
	return &netmap.Node{
		ID:       1,
		Index:    2,
		Capacity: 3,
		Price:    4,
		AttrMap: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		NodeInfo: NodeInfo(),
	}
}

// Nodes returns random netmap.Nodes.
func Nodes() netmap.Nodes {
	return netmap.Nodes{Node(), Node()}
}

// Netmap returns random netmap.Netmap.
func Netmap() *netmap.Netmap {
	nm, err := netmap.NewNetmap(Nodes())
	if err != nil {
		panic(err)
	}

	return nm
}
