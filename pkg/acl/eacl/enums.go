package eacl

import (
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Action taken if EACL record matched request.
type Action uint32

const (
	ActionUnknown Action = iota
	ActionAllow
	ActionDeny
)

// Operation is a object service method to match request.
type Operation uint32

const (
	OperationUnknown Operation = iota
	OperationGet
	OperationHead
	OperationPut
	OperationDelete
	OperationSearch
	OperationRange
	OperationRangeHash
)

// Role is a group of request senders to match request.
type Role uint32

const (
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
type Match uint32

const (
	MatchUnknown Match = iota
	MatchStringEqual
	MatchStringNotEqual
)

// FilterHeaderType indicates source of headers to make matches.
type FilterHeaderType uint32

const (
	HeaderTypeUnknown FilterHeaderType = iota
	HeaderFromRequest
	HeaderFromObject
)

func (a Action) ToV2() v2acl.Action {
	switch a {
	case ActionAllow:
		return v2acl.ActionAllow
	case ActionDeny:
		return v2acl.ActionDeny
	default:
		return v2acl.ActionUnknown
	}
}

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

func (o Operation) ToV2() v2acl.Operation {
	switch o {
	case OperationGet:
		return v2acl.OperationGet
	case OperationHead:
		return v2acl.OperationHead
	case OperationPut:
		return v2acl.OperationPut
	case OperationDelete:
		return v2acl.OperationDelete
	case OperationSearch:
		return v2acl.OperationSearch
	case OperationRange:
		return v2acl.OperationRange
	case OperationRangeHash:
		return v2acl.OperationRangeHash
	default:
		return v2acl.OperationUnknown
	}
}

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

func (r Role) ToV2() v2acl.Role {
	switch r {
	case RoleUser:
		return v2acl.RoleUser
	case RoleSystem:
		return v2acl.RoleSystem
	case RoleOthers:
		return v2acl.RoleOthers
	default:
		return v2acl.RoleUnknown
	}
}

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

func (m Match) ToV2() v2acl.MatchType {
	switch m {
	case MatchStringEqual:
		return v2acl.MatchTypeStringEqual
	case MatchStringNotEqual:
		return v2acl.MatchTypeStringNotEqual
	default:
		return v2acl.MatchTypeUnknown
	}
}

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

func (h FilterHeaderType) ToV2() v2acl.HeaderType {
	switch h {
	case HeaderFromRequest:
		return v2acl.HeaderTypeRequest
	case HeaderFromObject:
		return v2acl.HeaderTypeObject
	default:
		return v2acl.HeaderTypeUnknown
	}
}

func FilterHeaderTypeFromV2(header v2acl.HeaderType) (h FilterHeaderType) {
	switch header {
	case v2acl.HeaderTypeRequest:
		h = HeaderFromRequest
	case v2acl.HeaderTypeObject:
		h = HeaderFromObject
	default:
		h = HeaderTypeUnknown
	}

	return h
}
