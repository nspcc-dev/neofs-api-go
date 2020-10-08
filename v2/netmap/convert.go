package netmap

import (
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

func FilterToGRPCMessage(f *Filter) *netmap.Filter {
	if f == nil {
		return nil
	}

	m := new(netmap.Filter)

	m.SetName(f.GetName())
	m.SetKey(f.GetKey())
	m.SetValue(f.GetValue())
	m.SetOp(OperationToGRPCMessage(f.GetOp()))

	filters := make([]*netmap.Filter, 0, len(f.GetFilters()))
	for _, filter := range f.GetFilters() {
		filters = append(filters, FilterToGRPCMessage(filter))
	}
	m.SetFilters(filters)

	return m
}

func FilterFromGRPCMessage(m *netmap.Filter) *Filter {
	if m == nil {
		return nil
	}

	f := new(Filter)
	f.SetName(m.GetName())
	f.SetKey(m.GetKey())
	f.SetValue(m.GetValue())
	f.SetOp(OperationFromGRPCMessage(m.GetOp()))

	filters := make([]*Filter, 0, len(f.GetFilters()))
	for _, filter := range m.GetFilters() {
		filters = append(filters, FilterFromGRPCMessage(filter))
	}
	f.SetFilters(filters)

	return f
}

func SelectorToGRPCMessage(s *Selector) *netmap.Selector {
	if s == nil {
		return nil
	}

	m := new(netmap.Selector)

	m.SetName(s.GetName())
	m.SetCount(s.GetCount())
	m.SetClause(ClauseToGRPCMessage(s.GetClause()))
	m.SetFilter(s.GetFilter())
	m.SetAttribute(s.GetAttribute())

	return m
}

func SelectorFromGRPCMessage(m *netmap.Selector) *Selector {
	if m == nil {
		return nil
	}

	s := new(Selector)

	s.SetName(m.GetName())
	s.SetCount(m.GetCount())
	s.SetClause(ClauseFromGRPCMessage(m.GetClause()))
	s.SetFilter(m.GetFilter())
	s.SetAttribute(m.GetAttribute())

	return s
}

func ReplicaToGRPCMessage(r *Replica) *netmap.Replica {
	if r == nil {
		return nil
	}

	m := new(netmap.Replica)

	m.SetCount(r.GetCount())
	m.SetSelector(r.GetSelector())

	return m
}

func ReplicaFromGRPCMessage(m *netmap.Replica) *Replica {
	if m == nil {
		return nil
	}

	r := new(Replica)
	r.SetSelector(m.GetSelector())
	r.SetCount(m.GetCount())

	return r
}

func PlacementPolicyToGRPCMessage(p *PlacementPolicy) *netmap.PlacementPolicy {
	if p == nil {
		return nil
	}

	filters := make([]*netmap.Filter, 0, len(p.GetFilters()))
	for _, filter := range p.GetFilters() {
		filters = append(filters, FilterToGRPCMessage(filter))
	}

	selectors := make([]*netmap.Selector, 0, len(p.GetSelectors()))
	for _, selector := range p.GetSelectors() {
		selectors = append(selectors, SelectorToGRPCMessage(selector))
	}

	replicas := make([]*netmap.Replica, 0, len(p.GetReplicas()))
	for _, replica := range p.GetReplicas() {
		replicas = append(replicas, ReplicaToGRPCMessage(replica))
	}

	m := new(netmap.PlacementPolicy)

	m.SetContainerBackupFactor(p.GetContainerBackupFactor())
	m.SetFilters(filters)
	m.SetSelectors(selectors)
	m.SetReplicas(replicas)

	return m
}

func PlacementPolicyFromGRPCMessage(m *netmap.PlacementPolicy) *PlacementPolicy {
	if m == nil {
		return nil
	}

	filters := make([]*Filter, 0, len(m.GetFilters()))
	for _, filter := range m.GetFilters() {
		filters = append(filters, FilterFromGRPCMessage(filter))
	}

	selectors := make([]*Selector, 0, len(m.GetSelectors()))
	for _, selector := range m.GetSelectors() {
		selectors = append(selectors, SelectorFromGRPCMessage(selector))
	}

	replicas := make([]*Replica, 0, len(m.GetReplicas()))
	for _, replica := range m.GetReplicas() {
		replicas = append(replicas, ReplicaFromGRPCMessage(replica))
	}

	p := new(PlacementPolicy)

	p.SetContainerBackupFactor(m.GetContainerBackupFactor())
	p.SetFilters(filters)
	p.SetSelectors(selectors)
	p.SetReplicas(replicas)

	return p
}

func ClauseToGRPCMessage(n Clause) netmap.Clause {
	return netmap.Clause(n)
}

func ClauseFromGRPCMessage(n netmap.Clause) Clause {
	return Clause(n)
}

func OperationToGRPCMessage(n Operation) netmap.Operation {
	return netmap.Operation(n)
}

func OperationFromGRPCMessage(n netmap.Operation) Operation {
	return Operation(n)
}

func NodeStateToGRPCMessage(n NodeState) netmap.NodeInfo_State {
	return netmap.NodeInfo_State(n)
}

func NodeStateFromRPCMessage(n netmap.NodeInfo_State) NodeState {
	return NodeState(n)
}

func AttributeToGRPCMessage(a *Attribute) *netmap.NodeInfo_Attribute {
	if a == nil {
		return nil
	}

	m := new(netmap.NodeInfo_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())
	m.SetParents(a.GetParents())

	return m
}

func AttributeFromGRPCMessage(m *netmap.NodeInfo_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	a := new(Attribute)

	a.SetKey(m.GetKey())
	a.SetValue(m.GetValue())
	a.SetParents(m.GetParents())

	return a
}

func NodeInfoToGRPCMessage(n *NodeInfo) *netmap.NodeInfo {
	if n == nil {
		return nil
	}

	m := new(netmap.NodeInfo)

	m.SetPublicKey(n.GetPublicKey())
	m.SetAddress(n.GetAddress())
	m.SetState(NodeStateToGRPCMessage(n.GetState()))

	attr := n.GetAttributes()
	attrMsg := make([]*netmap.NodeInfo_Attribute, 0, len(attr))

	for i := range attr {
		attrMsg = append(attrMsg, AttributeToGRPCMessage(attr[i]))
	}

	m.SetAttributes(attrMsg)

	return m
}

func NodeInfoFromGRPCMessage(m *netmap.NodeInfo) *NodeInfo {
	if m == nil {
		return nil
	}

	a := new(NodeInfo)

	a.SetPublicKey(m.GetPublicKey())
	a.SetAddress(m.GetAddress())
	a.SetState(NodeStateFromRPCMessage(m.GetState()))

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	a.SetAttributes(attr)

	return a
}

func LocalNodeInfoRequestBodyToGRPCMessage(r *LocalNodeInfoRequestBody) *netmap.LocalNodeInfoRequest_Body {
	if r == nil {
		return nil
	}

	return new(netmap.LocalNodeInfoRequest_Body)
}

func LocalNodeInfoRequestBodyFromGRPCMessage(m *netmap.LocalNodeInfoRequest_Body) *LocalNodeInfoRequestBody {
	if m == nil {
		return nil
	}

	return new(LocalNodeInfoRequestBody)
}

func LocalNodeInfoResponseBodyToGRPCMessage(r *LocalNodeInfoResponseBody) *netmap.LocalNodeInfoResponse_Body {
	if r == nil {
		return nil
	}

	m := new(netmap.LocalNodeInfoResponse_Body)

	m.SetVersion(refs.VersionToGRPCMessage(r.GetVersion()))
	m.SetNodeInfo(NodeInfoToGRPCMessage(r.GetNodeInfo()))

	return m
}

func LocalNodeInfoResponseBodyFromGRPCMessage(m *netmap.LocalNodeInfoResponse_Body) *LocalNodeInfoResponseBody {
	if m == nil {
		return nil
	}

	r := new(LocalNodeInfoResponseBody)
	r.SetVersion(refs.VersionFromGRPCMessage(m.GetVersion()))
	r.SetNodeInfo(NodeInfoFromGRPCMessage(m.GetNodeInfo()))

	return r
}

func LocalNodeInfoRequestToGRPCMessage(r *LocalNodeInfoRequest) *netmap.LocalNodeInfoRequest {
	if r == nil {
		return nil
	}

	m := new(netmap.LocalNodeInfoRequest)
	m.SetBody(LocalNodeInfoRequestBodyToGRPCMessage(r.GetBody()))

	session.RequestHeadersToGRPC(r, m)

	return m
}

func LocalNodeInfoRequestFromGRPCMessage(m *netmap.LocalNodeInfoRequest) *LocalNodeInfoRequest {
	if m == nil {
		return nil
	}

	r := new(LocalNodeInfoRequest)
	r.SetBody(LocalNodeInfoRequestBodyFromGRPCMessage(m.GetBody()))

	session.RequestHeadersFromGRPC(m, r)

	return r
}

func LocalNodeInfoResponseToGRPCMessage(r *LocalNodeInfoResponse) *netmap.LocalNodeInfoResponse {
	if r == nil {
		return nil
	}

	m := new(netmap.LocalNodeInfoResponse)
	m.SetBody(LocalNodeInfoResponseBodyToGRPCMessage(r.GetBody()))

	session.ResponseHeadersToGRPC(r, m)

	return m
}

func LocalNodeInfoResponseFromGRPCMessage(m *netmap.LocalNodeInfoResponse) *LocalNodeInfoResponse {
	if m == nil {
		return nil
	}

	r := new(LocalNodeInfoResponse)
	r.SetBody(LocalNodeInfoResponseBodyFromGRPCMessage(m.GetBody()))

	session.ResponseHeadersFromGRPC(m, r)

	return r
}
