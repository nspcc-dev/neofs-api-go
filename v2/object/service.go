package object

import (
	"context"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type Service interface {
	Get(context.Context, *GetRequest) (GetObjectStreamer, error)
	Put(context.Context) (PutObjectStreamer, error)
	Head(context.Context, *HeadRequest) (*HeadResponse, error)
	Search(context.Context, *SearchRequest) (SearchObjectStreamer, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetRange(context.Context, *GetRangeRequest) (GetRangeObjectStreamer, error)
	GetRangeHash(context.Context, *GetRangeHashRequest) (*GetRangeHashResponse, error)
}

type GetRequestBody struct {
	addr *refs.Address

	raw bool
}

type GetRequest struct {
	body *GetRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetObjectPart interface {
	getObjectPart()
}

type GetObjectPartInit struct {
	id *refs.ObjectID

	sig *service.Signature

	hdr *Header
}

type GetObjectPartChunk struct {
	chunk []byte
}

type GetResponseBody struct {
	objPart GetObjectPart
}

type GetResponse struct {
	body *GetResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type PutObjectPart interface {
	putObjectPart()
}

type PutObjectPartInit struct {
	id *refs.ObjectID

	sig *service.Signature

	hdr *Header

	copyNum uint32
}

type PutObjectPartChunk struct {
	chunk []byte
}

type PutRequestBody struct {
	objPart PutObjectPart
}

type PutRequest struct {
	body *PutRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type PutResponseBody struct {
	id *refs.ObjectID
}

type PutResponse struct {
	body *PutResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type DeleteRequestBody struct {
	addr *refs.Address
}

type DeleteRequest struct {
	body *DeleteRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type DeleteResponseBody struct{}

type DeleteResponse struct {
	body *DeleteResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type HeadRequestBody struct {
	addr *refs.Address

	mainOnly, raw bool
}

type HeadRequest struct {
	body *HeadRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetHeaderPart interface {
	getHeaderPart()
}

type GetHeaderPartFull struct {
	hdr *Header
}

type GetHeaderPartShort struct {
	hdr *ShortHeader
}

type HeadResponseBody struct {
	hdrPart GetHeaderPart
}

type HeadResponse struct {
	body *HeadResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type SearchFilter struct {
	matchType MatchType

	name, val string
}

type SearchRequestBody struct {
	cid *refs.ContainerID

	version uint32

	filters []*SearchFilter
}

type SearchRequest struct {
	body *SearchRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type SearchResponseBody struct {
	idList []*refs.ObjectID
}

type SearchResponse struct {
	body *SearchResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type Range struct {
	off, len uint64
}

type GetRangeRequestBody struct {
	addr *refs.Address

	rng *Range
}

type GetRangeRequest struct {
	body *GetRangeRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetRangeResponseBody struct {
	chunk []byte
}

type GetRangeResponse struct {
	body *GetRangeResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

type GetRangeHashRequestBody struct {
	addr *refs.Address

	rngs []*Range

	salt []byte
}

type GetRangeHashRequest struct {
	body *GetRangeHashRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

type GetRangeHashResponseBody struct {
	hashList [][]byte
}

type GetRangeHashResponse struct {
	body *GetRangeHashResponseBody

	metaHeader *service.ResponseMetaHeader

	verifyHeader *service.ResponseVerificationHeader
}

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

func (r *GetRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
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

func (r *GetObjectPartInit) GetSignature() *service.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *GetObjectPartInit) SetSignature(v *service.Signature) {
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

func (r *GetResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (r *PutObjectPartInit) GetSignature() *service.Signature {
	if r != nil {
		return r.sig
	}

	return nil
}

func (r *PutObjectPartInit) SetSignature(v *service.Signature) {
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

func (r *PutRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
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

func (r *PutResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *PutResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *PutResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *PutResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (r *DeleteRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
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

func (r *DeleteResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *DeleteResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *DeleteResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *DeleteResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (r *HeadRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *HeadRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *HeadRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *HeadRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (h *GetHeaderPartFull) GetHeader() *Header {
	if h != nil {
		return h.hdr
	}

	return nil
}

func (h *GetHeaderPartFull) SetHeader(v *Header) {
	if h != nil {
		h.hdr = v
	}
}

func (*GetHeaderPartFull) getHeaderPart() {}

func (h *GetHeaderPartShort) GetShortHeader() *ShortHeader {
	if h != nil {
		return h.hdr
	}

	return nil
}

func (h *GetHeaderPartShort) SetShortHeader(v *ShortHeader) {
	if h != nil {
		h.hdr = v
	}
}

func (*GetHeaderPartShort) getHeaderPart() {}

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

func (r *HeadResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *HeadResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *HeadResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *HeadResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (f *SearchFilter) GetName() string {
	if f != nil {
		return f.name
	}

	return ""
}

func (f *SearchFilter) SetName(v string) {
	if f != nil {
		f.name = v
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

func (r *SearchRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SearchRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SearchRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SearchRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
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

func (r *SearchResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *SearchResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *SearchResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *SearchResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (r *GetRangeRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}

func (r *GetRangeResponseBody) GetChunk() []byte {
	if r != nil {
		return r.chunk
	}

	return nil
}

func (r *GetRangeResponseBody) SetChunk(v []byte) {
	if r != nil {
		r.chunk = v
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

func (r *GetRangeResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
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

func (r *GetRangeHashRequest) GetMetaHeader() *service.RequestMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeHashRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeHashRequest) GetVerificationHeader() *service.RequestVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeHashRequest) SetVerificationHeader(v *service.RequestVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
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

func (r *GetRangeHashResponse) GetMetaHeader() *service.ResponseMetaHeader {
	if r != nil {
		return r.metaHeader
	}

	return nil
}

func (r *GetRangeHashResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if r != nil {
		r.metaHeader = v
	}
}

func (r *GetRangeHashResponse) GetVerificationHeader() *service.ResponseVerificationHeader {
	if r != nil {
		return r.verifyHeader
	}

	return nil
}

func (r *GetRangeHashResponse) SetVerificationHeader(v *service.ResponseVerificationHeader) {
	if r != nil {
		r.verifyHeader = v
	}
}
