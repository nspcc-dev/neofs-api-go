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

// SetEaclTable sets eACL table of the bearer token.
func (m *BearerToken_Body) SetEaclTable(v *EACLTable) {
	if m != nil {
		m.EaclTable = v
	}
}

// SetOwnerId sets identifier of the bearer token owner.
func (m *BearerToken_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetLifetime sets lifetime of the bearer token.
func (m *BearerToken_Body) SetLifetime(v *BearerToken_Body_TokenLifetime) {
	if m != nil {
		m.Lifetime = v
	}
}

// SetBody sets bearer token body.
func (m *BearerToken) SetBody(v *BearerToken_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetSignature sets bearer token signature.
func (m *BearerToken) SetSignature(v *refs.Signature) {
	if m != nil {
		m.Signature = v
	}
}

// SetExp sets epoch number of the token expiration.
func (m *BearerToken_Body_TokenLifetime) SetExp(v uint64) {
	if m != nil {
		m.Exp = v
	}
}

// SetNbf sets starting epoch number of the token.
func (m *BearerToken_Body_TokenLifetime) SetNbf(v uint64) {
	if m != nil {
		m.Nbf = v
	}
}

// SetIat sets the number of the epoch in which the token was issued.
func (m *BearerToken_Body_TokenLifetime) SetIat(v uint64) {
	if m != nil {
		m.Iat = v
	}
}
