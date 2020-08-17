package acl

import (
	acl "github.com/nspcc-dev/neofs-api-go/v2/acl/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
)

func TargetToGRPCField(t Target) acl.Target {
	switch t {
	case TargetUser:
		return acl.Target_USER
	case TargetSystem:
		return acl.Target_SYSTEM
	case TargetOthers:
		return acl.Target_OTHERS
	default:
		return acl.Target_TARGET_UNSPECIFIED
	}
}

func TargetFromGRPCField(t acl.Target) Target {
	switch t {
	case acl.Target_USER:
		return TargetUser
	case acl.Target_SYSTEM:
		return TargetSystem
	case acl.Target_OTHERS:
		return TargetOthers
	default:
		return TargetUnknown
	}
}

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

func HeaderFilterToGRPCMessage(f *HeaderFilter) *acl.EACLRecord_FilterInfo {
	if f == nil {
		return nil
	}

	m := new(acl.EACLRecord_FilterInfo)

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

func HeaderFilterFromGRPCMessage(m *acl.EACLRecord_FilterInfo) *HeaderFilter {
	if m == nil {
		return nil
	}

	f := new(HeaderFilter)

	f.SetHeaderType(
		HeaderTypeFromGRPCField(m.GetHeader()),
	)

	f.SetMatchType(
		MatchTypeFromGRPCField(m.GetMatchType()),
	)

	f.SetName(m.GetHeaderName())
	f.SetValue(m.GetHeaderVal())

	return f
}

func TargetInfoToGRPCMessage(t *TargetInfo) *acl.EACLRecord_TargetInfo {
	if t == nil {
		return nil
	}

	m := new(acl.EACLRecord_TargetInfo)

	m.SetTarget(
		TargetToGRPCField(t.GetTarget()),
	)

	m.SetKeyList(t.GetKeyList())

	return m
}

func TargetInfoFromGRPCMessage(m *acl.EACLRecord_TargetInfo) *TargetInfo {
	if m == nil {
		return nil
	}

	t := new(TargetInfo)

	t.SetTarget(
		TargetFromGRPCField(m.GetTarget()),
	)

	t.SetKeyList(m.GetKeyList())

	return t
}

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
	filterMsg := make([]*acl.EACLRecord_FilterInfo, 0, len(filters))

	for i := range filters {
		filterMsg = append(filterMsg, HeaderFilterToGRPCMessage(filters[i]))
	}

	m.SetFilters(filterMsg)

	targets := r.GetTargets()
	targetMsg := make([]*acl.EACLRecord_TargetInfo, 0, len(targets))

	for i := range targets {
		targetMsg = append(targetMsg, TargetInfoToGRPCMessage(targets[i]))
	}

	m.SetTargets(targetMsg)

	return m
}

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
	targets := make([]*TargetInfo, 0, len(targetMsg))

	for i := range targetMsg {
		targets = append(targets, TargetInfoFromGRPCMessage(targetMsg[i]))
	}

	r.SetTargets(targets)

	return r
}

func TableToGRPCMessage(t *Table) *acl.EACLTable {
	if t == nil {
		return nil
	}

	m := new(acl.EACLTable)

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

func TableFromGRPCMessage(m *acl.EACLTable) *Table {
	if m == nil {
		return nil
	}

	t := new(Table)

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
