package eacl

import (
	"crypto/ecdsa"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

type (
	// Record of the EACL rule, that defines EACL action, targets for this action,
	// object service operation and filters for request headers.
	Record struct {
		action    Action
		operation Operation
		filters   []Filter
		targets   []Target
	}
)

func (r Record) Targets() []Target {
	return r.targets
}

func (r Record) Filters() []Filter {
	return r.filters
}

func (r Record) Operation() Operation {
	return r.operation
}

func (r *Record) SetOperation(operation Operation) {
	r.operation = operation
}

func (r Record) Action() Action {
	return r.action
}

func (r *Record) SetAction(action Action) {
	r.action = action
}

func (r *Record) AddTarget(role Role, keys ...ecdsa.PublicKey) {
	t := Target{
		role: role,
		keys: make([]ecdsa.PublicKey, 0, len(keys)),
	}

	for i := range keys {
		t.keys = append(t.keys, keys[i])
	}

	r.targets = append(r.targets, t)
}

func (r *Record) AddFilter(from FilterHeaderType, matcher Match, name, value string) {
	filter := Filter{
		from:    from,
		key:     name,
		matcher: matcher,
		value:   value,
	}

	r.filters = append(r.filters, filter)
}

func (r *Record) ToV2() *v2acl.Record {
	targets := make([]*v2acl.Target, 0, len(r.targets))
	for _, target := range r.targets {
		targets = append(targets, target.ToV2())
	}

	filters := make([]*v2acl.HeaderFilter, 0, len(r.filters))
	for _, filter := range r.filters {
		filters = append(filters, filter.ToV2())
	}

	v2 := new(v2acl.Record)

	v2.SetAction(r.action.ToV2())
	v2.SetOperation(r.operation.ToV2())
	v2.SetTargets(targets)
	v2.SetFilters(filters)

	return v2
}

func NewRecord() *Record {
	return new(Record)
}

func CreateRecord(action Action, operation Operation) *Record {
	r := NewRecord()
	r.action = action
	r.operation = operation
	r.targets = []Target{}
	r.filters = []Filter{}

	return r
}

func NewRecordFromV2(record *v2acl.Record) *Record {
	r := NewRecord()

	if record == nil {
		return r
	}

	r.action = ActionFromV2(record.GetAction())
	r.operation = OperationFromV2(record.GetOperation())

	v2targets := record.GetTargets()
	v2filters := record.GetFilters()

	r.targets = make([]Target, 0, len(v2targets))
	for i := range v2targets {
		r.targets = append(r.targets, *NewTargetFromV2(v2targets[i]))
	}

	r.filters = make([]Filter, 0, len(v2filters))
	for i := range v2filters {
		r.filters = append(r.filters, *NewFilterFromV2(v2filters[i]))
	}

	return r
}
