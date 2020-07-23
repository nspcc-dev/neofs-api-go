package eacl

// OperationType is an enumeration of operation types for extended ACL.
type OperationType uint32

// HeaderType is an enumeration of header types for extended ACL.
type HeaderType uint32

// MatchType is an enumeration of match types for extended ACL.
type MatchType uint32

// Action is an enumeration of extended ACL actions.
type Action uint32

// Group is an enumeration of access groups.
type Group uint32

// Header is an interface of string key-value pair,
type Header interface {
	// Must return string identifier of header.
	Name() string

	// Must return string value of header.
	Value() string
}

// TypedHeader is an interface of Header and HeaderType pair.
type TypedHeader interface {
	Header

	// Must return type of filtered header.
	HeaderType() HeaderType
}

// HeaderFilter is an interface of grouped information about filtered header.
type HeaderFilter interface {
	// Must return match type of filter.
	MatchType() MatchType

	TypedHeader
}

// Target is an interface of grouped information about extended ACL rule target.
type Target interface {
	// Must return ACL target type.
	Group() Group

	// Must return public key list of ACL targets.
	KeyList() [][]byte
}

// Record is an interface of record of extended ACL rule table.
type Record interface {
	// Must return operation type of extended ACL rule.
	OperationType() OperationType

	// Must return list of header filters of extended ACL rule.
	HeaderFilters() []HeaderFilter

	// Must return target list of extended ACL rule.
	TargetList() []Target

	// Must return action of extended ACL rule.
	Action() Action
}

// Table is an interface of extended ACL table.
type Table interface {
	// Must return list of extended ACL rules.
	Records() []Record
}


