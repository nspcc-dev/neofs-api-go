package netmap

import (
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
)

func PlacementPolicyToGRPCMessage(p *PlacementPolicy) *netmap.PlacementPolicy {
	if p == nil {
		return nil
	}

	// TODO: fill me
	return nil
}

func PlacementPolicyFromGRPCMessage(m *netmap.PlacementPolicy) *PlacementPolicy {
	if m == nil {
		return nil
	}

	// TODO: fill me
	return nil
}

func AttributeToGRPCMessage(a *Attribute) *netmap.NodeInfo_Attribute {
	if a == nil {
		return nil
	}

	m := new(netmap.NodeInfo_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())

	return m
}

func AttributeFromGRPCMessage(m *netmap.NodeInfo_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	a := new(Attribute)

	a.SetKey(m.GetKey())
	a.SetValue(m.GetValue())

	return a
}

func NodeInfoToGRPCMessage(n *NodeInfo) *netmap.NodeInfo {
	if n == nil {
		return nil
	}

	m := new(netmap.NodeInfo)

	m.SetPublicKey(n.GetPublicKey())
	m.SetAddress(n.GetAddress())

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

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	a.SetAttributes(attr)

	return a
}
