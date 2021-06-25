package eacl

import (
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/pkg"
	cid "github.com/nspcc-dev/neofs-api-go/pkg/container/id"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
)

// Record of the EACL rule, that defines EACL action, targets for this action,
// object service operation and filters for request headers.
//
// Record is compatible with v2 acl.EACLRecord message.
type Record struct {
	action    Action
	operation Operation
	filters   []*Filter
	targets   []*Target
}

// Targets returns list of target subjects to apply ACL rule to.
func (r Record) Targets() []*Target {
	return r.targets
}

// SetTargets sets list of target subjects to apply ACL rule to.
func (r *Record) SetTargets(targets ...*Target) {
	r.targets = targets
}

// Filters returns list of filters to match and see if rule is applicable.
func (r Record) Filters() []*Filter {
	return r.filters
}

// Operation returns NeoFS request verb to match.
func (r Record) Operation() Operation {
	return r.operation
}

// SetOperation sets NeoFS request verb to match.
func (r *Record) SetOperation(operation Operation) {
	r.operation = operation
}

// Action returns rule execution result.
func (r Record) Action() Action {
	return r.action
}

// SetAction sets rule execution result.
func (r *Record) SetAction(action Action) {
	r.action = action
}

// AddRecordTarget adds single Target to the Record.
func AddRecordTarget(r *Record, t *Target) {
	r.SetTargets(append(r.Targets(), t)...)
}

func (r *Record) addFilter(from FilterHeaderType, m Match, keyTyp filterKeyType, key string, val fmt.Stringer) {
	filter := &Filter{
		from: from,
		key: filterKey{
			typ: keyTyp,
			str: key,
		},
		matcher: m,
		value:   val,
	}

	r.filters = append(r.filters, filter)
}

func (r *Record) addObjectFilter(m Match, keyTyp filterKeyType, key string, val fmt.Stringer) {
	r.addFilter(HeaderFromObject, m, keyTyp, key, val)
}

func (r *Record) addObjectReservedFilter(m Match, typ filterKeyType, val fmt.Stringer) {
	r.addObjectFilter(m, typ, "", val)
}

// AddFilter adds generic filter.
func (r *Record) AddFilter(from FilterHeaderType, matcher Match, name, value string) {
	r.addFilter(from, matcher, 0, name, staticStringer(value))
}

// AddObjectAttributeFilter adds filter by object attribute.
func (r *Record) AddObjectAttributeFilter(m Match, key, value string) {
	r.addObjectFilter(m, 0, key, staticStringer(value))
}

// AddObjectVersionFilter adds filter by object version.
func (r *Record) AddObjectVersionFilter(m Match, v *pkg.Version) {
	r.addObjectReservedFilter(m, fKeyObjVersion, v)
}

// AddObjectContainerIDFilter adds filter by object container ID.
func (r *Record) AddObjectContainerIDFilter(m Match, id *cid.ID) {
	r.addObjectReservedFilter(m, fKeyObjContainerID, id)
}

// AddObjectOwnerIDFilter adds filter by object owner ID.
func (r *Record) AddObjectOwnerIDFilter(m Match, id *owner.ID) {
	r.addObjectReservedFilter(m, fKeyObjOwnerID, id)
}

// TODO: add remaining filters after neofs-api#72

// ToV2 converts Record to v2 acl.EACLRecord message.
//
// Nil Record converts to nil.
func (r *Record) ToV2() *v2acl.Record {
	if r == nil {
		return nil
	}

	v2 := new(v2acl.Record)

	if r.targets != nil {
		targets := make([]*v2acl.Target, 0, len(r.targets))
		for _, target := range r.targets {
			targets = append(targets, target.ToV2())
		}

		v2.SetTargets(targets)
	}

	if r.filters != nil {
		filters := make([]*v2acl.HeaderFilter, 0, len(r.filters))
		for _, filter := range r.filters {
			filters = append(filters, filter.ToV2())
		}

		v2.SetFilters(filters)
	}

	v2.SetAction(r.action.ToV2())
	v2.SetOperation(r.operation.ToV2())

	return v2
}

// NewRecord creates and returns blank Record instance.
//
// Defaults:
//  - action: ActionUnknown;
//  - operation: OperationUnknown;
//  - targets: nil,
//  - filters: nil.
func NewRecord() *Record {
	return new(Record)
}

// CreateRecord creates, initializes with parameters and returns Record instance.
func CreateRecord(action Action, operation Operation) *Record {
	r := NewRecord()
	r.action = action
	r.operation = operation
	r.targets = []*Target{}
	r.filters = []*Filter{}

	return r
}

// NewRecordFromV2 converts v2 acl.EACLRecord message to Record.
func NewRecordFromV2(record *v2acl.Record) *Record {
	r := NewRecord()

	if record == nil {
		return r
	}

	r.action = ActionFromV2(record.GetAction())
	r.operation = OperationFromV2(record.GetOperation())

	v2targets := record.GetTargets()
	v2filters := record.GetFilters()

	if v2targets != nil {
		r.targets = make([]*Target, 0, len(v2targets))
		for i := range v2targets {
			r.targets = append(r.targets, NewTargetFromV2(v2targets[i]))
		}
	}

	if v2filters != nil {
		r.filters = make([]*Filter, 0, len(v2filters))
		for i := range v2filters {
			r.filters = append(r.filters, NewFilterFromV2(v2filters[i]))
		}
	}

	return r
}

// Marshal marshals Record into a protobuf binary form.
//
// Buffer is allocated when the argument is empty.
// Otherwise, the first buffer is used.
func (r *Record) Marshal(b ...[]byte) ([]byte, error) {
	var buf []byte
	if len(b) > 0 {
		buf = b[0]
	}

	return r.ToV2().
		StableMarshal(buf)
}

// Unmarshal unmarshals protobuf binary representation of Record.
func (r *Record) Unmarshal(data []byte) error {
	fV2 := new(v2acl.Record)
	if err := fV2.Unmarshal(data); err != nil {
		return err
	}

	*r = *NewRecordFromV2(fV2)

	return nil
}

// MarshalJSON encodes Record to protobuf JSON format.
func (r *Record) MarshalJSON() ([]byte, error) {
	return r.ToV2().
		MarshalJSON()
}

// UnmarshalJSON decodes Record from protobuf JSON format.
func (r *Record) UnmarshalJSON(data []byte) error {
	tV2 := new(v2acl.Record)
	if err := tV2.UnmarshalJSON(data); err != nil {
		return err
	}

	*r = *NewRecordFromV2(tV2)

	return nil
}
