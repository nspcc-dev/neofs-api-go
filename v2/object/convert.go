package object

import (
	"fmt"

	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	refsGRPC "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

func TypeToGRPCField(t Type) object.ObjectType {
	return object.ObjectType(t)
}

func TypeFromGRPCField(t object.ObjectType) Type {
	return Type(t)
}

func MatchTypeToGRPCField(t MatchType) object.MatchType {
	return object.MatchType(t)
}

func MatchTypeFromGRPCField(t object.MatchType) MatchType {
	return MatchType(t)
}

func ShortHeaderToGRPCMessage(h *ShortHeader) *object.ShortHeader {
	if h == nil {
		return nil
	}

	m := new(object.ShortHeader)

	m.SetVersion(
		service.VersionToGRPCMessage(h.GetVersion()),
	)

	m.SetCreationEpoch(h.GetCreationEpoch())

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(h.GetOwnerID()),
	)

	m.SetObjectType(
		TypeToGRPCField(h.GetObjectType()),
	)

	m.SetPayloadLength(h.GetPayloadLength())

	return m
}

func ShortHeaderFromGRPCMessage(m *object.ShortHeader) *ShortHeader {
	if m == nil {
		return nil
	}

	h := new(ShortHeader)

	h.SetVersion(
		service.VersionFromGRPCMessage(m.GetVersion()),
	)

	h.SetCreationEpoch(m.GetCreationEpoch())

	h.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	h.SetObjectType(
		TypeFromGRPCField(m.GetObjectType()),
	)

	h.SetPayloadLength(m.GetPayloadLength())

	return h
}

func AttributeToGRPCMessage(a *Attribute) *object.Header_Attribute {
	if a == nil {
		return nil
	}

	m := new(object.Header_Attribute)

	m.SetKey(a.GetKey())
	m.SetValue(a.GetValue())

	return m
}

func AttributeFromGRPCMessage(m *object.Header_Attribute) *Attribute {
	if m == nil {
		return nil
	}

	h := new(Attribute)

	h.SetKey(m.GetKey())
	h.SetValue(m.GetValue())

	return h
}

func SplitHeaderToGRPCMessage(h *SplitHeader) *object.Header_Split {
	if h == nil {
		return nil
	}

	m := new(object.Header_Split)

	m.SetParent(
		refs.ObjectIDToGRPCMessage(h.GetParent()),
	)

	m.SetPrevious(
		refs.ObjectIDToGRPCMessage(h.GetPrevious()),
	)

	m.SetParentSignature(
		service.SignatureToGRPCMessage(h.GetParentSignature()),
	)

	m.SetParentHeader(
		HeaderToGRPCMessage(h.GetParentHeader()),
	)

	children := h.GetChildren()
	childMsg := make([]*refsGRPC.ObjectID, 0, len(children))

	for i := range children {
		childMsg = append(childMsg, refs.ObjectIDToGRPCMessage(children[i]))
	}

	m.SetChildren(childMsg)

	return m
}

func SplitHeaderFromGRPCMessage(m *object.Header_Split) *SplitHeader {
	if m == nil {
		return nil
	}

	h := new(SplitHeader)

	h.SetParent(
		refs.ObjectIDFromGRPCMessage(m.GetParent()),
	)

	h.SetPrevious(
		refs.ObjectIDFromGRPCMessage(m.GetPrevious()),
	)

	h.SetParentSignature(
		service.SignatureFromGRPCMessage(m.GetParentSignature()),
	)

	h.SetParentHeader(
		HeaderFromGRPCMessage(m.GetParentHeader()),
	)

	childMsg := m.GetChildren()
	children := make([]*refs.ObjectID, 0, len(childMsg))

	for i := range childMsg {
		children = append(children, refs.ObjectIDFromGRPCMessage(childMsg[i]))
	}

	h.SetChildren(children)

	return h
}

func HeaderToGRPCMessage(h *Header) *object.Header {
	if h == nil {
		return nil
	}

	m := new(object.Header)

	m.SetVersion(
		service.VersionToGRPCMessage(h.GetVersion()),
	)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(h.GetContainerID()),
	)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(h.GetOwnerID()),
	)

	m.SetCreationEpoch(h.GetCreationEpoch())

	m.SetPayloadLength(h.GetPayloadLength())

	m.SetPayloadHash(
		refs.ChecksumToGRPCMessage(h.GetPayloadHash()),
	)

	m.SetHomomorphicHash(
		refs.ChecksumToGRPCMessage(h.GetHomomorphicHash()),
	)

	m.SetObjectType(
		TypeToGRPCField(h.GetObjectType()),
	)

	m.SetSessionToken(
		service.SessionTokenToGRPCMessage(h.GetSessionToken()),
	)

	attr := h.GetAttributes()
	attrMsg := make([]*object.Header_Attribute, 0, len(attr))

	for i := range attr {
		attrMsg = append(attrMsg, AttributeToGRPCMessage(attr[i]))
	}

	m.SetAttributes(attrMsg)

	m.SetSplit(
		SplitHeaderToGRPCMessage(h.GetSplit()),
	)

	return m
}

func HeaderFromGRPCMessage(m *object.Header) *Header {
	if m == nil {
		return nil
	}

	h := new(Header)

	h.SetVersion(
		service.VersionFromGRPCMessage(m.GetVersion()),
	)

	h.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	h.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	h.SetCreationEpoch(m.GetCreationEpoch())

	h.SetPayloadLength(m.GetPayloadLength())

	h.SetPayloadHash(
		refs.ChecksumFromGRPCMessage(m.GetPayloadHash()),
	)

	h.SetHomomorphicHash(
		refs.ChecksumFromGRPCMessage(m.GetHomomorphicHash()),
	)

	h.SetObjectType(
		TypeFromGRPCField(m.GetObjectType()),
	)

	h.SetSessionToken(
		service.SessionTokenFromGRPCMessage(m.GetSessionToken()),
	)

	attrMsg := m.GetAttributes()
	attr := make([]*Attribute, 0, len(attrMsg))

	for i := range attrMsg {
		attr = append(attr, AttributeFromGRPCMessage(attrMsg[i]))
	}

	h.SetAttributes(attr)

	h.SetSplit(
		SplitHeaderFromGRPCMessage(m.GetSplit()),
	)

	return h
}

func ObjectToGRPCMessage(o *Object) *object.Object {
	if o == nil {
		return nil
	}

	m := new(object.Object)

	m.SetObjectId(
		refs.ObjectIDToGRPCMessage(o.GetObjectID()),
	)

	m.SetSignature(
		service.SignatureToGRPCMessage(o.GetSignature()),
	)

	m.SetHeader(
		HeaderToGRPCMessage(o.GetHeader()),
	)

	m.SetPayload(o.GetPayload())

	return m
}

func ObjectFromGRPCMessage(m *object.Object) *Object {
	if m == nil {
		return nil
	}

	o := new(Object)

	o.SetObjectID(
		refs.ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	o.SetSignature(
		service.SignatureFromGRPCMessage(m.GetSignature()),
	)

	o.SetHeader(
		HeaderFromGRPCMessage(m.GetHeader()),
	)

	o.SetPayload(m.GetPayload())

	return o
}

func GetRequestBodyToGRPCMessage(r *GetRequestBody) *object.GetRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetRequest_Body)

	m.SetAddress(
		refs.AddressToGRPCMessage(r.GetAddress()),
	)

	m.SetRaw(r.GetRaw())

	return m
}

func GetRequestBodyFromGRPCMessage(m *object.GetRequest_Body) *GetRequestBody {
	if m == nil {
		return nil
	}

	r := new(GetRequestBody)

	r.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	r.SetRaw(m.GetRaw())

	return r
}

func GetRequestToGRPCMessage(r *GetRequest) *object.GetRequest {
	if r == nil {
		return nil
	}

	m := new(object.GetRequest)

	m.SetBody(
		GetRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func GetRequestFromGRPCMessage(m *object.GetRequest) *GetRequest {
	if m == nil {
		return nil
	}

	r := new(GetRequest)

	r.SetBody(
		GetRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func GetObjectPartInitToGRPCMessage(r *GetObjectPartInit) *object.GetResponse_Body_Init {
	if r == nil {
		return nil
	}

	m := new(object.GetResponse_Body_Init)

	m.SetObjectId(
		refs.ObjectIDToGRPCMessage(r.GetObjectID()),
	)

	m.SetSignature(
		service.SignatureToGRPCMessage(r.GetSignature()),
	)

	m.SetHeader(
		HeaderToGRPCMessage(r.GetHeader()),
	)

	return m
}

func GetObjectPartInitFromGRPCMessage(m *object.GetResponse_Body_Init) *GetObjectPartInit {
	if m == nil {
		return nil
	}

	r := new(GetObjectPartInit)

	r.SetObjectID(
		refs.ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	r.SetSignature(
		service.SignatureFromGRPCMessage(m.GetSignature()),
	)

	r.SetHeader(
		HeaderFromGRPCMessage(m.GetHeader()),
	)

	return r
}

func GetObjectPartChunkToGRPCMessage(r *GetObjectPartChunk) *object.GetResponse_Body_Chunk {
	if r == nil {
		return nil
	}

	m := new(object.GetResponse_Body_Chunk)

	m.SetChunk(r.GetChunk())

	return m
}

func GetObjectPartChunkFromGRPCMessage(m *object.GetResponse_Body_Chunk) *GetObjectPartChunk {
	if m == nil {
		return nil
	}

	r := new(GetObjectPartChunk)

	r.SetChunk(m.GetChunk())

	return r
}

func GetResponseBodyToGRPCMessage(r *GetResponseBody) *object.GetResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetResponse_Body)

	switch v := r.GetObjectPart(); t := v.(type) {
	case nil:
	case *GetObjectPartInit:
		m.SetInit(
			GetObjectPartInitToGRPCMessage(t),
		)
	case *GetObjectPartChunk:
		m.SetChunk(
			GetObjectPartChunkToGRPCMessage(t),
		)
	default:
		panic(fmt.Sprintf("unknown object part %T", t))
	}

	return m
}

func GetResponseBodyFromGRPCMessage(m *object.GetResponse_Body) *GetResponseBody {
	if m == nil {
		return nil
	}

	r := new(GetResponseBody)

	switch v := m.GetObjectPart().(type) {
	case nil:
	case *object.GetResponse_Body_Init_:
		r.SetObjectPart(
			GetObjectPartInitFromGRPCMessage(v.Init),
		)
	case *object.GetResponse_Body_Chunk:
		r.SetObjectPart(
			GetObjectPartChunkFromGRPCMessage(v),
		)
	default:
		panic(fmt.Sprintf("unknown object part %T", v))
	}

	return r
}

func GetResponseToGRPCMessage(r *GetResponse) *object.GetResponse {
	if r == nil {
		return nil
	}

	m := new(object.GetResponse)

	m.SetBody(
		GetResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func GetResponseFromGRPCMessage(m *object.GetResponse) *GetResponse {
	if m == nil {
		return nil
	}

	r := new(GetResponse)

	r.SetBody(
		GetResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func PutObjectPartInitToGRPCMessage(r *PutObjectPartInit) *object.PutRequest_Body_Init {
	if r == nil {
		return nil
	}

	m := new(object.PutRequest_Body_Init)

	m.SetObjectId(
		refs.ObjectIDToGRPCMessage(r.GetObjectID()),
	)

	m.SetSignature(
		service.SignatureToGRPCMessage(r.GetSignature()),
	)

	m.SetHeader(
		HeaderToGRPCMessage(r.GetHeader()),
	)

	m.SetCopiesNumber(r.GetCopiesNumber())

	return m
}

func PutObjectPartInitFromGRPCMessage(m *object.PutRequest_Body_Init) *PutObjectPartInit {
	if m == nil {
		return nil
	}

	r := new(PutObjectPartInit)

	r.SetObjectID(
		refs.ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	r.SetSignature(
		service.SignatureFromGRPCMessage(m.GetSignature()),
	)

	r.SetHeader(
		HeaderFromGRPCMessage(m.GetHeader()),
	)

	r.SetCopiesNumber(m.GetCopiesNumber())

	return r
}

func PutObjectPartChunkToGRPCMessage(r *PutObjectPartChunk) *object.PutRequest_Body_Chunk {
	if r == nil {
		return nil
	}

	m := new(object.PutRequest_Body_Chunk)

	m.SetChunk(r.GetChunk())

	return m
}

func PutObjectPartChunkFromGRPCMessage(m *object.PutRequest_Body_Chunk) *PutObjectPartChunk {
	if m == nil {
		return nil
	}

	r := new(PutObjectPartChunk)

	r.SetChunk(m.GetChunk())

	return r
}

func PutRequestBodyToGRPCMessage(r *PutRequestBody) *object.PutRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.PutRequest_Body)

	switch v := r.GetObjectPart(); t := v.(type) {
	case nil:
	case *PutObjectPartInit:
		m.SetInit(
			PutObjectPartInitToGRPCMessage(t),
		)
	case *PutObjectPartChunk:
		m.SetChunk(
			PutObjectPartChunkToGRPCMessage(t),
		)
	default:
		panic(fmt.Sprintf("unknown object part %T", t))
	}

	return m
}

func PutRequestBodyFromGRPCMessage(m *object.PutRequest_Body) *PutRequestBody {
	if m == nil {
		return nil
	}

	r := new(PutRequestBody)

	switch v := m.GetObjectPart().(type) {
	case nil:
	case *object.PutRequest_Body_Init_:
		r.SetObjectPart(
			PutObjectPartInitFromGRPCMessage(v.Init),
		)
	case *object.PutRequest_Body_Chunk:
		r.SetObjectPart(
			PutObjectPartChunkFromGRPCMessage(v),
		)
	default:
		panic(fmt.Sprintf("unknown object part %T", v))
	}

	return r
}

func PutRequestToGRPCMessage(r *PutRequest) *object.PutRequest {
	if r == nil {
		return nil
	}

	m := new(object.PutRequest)

	m.SetBody(
		PutRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func PutRequestFromGRPCMessage(m *object.PutRequest) *PutRequest {
	if m == nil {
		return nil
	}

	r := new(PutRequest)

	r.SetBody(
		PutRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func PutResponseBodyToGRPCMessage(r *PutResponseBody) *object.PutResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.PutResponse_Body)

	m.SetObjectId(
		refs.ObjectIDToGRPCMessage(r.GetObjectID()),
	)

	return m
}

func PutResponseBodyFromGRPCMessage(m *object.PutResponse_Body) *PutResponseBody {
	if m == nil {
		return nil
	}

	r := new(PutResponseBody)

	r.SetObjectID(
		refs.ObjectIDFromGRPCMessage(m.GetObjectId()),
	)

	return r
}

func PutResponseToGRPCMessage(r *PutResponse) *object.PutResponse {
	if r == nil {
		return nil
	}

	m := new(object.PutResponse)

	m.SetBody(
		PutResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func PutResponseFromGRPCMessage(m *object.PutResponse) *PutResponse {
	if m == nil {
		return nil
	}

	r := new(PutResponse)

	r.SetBody(
		PutResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func DeleteRequestBodyToGRPCMessage(r *DeleteRequestBody) *object.DeleteRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.DeleteRequest_Body)

	m.SetAddress(
		refs.AddressToGRPCMessage(r.GetAddress()),
	)

	return m
}

func DeleteRequestBodyFromGRPCMessage(m *object.DeleteRequest_Body) *DeleteRequestBody {
	if m == nil {
		return nil
	}

	r := new(DeleteRequestBody)

	r.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	return r
}

func DeleteRequestToGRPCMessage(r *DeleteRequest) *object.DeleteRequest {
	if r == nil {
		return nil
	}

	m := new(object.DeleteRequest)

	m.SetBody(
		DeleteRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func DeleteRequestFromGRPCMessage(m *object.DeleteRequest) *DeleteRequest {
	if m == nil {
		return nil
	}

	r := new(DeleteRequest)

	r.SetBody(
		DeleteRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func DeleteResponseBodyToGRPCMessage(r *DeleteResponseBody) *object.DeleteResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.DeleteResponse_Body)

	return m
}

func DeleteResponseBodyFromGRPCMessage(m *object.DeleteResponse_Body) *DeleteResponseBody {
	if m == nil {
		return nil
	}

	r := new(DeleteResponseBody)

	return r
}

func DeleteResponseToGRPCMessage(r *DeleteResponse) *object.DeleteResponse {
	if r == nil {
		return nil
	}

	m := new(object.DeleteResponse)

	m.SetBody(
		DeleteResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func DeleteResponseFromGRPCMessage(m *object.DeleteResponse) *DeleteResponse {
	if m == nil {
		return nil
	}

	r := new(DeleteResponse)

	r.SetBody(
		DeleteResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func HeadRequestBodyToGRPCMessage(r *HeadRequestBody) *object.HeadRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.HeadRequest_Body)

	m.SetAddress(
		refs.AddressToGRPCMessage(r.GetAddress()),
	)

	m.SetMainOnly(r.GetMainOnly())

	m.SetRaw(r.GetRaw())

	return m
}

func HeadRequestBodyFromGRPCMessage(m *object.HeadRequest_Body) *HeadRequestBody {
	if m == nil {
		return nil
	}

	r := new(HeadRequestBody)

	r.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	r.SetMainOnly(m.GetMainOnly())

	r.SetRaw(m.GetRaw())

	return r
}

func HeadRequestToGRPCMessage(r *HeadRequest) *object.HeadRequest {
	if r == nil {
		return nil
	}

	m := new(object.HeadRequest)

	m.SetBody(
		HeadRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func HeadRequestFromGRPCMessage(m *object.HeadRequest) *HeadRequest {
	if m == nil {
		return nil
	}

	r := new(HeadRequest)

	r.SetBody(
		HeadRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func GetHeaderPartFullToGRPCMessage(r *GetHeaderPartFull) *object.HeadResponse_Body_Header {
	if r == nil {
		return nil
	}

	m := new(object.HeadResponse_Body_Header)

	m.SetHeader(
		HeaderToGRPCMessage(r.GetHeader()),
	)

	return m
}

func GetHeaderPartFullFromGRPCMessage(m *object.HeadResponse_Body_Header) *GetHeaderPartFull {
	if m == nil {
		return nil
	}

	r := new(GetHeaderPartFull)

	r.SetHeader(
		HeaderFromGRPCMessage(m.GetHeader()),
	)

	return r
}

func GetHeaderPartShortToGRPCMessage(r *GetHeaderPartShort) *object.HeadResponse_Body_ShortHeader {
	if r == nil {
		return nil
	}

	m := new(object.HeadResponse_Body_ShortHeader)

	m.SetShortHeader(
		ShortHeaderToGRPCMessage(r.GetShortHeader()),
	)

	return m
}

func GetHeaderPartShortFromGRPCMessage(m *object.HeadResponse_Body_ShortHeader) *GetHeaderPartShort {
	if m == nil {
		return nil
	}

	r := new(GetHeaderPartShort)

	r.SetShortHeader(
		ShortHeaderFromGRPCMessage(m.GetShortHeader()),
	)

	return r
}

func HeadResponseBodyToGRPCMessage(r *HeadResponseBody) *object.HeadResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.HeadResponse_Body)

	switch v := r.GetHeaderPart(); t := v.(type) {
	case nil:
	case *GetHeaderPartFull:
		m.SetHeader(
			GetHeaderPartFullToGRPCMessage(t),
		)
	case *GetHeaderPartShort:
		m.SetShortHeader(
			GetHeaderPartShortToGRPCMessage(t),
		)
	default:
		panic(fmt.Sprintf("unknown header part %T", t))
	}

	return m
}

func HeadResponseBodyFromGRPCMessage(m *object.HeadResponse_Body) *HeadResponseBody {
	if m == nil {
		return nil
	}

	r := new(HeadResponseBody)

	switch v := m.GetHead().(type) {
	case nil:
	case *object.HeadResponse_Body_Header:
		r.SetHeaderPart(
			GetHeaderPartFullFromGRPCMessage(v),
		)
	case *object.HeadResponse_Body_ShortHeader:
		r.SetHeaderPart(
			GetHeaderPartShortFromGRPCMessage(v),
		)
	default:
		panic(fmt.Sprintf("unknown header part %T", v))
	}

	return r
}

func HeadResponseToGRPCMessage(r *HeadResponse) *object.HeadResponse {
	if r == nil {
		return nil
	}

	m := new(object.HeadResponse)

	m.SetBody(
		HeadResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func HeadResponseFromGRPCMessage(m *object.HeadResponse) *HeadResponse {
	if m == nil {
		return nil
	}

	r := new(HeadResponse)

	r.SetBody(
		HeadResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func SearchFilterToGRPCMessage(f *SearchFilter) *object.SearchRequest_Body_Filter {
	if f == nil {
		return nil
	}

	m := new(object.SearchRequest_Body_Filter)

	m.SetMatchType(
		MatchTypeToGRPCField(f.GetMatchType()),
	)

	m.SetName(f.GetName())

	m.SetValue(f.GetValue())

	return m
}

func SearchFilterFromGRPCMessage(m *object.SearchRequest_Body_Filter) *SearchFilter {
	if m == nil {
		return nil
	}

	f := new(SearchFilter)

	f.SetMatchType(
		MatchTypeFromGRPCField(m.GetMatchType()),
	)

	f.SetName(m.GetName())

	f.SetValue(m.GetValue())

	return f
}

func SearchRequestBodyToGRPCMessage(r *SearchRequestBody) *object.SearchRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.SearchRequest_Body)

	m.SetContainerId(
		refs.ContainerIDToGRPCMessage(r.GetContainerID()),
	)

	m.SetVersion(r.GetVersion())

	filters := r.GetFilters()
	filterMsg := make([]*object.SearchRequest_Body_Filter, 0, len(filters))

	for i := range filters {
		filterMsg = append(filterMsg, SearchFilterToGRPCMessage(filters[i]))
	}

	m.SetFilters(filterMsg)

	return m
}

func SearchRequestBodyFromGRPCMessage(m *object.SearchRequest_Body) *SearchRequestBody {
	if m == nil {
		return nil
	}

	r := new(SearchRequestBody)

	r.SetContainerID(
		refs.ContainerIDFromGRPCMessage(m.GetContainerId()),
	)

	r.SetVersion(m.GetVersion())

	filterMsg := m.GetFilters()
	filters := make([]*SearchFilter, 0, len(filterMsg))

	for i := range filterMsg {
		filters = append(filters, SearchFilterFromGRPCMessage(filterMsg[i]))
	}

	r.SetFilters(filters)

	return r
}

func SearchRequestToGRPCMessage(r *SearchRequest) *object.SearchRequest {
	if r == nil {
		return nil
	}

	m := new(object.SearchRequest)

	m.SetBody(
		SearchRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func SearchRequestFromGRPCMessage(m *object.SearchRequest) *SearchRequest {
	if m == nil {
		return nil
	}

	r := new(SearchRequest)

	r.SetBody(
		SearchRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func SearchResponseBodyToGRPCMessage(r *SearchResponseBody) *object.SearchResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.SearchResponse_Body)

	ids := r.GetIDList()
	idMsg := make([]*refsGRPC.ObjectID, 0, len(ids))

	for i := range ids {
		idMsg = append(idMsg, refs.ObjectIDToGRPCMessage(ids[i]))
	}

	m.SetIdList(idMsg)

	return m
}

func SearchResponseBodyFromGRPCMessage(m *object.SearchResponse_Body) *SearchResponseBody {
	if m == nil {
		return nil
	}

	r := new(SearchResponseBody)

	idMsg := m.GetIdList()
	ids := make([]*refs.ObjectID, 0, len(idMsg))

	for i := range idMsg {
		ids = append(ids, refs.ObjectIDFromGRPCMessage(idMsg[i]))
	}

	r.SetIDList(ids)

	return r
}

func SearchResponseToGRPCMessage(r *SearchResponse) *object.SearchResponse {
	if r == nil {
		return nil
	}

	m := new(object.SearchResponse)

	m.SetBody(
		SearchResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func SearchResponseFromGRPCMessage(m *object.SearchResponse) *SearchResponse {
	if m == nil {
		return nil
	}

	r := new(SearchResponse)

	r.SetBody(
		SearchResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func RangeToGRPCMessage(r *Range) *object.Range {
	if r == nil {
		return nil
	}

	m := new(object.Range)

	m.SetOffset(r.GetOffset())
	m.SetLength(r.GetLength())

	return m
}

func RangeFromGRPCMessage(m *object.Range) *Range {
	if m == nil {
		return nil
	}

	r := new(Range)

	r.SetOffset(m.GetOffset())
	r.SetLength(m.GetLength())

	return r
}

func GetRangeRequestBodyToGRPCMessage(r *GetRangeRequestBody) *object.GetRangeRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeRequest_Body)

	m.SetAddress(
		refs.AddressToGRPCMessage(r.GetAddress()),
	)

	m.SetRange(
		RangeToGRPCMessage(r.GetRange()),
	)

	return m
}

func GetRangeRequestBodyFromGRPCMessage(m *object.GetRangeRequest_Body) *GetRangeRequestBody {
	if m == nil {
		return nil
	}

	r := new(GetRangeRequestBody)

	r.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	r.SetRange(
		RangeFromGRPCMessage(m.GetRange()),
	)

	return r
}

func GetRangeRequestToGRPCMessage(r *GetRangeRequest) *object.GetRangeRequest {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeRequest)

	m.SetBody(
		GetRangeRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func GetRangeRequestFromGRPCMessage(m *object.GetRangeRequest) *GetRangeRequest {
	if m == nil {
		return nil
	}

	r := new(GetRangeRequest)

	r.SetBody(
		GetRangeRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func GetRangeResponseBodyToGRPCMessage(r *GetRangeResponseBody) *object.GetRangeResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeResponse_Body)

	m.SetChunk(r.GetChunk())

	return m
}

func GetRangeResponseBodyFromGRPCMessage(m *object.GetRangeResponse_Body) *GetRangeResponseBody {
	if m == nil {
		return nil
	}

	r := new(GetRangeResponseBody)

	r.SetChunk(m.GetChunk())

	return r
}

func GetRangeResponseToGRPCMessage(r *GetRangeResponse) *object.GetRangeResponse {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeResponse)

	m.SetBody(
		GetRangeResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func GetRangeResponseFromGRPCMessage(m *object.GetRangeResponse) *GetRangeResponse {
	if m == nil {
		return nil
	}

	r := new(GetRangeResponse)

	r.SetBody(
		GetRangeResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}

func GetRangeHashRequestBodyToGRPCMessage(r *GetRangeHashRequestBody) *object.GetRangeHashRequest_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeHashRequest_Body)

	m.SetAddress(
		refs.AddressToGRPCMessage(r.GetAddress()),
	)

	m.SetSalt(r.GetSalt())

	rngs := r.GetRanges()
	rngMsg := make([]*object.Range, 0, len(rngs))

	for i := range rngs {
		rngMsg = append(rngMsg, RangeToGRPCMessage(rngs[i]))
	}

	m.SetRanges(rngMsg)

	m.SetType(refsGRPC.ChecksumType(r.GetType()))

	return m
}

func GetRangeHashRequestBodyFromGRPCMessage(m *object.GetRangeHashRequest_Body) *GetRangeHashRequestBody {
	if m == nil {
		return nil
	}

	r := new(GetRangeHashRequestBody)

	r.SetAddress(
		refs.AddressFromGRPCMessage(m.GetAddress()),
	)

	r.SetSalt(m.GetSalt())

	rngMsg := m.GetRanges()
	rngs := make([]*Range, 0, len(rngMsg))

	for i := range rngMsg {
		rngs = append(rngs, RangeFromGRPCMessage(rngMsg[i]))
	}

	r.SetRanges(rngs)

	r.SetType(refs.ChecksumType(m.GetType()))

	return r
}

func GetRangeHashRequestToGRPCMessage(r *GetRangeHashRequest) *object.GetRangeHashRequest {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeHashRequest)

	m.SetBody(
		GetRangeHashRequestBodyToGRPCMessage(r.GetBody()),
	)

	service.RequestHeadersToGRPC(r, m)

	return m
}

func GetRangeHashRequestFromGRPCMessage(m *object.GetRangeHashRequest) *GetRangeHashRequest {
	if m == nil {
		return nil
	}

	r := new(GetRangeHashRequest)

	r.SetBody(
		GetRangeHashRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, r)

	return r
}

func GetRangeHashResponseBodyToGRPCMessage(r *GetRangeHashResponseBody) *object.GetRangeHashResponse_Body {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeHashResponse_Body)

	m.SetType(refsGRPC.ChecksumType(r.GetType()))

	m.SetHashList(r.GetHashList())

	return m
}

func GetRangeHashResponseBodyFromGRPCMessage(m *object.GetRangeHashResponse_Body) *GetRangeHashResponseBody {
	if m == nil {
		return nil
	}

	r := new(GetRangeHashResponseBody)

	r.SetType(refs.ChecksumType(m.GetType()))

	r.SetHashList(m.GetHashList())

	return r
}

func GetRangeHashResponseToGRPCMessage(r *GetRangeHashResponse) *object.GetRangeHashResponse {
	if r == nil {
		return nil
	}

	m := new(object.GetRangeHashResponse)

	m.SetBody(
		GetRangeHashResponseBodyToGRPCMessage(r.GetBody()),
	)

	service.ResponseHeadersToGRPC(r, m)

	return m
}

func GetRangeHashResponseFromGRPCMessage(m *object.GetRangeHashResponse) *GetRangeHashResponse {
	if m == nil {
		return nil
	}

	r := new(GetRangeHashResponse)

	r.SetBody(
		GetRangeHashResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, r)

	return r
}
