package netmap

func newFilter(name string, k, v string, op Operation, fs ...*Filter) *Filter {
	f := NewFilter()
	f.SetName(name)
	f.SetKey(k)
	f.SetOperation(op)
	f.SetValue(v)
	f.SetInnerFilters(fs...)
	return f
}

func newSelector(name string, attr string, c Clause, count uint32, filter string) *Selector {
	s := NewSelector()
	s.SetName(name)
	s.SetAttribute(attr)
	s.SetCount(count)
	s.SetClause(c)
	s.SetFilter(filter)
	return s
}

func newPlacementPolicy(bf uint32, rs []*Replica, ss []*Selector, fs []*Filter) *PlacementPolicy {
	p := NewPlacementPolicy()
	p.SetContainerBackupFactor(bf)
	p.SetReplicas(rs...)
	p.SetSelectors(ss...)
	p.SetFilters(fs...)
	return p
}

func newReplica(c uint32, s string) *Replica {
	r := NewReplica()
	r.SetCount(c)
	r.SetSelector(s)
	return r
}

func nodeInfoFromAttributes(props ...string) NodeInfo {
	attrs := make([]*NodeAttribute, len(props)/2)
	for i := range attrs {
		attrs[i] = NewNodeAttribute()
		attrs[i].SetKey(props[i*2])
		attrs[i].SetValue(props[i*2+1])
	}
	n := NewNodeInfo()
	n.SetAttributes(attrs...)
	return *n
}

func getTestNode(props ...string) *Node {
	m := make(map[string]string, len(props)/2)
	for i := 0; i < len(props); i += 2 {
		m[props[i]] = props[i+1]
	}
	return &Node{AttrMap: m}
}
