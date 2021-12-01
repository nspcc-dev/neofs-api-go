package netmap

import refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"

// SetReplicas of placement policy.
func (m *PlacementPolicy) SetReplicas(v []*Replica) {
	if m != nil {
		m.Replicas = v
	}
}

// SetContainerBackupFactor of placement policy.
func (m *PlacementPolicy) SetContainerBackupFactor(v uint32) {
	if m != nil {
		m.ContainerBackupFactor = v
	}
}

// SetSelectors of placement policy.
func (m *PlacementPolicy) SetSelectors(v []*Selector) {
	if m != nil {
		m.Selectors = v
	}
}

// SetFilters of placement policy.
func (m *PlacementPolicy) SetFilters(v []*Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetSubnetID sets ID of subnet.
func (m *PlacementPolicy) SetSubnetID(v *refs.SubnetID) {
	if m != nil {
		m.SubnetId = v
	}
}

// SetName of placement filter.
func (m *Filter) SetName(v string) {
	if m != nil {
		m.Name = v
	}
}

// SetKey of placement filter.
func (m *Filter) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetOperation of placement filter.
func (m *Filter) SetOp(v Operation) {
	if m != nil {
		m.Op = v
	}
}

// SetValue of placement filter.
func (m *Filter) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetFilters sets sub-filters of placement filter.
func (m *Filter) SetFilters(v []*Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetName of placement selector.
func (m *Selector) SetName(v string) {
	if m != nil {
		m.Name = v
	}
}

// SetCount of nodes of placement selector.
func (m *Selector) SetCount(v uint32) {
	if m != nil {
		m.Count = v
	}
}

// SetAttribute of nodes of placement selector.
func (m *Selector) SetAttribute(v string) {
	if m != nil {
		m.Attribute = v
	}
}

// SetFilter of placement selector.
func (m *Selector) SetFilter(v string) {
	if m != nil {
		m.Filter = v
	}
}

// SetClause of placement selector.
func (m *Selector) SetClause(v Clause) {
	if m != nil {
		m.Clause = v
	}
}

// SetCount of object replica.
func (m *Replica) SetCount(v uint32) {
	if m != nil {
		m.Count = v
	}
}

// SetSelector of object replica.
func (m *Replica) SetSelector(v string) {
	if m != nil {
		m.Selector = v
	}
}

// SetKey sets key to the node attribute.
func (m *NodeInfo_Attribute) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetValue sets value of the node attribute.
func (m *NodeInfo_Attribute) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetParent sets value of the node parents.
func (m *NodeInfo_Attribute) SetParents(v []string) {
	if m != nil {
		m.Parents = v
	}
}

// SetAddress sets node network address.
//
// Deprecated: use SetAddresses.
func (m *NodeInfo) SetAddress(v string) {
	m.SetAddresses([]string{v})
}

// SetAddresses sets list of network addresses of the node.
func (m *NodeInfo) SetAddresses(v []string) {
	if m != nil {
		m.Addresses = v
	}
}

// SetPublicKey sets node public key in a binary format.
func (m *NodeInfo) SetPublicKey(v []byte) {
	if m != nil {
		m.PublicKey = v
	}
}

// SetAttributes sets list of the node attributes.
func (m *NodeInfo) SetAttributes(v []*NodeInfo_Attribute) {
	if m != nil {
		m.Attributes = v
	}
}

// SetState sets node state.
func (m *NodeInfo) SetState(v NodeInfo_State) {
	if m != nil {
		m.State = v
	}
}

// SetCurrentEpoch sets number of the current epoch.
func (x *NetworkInfo) SetCurrentEpoch(v uint64) {
	if x != nil {
		x.CurrentEpoch = v
	}
}

// SetMagicNumber sets magic number of the sidechain.
func (x *NetworkInfo) SetMagicNumber(v uint64) {
	if x != nil {
		x.MagicNumber = v
	}
}

// SetMsPerBlock sets MillisecondsPerBlock network parameter.
func (x *NetworkInfo) SetMsPerBlock(v int64) {
	if x != nil {
		x.MsPerBlock = v
	}
}

// SetNetworkConfig sets NeoFS network configuration.
func (x *NetworkInfo) SetNetworkConfig(v *NetworkConfig) {
	if x != nil {
		x.NetworkConfig = v
	}
}

// FromString parses Clause from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *Clause) FromString(s string) bool {
	i, ok := Clause_value[s]
	if ok {
		*x = Clause(i)
	}

	return ok
}

// FromString parses Operation from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *Operation) FromString(s string) bool {
	i, ok := Operation_value[s]
	if ok {
		*x = Operation(i)
	}

	return ok
}

// FromString parses NodeInfo_State from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *NodeInfo_State) FromString(s string) bool {
	i, ok := NodeInfo_State_value[s]
	if ok {
		*x = NodeInfo_State(i)
	}

	return ok
}

// SetKey sets parameter key.
func (x *NetworkConfig_Parameter) SetKey(v []byte) {
	if x != nil {
		x.Key = v
	}
}

// SetValue sets parameter value.
func (x *NetworkConfig_Parameter) SetValue(v []byte) {
	if x != nil {
		x.Value = v
	}
}

// SetParameters sets NeoFS network parameters.
func (x *NetworkConfig) SetParameters(v []*NetworkConfig_Parameter) {
	if x != nil {
		x.Parameters = v
	}
}
