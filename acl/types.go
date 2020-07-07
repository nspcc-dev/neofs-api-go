package acl

const (
	_ MatchType = iota
	stringEqual
	stringNotEqual
)

const (
	// ActionUndefined is ExtendedACLAction used to mark value as undefined.
	// Most of the tools consider ActionUndefined as incalculable.
	// Using ActionUndefined in ExtendedACLRecord is unsafe.
	ActionUndefined ExtendedACLAction = iota

	// ActionAllow is ExtendedACLAction used to mark an applicability of ACL rule.
	ActionAllow

	// ActionDeny is ExtendedACLAction used to mark an inapplicability of ACL rule.
	ActionDeny
)

const (
	_ HeaderType = iota

	// HdrTypeRequest is a HeaderType for request header.
	HdrTypeRequest

	// HdrTypeObjSys is a HeaderType for system headers of object.
	HdrTypeObjSys

	// HdrTypeObjUsr is a HeaderType for user headers of object.
	HdrTypeObjUsr
)

const (
	// HdrObjSysNameID is a name of ID field in system header of object.
	HdrObjSysNameID = "ID"

	// HdrObjSysNameCID is a name of CID field in system header of object.
	HdrObjSysNameCID = "CID"

	// HdrObjSysNameOwnerID is a name of OwnerID field in system header of object.
	HdrObjSysNameOwnerID = "OWNER_ID"

	// HdrObjSysNameVersion is a name of Version field in system header of object.
	HdrObjSysNameVersion = "VERSION"

	// HdrObjSysNamePayloadLength is a name of PayloadLength field in system header of object.
	HdrObjSysNamePayloadLength = "PAYLOAD_LENGTH"

	// HdrObjSysNameCreatedUnix is a name of CreatedAt.UnitTime field in system header of object.
	HdrObjSysNameCreatedUnix = "CREATED_UNIX"

	// HdrObjSysNameCreatedEpoch is a name of CreatedAt.Epoch field in system header of object.
	HdrObjSysNameCreatedEpoch = "CREATED_EPOCH"

	// HdrObjSysLinkPrev is a name of previous link header in extended headers of object.
	HdrObjSysLinkPrev = "LINK_PREV"

	// HdrObjSysLinkNext is a name of next link header in extended headers of object.
	HdrObjSysLinkNext = "LINK_NEXT"

	// HdrObjSysLinkChild is a name of child link header in extended headers of object.
	HdrObjSysLinkChild = "LINK_CHILD"

	// HdrObjSysLinkPar is a name of parent link header in extended headers of object.
	HdrObjSysLinkPar = "LINK_PAR"

	// HdrObjSysLinkSG is a name of storage group link header in extended headers of object.
	HdrObjSysLinkSG = "LINK_SG"
)

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
