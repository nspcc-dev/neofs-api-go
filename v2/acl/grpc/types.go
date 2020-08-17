package acl

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetContainerId sets container identifier of the eACL table.
func (m *EACLTable) SetContainerId(v *refs.ContainerID) {
	if m != nil {
		m.ContainerId = v
	}
}

// SetRecords sets record list of the eACL table.
func (m *EACLTable) SetRecords(v []*EACLRecord) {
	if m != nil {
		m.Records = v
	}
}

// SetOperation sets operation of the eACL record.
func (m *EACLRecord) SetOperation(v Operation) {
	if m != nil {
		m.Operation = v
	}
}

// SetAction sets action of the eACL record.
func (m *EACLRecord) SetAction(v Action) {
	if m != nil {
		m.Action = v
	}
}

// SetFilters sets filter list of the eACL record.
func (m *EACLRecord) SetFilters(v []*EACLRecord_FilterInfo) {
	if m != nil {
		m.Filters = v
	}
}

// SetTargets sets target list of the eACL record.
func (m *EACLRecord) SetTargets(v []*EACLRecord_TargetInfo) {
	if m != nil {
		m.Targets = v
	}
}

// SetHeader sets header type of the eACL filter.
func (m *EACLRecord_FilterInfo) SetHeader(v HeaderType) {
	if m != nil {
		m.Header = v
	}
}

// SetMatchType sets match type of the eACL filter.
func (m *EACLRecord_FilterInfo) SetMatchType(v MatchType) {
	if m != nil {
		m.MatchType = v
	}
}

// SetHeaderName sets header name of the eACL filter.
func (m *EACLRecord_FilterInfo) SetHeaderName(v string) {
	if m != nil {
		m.HeaderName = v
	}
}

// SetHeaderVal sets header value of the eACL filter.
func (m *EACLRecord_FilterInfo) SetHeaderVal(v string) {
	if m != nil {
		m.HeaderVal = v
	}
}

// SetTarget sets target group of the eACL target.
func (m *EACLRecord_TargetInfo) SetTarget(v Target) {
	if m != nil {
		m.Target = v
	}
}

// SetKeyList sets key list of the eACL target.
func (m *EACLRecord_TargetInfo) SetKeyList(v [][]byte) {
	if m != nil {
		m.KeyList = v
	}
}
