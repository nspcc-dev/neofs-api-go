package acl

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/refs"
)

// OperationType is an enumeration of operation types for extended ACL.
type OperationType uint32

// HeaderType is an enumeration of header types for extended ACL.
type HeaderType uint32

// MatchType is an enumeration of match types for extended ACL.
type MatchType uint32

// ExtendedACLAction is an enumeration of extended ACL actions.
type ExtendedACLAction uint32

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

// TypedHeaderSource is a various types of header set interface.
type TypedHeaderSource interface {
	// Must return list of Header of particular type.
	// Must return false if there is no ability to compose header list.
	HeadersOfType(HeaderType) ([]Header, bool)
}

// HeaderFilter is an interface of grouped information about filtered header.
type HeaderFilter interface {
	// Must return match type of filter.
	MatchType() MatchType

	TypedHeader
}

// ExtendedACLTarget is an interface of grouped information about extended ACL rule target.
type ExtendedACLTarget interface {
	// Must return ACL target type.
	Target() Target

	// Must return public key list of ACL targets.
	KeyList() [][]byte
}

// ExtendedACLRecord is an interface of record of extended ACL rule table.
type ExtendedACLRecord interface {
	// Must return operation type of extended ACL rule.
	OperationType() OperationType

	// Must return list of header filters of extended ACL rule.
	HeaderFilters() []HeaderFilter

	// Must return target list of extended ACL rule.
	TargetList() []ExtendedACLTarget

	// Must return action of extended ACL rule.
	Action() ExtendedACLAction
}

// ExtendedACLTable is an interface of extended ACL table.
type ExtendedACLTable interface {
	// Must return list of extended ACL rules.
	Records() []ExtendedACLRecord
}

// ExtendedACLSource is an interface of storage of extended ACL tables with read access.
type ExtendedACLSource interface {
	// Must return extended ACL table by container ID key.
	GetExtendedACLTable(context.Context, refs.CID) (ExtendedACLTable, error)
}

// ExtendedACLStore is an interface of storage of extended ACL tables.
type ExtendedACLStore interface {
	ExtendedACLSource

	// Must store extended ACL table for container ID key.
	PutExtendedACLTable(context.Context, refs.CID, ExtendedACLTable) error
}

const (
	_ OperationType = iota

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
