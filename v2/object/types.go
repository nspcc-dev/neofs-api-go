package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type Type uint32

type MatchType uint32

type ShortHeader struct {
	version *service.Version

	creatEpoch uint64

	ownerID *refs.OwnerID

	typ Type

	payloadLen uint64
}

type Attribute struct {
	key, val string
}

type SplitHeader struct {
	par, prev *refs.ObjectID

	parSig *service.Signature

	parHdr *Header

	children []*refs.ObjectID
}

type Header struct {
	version *service.Version

	cid *refs.ContainerID

	ownerID *refs.OwnerID

	creatEpoch uint64

	payloadLen uint64

	payloadHash, homoHash []byte

	typ Type

	sessionToken *service.SessionToken

	attr []*Attribute

	split *SplitHeader
}

type Object struct {
	objectID *refs.ObjectID

	idSig *service.Signature

	header *Header

	payload []byte
}

const (
	TypeRegular Type = iota
	TypeTombstone
	TypeStorageGroup
)

const (
	MatchUnknown MatchType = iota
	MatchStringEqual
)

func (h *ShortHeader) GetVersion() *service.Version {
	if h != nil {
		return h.version
	}

	return nil
}

func (h *ShortHeader) SetVersion(v *service.Version) {
	if h != nil {
		h.SetVersion(v)
	}
}

func (h *ShortHeader) GetCreationEpoch() uint64 {
	if h != nil {
		return h.creatEpoch
	}

	return 0
}

func (h *ShortHeader) SetCreationEpoch(v uint64) {
	if h != nil {
		h.creatEpoch = v
	}
}

func (h *ShortHeader) GetOwnerID() *refs.OwnerID {
	if h != nil {
		return h.ownerID
	}

	return nil
}

func (h *ShortHeader) SetOwnerID(v *refs.OwnerID) {
	if h != nil {
		h.ownerID = v
	}
}

func (h *ShortHeader) GetObjectType() Type {
	if h != nil {
		return h.typ
	}

	return TypeRegular
}

func (h *ShortHeader) SetObjectType(v Type) {
	if h != nil {
		h.typ = v
	}
}

func (h *ShortHeader) GeyPayloadLength() uint64 {
	if h != nil {
		return h.payloadLen
	}

	return 0
}

func (h *ShortHeader) SetPayloadLength(v uint64) {
	if h != nil {
		h.payloadLen = v
	}
}

func (a *Attribute) GetKey() string {
	if a != nil {
		return a.key
	}

	return ""
}

func (a *Attribute) SetKey(v string) {
	if a != nil {
		a.key = v
	}
}

func (a *Attribute) GetValue() string {
	if a != nil {
		return a.val
	}

	return ""
}

func (a *Attribute) SetValue(v string) {
	if a != nil {
		a.val = v
	}
}

func (h *SplitHeader) GetParent() *refs.ObjectID {
	if h != nil {
		return h.par
	}

	return nil
}

func (h *SplitHeader) SetParent(v *refs.ObjectID) {
	if h != nil {
		h.par = v
	}
}

func (h *SplitHeader) GetPrevious() *refs.ObjectID {
	if h != nil {
		return h.prev
	}

	return nil
}

func (h *SplitHeader) SetPrevious(v *refs.ObjectID) {
	if h != nil {
		h.prev = v
	}
}

func (h *SplitHeader) GetParentSignature() *service.Signature {
	if h != nil {
		return h.parSig
	}

	return nil
}

func (h *SplitHeader) SetParentSignature(v *service.Signature) {
	if h != nil {
		h.parSig = v
	}
}

func (h *SplitHeader) GetParentHeader() *Header {
	if h != nil {
		return h.parHdr
	}

	return nil
}

func (h *SplitHeader) SetParentHeader(v *Header) {
	if h != nil {
		h.parHdr = v
	}
}

func (h *SplitHeader) GetChildren() []*refs.ObjectID {
	if h != nil {
		return h.children
	}

	return nil
}

func (h *SplitHeader) SetChildren(v []*refs.ObjectID) {
	if h != nil {
		h.children = v
	}
}

func (h *Header) GetVersion() *service.Version {
	if h != nil {
		return h.version
	}

	return nil
}

func (h *Header) SetVersion(v *service.Version) {
	if h != nil {
		h.version = v
	}
}

func (h *Header) GetContainerID() *refs.ContainerID {
	if h != nil {
		return h.cid
	}

	return nil
}

func (h *Header) SetContainerID(v *refs.ContainerID) {
	if h != nil {
		h.cid = v
	}
}

func (h *Header) GetOwnerID() *refs.OwnerID {
	if h != nil {
		return h.ownerID
	}

	return nil
}

func (h *Header) SetOwnerID(v *refs.OwnerID) {
	if h != nil {
		h.ownerID = v
	}
}

func (h *Header) GetCreationEpoch() uint64 {
	if h != nil {
		return h.creatEpoch
	}

	return 0
}

func (h *Header) SetCreationEpoch(v uint64) {
	if h != nil {
		h.creatEpoch = v
	}
}

func (h *Header) GetPayloadLength() uint64 {
	if h != nil {
		return h.payloadLen
	}

	return 0
}

func (h *Header) SetPayloadLength(v uint64) {
	if h != nil {
		h.payloadLen = v
	}
}

func (h *Header) GetPayloadHash() []byte {
	if h != nil {
		return h.payloadHash
	}

	return nil
}

func (h *Header) SetPayloadHash(v []byte) {
	if h != nil {
		h.payloadHash = v
	}
}

func (h *Header) GetObjectType() Type {
	if h != nil {
		return h.typ
	}

	return TypeRegular
}

func (h *Header) SetObjectType(v Type) {
	if h != nil {
		h.typ = v
	}
}

func (h *Header) GetHomomorphicHash() []byte {
	if h != nil {
		return h.homoHash
	}

	return nil
}

func (h *Header) SetHomomorphicHash(v []byte) {
	if h != nil {
		h.homoHash = v
	}
}

func (h *Header) GetSessionToken() *service.SessionToken {
	if h != nil {
		return h.sessionToken
	}

	return nil
}

func (h *Header) SetSessionToken(v *service.SessionToken) {
	if h != nil {
		h.sessionToken = v
	}
}

func (h *Header) GetAttributes() []*Attribute {
	if h != nil {
		return h.attr
	}

	return nil
}

func (h *Header) SetAttributes(v []*Attribute) {
	if h != nil {
		h.attr = v
	}
}

func (h *Header) GetSplit() *SplitHeader {
	if h != nil {
		return h.split
	}

	return nil
}

func (h *Header) SetSplit(v *SplitHeader) {
	if h != nil {
		h.split = v
	}
}

func (o *Object) GetObjectID() *refs.ObjectID {
	if o != nil {
		return o.objectID
	}

	return nil
}

func (o *Object) SetObjectID(v *refs.ObjectID) {
	if o != nil {
		o.objectID = v
	}
}

func (o *Object) GetSignature() *service.Signature {
	if o != nil {
		return o.idSig
	}

	return nil
}

func (o *Object) SetSignature(v *service.Signature) {
	if o != nil {
		o.idSig = v
	}
}

func (o *Object) GetHeader() *Header {
	if o != nil {
		return o.header
	}

	return nil
}

func (o *Object) SetHeader(v *Header) {
	if o != nil {
		o.header = v
	}
}

func (o *Object) GetPayload() []byte {
	if o != nil {
		return o.payload
	}

	return nil
}

func (o *Object) SetPayload(v []byte) {
	if o != nil {
		o.payload = v
	}
}
