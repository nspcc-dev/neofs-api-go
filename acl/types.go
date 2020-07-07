package acl

// SetMatchType is MatchType field setter.
func (m *EACLRecord_FilterInfo) SetMatchType(v EACLRecord_FilterInfo_MatchType) {
	m.MatchType = v
}

// SetHeader is a Header field setter.
func (m *EACLRecord_FilterInfo) SetHeader(v EACLRecord_FilterInfo_Header) {
	m.Header = v
}

// SetHeaderName is a HeaderName field setter.
func (m *EACLRecord_FilterInfo) SetHeaderName(v string) {
	m.HeaderName = v
}

// SetHeaderVal is a HeaderVal field setter.
func (m *EACLRecord_FilterInfo) SetHeaderVal(v string) {
	m.HeaderVal = v
}

// SetTarget is a Target field setter.
func (m *EACLRecord_TargetInfo) SetTarget(v Target) {
	m.Target = v
}

// SetKeyList is a KeyList field setter.
func (m *EACLRecord_TargetInfo) SetKeyList(v [][]byte) {
	m.KeyList = v
}

// SetOperation is an Operation field setter.
func (m *EACLRecord) SetOperation(v EACLRecord_Operation) {
	m.Operation = v
}

// SetAction is an Action field setter.
func (m *EACLRecord) SetAction(v EACLRecord_Action) {
	m.Action = v
}

// SetFilters is a Filters field setter.
func (m *EACLRecord) SetFilters(v []*EACLRecord_FilterInfo) {
	m.Filters = v
}

// SetTargets is a Targets field setter.
func (m *EACLRecord) SetTargets(v []*EACLRecord_TargetInfo) {
	m.Targets = v
}

// SetRecords is a Records field setter.
func (m *EACLTable) SetRecords(v []*EACLRecord) {
	m.Records = v
}
