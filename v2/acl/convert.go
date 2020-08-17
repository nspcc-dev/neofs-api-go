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

func OperationToGRPCField(t Operation) acl.EACLRecord_Operation {
	switch t {
	case OperationPut:
		return acl.EACLRecord_PUT
	case OperationDelete:
		return acl.EACLRecord_DELETE
	case OperationGet:
		return acl.EACLRecord_GET
	case OperationHead:
		return acl.EACLRecord_HEAD
	case OperationSearch:
		return acl.EACLRecord_SEARCH
	case OperationRange:
		return acl.EACLRecord_GETRANGE
	case OperationRangeHash:
		return acl.EACLRecord_GETRANGEHASH
	default:
		return acl.EACLRecord_OPERATION_UNSPECIFIED
	}
}

func OperationFromGRPCField(t acl.EACLRecord_Operation) Operation {
	switch t {
	case acl.EACLRecord_PUT:
		return OperationPut
	case acl.EACLRecord_DELETE:
		return OperationDelete
	case acl.EACLRecord_GET:
		return OperationGet
	case acl.EACLRecord_HEAD:
		return OperationHead
	case acl.EACLRecord_SEARCH:
		return OperationSearch
	case acl.EACLRecord_GETRANGE:
		return OperationRange
	case acl.EACLRecord_GETRANGEHASH:
		return OperationRangeHash
	default:
		return OperationUnknown
	}
}

func ActionToGRPCField(t Action) acl.EACLRecord_Action {
	switch t {
	case ActionDeny:
		return acl.EACLRecord_DENY
	case ActionAllow:
		return acl.EACLRecord_ALLOW
	default:
		return acl.EACLRecord_ACTION_UNSPECIFIED
	}
}

func ActionFromGRPCField(t acl.EACLRecord_Action) Action {
	switch t {
	case acl.EACLRecord_DENY:
		return ActionDeny
	case acl.EACLRecord_ALLOW:
		return ActionAllow
	default:
		return ActionUnknown
	}
}

func HeaderTypeToGRPCField(t HeaderType) acl.EACLRecord_FilterInfo_Header {
	switch t {
	case HeaderTypeRequest:
		return acl.EACLRecord_FilterInfo_REQUEST
	case HeaderTypeObject:
		return acl.EACLRecord_FilterInfo_OBJECT
	default:
		return acl.EACLRecord_FilterInfo_HEADER_UNSPECIFIED
	}
}

func HeaderTypeFromGRPCField(t acl.EACLRecord_FilterInfo_Header) HeaderType {
	switch t {
	case acl.EACLRecord_FilterInfo_REQUEST:
		return HeaderTypeRequest
	case acl.EACLRecord_FilterInfo_OBJECT:
		return HeaderTypeObject
	default:
		return HeaderTypeUnknown
	}
}

func MatchTypeToGRPCField(t MatchType) acl.EACLRecord_FilterInfo_MatchType {
	switch t {
	case MatchTypeStringEqual:
		return acl.EACLRecord_FilterInfo_STRING_EQUAL
	case MatchTypeStringNotEqual:
		return acl.EACLRecord_FilterInfo_STRING_NOT_EQUAL
	default:
		return acl.EACLRecord_FilterInfo_MATCH_TYPE_UNSPECIFIED
	}
}

func MatchTypeFromGRPCField(t acl.EACLRecord_FilterInfo_MatchType) MatchType {
	switch t {
	case acl.EACLRecord_FilterInfo_STRING_EQUAL:
		return MatchTypeStringEqual
	case acl.EACLRecord_FilterInfo_STRING_NOT_EQUAL:
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
