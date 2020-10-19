package acl

import "github.com/nspcc-dev/neofs-api-go/v2/refs"

// HeaderFilter is a unified structure of FilterInfo
// message from proto definition.
type HeaderFilter struct {
	hdrType HeaderType

	matchType MatchType

	name, value string
}

// Target is a unified structure of Target
// message from proto definition.
type Target struct {
	role Role

	keys [][]byte
}

// Record is a unified structure of EACLRecord
// message from proto definition.
type Record struct {
	op Operation

	action Action

	filters []*HeaderFilter

	targets []*Target
}

// Table is a unified structure of EACLTable
// message from proto definition.
type Table struct {
	version *refs.Version

	cid *refs.ContainerID

	records []*Record
}

type TokenLifetime struct {
	exp, nbf, iat uint64
}

type BearerTokenBody struct {
	eacl *Table

	ownerID *refs.OwnerID

	lifetime *TokenLifetime
}

type BearerToken struct {
	body *BearerTokenBody

	sig *refs.Signature
}

// Target is a unified enum of MatchType enum from proto definition.
type MatchType uint32

// HeaderType is a unified enum of HeaderType enum from proto definition.
type HeaderType uint32

// Action is a unified enum of Action enum from proto definition.
type Action uint32

// Operation is a unified enum of Operation enum from proto definition.
type Operation uint32

// Role is a unified enum of Role enum from proto definition.
type Role uint32

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
	ActionAllow
	ActionDeny
)

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

const (
	RoleUnknown Role = iota
	RoleUser
	RoleSystem
	RoleOthers
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

func (t *Target) GetRole() Role {
	if t != nil {
		return t.role
	}

	return RoleUnknown
}

func (t *Target) SetRole(v Role) {
	if t != nil {
		t.role = v
	}
}

func (t *Target) GetKeys() [][]byte {
	if t != nil {
		return t.keys
	}

	return nil
}

func (t *Target) SetKeys(v [][]byte) {
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

func (r *Record) GetTargets() []*Target {
	if r != nil {
		return r.targets
	}

	return nil
}

func (r *Record) SetTargets(v []*Target) {
	if r != nil {
		r.targets = v
	}
}

func (t *Table) GetVersion() *refs.Version {
	if t != nil {
		return t.version
	}

	return nil
}

func (t *Table) SetVersion(v *refs.Version) {
	if t != nil {
		t.version = v
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

func (l *TokenLifetime) GetExp() uint64 {
	if l != nil {
		return l.exp
	}

	return 0
}

func (l *TokenLifetime) SetExp(v uint64) {
	if l != nil {
		l.exp = v
	}
}

func (l *TokenLifetime) GetNbf() uint64 {
	if l != nil {
		return l.nbf
	}

	return 0
}

func (l *TokenLifetime) SetNbf(v uint64) {
	if l != nil {
		l.nbf = v
	}
}

func (l *TokenLifetime) GetIat() uint64 {
	if l != nil {
		return l.iat
	}

	return 0
}

func (l *TokenLifetime) SetIat(v uint64) {
	if l != nil {
		l.iat = v
	}
}

func (bt *BearerTokenBody) GetEACL() *Table {
	if bt != nil {
		return bt.eacl
	}

	return nil
}

func (bt *BearerTokenBody) SetEACL(v *Table) {
	if bt != nil {
		bt.eacl = v
	}
}

func (bt *BearerTokenBody) GetOwnerID() *refs.OwnerID {
	if bt != nil {
		return bt.ownerID
	}

	return nil
}

func (bt *BearerTokenBody) SetOwnerID(v *refs.OwnerID) {
	if bt != nil {
		bt.ownerID = v
	}
}

func (bt *BearerTokenBody) GetLifetime() *TokenLifetime {
	if bt != nil {
		return bt.lifetime
	}

	return nil
}

func (bt *BearerTokenBody) SetLifetime(v *TokenLifetime) {
	if bt != nil {
		bt.lifetime = v
	}
}

func (bt *BearerToken) GetBody() *BearerTokenBody {
	if bt != nil {
		return bt.body
	}

	return nil
}

func (bt *BearerToken) SetBody(v *BearerTokenBody) {
	if bt != nil {
		bt.body = v
	}
}

func (bt *BearerToken) GetSignature() *refs.Signature {
	if bt != nil {
		return bt.sig
	}

	return nil
}

func (bt *BearerToken) SetSignature(v *refs.Signature) {
	if bt != nil {
		bt.sig = v
	}
}
