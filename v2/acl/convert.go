package acl

import (
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

// RoleToGRPCField converts unified role enum into grpc enum.
func RoleToGRPCField(t Role) acl.Role {
	switch t {
	case RoleUser:
		return acl.Role_USER
	case RoleSystem:
		return acl.Role_SYSTEM
	case RoleOthers:
		return acl.Role_OTHERS
	default:
		return acl.Role_ROLE_UNSPECIFIED
	}
}

// RoleFromGRPCField converts grpc enum into unified role enum.
func RoleFromGRPCField(t acl.Role) Role {
	switch t {
	case acl.Role_USER:
		return RoleUser
	case acl.Role_SYSTEM:
		return RoleSystem
	case acl.Role_OTHERS:
		return RoleOthers
	default:
		return RoleUnknown
	}
}

// OperationToGRPCField converts unified operation enum into grpc enum.
func OperationToGRPCField(t Operation) acl.Operation {
	switch t {
	case OperationPut:
		return acl.Operation_PUT
	case OperationDelete:
		return acl.Operation_DELETE
	case OperationGet:
		return acl.Operation_GET
	case OperationHead:
		return acl.Operation_HEAD
	case OperationSearch:
		return acl.Operation_SEARCH
	case OperationRange:
		return acl.Operation_GETRANGE
	case OperationRangeHash:
		return acl.Operation_GETRANGEHASH
	default:
		return acl.Operation_OPERATION_UNSPECIFIED
	}
}

// OperationFromGRPCField converts grpc enum into unified operation enum.
func OperationFromGRPCField(t acl.Operation) Operation {
	switch t {
	case acl.Operation_PUT:
		return OperationPut
	case acl.Operation_DELETE:
		return OperationDelete
	case acl.Operation_GET:
		return OperationGet
	case acl.Operation_HEAD:
		return OperationHead
	case acl.Operation_SEARCH:
		return OperationSearch
	case acl.Operation_GETRANGE:
		return OperationRange
	case acl.Operation_GETRANGEHASH:
		return OperationRangeHash
	default:
		return OperationUnknown
	}
}

// ActionToGRPCField converts unified action enum into grpc enum.
func ActionToGRPCField(t Action) acl.Action {
	switch t {
	case ActionDeny:
		return acl.Action_DENY
	case ActionAllow:
		return acl.Action_ALLOW
	default:
		return acl.Action_ACTION_UNSPECIFIED
	}
}

// ActionFromGRPCField converts grpc enum into unified action enum.
func ActionFromGRPCField(t acl.Action) Action {
	switch t {
	case acl.Action_DENY:
		return ActionDeny
	case acl.Action_ALLOW:
		return ActionAllow
	default:
		return ActionUnknown
	}
}

// HeaderTypeToGRPCField converts unified header type enum into grpc enum.
func HeaderTypeToGRPCField(t HeaderType) acl.HeaderType {
	switch t {
	case HeaderTypeRequest:
		return acl.HeaderType_REQUEST
	case HeaderTypeObject:
		return acl.HeaderType_OBJECT
	default:
		return acl.HeaderType_HEADER_UNSPECIFIED
	}
}

// HeaderTypeFromGRPCField converts grpc enum into unified header type enum.
func HeaderTypeFromGRPCField(t acl.HeaderType) HeaderType {
	switch t {
	case acl.HeaderType_REQUEST:
		return HeaderTypeRequest
	case acl.HeaderType_OBJECT:
		return HeaderTypeObject
	default:
		return HeaderTypeUnknown
	}
}

// MatchTypeToGRPCField converts unified match type enum into grpc enum.
func MatchTypeToGRPCField(t MatchType) acl.MatchType {
	switch t {
	case MatchTypeStringEqual:
		return acl.MatchType_STRING_EQUAL
	case MatchTypeStringNotEqual:
		return acl.MatchType_STRING_NOT_EQUAL
	default:
		return acl.MatchType_MATCH_TYPE_UNSPECIFIED
	}
}

// MatchTypeFromGRPCField converts grpc enum into unified match type enum.
func MatchTypeFromGRPCField(t acl.MatchType) MatchType {
	switch t {
	case acl.MatchType_STRING_EQUAL:
		return MatchTypeStringEqual
	case acl.MatchType_STRING_NOT_EQUAL:
		return MatchTypeStringNotEqual
	default:
		return MatchTypeUnknown
	}
}

// HeaderFilterToGRPCMessage converts unified header filter struct into grpc struct.
func HeaderFilterToGRPCMessage(f *HeaderFilter) *acl.EACLRecord_Filter {
	if f == nil {
		return nil
	}

	m := new(acl.EACLRecord_Filter)

	m.SetHeader(
		HeaderTypeToGRPCField(f.GetHeaderType()),
	)

	m.SetMatchType(
		MatchTypeToGRPCField(f.GetMatchType()),
	)

	m.SetHeaderName(f.GetName())
	m.SetHeaderVal(f.GetValue())

	return m
}

// HeaderFilterFromGRPCMessage converts grpc struct into unified header filter struct.
func HeaderFilterFromGRPCMessage(m *acl.EACLRecord_Filter) *HeaderFilter {
	if m == nil {
		return nil
	}

	f := new(HeaderFilter)

	f.SetHeaderType(
		HeaderTypeFromGRPCField(m.GetHeaderType()),
	)

	f.SetMatchType(
		MatchTypeFromGRPCField(m.GetMatchType()),
	)

	f.SetName(m.GetHeaderName())
	f.SetValue(m.GetHeaderVal())

	return f
}

// TargetToGRPCMessage converts unified role info struct into grpc struct.
func TargetToGRPCMessage(t *Target) *acl.EACLRecord_Target {
	if t == nil {
		return nil
	}

	m := new(acl.EACLRecord_Target)

	m.SetRole(
		RoleToGRPCField(t.GetRole()),
	)

	m.SetKeys(t.GetKeys())

	return m
}

// TargetInfoFromGRPCMessage converts grpc struct into unified role info struct.
func TargetInfoFromGRPCMessage(m *acl.EACLRecord_Target) *Target {
	if m == nil {
		return nil
	}

	t := new(Target)

	t.SetRole(
		RoleFromGRPCField(m.GetRole()),
	)

	t.SetKeys(m.GetKeys())

	return t
}

// RecordToGRPCMessage converts unified acl record struct into grpc struct.
func RecordToGRPCMessage(r *Record) *acl.EACLRecord {
	if r == nil {
		return nil
	}

	m := new(acl.EACLRecord)

	m.SetOperation(
		OperationToGRPCField(r.GetOperation()),
	)

	m.SetAction(
		ActionToGRPCField(r.GetAction()),
	)

	filters := r.GetFilters()
	filterMsg := make([]*acl.EACLRecord_Filter, 0, len(filters))

	for i := range filters {
		filterMsg = append(filterMsg, HeaderFilterToGRPCMessage(filters[i]))
	}

	m.SetFilters(filterMsg)

	targets := r.GetTargets()
	targetMsg := make([]*acl.EACLRecord_Target, 0, len(targets))

	for i := range targets {
		targetMsg = append(targetMsg, TargetToGRPCMessage(targets[i]))
	}

	m.SetTargets(targetMsg)

	return m
}

// RecordFromGRPCMessage converts grpc struct into unified acl record struct.
func RecordFromGRPCMessage(m *acl.EACLRecord) *Record {
	if m == nil {
		return nil
	}

	r := new(Record)

	r.SetOperation(
		OperationFromGRPCField(m.GetOperation()),
	)

	r.SetAction(
		ActionFromGRPCField(m.GetAction()),
	)

	filterMsg := m.GetFilters()
	filters := make([]*HeaderFilter, 0, len(filterMsg))

	for i := range filterMsg {
		filters = append(filters, HeaderFilterFromGRPCMessage(filterMsg[i]))
	}

	r.SetFilters(filters)

	targetMsg := m.GetTargets()
	targets := make([]*Target, 0, len(targetMsg))

	for i := range targetMsg {
		targets = append(targets, TargetInfoFromGRPCMessage(targetMsg[i]))
	}

	r.SetTargets(targets)

	return r
}

// TableToGRPCMessage converts unified acl table struct into grpc struct.
func TableToGRPCMessage(t *Table) *acl.EACLTable {
	if t == nil {
		return nil
	}

	m := new(acl.EACLTable)

	m.SetVersion(
		refs.VersionToGRPCMessage(t.GetVersion()),
	)
	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(t.GetContainerID()),
	)

	records := t.GetRecords()
	recordMsg := make([]*acl.EACLRecord, 0, len(records))

	for i := range records {
		recordMsg = append(recordMsg, RecordToGRPCMessage(records[i]))
	}

	m.SetRecords(recordMsg)

	return m
}

// TableFromGRPCMessage converts grpc struct into unified acl table struct.
func TableFromGRPCMessage(m *acl.EACLTable) *Table {
	if m == nil {
		return nil
	}

	t := new(Table)

	t.SetVersion(
		refs.VersionFromGRPCMessage(m.GetVersion()),
	)
	t.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	recordMsg := m.GetRecords()
	records := make([]*Record, 0, len(recordMsg))

	for i := range recordMsg {
		records = append(records, RecordFromGRPCMessage(recordMsg[i]))
	}

	t.SetRecords(records)

	return t
}

func TokenLifetimeToGRPCMessage(tl *TokenLifetime) *acl.BearerToken_Body_TokenLifetime {
	if tl == nil {
		return nil
	}

	m := new(acl.BearerToken_Body_TokenLifetime)

	m.SetExp(tl.GetExp())
	m.SetNbf(tl.GetNbf())
	m.SetIat(tl.GetIat())

	return m
}

func TokenLifetimeFromGRPCMessage(m *acl.BearerToken_Body_TokenLifetime) *TokenLifetime {
	if m == nil {
		return nil
	}

	tl := new(TokenLifetime)

	tl.SetExp(m.GetExp())
	tl.SetNbf(m.GetNbf())
	tl.SetIat(m.GetIat())

	return tl
}

func BearerTokenBodyToGRPCMessage(v *BearerTokenBody) *acl.BearerToken_Body {
	if v == nil {
		return nil
	}

	m := new(acl.BearerToken_Body)

	m.SetEaclTable(
		TableToGRPCMessage(v.GetEACL()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(v.GetOwnerID()),
	)

	m.SetLifetime(
		TokenLifetimeToGRPCMessage(v.GetLifetime()),
	)

	return m
}

func BearerTokenBodyFromGRPCMessage(m *acl.BearerToken_Body) *BearerTokenBody {
	if m == nil {
		return nil
	}

	bt := new(BearerTokenBody)

	bt.SetEACL(
		TableFromGRPCMessage(m.GetEaclTable()),
	)

	bt.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	bt.SetLifetime(
		TokenLifetimeFromGRPCMessage(m.GetLifetime()),
	)

	return bt
}

func BearerTokenToGRPCMessage(t *BearerToken) *acl.BearerToken {
	if t == nil {
		return nil
	}

	m := new(acl.BearerToken)

	m.SetBody(
		BearerTokenBodyToGRPCMessage(t.GetBody()),
	)

	m.SetSignature(
		refs.SignatureToGRPCMessage(t.GetSignature()),
	)

	return m
}

func BearerTokenFromGRPCMessage(m *acl.BearerToken) *BearerToken {
	if m == nil {
		return nil
	}

	bt := new(BearerToken)

	bt.SetBody(
		BearerTokenBodyFromGRPCMessage(m.GetBody()),
	)

	bt.SetSignature(
		refs.SignatureFromGRPCMessage(m.GetSignature()),
	)

	return bt
}
