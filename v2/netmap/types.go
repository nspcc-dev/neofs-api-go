package netmap

// SetOp sets operation of the simple filter.
func (m *PlacementPolicy_FilterGroup_Filter_SimpleFilter) SetOp(v PlacementPolicy_FilterGroup_Filter_SimpleFilter_Operation) {
	if m != nil {
		m.Op = v
	}
}

// SetValue sets value of the simple filter.
func (m *PlacementPolicy_FilterGroup_Filter_SimpleFilter) SetValue(v string) {
	if m != nil {
		m.Args = &PlacementPolicy_FilterGroup_Filter_SimpleFilter_Value{
			Value: v,
		}
	}
}

// SetFArgs sets filter args of the simple filter.
func (m *PlacementPolicy_FilterGroup_Filter_SimpleFilter) SetFArgs(v *PlacementPolicy_FilterGroup_Filter_SimpleFilter_SimpleFilters) {
	if m != nil {
		m.Args = &PlacementPolicy_FilterGroup_Filter_SimpleFilter_FArgs{
			FArgs: v,
		}
	}
}

// SetFilters sets list of the simple filters.
func (m *PlacementPolicy_FilterGroup_Filter_SimpleFilter_SimpleFilters) SetFilters(v []*PlacementPolicy_FilterGroup_Filter_SimpleFilter) {
	if m != nil {
		m.Filters = v
	}
}

// SeyKey sets key of the filter.
func (m *PlacementPolicy_FilterGroup_Filter) SeyKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetF sets simple filter of the filter.
func (m *PlacementPolicy_FilterGroup_Filter) SetF(v *PlacementPolicy_FilterGroup_Filter_SimpleFilter) {
	if m != nil {
		m.F = v
	}
}

// SetCount sets count value of the selector.
func (m *PlacementPolicy_FilterGroup_Selector) SetCount(v uint32) {
	if m != nil {
		m.Count = v
	}
}

// SetKey sets key of the selector.
func (m *PlacementPolicy_FilterGroup_Selector) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetFilters sets list of the filters.
func (m *PlacementPolicy_FilterGroup) SetFilters(v []*PlacementPolicy_FilterGroup_Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetSelectors sets list of the selectors.
func (m *PlacementPolicy_FilterGroup) SetSelectors(v []*PlacementPolicy_FilterGroup_Selector) {
	if m != nil {
		m.Selectors = v
	}
}

// SetExclude sets exclude list.
func (m *PlacementPolicy_FilterGroup) SetExclude(v []uint32) {
	if m != nil {
		m.Exclude = v
	}
}

// SetReplFactor sets replication factor of the placement rule.
func (m *PlacementPolicy) SetReplFactor(v uint32) {
	if m != nil {
		m.ReplFactor = v
	}
}

// SetSfGroups sets list of the selector-filter groups.
func (m *PlacementPolicy) SetSfGroups(v []*PlacementPolicy_FilterGroup) {
	if m != nil {
		m.FilterGroups = v
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

// SetAddress sets node network address.
func (m *NodeInfo) SetAddress(v string) {
	if m != nil {
		m.Address = v
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
