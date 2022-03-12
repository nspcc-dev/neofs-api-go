package object

import (
	object "github.com/nspcc-dev/neofs-api-go/v2/object/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

const (
	shortHdrVersionField    = 1
	shortHdrEpochField      = 2
	shortHdrOwnerField      = 3
	shortHdrObjectTypeField = 4
	shortHdrPayloadLength   = 5
	shortHdrHashField       = 6
	shortHdrHomoHashField   = 7

	attributeKeyField   = 1
	attributeValueField = 2

	splitHdrParentField          = 1
	splitHdrPreviousField        = 2
	splitHdrParentSignatureField = 3
	splitHdrParentHeaderField    = 4
	splitHdrChildrenField        = 5
	splitHdrSplitIDField         = 6

	hdrVersionField         = 1
	hdrContainerIDField     = 2
	hdrOwnerIDField         = 3
	hdrEpochField           = 4
	hdrPayloadLengthField   = 5
	hdrPayloadHashField     = 6
	hdrObjectTypeField      = 7
	hdrHomomorphicHashField = 8
	hdrSessionTokenField    = 9
	hdrAttributesField      = 10
	hdrSplitField           = 11

	hdrWithSigHeaderField    = 1
	hdrWithSigSignatureField = 2

	objIDField        = 1
	objSignatureField = 2
	objHeaderField    = 3
	objPayloadField   = 4

	splitInfoSplitIDField  = 1
	splitInfoLastPartField = 2
	splitInfoLinkField     = 3

	getReqBodyAddressField = 1
	getReqBodyRawFlagField = 2

	getRespInitObjectIDField  = 1
	getRespInitSignatureField = 2
	getRespInitHeaderField    = 3

	getRespBodyInitField      = 1
	getRespBodyChunkField     = 2
	getRespBodySplitInfoField = 3

	putReqInitObjectIDField  = 1
	putReqInitSignatureField = 2
	putReqInitHeaderField    = 3
	putReqInitCopiesNumField = 4

	putReqBodyInitField  = 1
	putReqBodyChunkField = 2

	putRespBodyObjectIDField = 1

	deleteReqBodyAddressField = 1

	deleteRespBodyTombstoneFNum = 1

	headReqBodyAddressField  = 1
	headReqBodyMainFlagField = 2
	headReqBodyRawFlagField  = 3

	headRespBodyHeaderField      = 1
	headRespBodyShortHeaderField = 2
	headRespBodySplitInfoField   = 3

	searchFilterMatchField = 1
	searchFilterNameField  = 2
	searchFilterValueField = 3

	searchReqBodyContainerIDField = 1
	searchReqBodyVersionField     = 2
	searchReqBodyFiltersField     = 3

	searchRespBodyObjectIDsField = 1

	rangeOffsetField = 1
	rangeLengthField = 2

	getRangeReqBodyAddressField = 1
	getRangeReqBodyRangeField   = 2
	getRangeReqBodyRawField     = 3

	getRangeRespChunkField     = 1
	getRangeRespSplitInfoField = 2

	getRangeHashReqBodyAddressField = 1
	getRangeHashReqBodyRangesField  = 2
	getRangeHashReqBodySaltField    = 3
	getRangeHashReqBodyTypeField    = 4

	getRangeHashRespBodyTypeField     = 1
	getRangeHashRespBodyHashListField = 2
)

func (h *ShortHeader) StableMarshal(buf []byte) ([]byte, error) {
	if h == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, h.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(shortHdrVersionField, buf[offset:], h.version)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.UInt64Marshal(shortHdrEpochField, buf[offset:], h.creatEpoch)

	n, err = proto.NestedStructureMarshal(shortHdrOwnerField, buf[offset:], h.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.EnumMarshal(shortHdrObjectTypeField, buf[offset:], int32(h.typ))
	offset += proto.UInt64Marshal(shortHdrPayloadLength, buf[offset:], h.payloadLen)

	n, err = proto.NestedStructureMarshal(shortHdrHashField, buf[offset:], h.payloadHash)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(shortHdrHomoHashField, buf[offset:], h.homoHash)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *ShortHeader) StableSize() (size int) {
	if h == nil {
		return 0
	}

	size += proto.NestedStructureSize(shortHdrVersionField, h.version)
	size += proto.UInt64Size(shortHdrEpochField, h.creatEpoch)
	size += proto.NestedStructureSize(shortHdrOwnerField, h.ownerID)
	size += proto.EnumSize(shortHdrObjectTypeField, int32(h.typ))
	size += proto.UInt64Size(shortHdrPayloadLength, h.payloadLen)
	size += proto.NestedStructureSize(shortHdrHashField, h.payloadHash)
	size += proto.NestedStructureSize(shortHdrHomoHashField, h.homoHash)

	return size
}

func (h *ShortHeader) Unmarshal(data []byte) error {
	return message.Unmarshal(h, data, new(object.ShortHeader))
}

func (a *Attribute) StableMarshal(buf []byte) ([]byte, error) {
	if a == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, a.StableSize())
	}

	var offset int

	offset += proto.StringMarshal(attributeKeyField, buf[offset:], a.key)
	proto.StringMarshal(attributeValueField, buf[offset:], a.val)

	return buf, nil
}

func (a *Attribute) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.StringSize(shortHdrVersionField, a.key)
	size += proto.StringSize(shortHdrEpochField, a.val)

	return size
}

func (a *Attribute) Unmarshal(data []byte) error {
	return message.Unmarshal(a, data, new(object.Header_Attribute))
}

func (h *SplitHeader) StableMarshal(buf []byte) ([]byte, error) {
	if h == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, h.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(splitHdrParentField, buf[offset:], h.par)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(splitHdrPreviousField, buf[offset:], h.prev)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(splitHdrParentSignatureField, buf[offset:], h.parSig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(splitHdrParentHeaderField, buf[offset:], h.parHdr)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = refs.ObjectIDNestedListMarshal(splitHdrChildrenField, buf[offset:], h.children)
	if err != nil {
		return nil, err
	}

	offset += n

	proto.BytesMarshal(splitHdrSplitIDField, buf[offset:], h.splitID)

	return buf, nil
}

func (h *SplitHeader) StableSize() (size int) {
	if h == nil {
		return 0
	}

	size += proto.NestedStructureSize(splitHdrParentField, h.par)
	size += proto.NestedStructureSize(splitHdrPreviousField, h.prev)
	size += proto.NestedStructureSize(splitHdrParentSignatureField, h.parSig)
	size += proto.NestedStructureSize(splitHdrParentHeaderField, h.parHdr)
	size += refs.ObjectIDNestedListSize(splitHdrChildrenField, h.children)
	size += proto.BytesSize(splitHdrSplitIDField, h.splitID)

	return size
}

func (h *SplitHeader) Unmarshal(data []byte) error {
	return message.Unmarshal(h, data, new(object.Header_Split))
}

func (h *Header) StableMarshal(buf []byte) ([]byte, error) {
	if h == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, h.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(hdrVersionField, buf[offset:], h.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(hdrContainerIDField, buf[offset:], h.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(hdrOwnerIDField, buf[offset:], h.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.UInt64Marshal(hdrEpochField, buf[offset:], h.creatEpoch)
	offset += proto.UInt64Marshal(hdrPayloadLengthField, buf[offset:], h.payloadLen)

	n, err = proto.NestedStructureMarshal(hdrPayloadHashField, buf[offset:], h.payloadHash)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.EnumMarshal(hdrObjectTypeField, buf[offset:], int32(h.typ))

	n, err = proto.NestedStructureMarshal(hdrHomomorphicHashField, buf[offset:], h.homoHash)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(hdrSessionTokenField, buf[offset:], h.sessionToken)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range h.attr {
		n, err = proto.NestedStructureMarshal(hdrAttributesField, buf[offset:], &h.attr[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	_, err = proto.NestedStructureMarshal(hdrSplitField, buf[offset:], h.split)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *Header) StableSize() (size int) {
	if h == nil {
		return 0
	}

	size += proto.NestedStructureSize(hdrVersionField, h.version)
	size += proto.NestedStructureSize(hdrContainerIDField, h.cid)
	size += proto.NestedStructureSize(hdrOwnerIDField, h.ownerID)
	size += proto.UInt64Size(hdrEpochField, h.creatEpoch)
	size += proto.UInt64Size(hdrPayloadLengthField, h.payloadLen)
	size += proto.NestedStructureSize(hdrPayloadHashField, h.payloadHash)
	size += proto.EnumSize(hdrObjectTypeField, int32(h.typ))
	size += proto.NestedStructureSize(hdrHomomorphicHashField, h.homoHash)
	size += proto.NestedStructureSize(hdrSessionTokenField, h.sessionToken)
	for i := range h.attr {
		size += proto.NestedStructureSize(hdrAttributesField, &h.attr[i])
	}
	size += proto.NestedStructureSize(hdrSplitField, h.split)

	return size
}

func (h *Header) Unmarshal(data []byte) error {
	return message.Unmarshal(h, data, new(object.Header))
}

func (h *HeaderWithSignature) StableMarshal(buf []byte) ([]byte, error) {
	if h == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, h.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(hdrWithSigHeaderField, buf[offset:], h.header)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(hdrWithSigSignatureField, buf[offset:], h.signature)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (h *HeaderWithSignature) StableSize() (size int) {
	if h == nil {
		return 0
	}

	size += proto.NestedStructureSize(hdrVersionField, h.header)
	size += proto.NestedStructureSize(hdrContainerIDField, h.signature)

	return size
}

func (h *HeaderWithSignature) Unmarshal(data []byte) error {
	return message.Unmarshal(h, data, new(object.HeaderWithSignature))
}

func (o *Object) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(objIDField, buf[offset:], o.objectID)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(objSignatureField, buf[offset:], o.idSig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(objHeaderField, buf[offset:], o.header)
	if err != nil {
		return nil, err
	}

	offset += n

	proto.BytesMarshal(objPayloadField, buf[offset:], o.payload)

	return buf, nil
}

func (o *Object) StableSize() (size int) {
	if o == nil {
		return 0
	}

	size += proto.NestedStructureSize(objIDField, o.objectID)
	size += proto.NestedStructureSize(objSignatureField, o.idSig)
	size += proto.NestedStructureSize(objHeaderField, o.header)
	size += proto.BytesSize(objPayloadField, o.payload)

	return size
}

func (o *Object) Unmarshal(data []byte) error {
	return message.Unmarshal(o, data, new(object.Object))
}

func (s *SplitInfo) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	offset += proto.BytesMarshal(splitInfoSplitIDField, buf[offset:], s.splitID)

	n, err = proto.NestedStructureMarshal(splitInfoLastPartField, buf[offset:], s.lastPart)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(splitInfoLinkField, buf[offset:], s.link)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *SplitInfo) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.BytesSize(splitInfoSplitIDField, s.splitID)
	size += proto.NestedStructureSize(splitInfoLastPartField, s.lastPart)
	size += proto.NestedStructureSize(splitInfoLinkField, s.link)

	return size
}

func (s *SplitInfo) Unmarshal(data []byte) error {
	return message.Unmarshal(s, data, new(object.SplitInfo))
}

func (r *GetRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(getReqBodyAddressField, buf[offset:], r.addr)
	if err != nil {
		return nil, err
	}

	offset += n

	proto.BoolMarshal(getReqBodyRawFlagField, buf[offset:], r.raw)

	return buf, nil
}

func (r *GetRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getReqBodyAddressField, r.addr)
	size += proto.BoolSize(getReqBodyRawFlagField, r.raw)

	return size
}

func (r *GetRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetRequest_Body))
}

func (r *GetObjectPartInit) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(getRespInitObjectIDField, buf[offset:], r.id)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(getRespInitSignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(getRespInitHeaderField, buf[offset:], r.hdr)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetObjectPartInit) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getRespInitObjectIDField, r.id)
	size += proto.NestedStructureSize(getRespInitSignatureField, r.sig)
	size += proto.NestedStructureSize(getRespInitHeaderField, r.hdr)

	return size
}

func (r *GetObjectPartInit) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetResponse_Body_Init))
}

func (r *GetResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	switch v := r.objPart.(type) {
	case nil:
	case *GetObjectPartInit:
		_, err := proto.NestedStructureMarshal(getRespBodyInitField, buf, v)
		if err != nil {
			return nil, err
		}
	case *GetObjectPartChunk:
		if v != nil {
			proto.BytesMarshal(getRespBodyChunkField, buf, v.chunk)
		}
	case *SplitInfo:
		_, err := proto.NestedStructureMarshal(getRespBodySplitInfoField, buf, v)
		if err != nil {
			return nil, err
		}
	default:
		panic("unknown one of object get response body type")
	}

	return buf, nil
}

func (r *GetResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	switch v := r.objPart.(type) {
	case nil:
	case *GetObjectPartInit:
		size += proto.NestedStructureSize(getRespBodyInitField, v)
	case *GetObjectPartChunk:
		if v != nil {
			size += proto.BytesSize(getRespBodyChunkField, v.chunk)
		}
	case *SplitInfo:
		size += proto.NestedStructureSize(getRespBodySplitInfoField, v)
	default:
		panic("unknown one of object get response body type")
	}

	return
}

func (r *GetResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetResponse_Body))
}

func (r *PutObjectPartInit) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(putReqInitObjectIDField, buf[offset:], r.id)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(putReqInitSignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(putReqInitHeaderField, buf[offset:], r.hdr)
	if err != nil {
		return nil, err
	}

	offset += n

	proto.UInt32Marshal(putReqInitCopiesNumField, buf[offset:], r.copyNum)

	return buf, nil
}

func (r *PutObjectPartInit) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(putReqInitObjectIDField, r.id)
	size += proto.NestedStructureSize(putReqInitSignatureField, r.sig)
	size += proto.NestedStructureSize(putReqInitHeaderField, r.hdr)
	size += proto.UInt32Size(putReqInitCopiesNumField, r.copyNum)

	return size
}

func (r *PutObjectPartInit) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.PutRequest_Body_Init))
}

func (r *PutRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	switch v := r.objPart.(type) {
	case nil:
	case *PutObjectPartInit:
		_, err := proto.NestedStructureMarshal(putReqBodyInitField, buf, v)
		if err != nil {
			return nil, err
		}
	case *PutObjectPartChunk:
		if v != nil {
			proto.BytesMarshal(putReqBodyChunkField, buf, v.chunk)
		}
	default:
		panic("unknown one of object put request body type")
	}

	return buf, nil
}

func (r *PutRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	switch v := r.objPart.(type) {
	case nil:
	case *PutObjectPartInit:
		size += proto.NestedStructureSize(putReqBodyInitField, v)
	case *PutObjectPartChunk:
		if v != nil {
			size += proto.BytesSize(putReqBodyChunkField, v.chunk)
		}
	default:
		panic("unknown one of object get response body type")
	}

	return size
}

func (r *PutRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.PutRequest_Body))
}

func (r *PutResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(putRespBodyObjectIDField, buf, r.id)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *PutResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(putRespBodyObjectIDField, r.id)

	return size
}

func (r *PutResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.PutResponse_Body))
}

func (r *DeleteRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(deleteReqBodyAddressField, buf, r.addr)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *DeleteRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(deleteReqBodyAddressField, r.addr)

	return size
}

func (r *DeleteRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.DeleteRequest_Body))
}

func (r *DeleteResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(deleteRespBodyTombstoneFNum, buf, r.tombstone)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *DeleteResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(deleteRespBodyTombstoneFNum, r.tombstone)

	return size
}

func (r *DeleteResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.DeleteResponse_Body))
}

func (r *HeadRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(headReqBodyAddressField, buf[offset:], r.addr)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.BoolMarshal(headReqBodyMainFlagField, buf[offset:], r.mainOnly)
	proto.BoolMarshal(headReqBodyRawFlagField, buf[offset:], r.raw)

	return buf, nil
}

func (r *HeadRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(headReqBodyAddressField, r.addr)
	size += proto.BoolSize(headReqBodyMainFlagField, r.mainOnly)
	size += proto.BoolSize(headReqBodyRawFlagField, r.raw)

	return size
}

func (r *HeadRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.HeadRequest_Body))
}

func (r *HeadResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	switch v := r.hdrPart.(type) {
	case nil:
	case *HeaderWithSignature:
		if v != nil {
			_, err := proto.NestedStructureMarshal(headRespBodyHeaderField, buf, v)
			if err != nil {
				return nil, err
			}
		}
	case *ShortHeader:
		if v != nil {
			_, err := proto.NestedStructureMarshal(headRespBodyShortHeaderField, buf, v)
			if err != nil {
				return nil, err
			}
		}
	case *SplitInfo:
		if v != nil {
			_, err := proto.NestedStructureMarshal(headRespBodySplitInfoField, buf, v)
			if err != nil {
				return nil, err
			}
		}
	default:
		panic("unknown one of object put request body type")
	}

	return buf, nil
}

func (r *HeadResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	switch v := r.hdrPart.(type) {
	case nil:
	case *HeaderWithSignature:
		if v != nil {
			size += proto.NestedStructureSize(headRespBodyHeaderField, v)
		}
	case *ShortHeader:
		if v != nil {
			size += proto.NestedStructureSize(headRespBodyShortHeaderField, v)
		}
	case *SplitInfo:
		if v != nil {
			size += proto.NestedStructureSize(headRespBodySplitInfoField, v)
		}
	default:
		panic("unknown one of object put request body type")
	}

	return
}

func (r *HeadResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.HeadResponse_Body))
}

func (f *SearchFilter) StableMarshal(buf []byte) ([]byte, error) {
	if f == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, f.StableSize())
	}

	var offset int

	offset += proto.EnumMarshal(searchFilterMatchField, buf[offset:], int32(f.matchType))
	offset += proto.StringMarshal(searchFilterNameField, buf[offset:], f.key)
	proto.StringMarshal(searchFilterValueField, buf[offset:], f.val)

	return buf, nil
}

func (f *SearchFilter) StableSize() (size int) {
	if f == nil {
		return 0
	}

	size += proto.EnumSize(searchFilterMatchField, int32(f.matchType))
	size += proto.StringSize(searchFilterNameField, f.key)
	size += proto.StringSize(searchFilterValueField, f.val)

	return size
}

func (f *SearchFilter) Unmarshal(data []byte) error {
	return message.Unmarshal(f, data, new(object.SearchRequest_Body_Filter))
}

func (r *SearchRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(searchReqBodyContainerIDField, buf[offset:], r.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	offset += proto.UInt32Marshal(searchReqBodyVersionField, buf[offset:], r.version)

	for i := range r.filters {
		n, err = proto.NestedStructureMarshal(searchReqBodyFiltersField, buf[offset:], &r.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (r *SearchRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(searchReqBodyContainerIDField, r.cid)
	size += proto.UInt32Size(searchReqBodyVersionField, r.version)

	for i := range r.filters {
		size += proto.NestedStructureSize(searchReqBodyFiltersField, &r.filters[i])
	}

	return size
}

func (r *SearchRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.SearchRequest_Body))
}

func (r *SearchResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset int
		err    error
	)

	_, err = refs.ObjectIDNestedListMarshal(searchRespBodyObjectIDsField, buf[offset:], r.idList)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *SearchResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += refs.ObjectIDNestedListSize(searchRespBodyObjectIDsField, r.idList)

	return size
}

func (r *SearchResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.SearchResponse_Body))
}

func (r *Range) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var offset int

	offset += proto.UInt64Marshal(rangeOffsetField, buf[offset:], r.off)
	proto.UInt64Marshal(rangeLengthField, buf[offset:], r.len)

	return buf, nil
}

func (r *Range) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.UInt64Size(rangeOffsetField, r.off)
	size += proto.UInt64Size(rangeLengthField, r.len)

	return size
}

func (r *Range) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.Range))
}

func (r *GetRangeRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(getRangeReqBodyAddressField, buf[offset:], r.addr)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(getRangeReqBodyRangeField, buf[offset:], r.rng)
	if err != nil {
		return nil, err
	}

	offset += n

	proto.BoolMarshal(getRangeReqBodyRawField, buf[offset:], r.raw)

	return buf, nil
}

func (r *GetRangeRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getRangeReqBodyAddressField, r.addr)
	size += proto.NestedStructureSize(getRangeReqBodyRangeField, r.rng)
	size += proto.BoolSize(getRangeReqBodyRawField, r.raw)

	return size
}

func (r *GetRangeRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetRangeRequest_Body))
}

func (r *GetRangeResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	switch v := r.rngPart.(type) {
	case nil:
	case *GetRangePartChunk:
		if v != nil {
			proto.BytesMarshal(getRangeRespChunkField, buf, v.chunk)
		}
	case *SplitInfo:
		if v != nil {
			_, err := proto.NestedStructureMarshal(getRangeRespSplitInfoField, buf, v)
			if err != nil {
				return nil, err
			}
		}
	default:
		panic("unknown one of object get range request body type")
	}

	return buf, nil
}

func (r *GetRangeResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	switch v := r.rngPart.(type) {
	case nil:
	case *GetRangePartChunk:
		if v != nil {
			size += proto.BytesSize(getRangeRespChunkField, v.chunk)
		}
	case *SplitInfo:
		if v != nil {
			size = proto.NestedStructureSize(getRangeRespSplitInfoField, v)
		}
	default:
		panic("unknown one of object get range request body type")
	}

	return
}

func (r *GetRangeResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetRangeResponse_Body))
}

func (r *GetRangeHashRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(getRangeHashReqBodyAddressField, buf[offset:], r.addr)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range r.rngs {
		n, err = proto.NestedStructureMarshal(getRangeHashReqBodyRangesField, buf[offset:], &r.rngs[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	offset += proto.BytesMarshal(getRangeHashReqBodySaltField, buf[offset:], r.salt)
	proto.EnumMarshal(getRangeHashReqBodyTypeField, buf[offset:], int32(r.typ))

	return buf, nil
}

func (r *GetRangeHashRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getRangeHashReqBodyAddressField, r.addr)

	for i := range r.rngs {
		size += proto.NestedStructureSize(getRangeHashReqBodyRangesField, &r.rngs[i])
	}

	size += proto.BytesSize(getRangeHashReqBodySaltField, r.salt)
	size += proto.EnumSize(getRangeHashReqBodyTypeField, int32(r.typ))

	return size
}

func (r *GetRangeHashRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetRangeHashRequest_Body))
}

func (r *GetRangeHashResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var offset int

	offset += proto.EnumMarshal(getRangeHashRespBodyTypeField, buf, int32(r.typ))
	proto.RepeatedBytesMarshal(getRangeHashRespBodyHashListField, buf[offset:], r.hashList)

	return buf, nil
}

func (r *GetRangeHashResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.EnumSize(getRangeHashRespBodyTypeField, int32(r.typ))
	size += proto.RepeatedBytesSize(getRangeHashRespBodyHashListField, r.hashList)

	return size
}

func (r *GetRangeHashResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(object.GetRangeHashResponse_Body))
}
