package object

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
)

type Type uint32

type MatchType uint32

type ShortHeader struct {
	version *refs.Version

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

	parSig *refs.Signature

	parHdr *Header

	children []*refs.ObjectID

	splitID []byte
}

type Header struct {
	version *refs.Version

	cid *refs.ContainerID

	ownerID *refs.OwnerID

	creatEpoch uint64

	payloadLen uint64

	payloadHash, homoHash *refs.Checksum

	typ Type

	sessionToken *session.SessionToken

	attr []*Attribute

	split *SplitHeader
}

type HeaderWithSignature struct {
	header *Header

	signature *refs.Signature
}

type Object struct {
	objectID *refs.ObjectID

	idSig *refs.Signature

	header *Header

	payload []byte
}

type SplitInfo struct {
	splitID []byte

	lastPart *refs.ObjectID

	link *refs.ObjectID
}

type GetRequestBody struct {
	addr *refs.Address

	raw bool
}

type GetObjectPart interface {
	getObjectPart()
}

type GetObjectPartInit struct {
	id *refs.ObjectID

	sig *refs.Signature

	hdr *Header
}

type GetObjectPartChunk struct {
	chunk []byte
}

type GetResponseBody struct {
	objPart GetObjectPart
}

type PutObjectPart interface {
	putObjectPart()
}

type PutObjectPartInit struct {
	id *refs.ObjectID

	sig *refs.Signature

	hdr *Header

	copyNum uint32
}

type PutObjectPartChunk struct {
	chunk []byte
}

type PutRequestBody struct {
	objPart PutObjectPart
}

type PutResponseBody struct {
	id *refs.ObjectID
}

type DeleteRequestBody struct {
	addr *refs.Address
}

type DeleteResponseBody struct{}

type HeadRequestBody struct {
	addr *refs.Address

	mainOnly, raw bool
}

type GetHeaderPart interface {
	getHeaderPart()
}

type HeadResponseBody struct {
	hdrPart GetHeaderPart
}

type SearchFilter struct {
	matchType MatchType

	key, val string
}

type SearchRequestBody struct {
	cid *refs.ContainerID

	version uint32

	filters []*SearchFilter
}

type SearchResponseBody struct {
	idList []*refs.ObjectID
}

type Range struct {
	off, len uint64
}

type GetRangeRequestBody struct {
	addr *refs.Address

	rng *Range

	raw bool
}

type GetRangePart interface {
	getRangePart()
}

type GetRangePartChunk struct {
	chunk []byte
}

type GetRangeResponseBody struct {
	rngPart GetRangePart
}

type GetRangeHashRequestBody struct {
	addr *refs.Address

	rngs []*Range

	salt []byte

	typ refs.ChecksumType
}

type GetRangeHashResponseBody struct {
	typ refs.ChecksumType

	hashList [][]byte
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

func (h *ShortHeader) GetVersion() *refs.Version {
	if h != nil {
		return h.version
	}

	return nil
}

func (h *ShortHeader) SetVersion(v *refs.Version) {
	if h != nil {
		h.version = v
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

func (h *ShortHeader) GetPayloadLength() uint64 {
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

func (h *ShortHeader) getHeaderPart() {}

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

func (h *SplitHeader) GetParentSignature() *refs.Signature {
	if h != nil {
		return h.parSig
	}

	return nil
}

func (h *SplitHeader) SetParentSignature(v *refs.Signature) {
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

func (h *SplitHeader) GetSplitID() []byte {
	if h != nil {
		return h.splitID
	}

	return nil
}

func (h *SplitHeader) SetSplitID(v []byte) {
	if h != nil {
		h.splitID = v
	}
}

func (h *Header) GetVersion() *refs.Version {
	if h != nil {
		return h.version
	}

	return nil
}

func (h *Header) SetVersion(v *refs.Version) {
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

func (h *Header) GetPayloadHash() *refs.Checksum {
	if h != nil {
		return h.payloadHash
	}

	return nil
}

func (h *Header) SetPayloadHash(v *refs.Checksum) {
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

func (h *Header) GetHomomorphicHash() *refs.Checksum {
	if h != nil {
		return h.homoHash
	}

	return nil
}

func (h *Header) SetHomomorphicHash(v *refs.Checksum) {
	if h != nil {
		h.homoHash = v
	}
}

func (h *Header) GetSessionToken() *session.SessionToken {
	if h != nil {
		return h.sessionToken
	}

	return nil
}

func (h *Header) SetSessionToken(v *session.SessionToken) {
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

func (h *HeaderWithSignature) GetHeader() *Header {
	if h != nil {
		return h.header
	}

	return nil
}

func (h *HeaderWithSignature) SetHeader(v *Header) {
	if h != nil {
		h.header = v
	}
}

func (h *HeaderWithSignature) GetSignature() *refs.Signature {
	if h != nil {
		return h.signature
	}

	return nil
}

func (h *HeaderWithSignature) SetSignature(v *refs.Signature) {
	if h != nil {
		h.signature = v
	}
}

func (h *HeaderWithSignature) getHeaderPart() {}

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

func (o *Object) GetSignature() *refs.Signature {
	if o != nil {
		return o.idSig
	}

	return nil
}

func (o *Object) SetSignature(v *refs.Signature) {
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

func (s *SplitInfo) GetSplitID() []byte {
	if s.splitID != nil {
		return s.splitID
	}

	return nil
}

func (s *SplitInfo) SetSplitID(v []byte) {
	if s != nil {
		s.splitID = v
	}
}

func (s *SplitInfo) GetLastPart() *refs.ObjectID {
	if s != nil {
		return s.lastPart
	}

	return nil
}

func (s *SplitInfo) SetLastPart(v *refs.ObjectID) {
	if s != nil {
		s.lastPart = v
	}
}

func (s *SplitInfo) GetLink() *refs.ObjectID {
	if s != nil {
		return s.link
	}

	return nil
}

func (s *SplitInfo) SetLink(v *refs.ObjectID) {
	if s != nil {
		s.link = v
	}
}

func (s *SplitInfo) getObjectPart() {}

func (s *SplitInfo) getHeaderPart() {}

func (s *SplitInfo) getRangePart() {}

func (r *GetRequestBody) GetAddress() *refs.Address {
	if r != nil {
		return r.addr
	}

	return nil
}

func (r *GetRequestBody) SetAddress(v *refs.Address) {
	if r != nil {
		r.addr = v
	}
}

func (r *GetRequestBody) GetRaw() bool {
	if r != nil {
		return r.raw
	}

	return false
}

func (r *GetRequestBody) SetRaw(v bool) {
	if r != nil {
		r.raw = v
	}
}

func (r *GetRequest) GetBody() *GetRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRequest) SetBody(v *GetRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetObjectPartInit) GetObjectID() *refs.ObjectID {
	if r != nil {
		return r.id
	}

	return nil
}

func (r *GetObjectPartInit) SetObjectID(v *refs.ObjectID) {
	if r != nil {
		r.id = v
	}
}

func (r *GetObjectPartInit) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *GetObjectPartInit) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *GetObjectPartInit) GetHeader() *Header {
	if r != nil {
		return r.hdr
	}

	return nil
}

func (r *GetObjectPartInit) SetHeader(v *Header) {
	if r != nil {
		r.hdr = v
	}
}

func (r *GetObjectPartInit) getObjectPart() {}

func (r *GetObjectPartChunk) GetChunk() []byte {
	if r != nil {
		return r.chunk
	}

	return nil
}

func (r *GetObjectPartChunk) SetChunk(v []byte) {
	if r != nil {
		r.chunk = v
	}
}

func (r *GetObjectPartChunk) getObjectPart() {}

func (r *GetResponseBody) GetObjectPart() GetObjectPart {
	if r != nil {
		return r.objPart
	}

	return nil
}

func (r *GetResponseBody) SetObjectPart(v GetObjectPart) {
	if r != nil {
		r.objPart = v
	}
}

func (r *GetResponse) GetBody() *GetResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetResponse) SetBody(v *GetResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *PutObjectPartInit) GetObjectID() *refs.ObjectID {
	if r != nil {
		return r.id
	}

	return nil
}

func (r *PutObjectPartInit) SetObjectID(v *refs.ObjectID) {
	if r != nil {
		r.id = v
	}
}

func (r *PutObjectPartInit) GetSignature() *refs.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *PutObjectPartInit) SetSignature(v *refs.Signature) {
	if r != nil {
		r.sig = v
	}
}

func (r *PutObjectPartInit) GetHeader() *Header {
	if r != nil {
		return r.hdr
	}

	return nil
}

func (r *PutObjectPartInit) SetHeader(v *Header) {
	if r != nil {
		r.hdr = v
	}
}

func (r *PutObjectPartInit) GetCopiesNumber() uint32 {
	if r != nil {
		return r.copyNum
	}

	return 0
}

func (r *PutObjectPartInit) SetCopiesNumber(v uint32) {
	if r != nil {
		r.copyNum = v
	}
}

func (r *PutObjectPartInit) putObjectPart() {}

func (r *PutObjectPartChunk) GetChunk() []byte {
	if r != nil {
		return r.chunk
	}

	return nil
}

func (r *PutObjectPartChunk) SetChunk(v []byte) {
	if r != nil {
		r.chunk = v
	}
}

func (r *PutObjectPartChunk) putObjectPart() {}

func (r *PutRequestBody) GetObjectPart() PutObjectPart {
	if r != nil {
		return r.objPart
	}

	return nil
}

func (r *PutRequestBody) SetObjectPart(v PutObjectPart) {
	if r != nil {
		r.objPart = v
	}
}

func (r *PutRequest) GetBody() *PutRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *PutRequest) SetBody(v *PutRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *PutRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *PutResponseBody) GetObjectID() *refs.ObjectID {
	if r != nil {
		return r.id
	}

	return nil
}

func (r *PutResponseBody) SetObjectID(v *refs.ObjectID) {
	if r != nil {
		r.id = v
	}
}

func (r *PutResponse) GetBody() *PutResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *PutResponse) SetBody(v *PutResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *PutResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *DeleteRequestBody) GetAddress() *refs.Address {
	if r != nil {
		return r.addr
	}

	return nil
}

func (r *DeleteRequestBody) SetAddress(v *refs.Address) {
	if r != nil {
		r.addr = v
	}
}

func (r *DeleteRequest) GetBody() *DeleteRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *DeleteRequest) SetBody(v *DeleteRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *DeleteRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *DeleteResponse) GetBody() *DeleteResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *DeleteResponse) SetBody(v *DeleteResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *DeleteResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *HeadRequestBody) GetAddress() *refs.Address {
	if r != nil {
		return r.addr
	}

	return nil
}

func (r *HeadRequestBody) SetAddress(v *refs.Address) {
	if r != nil {
		r.addr = v
	}
}

func (r *HeadRequestBody) GetMainOnly() bool {
	if r != nil {
		return r.mainOnly
	}

	return false
}

func (r *HeadRequestBody) SetMainOnly(v bool) {
	if r != nil {
		r.mainOnly = v
	}
}

func (r *HeadRequestBody) GetRaw() bool {
	if r != nil {
		return r.raw
	}

	return false
}

func (r *HeadRequestBody) SetRaw(v bool) {
	if r != nil {
		r.raw = v
	}
}

func (r *HeadRequest) GetBody() *HeadRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *HeadRequest) SetBody(v *HeadRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *HeadRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *HeadRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *HeadRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *HeadRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *HeadResponseBody) GetHeaderPart() GetHeaderPart {
	if r != nil {
		return r.hdrPart
	}

	return nil
}

func (r *HeadResponseBody) SetHeaderPart(v GetHeaderPart) {
	if r != nil {
		r.hdrPart = v
	}
}

func (r *HeadResponse) GetBody() *HeadResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *HeadResponse) SetBody(v *HeadResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *HeadResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *HeadResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *HeadResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *HeadResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (f *SearchFilter) GetMatchType() MatchType {
	if f != nil {
		return f.matchType
	}

	return MatchUnknown
}

func (f *SearchFilter) SetMatchType(v MatchType) {
	if f != nil {
		f.matchType = v
	}
}

func (f *SearchFilter) GetKey() string {
	if f != nil {
		return f.key
	}

	return ""
}

func (f *SearchFilter) SetKey(v string) {
	if f != nil {
		f.key = v
	}
}

func (f *SearchFilter) GetValue() string {
	if f != nil {
		return f.val
	}

	return ""
}

func (f *SearchFilter) SetValue(v string) {
	if f != nil {
		f.val = v
	}
}

func (r *SearchRequestBody) GetContainerID() *refs.ContainerID {
	if r != nil {
		return r.cid
	}

	return nil
}

func (r *SearchRequestBody) SetContainerID(v *refs.ContainerID) {
	if r != nil {
		r.cid = v
	}
}

func (r *SearchRequestBody) GetVersion() uint32 {
	if r != nil {
		return r.version
	}

	return 0
}

func (r *SearchRequestBody) SetVersion(v uint32) {
	if r != nil {
		r.version = v
	}
}

func (r *SearchRequestBody) GetFilters() []*SearchFilter {
	if r != nil {
		return r.filters
	}

	return nil
}

func (r *SearchRequestBody) SetFilters(v []*SearchFilter) {
	if r != nil {
		r.filters = v
	}
}

func (r *SearchRequest) GetBody() *SearchRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *SearchRequest) SetBody(v *SearchRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *SearchRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SearchRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SearchRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SearchRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *SearchResponseBody) GetIDList() []*refs.ObjectID {
	if r != nil {
		return r.idList
	}

	return nil
}

func (r *SearchResponseBody) SetIDList(v []*refs.ObjectID) {
	if r != nil {
		r.idList = v
	}
}

func (r *SearchResponse) GetBody() *SearchResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *SearchResponse) SetBody(v *SearchResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *SearchResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SearchResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SearchResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SearchResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *Range) GetOffset() uint64 {
	if r != nil {
		return r.off
	}

	return 0
}

func (r *Range) SetOffset(v uint64) {
	if r != nil {
		r.off = v
	}
}

func (r *Range) GetLength() uint64 {
	if r != nil {
		return r.len
	}

	return 0
}

func (r *Range) SetLength(v uint64) {
	if r != nil {
		r.len = v
	}
}

func (r *GetRangeRequestBody) GetAddress() *refs.Address {
	if r != nil {
		return r.addr
	}

	return nil
}

func (r *GetRangeRequestBody) SetAddress(v *refs.Address) {
	if r != nil {
		r.addr = v
	}
}

func (r *GetRangeRequestBody) GetRange() *Range {
	if r != nil {
		return r.rng
	}

	return nil
}

func (r *GetRangeRequestBody) SetRange(v *Range) {
	if r != nil {
		r.rng = v
	}
}

func (r *GetRangeRequestBody) GetRaw() bool {
	if r != nil {
		return r.raw
	}

	return false
}

func (r *GetRangeRequestBody) SetRaw(v bool) {
	if r != nil {
		r.raw = v
	}
}

func (r *GetRangeRequest) GetBody() *GetRangeRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRangeRequest) SetBody(v *GetRangeRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRangeRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetRangePartChunk) GetChunk() []byte {
	if r != nil {
		return r.chunk
	}

	return nil
}

func (r *GetRangePartChunk) SetChunk(v []byte) {
	if r != nil {
		r.chunk = v
	}
}

func (r *GetRangePartChunk) getRangePart() {}

func (r *GetRangeResponseBody) GetRangePart() GetRangePart {
	if r != nil {
		return r.rngPart
	}

	return nil
}

func (r *GetRangeResponseBody) SetRangePart(v GetRangePart) {
	if r != nil {
		r.rngPart = v
	}
}

func (r *GetRangeResponse) GetBody() *GetRangeResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRangeResponse) SetBody(v *GetRangeResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRangeResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetRangeHashRequestBody) GetAddress() *refs.Address {
	if r != nil {
		return r.addr
	}

	return nil
}

func (r *GetRangeHashRequestBody) SetAddress(v *refs.Address) {
	if r != nil {
		r.addr = v
	}
}

func (r *GetRangeHashRequestBody) GetRanges() []*Range {
	if r != nil {
		return r.rngs
	}

	return nil
}

func (r *GetRangeHashRequestBody) SetRanges(v []*Range) {
	if r != nil {
		r.rngs = v
	}
}

func (r *GetRangeHashRequestBody) GetSalt() []byte {
	if r != nil {
		return r.salt
	}

	return nil
}

func (r *GetRangeHashRequestBody) SetSalt(v []byte) {
	if r != nil {
		r.salt = v
	}
}

func (r *GetRangeHashRequestBody) GetType() refs.ChecksumType {
	if r != nil {
		return r.typ
	}

	return refs.UnknownChecksum
}

func (r *GetRangeHashRequestBody) SetType(v refs.ChecksumType) {
	if r != nil {
		r.typ = v
	}
}

func (r *GetRangeHashRequest) GetBody() *GetRangeHashRequestBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRangeHashRequest) SetBody(v *GetRangeHashRequestBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRangeHashRequest) GetMetaHeader() *session.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeHashRequest) SetMetaHeader(v *session.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeHashRequest) GetVerificationHeader() *session.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeHashRequest) SetVerificationHeader(v *session.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetRangeHashResponseBody) GetType() refs.ChecksumType {
	if r != nil {
		return r.typ
	}

	return refs.UnknownChecksum
}

func (r *GetRangeHashResponseBody) SetType(v refs.ChecksumType) {
	if r != nil {
		r.typ = v
	}
}

func (r *GetRangeHashResponseBody) GetHashList() [][]byte {
	if r != nil {
		return r.hashList
	}

	return nil
}

func (r *GetRangeHashResponseBody) SetHashList(v [][]byte) {
	if r != nil {
		r.hashList = v
	}
}

func (r *GetRangeHashResponse) GetBody() *GetRangeHashResponseBody {
	if r != nil {
		return r.body
	}

	return nil
}

func (r *GetRangeHashResponse) SetBody(v *GetRangeHashResponseBody) {
	if r != nil {
		r.body = v
	}
}

func (r *GetRangeHashResponse) GetMetaHeader() *session.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeHashResponse) SetMetaHeader(v *session.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeHashResponse) GetVerificationHeader() *session.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeHashResponse) SetVerificationHeader(v *session.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}
