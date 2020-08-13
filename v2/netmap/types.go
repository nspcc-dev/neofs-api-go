package netmap

// SetOp sets operation of the simple filter.
func (m *PlacementRule_SFGroup_Filter_SimpleFilter) SetOp(v PlacementRule_SFGroup_Filter_SimpleFilter_Operation) {
	if m != nil {
		m.Op = v
	}
}

// SetValue sets value of the simple filter.
func (m *PlacementRule_SFGroup_Filter_SimpleFilter) SetValue(v string) {
	if m != nil {
		m.Args = &PlacementRule_SFGroup_Filter_SimpleFilter_Value{
			Value: v,
		}
	}
}

// SetFArgs sets filter args of the simple filter.
func (m *PlacementRule_SFGroup_Filter_SimpleFilter) SetFArgs(v *PlacementRule_SFGroup_Filter_SimpleFilters) {
	if m != nil {
		m.Args = &PlacementRule_SFGroup_Filter_SimpleFilter_FArgs{
			FArgs: v,
		}
	}
}

// SetFilters sets list of the simple filters.
func (m *PlacementRule_SFGroup_Filter_SimpleFilters) SetFilters(v []*PlacementRule_SFGroup_Filter_SimpleFilter) {
	if m != nil {
		m.Filters = v
	}
}

// SeyKey sets key of the filter.
func (m *PlacementRule_SFGroup_Filter) SeyKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetF sets simple filter of the filter.
func (m *PlacementRule_SFGroup_Filter) SetF(v *PlacementRule_SFGroup_Filter_SimpleFilter) {
	if m != nil {
		m.F = v
	}
}

// SetCount sets count value of the selector.
func (m *PlacementRule_SFGroup_Selector) SetCount(v uint32) {
	if m != nil {
		m.Count = v
	}
}

// SetKey sets key of the selector.
func (m *PlacementRule_SFGroup_Selector) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetFilters sets list of the filters.
func (m *PlacementRule_SFGroup) SetFilters(v []*PlacementRule_SFGroup_Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetSelectors sets list of the selectors.
func (m *PlacementRule_SFGroup) SetSelectors(v []*PlacementRule_SFGroup_Selector) {
	if m != nil {
		m.Selectors = v
	}
}

// SetExclude sets exclude list.
func (m *PlacementRule_SFGroup) SetExclude(v []uint32) {
	if m != nil {
		m.Exclude = v
	}
}

// SetReplFactor sets replication factor of the placement rule.
func (m *PlacementRule) SetReplFactor(v uint32) {
	if m != nil {
		m.ReplFactor = v
	}
}

// SetSfGroups sets list of the selector-filter groups.
func (m *PlacementRule) SetSfGroups(v []*PlacementRule_SFGroup) {
	if m != nil {
		m.SfGroups = v
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
