package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/v2/netmap"
)

func newFilter(name string, k, v string, op netmap.Operation, fs ...*netmap.Filter) *netmap.Filter {
	f := new(netmap.Filter)
	f.SetName(name)
	f.SetKey(k)
	f.SetOp(op)
	f.SetValue(v)
	f.SetFilters(fs)
	return f
}

func newSelector(name string, attr string, c netmap.Clause, count uint32, filter string) *netmap.Selector {
	s := new(netmap.Selector)
	s.SetName(name)
	s.SetAttribute(attr)
	s.SetCount(count)
	s.SetClause(c)
	s.SetFilter(filter)
	return s
}

func newPlacementPolicy(bf uint32, rs []*netmap.Replica, ss []*netmap.Selector, fs []*netmap.Filter) *netmap.PlacementPolicy {
	p := new(netmap.PlacementPolicy)
	p.SetContainerBackupFactor(bf)
	p.SetReplicas(rs)
	p.SetSelectors(ss)
	p.SetFilters(fs)
	return p
}

func newReplica(c uint32, s string) *netmap.Replica {
	r := new(netmap.Replica)
	r.SetCount(c)
	r.SetSelector(s)
	return r
}

func nodeInfoFromAttributes(props ...string) netmap.NodeInfo {
	attrs := make([]*netmap.Attribute, len(props)/2)
	for i := range attrs {
		attrs[i] = new(netmap.Attribute)
		attrs[i].SetKey(props[i*2])
		attrs[i].SetValue(props[i*2+1])
	}
	var n netmap.NodeInfo
	n.SetAttributes(attrs)
	return n
}

func getTestNode(props ...string) *Node {
	m := make(map[string]string, len(props)/2)
	for i := 0; i < len(props); i += 2 {
		m[props[i]] = props[i+1]
	}
	return &Node{AttrMap: m}
}
