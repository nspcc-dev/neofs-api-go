package eacl

import (
	"github.com/nspcc-dev/neofs-api-go/acl"
)

// FilterWrapper is a wrapper over acl.EACLRecord_FilterInfo pointer.
type FilterWrapper struct {
	filter *acl.EACLRecord_FilterInfo
}

// TargetWrapper is a wrapper over acl.EACLRecord_TargetInfo pointer.
type TargetWrapper struct {
	target *acl.EACLRecord_TargetInfo
}

// RecordWrapper is a wrapper over acl.EACLRecord pointer.
type RecordWrapper struct {
	record *acl.EACLRecord
}

// TableWrapper is a wrapper over acl.EACLTable pointer.
type TableWrapper struct {
	table *acl.EACLTable
}

// WrapFilterInfo wraps EACLRecord_FilterInfo pointer.
//
// If argument is nil, new EACLRecord_FilterInfo is initialized.
func WrapFilterInfo(v *acl.EACLRecord_FilterInfo) FilterWrapper {
	if v == nil {
		v = new(acl.EACLRecord_FilterInfo)
	}

	return FilterWrapper{
		filter: v,
	}
}

// WrapTarget wraps EACLRecord_TargetInfo pointer.
//
// If argument is nil, new EACLRecord_TargetInfo is initialized.
func WrapTarget(v *acl.EACLRecord_TargetInfo) TargetWrapper {
	if v == nil {
		v = new(acl.EACLRecord_TargetInfo)
	}

	return TargetWrapper{
		target: v,
	}
}

// WrapRecord wraps EACLRecord pointer.
//
// If argument is nil, new EACLRecord is initialized.
func WrapRecord(v *acl.EACLRecord) RecordWrapper {
	if v == nil {
		v = new(acl.EACLRecord)
	}

	return RecordWrapper{
		record: v,
	}
}

// WrapTable wraps EACLTable pointer.
//
// If argument is nil, new EACLTable is initialized.
func WrapTable(v *acl.EACLTable) TableWrapper {
	if v == nil {
		v = new(acl.EACLTable)
	}

	return TableWrapper{
		table: v,
	}
}

// MatchType returns the match type of the filter.
//
// If filter is not initialized, 0 returns.
//
// Returns 0 if MatchType is not one of:
//  - EACLRecord_FilterInfo_StringEqual;
//  - EACLRecord_FilterInfo_StringNotEqual.
func (s FilterWrapper) MatchType() (res MatchType) {
	if s.filter != nil {
		switch s.filter.GetMatchType() {
		case acl.EACLRecord_FilterInfo_StringEqual:
			res = StringEqual
		case acl.EACLRecord_FilterInfo_StringNotEqual:
			res = StringNotEqual
		}
	}

	return
}

// SetMatchType sets the match type of the filter.
//
// If filter is not initialized, nothing changes.
//
// MatchType is set to EACLRecord_FilterInfo_MatchUnknown if argument is not one of:
//  - StringEqual;
//  - StringNotEqual.
func (s FilterWrapper) SetMatchType(v MatchType) {
	if s.filter != nil {
		switch v {
		case StringEqual:
			s.filter.SetMatchType(acl.EACLRecord_FilterInfo_StringEqual)
		case StringNotEqual:
			s.filter.SetMatchType(acl.EACLRecord_FilterInfo_StringNotEqual)
		default:
			s.filter.SetMatchType(acl.EACLRecord_FilterInfo_MatchUnknown)
		}
	}
}

// Name returns the name of filtering header.
//
// If filter is not initialized, empty string returns.
func (s FilterWrapper) Name() string {
	if s.filter == nil {
		return ""
	}

	return s.filter.GetHeaderName()
}

// SetName sets the name of the filtering header.
//
// If filter is not initialized, nothing changes.
func (s FilterWrapper) SetName(v string) {
	if s.filter != nil {
		s.filter.SetHeaderName(v)
	}
}

// Value returns the value of filtering header.
//
// If filter is not initialized, empty string returns.
func (s FilterWrapper) Value() string {
	if s.filter == nil {
		return ""
	}

	return s.filter.GetHeaderVal()
}

// SetValue sets the value of filtering header.
//
// If filter is not initialized, nothing changes.
func (s FilterWrapper) SetValue(v string) {
	if s.filter != nil {
		s.filter.SetHeaderVal(v)
	}
}

// HeaderType returns the header type of the filter.
//
// If filter is not initialized, HdrTypeUnknown returns.
//
// Returns HdrTypeUnknown if Header is not one of:
//  - EACLRecord_FilterInfo_Request;
//  - EACLRecord_FilterInfo_ObjectSystem;
//  - EACLRecord_FilterInfo_ObjectUser.
func (s FilterWrapper) HeaderType() (res HeaderType) {
	res = HdrTypeUnknown

	if s.filter != nil {
		switch s.filter.GetHeader() {
		case acl.EACLRecord_FilterInfo_Request:
			res = HdrTypeRequest
		case acl.EACLRecord_FilterInfo_ObjectSystem:
			res = HdrTypeObjSys
		case acl.EACLRecord_FilterInfo_ObjectUser:
			res = HdrTypeObjUsr
		}
	}

	return
}

// SetHeaderType sets the header type of the filter.
//
// If filter is not initialized, nothing changes.
//
// Header is set to EACLRecord_FilterInfo_HeaderUnknown if argument is not one of:
//  - HdrTypeRequest;
//  - HdrTypeObjSys;
//  - HdrTypeObjUsr.
func (s FilterWrapper) SetHeaderType(t HeaderType) {
	if s.filter != nil {
		switch t {
		case HdrTypeRequest:
			s.filter.SetHeader(acl.EACLRecord_FilterInfo_Request)
		case HdrTypeObjSys:
			s.filter.SetHeader(acl.EACLRecord_FilterInfo_ObjectSystem)
		case HdrTypeObjUsr:
			s.filter.SetHeader(acl.EACLRecord_FilterInfo_ObjectUser)
		default:
			s.filter.SetHeader(acl.EACLRecord_FilterInfo_HeaderUnknown)
		}
	}
}

// Group returns the access group of the target.
//
// If target is not initialized, GroupUnknown returns.
//
// Returns GroupUnknown if Target is not one of:
//  - Target_User;
//  - GroupSystem;
//  - GroupOthers.
func (s TargetWrapper) Group() (res Group) {
	res = GroupUnknown

	if s.target != nil {
		switch s.target.GetTarget() {
		case acl.Target_User:
			res = GroupUser
		case acl.Target_System:
			res = GroupSystem
		case acl.Target_Others:
			res = GroupOthers
		}
	}

	return
}

// SetGroup sets the access group of the target.
//
// If target is not initialized, nothing changes.
//
// Target is set to Target_Unknown if argument is not one of:
//  - GroupUser;
//  - GroupSystem;
//  - GroupOthers.
func (s TargetWrapper) SetGroup(g Group) {
	if s.target != nil {
		switch g {
		case GroupUser:
			s.target.SetTarget(acl.Target_User)
		case GroupSystem:
			s.target.SetTarget(acl.Target_System)
		case GroupOthers:
			s.target.SetTarget(acl.Target_Others)
		default:
			s.target.SetTarget(acl.Target_Unknown)
		}
	}
}

// KeyList returns the key list of the target.
//
// If target is not initialized, nil returns.
func (s TargetWrapper) KeyList() [][]byte {
	if s.target == nil {
		return nil
	}

	return s.target.GetKeyList()
}

// SetKeyList sets the key list of the target.
//
// If target is not initialized, nothing changes.
func (s TargetWrapper) SetKeyList(v [][]byte) {
	if s.target != nil {
		s.target.SetKeyList(v)
	}
}

// OperationType returns the operation type of the record.
//
// If record is not initialized, OpTypeUnknown returns.
//
// Returns OpTypeUnknown if Operation is not one of:
//  - EACLRecord_HEAD;
//  - EACLRecord_PUT;
//  - EACLRecord_SEARCH;
//  - EACLRecord_GET;
//  - EACLRecord_GETRANGE;
//  - EACLRecord_GETRANGEHASH;
//  - EACLRecord_DELETE.
func (s RecordWrapper) OperationType() (res OperationType) {
	res = OpTypeUnknown

	if s.record != nil {
		switch s.record.GetOperation() {
		case acl.EACLRecord_HEAD:
			res = OpTypeHead
		case acl.EACLRecord_PUT:
			res = OpTypePut
		case acl.EACLRecord_SEARCH:
			res = OpTypeSearch
		case acl.EACLRecord_GET:
			res = OpTypeGet
		case acl.EACLRecord_GETRANGE:
			res = OpTypeRange
		case acl.EACLRecord_GETRANGEHASH:
			res = OpTypeRangeHash
		case acl.EACLRecord_DELETE:
			res = OpTypeDelete
		}
	}

	return
}

// SetOperationType sets the operation type of the record.
//
// If record is not initialized, nothing changes.
//
// Operation is set to EACLRecord_OPERATION_UNKNOWN if argument is not one of:
//  - OpTypeHead;
//  - OpTypePut;
//  - OpTypeSearch;
//  - OpTypeGet;
//  - OpTypeRange;
//  - OpTypeRangeHash;
//  - OpTypeDelete.
func (s RecordWrapper) SetOperationType(v OperationType) {
	if s.record != nil {
		switch v {
		case OpTypeHead:
			s.record.SetOperation(acl.EACLRecord_HEAD)
		case OpTypePut:
			s.record.SetOperation(acl.EACLRecord_PUT)
		case OpTypeSearch:
			s.record.SetOperation(acl.EACLRecord_SEARCH)
		case OpTypeGet:
			s.record.SetOperation(acl.EACLRecord_GET)
		case OpTypeRange:
			s.record.SetOperation(acl.EACLRecord_GETRANGE)
		case OpTypeRangeHash:
			s.record.SetOperation(acl.EACLRecord_GETRANGEHASH)
		case OpTypeDelete:
			s.record.SetOperation(acl.EACLRecord_DELETE)
		default:
			s.record.SetOperation(acl.EACLRecord_OPERATION_UNKNOWN)
		}
	}
}

// Action returns the action of the record.
//
// If record is not initialized, ActionUnknown returns.
//
// Returns ActionUnknown if Action is not one of:
//  - EACLRecord_Deny;
//  - EACLRecord_Allow.
func (s RecordWrapper) Action() (res Action) {
	res = ActionUnknown

	if s.record != nil {
		switch s.record.GetAction() {
		case acl.EACLRecord_Deny:
			res = ActionDeny
		case acl.EACLRecord_Allow:
			res = ActionAllow
		}
	}

	return
}

// SetAction sets the action of the record.
//
// If record is not initialized, nothing changes.
//
// Action is set to EACLRecord_ActionUnknown if argument is not one of:
//  - ActionDeny;
//  - ActionAllow.
func (s RecordWrapper) SetAction(v Action) {
	if s.record != nil {
		switch v {
		case ActionDeny:
			s.record.SetAction(acl.EACLRecord_Deny)
		case ActionAllow:
			s.record.SetAction(acl.EACLRecord_Allow)
		default:
			s.record.SetAction(acl.EACLRecord_ActionUnknown)
		}
	}
}

// HeaderFilters returns the header filter list of the record.
//
// If record is not initialized, nil returns.
func (s RecordWrapper) HeaderFilters() []HeaderFilter {
	if s.record == nil {
		return nil
	}

	filters := s.record.GetFilters()

	res := make([]HeaderFilter, 0, len(filters))

	for i := range filters {
		res = append(res, WrapFilterInfo(filters[i]))
	}

	return res
}

// SetHeaderFilters sets the header filter list of the record.
//
// Ignores nil elements of argument.
// If record is not initialized, nothing changes.
func (s RecordWrapper) SetHeaderFilters(v []HeaderFilter) {
	if s.record == nil {
		return
	}

	filters := make([]*acl.EACLRecord_FilterInfo, 0, len(v))

	for i := range v {
		if v[i] == nil {
			continue
		}

		w := WrapFilterInfo(nil)
		w.SetMatchType(v[i].MatchType())
		w.SetHeaderType(v[i].HeaderType())
		w.SetName(v[i].Name())
		w.SetValue(v[i].Value())

		filters = append(filters, w.filter)
	}

	s.record.SetFilters(filters)
}

// TargetList returns the target list of the record.
//
// If record is not initialized, nil returns.
func (s RecordWrapper) TargetList() []Target {
	if s.record == nil {
		return nil
	}

	targets := s.record.GetTargets()

	res := make([]Target, 0, len(targets))

	for i := range targets {
		res = append(res, WrapTarget(targets[i]))
	}

	return res
}

// SetTargetList sets the target list of the record.
//
// Ignores nil elements of argument.
// If record is not initialized, nothing changes.
func (s RecordWrapper) SetTargetList(v []Target) {
	if s.record == nil {
		return
	}

	targets := make([]*acl.EACLRecord_TargetInfo, 0, len(v))

	for i := range v {
		if v[i] == nil {
			continue
		}

		w := WrapTarget(nil)
		w.SetGroup(v[i].Group())
		w.SetKeyList(v[i].KeyList())

		targets = append(targets, w.target)
	}

	s.record.SetTargets(targets)
}

// Records returns the record list of the table.
//
// If table is not initialized, nil returns.
func (s TableWrapper) Records() []Record {
	if s.table == nil {
		return nil
	}

	records := s.table.GetRecords()

	res := make([]Record, 0, len(records))

	for i := range records {
		res = append(res, WrapRecord(records[i]))
	}

	return res
}

// SetRecords sets the record list of the table.
//
// Ignores nil elements of argument.
// If table is not initialized, nothing changes.
func (s TableWrapper) SetRecords(v []Record) {
	if s.table == nil {
		return
	}

	records := make([]*acl.EACLRecord, 0, len(v))

	for i := range v {
		if v[i] == nil {
			continue
		}

		w := WrapRecord(nil)
		w.SetOperationType(v[i].OperationType())
		w.SetAction(v[i].Action())
		w.SetHeaderFilters(v[i].HeaderFilters())
		w.SetTargetList(v[i].TargetList())

		records = append(records, w.record)
	}

	s.table.SetRecords(records)
}


