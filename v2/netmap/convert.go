package netmap

import (
	neofsgrpc "github.com/nspcc-dev/neofs-api-go/rpc/grpc"
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

func (f *Filter) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.Filter

	if f != nil {
		m = new(netmap.Filter)

		m.SetKey(f.key)
		m.SetValue(f.value)
		m.SetName(f.name)
		m.SetOp(OperationToGRPCMessage(f.op))
		m.SetFilters(FiltersToGRPC(f.filters))
	}

	return m
}

func (f *Filter) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.Filter)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	f.filters, err = FiltersFromGRPC(v.GetFilters())
	if err != nil {
		return err
	}

	f.key = v.GetKey()
	f.value = v.GetValue()
	f.name = v.GetName()
	f.op = OperationFromGRPCMessage(v.GetOp())

	return nil
}

func FiltersToGRPC(fs []*Filter) (res []*netmap.Filter) {
	if fs != nil {
		res = make([]*netmap.Filter, 0, len(fs))

		for i := range fs {
			res = append(res, fs[i].ToGRPCMessage().(*netmap.Filter))
		}
	}

	return
}

func FiltersFromGRPC(fs []*netmap.Filter) (res []*Filter, err error) {
	if fs != nil {
		res = make([]*Filter, 0, len(fs))

		for i := range fs {
			var f *Filter

			if fs[i] != nil {
				f = new(Filter)

				err = f.FromGRPCMessage(fs[i])
				if err != nil {
					return
				}
			}

			res = append(res, f)
		}
	}

	return
}

func (s *Selector) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.Selector

	if s != nil {
		m = new(netmap.Selector)

		m.SetName(s.name)
		m.SetAttribute(s.attribute)
		m.SetFilter(s.filter)
		m.SetCount(s.count)
		m.SetClause(ClauseToGRPCMessage(s.clause))
	}

	return m
}

func (s *Selector) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.Selector)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	s.name = v.GetName()
	s.attribute = v.GetAttribute()
	s.filter = v.GetFilter()
	s.count = v.GetCount()
	s.clause = ClauseFromGRPCMessage(v.GetClause())

	return nil
}

func SelectorsToGRPC(ss []*Selector) (res []*netmap.Selector) {
	if ss != nil {
		res = make([]*netmap.Selector, 0, len(ss))

		for i := range ss {
			res = append(res, ss[i].ToGRPCMessage().(*netmap.Selector))
		}
	}

	return
}

func SelectorsFromGRPC(ss []*netmap.Selector) (res []*Selector, err error) {
	if ss != nil {
		res = make([]*Selector, 0, len(ss))

		for i := range ss {
			var s *Selector

			if ss[i] != nil {
				s = new(Selector)

				err = s.FromGRPCMessage(ss[i])
				if err != nil {
					return
				}
			}

			res = append(res, s)
		}
	}

	return
}

func (r *Replica) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.Replica

	if r != nil {
		m = new(netmap.Replica)

		m.SetSelector(r.selector)
		m.SetCount(r.count)
	}

	return m
}

func (r *Replica) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.Replica)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	r.selector = v.GetSelector()
	r.count = v.GetCount()

	return nil
}

func ReplicasToGRPC(rs []*Replica) (res []*netmap.Replica) {
	if rs != nil {
		res = make([]*netmap.Replica, 0, len(rs))

		for i := range rs {
			res = append(res, rs[i].ToGRPCMessage().(*netmap.Replica))
		}
	}

	return
}

func ReplicasFromGRPC(rs []*netmap.Replica) (res []*Replica, err error) {
	if rs != nil {
		res = make([]*Replica, 0, len(rs))

		for i := range rs {
			var r *Replica

			if rs[i] != nil {
				r = new(Replica)

				err = r.FromGRPCMessage(rs[i])
				if err != nil {
					return
				}
			}

			res = append(res, r)
		}
	}

	return
}

func (p *PlacementPolicy) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.PlacementPolicy

	if p != nil {
		m = new(netmap.PlacementPolicy)

		m.SetFilters(FiltersToGRPC(p.filters))
		m.SetSelectors(SelectorsToGRPC(p.selectors))
		m.SetReplicas(ReplicasToGRPC(p.replicas))
		m.SetContainerBackupFactor(p.backupFactor)
	}

	return m
}

func (p *PlacementPolicy) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.PlacementPolicy)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	p.filters, err = FiltersFromGRPC(v.GetFilters())
	if err != nil {
		return err
	}

	p.selectors, err = SelectorsFromGRPC(v.GetSelectors())
	if err != nil {
		return err
	}

	p.replicas, err = ReplicasFromGRPC(v.GetReplicas())
	if err != nil {
		return err
	}

	p.backupFactor = v.GetContainerBackupFactor()

	return nil
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

func (a *Attribute) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NodeInfo_Attribute

	if a != nil {
		m = new(netmap.NodeInfo_Attribute)

		m.SetKey(a.key)
		m.SetValue(a.value)
		m.SetParents(a.parents)
	}

	return m
}

func (a *Attribute) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NodeInfo_Attribute)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	a.key = v.GetKey()
	a.value = v.GetValue()
	a.parents = v.GetParents()

	return nil
}

func AttributesToGRPC(as []*Attribute) (res []*netmap.NodeInfo_Attribute) {
	if as != nil {
		res = make([]*netmap.NodeInfo_Attribute, 0, len(as))

		for i := range as {
			res = append(res, as[i].ToGRPCMessage().(*netmap.NodeInfo_Attribute))
		}
	}

	return
}

func AttributesFromGRPC(as []*netmap.NodeInfo_Attribute) (res []*Attribute, err error) {
	if as != nil {
		res = make([]*Attribute, 0, len(as))

		for i := range as {
			var a *Attribute

			if as[i] != nil {
				a = new(Attribute)

				err = a.FromGRPCMessage(as[i])
				if err != nil {
					return
				}
			}

			res = append(res, a)
		}
	}

	return
}

func (ni *NodeInfo) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NodeInfo

	if ni != nil {
		m = new(netmap.NodeInfo)

		m.SetPublicKey(ni.publicKey)
		m.SetAddresses(ni.addresses)
		m.SetState(NodeStateToGRPCMessage(ni.state))
		m.SetAttributes(AttributesToGRPC(ni.attributes))
	}

	return m
}

func (ni *NodeInfo) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NodeInfo)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	ni.attributes, err = AttributesFromGRPC(v.GetAttributes())
	if err != nil {
		return err
	}

	ni.publicKey = v.GetPublicKey()
	ni.addresses = v.GetAddresses()
	ni.state = NodeStateFromRPCMessage(v.GetState())

	return nil
}

func (l *LocalNodeInfoRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.LocalNodeInfoRequest_Body

	if l != nil {
		m = new(netmap.LocalNodeInfoRequest_Body)
	}

	return m
}

func (l *LocalNodeInfoRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.LocalNodeInfoRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

func (l *LocalNodeInfoRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.LocalNodeInfoRequest

	if l != nil {
		m = new(netmap.LocalNodeInfoRequest)

		m.SetBody(l.body.ToGRPCMessage().(*netmap.LocalNodeInfoRequest_Body))
		l.RequestHeaders.ToMessage(m)
	}

	return m
}

func (l *LocalNodeInfoRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.LocalNodeInfoRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		l.body = nil
	} else {
		if l.body == nil {
			l.body = new(LocalNodeInfoRequestBody)
		}

		err = l.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return l.RequestHeaders.FromMessage(v)
}

func (l *LocalNodeInfoResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.LocalNodeInfoResponse_Body

	if l != nil {
		m = new(netmap.LocalNodeInfoResponse_Body)

		m.SetVersion(l.version.ToGRPCMessage().(*refsGRPC.Version))
		m.SetNodeInfo(l.nodeInfo.ToGRPCMessage().(*netmap.NodeInfo))
	}

	return m
}

func (l *LocalNodeInfoResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.LocalNodeInfoResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	version := v.GetVersion()
	if version == nil {
		l.version = nil
	} else {
		if l.version == nil {
			l.version = new(refs.Version)
		}

		err = l.version.FromGRPCMessage(version)
		if err != nil {
			return err
		}
	}

	nodeInfo := v.GetNodeInfo()
	if nodeInfo == nil {
		l.nodeInfo = nil
	} else {
		if l.nodeInfo == nil {
			l.nodeInfo = new(NodeInfo)
		}

		err = l.nodeInfo.FromGRPCMessage(nodeInfo)
	}

	return err
}

func (l *LocalNodeInfoResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.LocalNodeInfoResponse

	if l != nil {
		m = new(netmap.LocalNodeInfoResponse)

		m.SetBody(l.body.ToGRPCMessage().(*netmap.LocalNodeInfoResponse_Body))
		l.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (l *LocalNodeInfoResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.LocalNodeInfoResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		l.body = nil
	} else {
		if l.body == nil {
			l.body = new(LocalNodeInfoResponseBody)
		}

		err = l.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return l.ResponseHeaders.FromMessage(v)
}

func (i *NetworkInfo) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NetworkInfo

	if i != nil {
		m = new(netmap.NetworkInfo)

		m.SetMagicNumber(i.magicNum)
		m.SetCurrentEpoch(i.curEpoch)
	}

	return m
}

func (i *NetworkInfo) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NetworkInfo)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	i.magicNum = v.GetMagicNumber()
	i.curEpoch = v.GetCurrentEpoch()

	return nil
}

func (l *NetworkInfoRequestBody) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NetworkInfoRequest_Body

	if l != nil {
		m = new(netmap.NetworkInfoRequest_Body)
	}

	return m
}

func (l *NetworkInfoRequestBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NetworkInfoRequest_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	return nil
}

func (l *NetworkInfoRequest) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NetworkInfoRequest

	if l != nil {
		m = new(netmap.NetworkInfoRequest)

		m.SetBody(l.body.ToGRPCMessage().(*netmap.NetworkInfoRequest_Body))
		l.RequestHeaders.ToMessage(m)
	}

	return m
}

func (l *NetworkInfoRequest) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NetworkInfoRequest)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		l.body = nil
	} else {
		if l.body == nil {
			l.body = new(NetworkInfoRequestBody)
		}

		err = l.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return l.RequestHeaders.FromMessage(v)
}

func (i *NetworkInfoResponseBody) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NetworkInfoResponse_Body

	if i != nil {
		m = new(netmap.NetworkInfoResponse_Body)

		m.SetNetworkInfo(i.netInfo.ToGRPCMessage().(*netmap.NetworkInfo))
	}

	return m
}

func (i *NetworkInfoResponseBody) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NetworkInfoResponse_Body)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	netInfo := v.GetNetworkInfo()
	if netInfo == nil {
		i.netInfo = nil
	} else {
		if i.netInfo == nil {
			i.netInfo = new(NetworkInfo)
		}

		err = i.netInfo.FromGRPCMessage(netInfo)
	}

	return err
}

func (l *NetworkInfoResponse) ToGRPCMessage() neofsgrpc.Message {
	var m *netmap.NetworkInfoResponse

	if l != nil {
		m = new(netmap.NetworkInfoResponse)

		m.SetBody(l.body.ToGRPCMessage().(*netmap.NetworkInfoResponse_Body))
		l.ResponseHeaders.ToMessage(m)
	}

	return m
}

func (l *NetworkInfoResponse) FromGRPCMessage(m neofsgrpc.Message) error {
	v, ok := m.(*netmap.NetworkInfoResponse)
	if !ok {
		return message.NewUnexpectedMessageType(m, v)
	}

	var err error

	body := v.GetBody()
	if body == nil {
		l.body = nil
	} else {
		if l.body == nil {
			l.body = new(NetworkInfoResponseBody)
		}

		err = l.body.FromGRPCMessage(body)
		if err != nil {
			return err
		}
	}

	return l.ResponseHeaders.FromMessage(v)
}
