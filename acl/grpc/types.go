package acl

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
)

// SetVersion sets version of EACL rules in table.
func (m *EACLTable) SetVersion(v *refs.Version) {
	if m != nil {
		m.Version = v
	}
}

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
func (m *EACLRecord) SetFilters(v []*EACLRecord_Filter) {
	if m != nil {
		m.Filters = v
	}
}

// SetTargets sets target list of the eACL record.
func (m *EACLRecord) SetTargets(v []*EACLRecord_Target) {
	if m != nil {
		m.Targets = v
	}
}

// SetHeader sets header type of the eACL filter.
func (m *EACLRecord_Filter) SetHeader(v HeaderType) {
	if m != nil {
		m.HeaderType = v
	}
}

// SetMatchType sets match type of the eACL filter.
func (m *EACLRecord_Filter) SetMatchType(v MatchType) {
	if m != nil {
		m.MatchType = v
	}
}

// SetKey sets key of the eACL filter.
func (m *EACLRecord_Filter) SetKey(v string) {
	if m != nil {
		m.Key = v
	}
}

// SetValue sets value of the eACL filter.
func (m *EACLRecord_Filter) SetValue(v string) {
	if m != nil {
		m.Value = v
	}
}

// SetRole sets target group of the eACL target.
func (m *EACLRecord_Target) SetRole(v Role) {
	if m != nil {
		m.Role = v
	}
}

// SetKeys of the eACL target.
func (m *EACLRecord_Target) SetKeys(v [][]byte) {
	if m != nil {
		m.Keys = v
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

// FromString parses Action from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *Action) FromString(s string) bool {
	i, ok := Action_value[s]
	if ok {
		*x = Action(i)
	}

	return ok
}

// FromString parses Role from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *Role) FromString(s string) bool {
	i, ok := Role_value[s]
	if ok {
		*x = Role(i)
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

// FromString parses MatchType from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *MatchType) FromString(s string) bool {
	i, ok := MatchType_value[s]
	if ok {
		*x = MatchType(i)
	}

	return ok
}

// FromString parses HeaderType from a string representation,
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
func (x *HeaderType) FromString(s string) bool {
	i, ok := HeaderType_value[s]
	if ok {
		*x = HeaderType(i)
	}

	return ok
}
