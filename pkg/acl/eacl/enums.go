package eacl

import (
	"errors"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Action taken if EACL record matched request.
// Action is compatible with v2 acl.Action enum.
type Action uint32

const (
	// ActionUnknown is an Action value used to mark action as undefined.
	ActionUnknown Action = iota

	// ActionAllow is an Action value that allows access to the operation from context.
	ActionAllow

	// ActionDeny is an Action value that denies access to the operation from context.
	ActionDeny
)

// Operation is a object service method to match request.
// Operation is compatible with v2 acl.Operation enum.
type Operation uint32

const (
	// OperationUnknown is an Operation value used to mark operation as undefined.
	OperationUnknown Operation = iota

	// OperationGet is an object get Operation.
	OperationGet

	// OperationHead is an Operation of getting the object header.
	OperationHead

	// OperationPut is an object put Operation.
	OperationPut

	// OperationDelete is an object delete Operation.
	OperationDelete

	// OperationSearch is an object search Operation.
	OperationSearch

	// OperationRange is an object payload range retrieval Operation.
	OperationRange

	// OperationRangeHash is an object payload range hashing Operation.
	OperationRangeHash
)

// Role is a group of request senders to match request.
// Role is compatible with v2 acl.Role enum.
type Role uint32

const (
	// RoleUnknown is a Role value used to mark role as undefined.
	RoleUnknown Role = iota

	// RoleUser is a group of senders that contains only key of container owner.
	RoleUser

	// RoleSystem is a group of senders that contains keys of container nodes and
	// inner ring nodes.
	RoleSystem

	// RoleOthers is a group of senders that contains none of above keys.
	RoleOthers
)

// Match is binary operation on filer name and value to check if request is matched.
// Match is compatible with v2 acl.MatchType enum.
type Match uint32

const (
	// MatchUnknown is a Match value used to mark matcher as undefined.
	MatchUnknown Match = iota

	// MatchStringEqual is a Match of string equality.
	MatchStringEqual

	// MatchStringNotEqual is a Match of string inequality.
	MatchStringNotEqual
)

// FilterHeaderType indicates source of headers to make matches.
// FilterHeaderType is compatible with v2 acl.HeaderType enum.
type FilterHeaderType uint32

const (
	// HeaderTypeUnknown is a FilterHeaderType value used to mark header type as undefined.
	HeaderTypeUnknown FilterHeaderType = iota

	// HeaderFromRequest is a FilterHeaderType for request X-Header.
	HeaderFromRequest

	// HeaderFromObject is a FilterHeaderType for object header.
	HeaderFromObject

	// HeaderFromService is a FilterHeaderType for service header.
	HeaderFromService
)

// ToV2 converts Action to v2 Action enum value.
func (a Action) ToV2() v2acl.Action {
	if a2, ok := actionToV2(a); ok {
		return a2
	}

	return v2acl.ActionUnknown
}

// converts Action to v2 Action enum value. Returns false if value is not a named constant.
func actionToV2(a Action) (v2acl.Action, bool) {
	switch a {
	default:
		return 0, false
	case ActionUnknown:
		return v2acl.ActionUnknown, true
	case ActionAllow:
		return v2acl.ActionAllow, true
	case ActionDeny:
		return v2acl.ActionDeny, true
	}
}

// ActionFromV2 converts v2 Action enum value to Action.
func ActionFromV2(action v2acl.Action) (a Action) {
	switch action {
	case v2acl.ActionAllow:
		a = ActionAllow
	case v2acl.ActionDeny:
		a = ActionDeny
	default:
		a = ActionUnknown
	}

	return a
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (a Action) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := a.MarshalText()
	return string(txt)
}

var errUnsupportedAction = errors.New("unsupported Action")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * ActionAllow: ALLOW;
//  * ActionDeny: DENY;
//  * ActionUnknown: ACTION_UNSPECIFIED.
func (a Action) MarshalText() ([]byte, error) {
	a2, ok := actionToV2(a)
	if !ok {
		return nil, errUnsupportedAction
	}

	return []byte(a2.String()), nil
}

func (a *Action) UnmarshalText(text []byte) error {
	var a2 v2acl.Action

	ok := a2.FromString(string(text))
	if !ok {
		return errUnsupportedAction
	}

	*a = ActionFromV2(a2)

	return nil
}

// FromString parses Action from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (a *Action) FromString(s string) bool {
	return a.UnmarshalText([]byte(s)) == nil
}

// ToV2 converts Operation to v2 Operation enum value.
func (o Operation) ToV2() v2acl.Operation {
	if o2, ok := operationToV2(o); ok {
		return o2
	}

	return v2acl.OperationUnknown
}

// converts Operation to v2 Operation enum value. Returns false if value is not a named constant.
func operationToV2(o Operation) (v2acl.Operation, bool) {
	switch o {
	default:
		return 0, false
	case OperationUnknown:
		return v2acl.OperationUnknown, true
	case OperationGet:
		return v2acl.OperationGet, true
	case OperationHead:
		return v2acl.OperationHead, true
	case OperationPut:
		return v2acl.OperationPut, true
	case OperationDelete:
		return v2acl.OperationDelete, true
	case OperationSearch:
		return v2acl.OperationSearch, true
	case OperationRange:
		return v2acl.OperationRange, true
	case OperationRangeHash:
		return v2acl.OperationRangeHash, true
	}
}

// OperationFromV2 converts v2 Operation enum value to Operation.
func OperationFromV2(operation v2acl.Operation) (o Operation) {
	switch operation {
	case v2acl.OperationGet:
		o = OperationGet
	case v2acl.OperationHead:
		o = OperationHead
	case v2acl.OperationPut:
		o = OperationPut
	case v2acl.OperationDelete:
		o = OperationDelete
	case v2acl.OperationSearch:
		o = OperationSearch
	case v2acl.OperationRange:
		o = OperationRange
	case v2acl.OperationRangeHash:
		o = OperationRangeHash
	default:
		o = OperationUnknown
	}

	return o
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (o Operation) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := o.MarshalText()
	return string(txt)
}

var errUnsupportedOperation = errors.New("unsupported Operation")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * OperationGet: GET;
//  * OperationHead: HEAD;
//  * OperationPut: PUT;
//  * OperationDelete: DELETE;
//  * OperationSearch: SEARCH;
//  * OperationRange: GETRANGE;
//  * OperationRangeHash: GETRANGEHASH;
//  * OperationUnknown: OPERATION_UNSPECIFIED.
func (o Operation) MarshalText() ([]byte, error) {
	o2, ok := operationToV2(o)
	if !ok {
		return nil, errUnsupportedOperation
	}

	return []byte(o2.String()), nil
}

func (o *Operation) UnmarshalText(text []byte) error {
	var o2 v2acl.Operation

	ok := o2.FromString(string(text))
	if !ok {
		return errUnsupportedAction
	}

	*o = OperationFromV2(o2)

	return nil
}

// FromString parses Operation from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (o *Operation) FromString(s string) bool {
	return o.UnmarshalText([]byte(s)) == nil
}

// ToV2 converts Role to v2 Role enum value.
func (r Role) ToV2() v2acl.Role {
	if r2, ok := roleToV2(r); ok {
		return r2
	}

	return v2acl.RoleUnknown
}

// converts Role to v2 Role enum value. Returns false if value is not a named constant.
func roleToV2(r Role) (v2acl.Role, bool) {
	switch r {
	default:
		return 0, false
	case RoleUnknown:
		return v2acl.RoleUnknown, true
	case RoleUser:
		return v2acl.RoleUser, true
	case RoleSystem:
		return v2acl.RoleSystem, true
	case RoleOthers:
		return v2acl.RoleOthers, true
	}
}

// RoleFromV2 converts v2 Role enum value to Role.
func RoleFromV2(role v2acl.Role) (r Role) {
	switch role {
	case v2acl.RoleUser:
		r = RoleUser
	case v2acl.RoleSystem:
		r = RoleSystem
	case v2acl.RoleOthers:
		r = RoleOthers
	default:
		r = RoleUnknown
	}

	return r
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (r Role) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := r.MarshalText()
	return string(txt)
}

var errUnsupportedRole = errors.New("unsupported Role")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * RoleUser: USER;
//  * RoleSystem: SYSTEM;
//  * RoleOthers: OTHERS;
//  * RoleUnknown: ROLE_UNKNOWN.
func (r Role) MarshalText() ([]byte, error) {
	r2, ok := roleToV2(r)
	if !ok {
		return nil, errUnsupportedRole
	}

	return []byte(r2.String()), nil
}

func (r *Role) UnmarshalText(text []byte) error {
	var r2 v2acl.Role

	ok := r2.FromString(string(text))
	if !ok {
		return errUnsupportedRole
	}

	*r = RoleFromV2(r2)

	return nil
}

// FromString parses Role from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (r *Role) FromString(s string) bool {
	return r.UnmarshalText([]byte(s)) == nil
}

// ToV2 converts Match to v2 MatchType enum value.
func (m Match) ToV2() v2acl.MatchType {
	if m2, ok := matchToV2(m); ok {
		return m2
	}

	return v2acl.MatchTypeUnknown
}

// converts Match to v2 MatchType enum value. Returns false if value is not a named constant.
func matchToV2(m Match) (v2acl.MatchType, bool) {
	switch m {
	default:
		return 0, false
	case MatchUnknown:
		return v2acl.MatchTypeUnknown, true
	case MatchStringEqual:
		return v2acl.MatchTypeStringEqual, true
	case MatchStringNotEqual:
		return v2acl.MatchTypeStringNotEqual, true
	}
}

// MatchFromV2 converts v2 MatchType enum value to Match.
func MatchFromV2(match v2acl.MatchType) (m Match) {
	switch match {
	case v2acl.MatchTypeStringEqual:
		m = MatchStringEqual
	case v2acl.MatchTypeStringNotEqual:
		m = MatchStringNotEqual
	default:
		m = MatchUnknown
	}

	return m
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (m Match) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := m.MarshalText()
	return string(txt)
}

var errUnsupportedMatch = errors.New("unsupported Match")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * MatchStringEqual: STRING_EQUAL;
//  * MatchStringNotEqual: STRING_NOT_EQUAL;
//  * MatchUnknown: MATCH_TYPE_UNSPECIFIED.
func (m Match) MarshalText() ([]byte, error) {
	m2, ok := matchToV2(m)
	if !ok {
		return nil, errUnsupportedMatch
	}

	return []byte(m2.String()), nil
}

func (m *Match) UnmarshalText(text []byte) error {
	var m2 v2acl.MatchType

	ok := m2.FromString(string(text))
	if !ok {
		return errUnsupportedMatch
	}

	*m = MatchFromV2(m2)

	return nil
}

// FromString parses Match from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (m *Match) FromString(s string) bool {
	return m.UnmarshalText([]byte(s)) == nil
}

// ToV2 converts FilterHeaderType to v2 HeaderType enum value.
func (h FilterHeaderType) ToV2() v2acl.HeaderType {
	if h2, ok := filterHeaderTypeToV2(h); ok {
		return h2
	}

	return v2acl.HeaderTypeUnknown
}

// converts FilterHeaderType to v2 HeaderType enum value. Returns false if value is not a named constant.
func filterHeaderTypeToV2(h FilterHeaderType) (v2acl.HeaderType, bool) {
	switch h {
	default:
		return 0, false
	case HeaderTypeUnknown:
		return v2acl.HeaderTypeUnknown, true
	case HeaderFromRequest:
		return v2acl.HeaderTypeRequest, true
	case HeaderFromObject:
		return v2acl.HeaderTypeObject, true
	case HeaderFromService:
		return v2acl.HeaderTypeService, true
	}
}

// FilterHeaderTypeFromV2 converts v2 HeaderType enum value to FilterHeaderType.
func FilterHeaderTypeFromV2(header v2acl.HeaderType) (h FilterHeaderType) {
	switch header {
	case v2acl.HeaderTypeRequest:
		h = HeaderFromRequest
	case v2acl.HeaderTypeObject:
		h = HeaderFromObject
	case v2acl.HeaderTypeService:
		h = HeaderFromService
	default:
		h = HeaderTypeUnknown
	}

	return h
}

// String implements fmt.Stringer.
//
// Use MarshalText to get the canonical text format.
func (h FilterHeaderType) String() string {
	// TODO: simplify stringer after FromString will be removed (neofs-api-go#346)
	txt, _ := h.MarshalText()
	return string(txt)
}

var errUnsupportedHeaderType = errors.New("unsupported FilterHeaderType")

// MarshalText implements encoding.TextMarshaler.
//
// Text mapping:
//  * HeaderFromRequest: REQUEST;
//  * HeaderFromObject: OBJECT;
//  * HeaderTypeUnknown: HEADER_UNSPECIFIED.
func (h FilterHeaderType) MarshalText() ([]byte, error) {
	m2, ok := filterHeaderTypeToV2(h)
	if !ok {
		return nil, errUnsupportedHeaderType
	}

	return []byte(m2.String()), nil
}

func (h *FilterHeaderType) UnmarshalText(text []byte) error {
	var h2 v2acl.HeaderType

	ok := h2.FromString(string(text))
	if !ok {
		return errUnsupportedHeaderType
	}

	*h = FilterHeaderTypeFromV2(h2)

	return nil
}

// FromString parses FilterHeaderType from a string representation.
// It is a reverse action to String().
//
// Returns true if s was parsed successfully.
//
// Deprecated: use UnmarshalText instead.
func (h *FilterHeaderType) FromString(s string) bool {
	return h.UnmarshalText([]byte(s)) == nil
}
