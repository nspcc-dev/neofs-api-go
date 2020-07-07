package acl

// EACLFilterWrapper is a wrapper over EACLRecord_FilterInfo pointer.
type EACLFilterWrapper struct {
	filter *EACLRecord_FilterInfo
}

// EACLTargetWrapper is a wrapper over EACLRecord_TargetInfo pointer.
type EACLTargetWrapper struct {
	target *EACLRecord_TargetInfo
}

// EACLRecordWrapper is a wrapper over EACLRecord pointer.
type EACLRecordWrapper struct {
	record *EACLRecord
}

// EACLTableWrapper is a wrapper over EACLTable pointer.
type EACLTableWrapper struct {
	table *EACLTable
}

// WrapFilterInfo wraps EACLRecord_FilterInfo pointer.
//
// If argument is nil, new EACLRecord_FilterInfo is initialized.
func WrapFilterInfo(v *EACLRecord_FilterInfo) EACLFilterWrapper {
	if v == nil {
		v = new(EACLRecord_FilterInfo)
	}

	return EACLFilterWrapper{
		filter: v,
	}
}

// WrapEACLTarget wraps EACLRecord_TargetInfo pointer.
//
// If argument is nil, new EACLRecord_TargetInfo is initialized.
func WrapEACLTarget(v *EACLRecord_TargetInfo) EACLTargetWrapper {
	if v == nil {
		v = new(EACLRecord_TargetInfo)
	}

	return EACLTargetWrapper{
		target: v,
	}
}

// WrapEACLRecord wraps EACLRecord pointer.
//
// If argument is nil, new EACLRecord is initialized.
func WrapEACLRecord(v *EACLRecord) EACLRecordWrapper {
	if v == nil {
		v = new(EACLRecord)
	}

	return EACLRecordWrapper{
		record: v,
	}
}

// WrapEACLTable wraps EACLTable pointer.
//
// If argument is nil, new EACLTable is initialized.
func WrapEACLTable(v *EACLTable) EACLTableWrapper {
	if v == nil {
		v = new(EACLTable)
	}

	return EACLTableWrapper{
		table: v,
	}
}

// MatchType returns casted result of MatchType field getter.
//
// If filter is not initialized, 0 returns.
//
// Returns 0 if MatchType is not one of:
//  - EACLRecord_FilterInfo_StringEqual;
//  - EACLRecord_FilterInfo_StringNotEqual.
func (s EACLFilterWrapper) MatchType() (res MatchType) {
	if s.filter != nil {
		switch s.filter.GetMatchType() {
		case EACLRecord_FilterInfo_StringEqual:
			res = stringEqual
		case EACLRecord_FilterInfo_StringNotEqual:
			res = stringNotEqual
		}
	}

	return
}

// SetMatchType passes casted argument to MatchType field setter.
//
// If filter is not initialized, nothing changes.
//
// MatchType is set to EACLRecord_FilterInfo_MatchUnknown if argument is not one of:
//  - StringEqual;
//  - StringNotEqual.
func (s EACLFilterWrapper) SetMatchType(v MatchType) {
	if s.filter != nil {
		switch v {
		case stringEqual:
			s.filter.SetMatchType(EACLRecord_FilterInfo_StringEqual)
		case stringNotEqual:
			s.filter.SetMatchType(EACLRecord_FilterInfo_StringNotEqual)
		default:
			s.filter.SetMatchType(EACLRecord_FilterInfo_MatchUnknown)
		}
	}
}

// Name returns the result of HeaderName field getter.
//
// If filter is not initialized, empty string returns.
func (s EACLFilterWrapper) Name() string {
	if s.filter == nil {
		return ""
	}

	return s.filter.GetHeaderName()
}

// SetName passes argument to HeaderName field setter.
//
// If filter is not initialized, nothing changes.
func (s EACLFilterWrapper) SetName(v string) {
	if s.filter != nil {
		s.filter.SetHeaderName(v)
	}
}

// Value returns the result of HeaderVal field getter.
//
// If filter is not initialized, empty string returns.
func (s EACLFilterWrapper) Value() string {
	if s.filter == nil {
		return ""
	}

	return s.filter.GetHeaderVal()
}

// SetValue passes argument to HeaderVal field setter.
//
// If filter is not initialized, nothing changes.
func (s EACLFilterWrapper) SetValue(v string) {
	if s.filter != nil {
		s.filter.SetHeaderVal(v)
	}
}

// HeaderType returns the result of Header field getter.
//
// If filter is not initialized, 0 returns.
//
// Returns 0 if Header is not one of:
//  - EACLRecord_FilterInfo_Request;
//  - EACLRecord_FilterInfo_ObjectSystem;
//  - EACLRecord_FilterInfo_ObjectUser.
func (s EACLFilterWrapper) HeaderType() (res HeaderType) {
	if s.filter != nil {
		switch s.filter.GetHeader() {
		case EACLRecord_FilterInfo_Request:
			res = HdrTypeRequest
		case EACLRecord_FilterInfo_ObjectSystem:
			res = HdrTypeObjSys
		case EACLRecord_FilterInfo_ObjectUser:
			res = HdrTypeObjUsr
		}
	}

	return
}

// SetHeaderType passes casted argument to Header field setter.
//
// If filter is not initialized, nothing changes.
//
// Header is set to EACLRecord_FilterInfo_HeaderUnknown if argument is not one of:
//  - HdrTypeRequest;
//  - HdrTypeObjSys;
//  - HdrTypeObjUsr.
func (s EACLFilterWrapper) SetHeaderType(t HeaderType) {
	if s.filter != nil {
		switch t {
		case HdrTypeRequest:
			s.filter.SetHeader(EACLRecord_FilterInfo_Request)
		case HdrTypeObjSys:
			s.filter.SetHeader(EACLRecord_FilterInfo_ObjectSystem)
		case HdrTypeObjUsr:
			s.filter.SetHeader(EACLRecord_FilterInfo_ObjectUser)
		default:
			s.filter.SetHeader(EACLRecord_FilterInfo_HeaderUnknown)
		}
	}
}

// Target returns the result of Target field getter.
//
// If target is not initialized, Target_Unknown returns.
func (s EACLTargetWrapper) Target() Target {
	if s.target == nil {
		return Target_Unknown
	}

	return s.target.GetTarget()
}

// SetTarget passes argument to Target field setter.
//
// If target is not initialized, nothing changes.
func (s EACLTargetWrapper) SetTarget(v Target) {
	if s.target != nil {
		s.target.SetTarget(v)
	}
}

// KeyList returns the result of KeyList field getter.
//
// If target is not initialized, nil returns.
func (s EACLTargetWrapper) KeyList() [][]byte {
	if s.target == nil {
		return nil
	}

	return s.target.GetKeyList()
}

// SetKeyList passes argument to KeyList field setter.
//
// If target is not initialized, nothing changes.
func (s EACLTargetWrapper) SetKeyList(v [][]byte) {
	if s.target != nil {
		s.target.SetKeyList(v)
	}
}

// OperationType returns casted result of Operation field getter.
//
// If record is not initialized, 0 returns.
//
// Returns 0 if Operation is not one of:
//  - EACLRecord_HEAD;
//  - EACLRecord_PUT;
//  - EACLRecord_SEARCH;
//  - EACLRecord_GET;
//  - EACLRecord_GETRANGE;
//  - EACLRecord_GETRANGEHASH;
//  - EACLRecord_DELETE.
func (s EACLRecordWrapper) OperationType() (res OperationType) {
	if s.record != nil {
		switch s.record.GetOperation() {
		case EACLRecord_HEAD:
			res = OpTypeHead
		case EACLRecord_PUT:
			res = OpTypePut
		case EACLRecord_SEARCH:
			res = OpTypeSearch
		case EACLRecord_GET:
			res = OpTypeGet
		case EACLRecord_GETRANGE:
			res = OpTypeRange
		case EACLRecord_GETRANGEHASH:
			res = OpTypeRangeHash
		case EACLRecord_DELETE:
			res = OpTypeDelete
		}
	}

	return
}

// SetOperationType passes casted argument to Operation field setter.
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
func (s EACLRecordWrapper) SetOperationType(v OperationType) {
	if s.record != nil {
		switch v {
		case OpTypeHead:
			s.record.SetOperation(EACLRecord_HEAD)
		case OpTypePut:
			s.record.SetOperation(EACLRecord_PUT)
		case OpTypeSearch:
			s.record.SetOperation(EACLRecord_SEARCH)
		case OpTypeGet:
			s.record.SetOperation(EACLRecord_GET)
		case OpTypeRange:
			s.record.SetOperation(EACLRecord_GETRANGE)
		case OpTypeRangeHash:
			s.record.SetOperation(EACLRecord_GETRANGEHASH)
		case OpTypeDelete:
			s.record.SetOperation(EACLRecord_DELETE)
		default:
			s.record.SetOperation(EACLRecord_OPERATION_UNKNOWN)
		}
	}
}

// Action returns casted result of Action field getter.
//
// If record is not initialized, 0 returns.
//
// Returns 0 if Action is not one of:
//  - EACLRecord_Deny;
//  - EACLRecord_Allow.
func (s EACLRecordWrapper) Action() (res ExtendedACLAction) {
	if s.record != nil {
		switch s.record.GetAction() {
		case EACLRecord_Deny:
			res = ActionDeny
		case EACLRecord_Allow:
			res = ActionAllow
		}
	}

	return
}

// SetAction passes casted argument to Action field setter.
//
// If record is not initialized, nothing changes.
//
// Action is set to EACLRecord_ActionUnknown if argument is not one of:
//  - ActionDeny;
//  - ActionAllow.
func (s EACLRecordWrapper) SetAction(v ExtendedACLAction) {
	if s.record != nil {
		switch v {
		case ActionDeny:
			s.record.SetAction(EACLRecord_Deny)
		case ActionAllow:
			s.record.SetAction(EACLRecord_Allow)
		default:
			s.record.SetAction(EACLRecord_ActionUnknown)
		}
	}
}

// HeaderFilters wraps all elements from Filters field getter result and returns HeaderFilter list.
//
// If record is not initialized, nil returns.
func (s EACLRecordWrapper) HeaderFilters() []HeaderFilter {
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

// SetHeaderFilters converts HeaderFilter list to EACLRecord_FilterInfo list and passes it to Filters field setter.
//
// Ignores nil elements of argument.
// If record is not initialized, nothing changes.
func (s EACLRecordWrapper) SetHeaderFilters(v []HeaderFilter) {
	if s.record == nil {
		return
	}

	filters := make([]*EACLRecord_FilterInfo, 0, len(v))

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

// TargetList wraps all elements from Targets field getter result and returns ExtendedACLTarget list.
//
// If record is not initialized, nil returns.
func (s EACLRecordWrapper) TargetList() []ExtendedACLTarget {
	if s.record == nil {
		return nil
	}

	targets := s.record.GetTargets()

	res := make([]ExtendedACLTarget, 0, len(targets))

	for i := range targets {
		res = append(res, WrapEACLTarget(targets[i]))
	}

	return res
}

// SetTargetList converts ExtendedACLTarget list to EACLRecord_TargetInfo list and passes it to Targets field setter.
//
// Ignores nil elements of argument.
// If record is not initialized, nothing changes.
func (s EACLRecordWrapper) SetTargetList(v []ExtendedACLTarget) {
	if s.record == nil {
		return
	}

	targets := make([]*EACLRecord_TargetInfo, 0, len(v))

	for i := range v {
		if v[i] == nil {
			continue
		}

		w := WrapEACLTarget(nil)
		w.SetTarget(v[i].Target())
		w.SetKeyList(v[i].KeyList())

		targets = append(targets, w.target)
	}

	s.record.SetTargets(targets)
}

// Records wraps all elements from Records field getter result and returns ExtendedACLRecord list.
//
// If table is not initialized, nil returns.
func (s EACLTableWrapper) Records() []ExtendedACLRecord {
	if s.table == nil {
		return nil
	}

	records := s.table.GetRecords()

	res := make([]ExtendedACLRecord, 0, len(records))

	for i := range records {
		res = append(res, WrapEACLRecord(records[i]))
	}

	return res
}

// SetRecords converts ExtendedACLRecord list to EACLRecord list and passes it to Records field setter.
//
// Ignores nil elements of argument.
// If table is not initialized, nothing changes.
func (s EACLTableWrapper) SetRecords(v []ExtendedACLRecord) {
	if s.table == nil {
		return
	}

	records := make([]*EACLRecord, 0, len(v))

	for i := range v {
		if v[i] == nil {
			continue
		}

		w := WrapEACLRecord(nil)
		w.SetOperationType(v[i].OperationType())
		w.SetAction(v[i].Action())
		w.SetHeaderFilters(v[i].HeaderFilters())
		w.SetTargetList(v[i].TargetList())

		records = append(records, w.record)
	}

	s.table.SetRecords(records)
}

// MarshalBinary returns the result of Marshal method.
func (s EACLTableWrapper) MarshalBinary() ([]byte, error) {
	return s.table.Marshal()
}

// UnmarshalBinary passes argument to Unmarshal method and returns its result.
func (s EACLTableWrapper) UnmarshalBinary(data []byte) error {
	return s.table.Unmarshal(data)
}
