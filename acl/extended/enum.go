package eacl

const (
	// MatchUnknown is a MatchType value used to mark value as undefined.
	// Most of the tools consider MatchUnknown as incalculable.
	// Using MatchUnknown in HeaderFilter is unsafe.
	MatchUnknown MatchType = iota

	// StringEqual is a MatchType of string equality.
	StringEqual

	// StringNotEqual is a MatchType of string inequality.
	StringNotEqual
)

const (
	// ActionUnknown is Action used to mark value as undefined.
	// Most of the tools consider ActionUnknown as incalculable.
	// Using ActionUnknown in Record is unsafe.
	ActionUnknown Action = iota

	// ActionAllow is Action used to mark an applicability of ACL rule.
	ActionAllow

	// ActionDeny is Action used to mark an inapplicability of ACL rule.
	ActionDeny
)

const (
	// GroupUnknown is a Group value used to mark value as undefined.
	// Most of the tools consider GroupUnknown as incalculable.
	// Using GroupUnknown in Target is unsafe.
	GroupUnknown Group = iota

	// GroupUser is a Group value for User access group.
	GroupUser

	// GroupSystem is a Group value for System access group.
	GroupSystem

	// GroupOthers is a Group value for Others access group.
	GroupOthers
)

const (
	// HdrTypeUnknown is a HeaderType value used to mark value as undefined.
	// Most of the tools consider HdrTypeUnknown as incalculable.
	// Using HdrTypeUnknown in HeaderFilter is unsafe.
	HdrTypeUnknown HeaderType = iota

	// HdrTypeRequest is a HeaderType for request header.
	HdrTypeRequest

	// HdrTypeObjSys is a HeaderType for system headers of object.
	HdrTypeObjSys

	// HdrTypeObjUsr is a HeaderType for user headers of object.
	HdrTypeObjUsr
)

const (
	// OpTypeUnknown is a OperationType value used to mark value as undefined.
	// Most of the tools consider OpTypeUnknown as incalculable.
	// Using OpTypeUnknown in Record is unsafe.
	OpTypeUnknown OperationType = iota

	// OpTypeGet is an OperationType for object.Get RPC
	OpTypeGet

	// OpTypePut is an OperationType for object.Put RPC
	OpTypePut

	// OpTypeHead is an OperationType for object.Head RPC
	OpTypeHead

	// OpTypeSearch is an OperationType for object.Search RPC
	OpTypeSearch

	// OpTypeDelete is an OperationType for object.Delete RPC
	OpTypeDelete

	// OpTypeRange is an OperationType for object.GetRange RPC
	OpTypeRange

	// OpTypeRangeHash is an OperationType for object.GetRangeHash RPC
	OpTypeRangeHash
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

