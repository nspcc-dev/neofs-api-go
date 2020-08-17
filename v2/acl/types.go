package acl

import "github.com/nspcc-dev/neofs-api-go/v2/refs"

type HeaderFilter struct {
	hdrType HeaderType

	matchType MatchType

	name, value string
}

type TargetInfo struct {
	target Target

	keys [][]byte
}

type Record struct {
	op Operation

	action Action

	filters []*HeaderFilter

	targets []*TargetInfo
}

type Table struct {
	cid *refs.ContainerID

	records []*Record
}

type MatchType uint32

type HeaderType uint32

type Action uint32

type Operation uint32

type Target uint32

const (
	MatchTypeUnknown MatchType = iota
	MatchTypeStringEqual
	MatchTypeStringNotEqual
)

const (
	HeaderTypeUnknown HeaderType = iota
	HeaderTypeRequest
	HeaderTypeObject
)

const (
	ActionUnknown Action = iota
	ActionDeny
	ActionAllow
)

const (
	OperationUnknown Operation = iota
	OperationGet
	OperationPut
	OperationHead
	OperationDelete
	OperationSearch
	OperationRange
	OperationRangeHash
)

const (
	TargetUnknown Target = iota
	TargetUser
	TargetSystem
	TargetOthers
)

func (f *HeaderFilter) GetHeaderType() HeaderType {
	if f != nil {
		return f.hdrType
	}

	return HeaderTypeUnknown
}

func (f *HeaderFilter) SetHeaderType(v HeaderType) {
	if f != nil {
		f.hdrType = v
	}
}

func (f *HeaderFilter) GetMatchType() MatchType {
	if f != nil {
		return f.matchType
	}

	return MatchTypeUnknown
}

func (f *HeaderFilter) SetMatchType(v MatchType) {
	if f != nil {
		f.matchType = v
	}
}

func (f *HeaderFilter) GetName() string {
	if f != nil {
		return f.name
	}

	return ""
}

func (f *HeaderFilter) SetName(v string) {
	if f != nil {
		f.name = v
	}
}

func (f *HeaderFilter) GetValue() string {
	if f != nil {
		return f.value
	}

	return ""
}

func (f *HeaderFilter) SetValue(v string) {
	if f != nil {
		f.value = v
	}
}

func (t *TargetInfo) GetTarget() Target {
	if t != nil {
		return t.target
	}

	return TargetUnknown
}

func (t *TargetInfo) SetTarget(v Target) {
	if t != nil {
		t.target = v
	}
}

func (t *TargetInfo) GetKeyList() [][]byte {
	if t != nil {
		return t.keys
	}

	return nil
}

func (t *TargetInfo) SetKeyList(v [][]byte) {
	if t != nil {
		t.keys = v
	}
}

func (r *Record) GetOperation() Operation {
	if r != nil {
		return r.op
	}

	return OperationUnknown
}

func (r *Record) SetOperation(v Operation) {
	if r != nil {
		r.op = v
	}
}

func (r *Record) GetAction() Action {
	if r != nil {
		return r.action
	}

	return ActionUnknown
}

func (r *Record) SetAction(v Action) {
	if r != nil {
		r.action = v
	}
}

func (r *Record) GetFilters() []*HeaderFilter {
	if r != nil {
		return r.filters
	}

	return nil
}

func (r *Record) SetFilters(v []*HeaderFilter) {
	if r != nil {
		r.filters = v
	}
}

func (r *Record) GetTargets() []*TargetInfo {
	if r != nil {
		return r.targets
	}

	return nil
}

func (r *Record) SetTargets(v []*TargetInfo) {
	if r != nil {
		r.targets = v
	}
}

func (t *Table) GetContainerID() *refs.ContainerID {
	if t != nil {
		return t.cid
	}

	return nil
}

func (t *Table) SetContainerID(v *refs.ContainerID) {
	if t != nil {
		t.cid = v
	}
}

func (t *Table) GetRecords() []*Record {
	if t != nil {
		return t.records
	}

	return nil
}

func (t *Table) SetRecords(v []*Record) {
	if t != nil {
		t.records = v
	}
}
