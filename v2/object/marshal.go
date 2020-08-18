package object

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	shortHdrVersionField    = 1
	shortHdrEpochField      = 2
	shortHdrOwnerField      = 3
	shortHdrObjectTypeField = 4
	shortHdrPayloadLength   = 5

	attributeKeyField   = 1
	attributeValueField = 2

	splitHdrParentField          = 1
	splitHdrPreviousField        = 2
	splitHdrParentSignatureField = 3
	splitHdrParentHeaderField    = 4
	splitHdrChildrenField        = 5

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

	objIDField        = 1
	objSignatureField = 2
	objHeaderField    = 3
	objPayloadField   = 4
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

	n, err = proto.UInt64Marshal(shortHdrEpochField, buf[offset:], h.creatEpoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(shortHdrOwnerField, buf[offset:], h.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(shortHdrObjectTypeField, buf[offset:], int32(h.typ))
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt64Marshal(shortHdrPayloadLength, buf[offset:], h.payloadLen)
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

	return size
}

func (a *Attribute) StableMarshal(buf []byte) ([]byte, error) {
	if a == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, a.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.StringMarshal(attributeKeyField, buf[offset:], a.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.StringMarshal(attributeValueField, buf[offset:], a.val)
	if err != nil {
		return nil, err
	}

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

	for i := range h.children {
		n, err = proto.NestedStructureMarshal(splitHdrChildrenField, buf[offset:], h.children[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

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

	for i := range h.children {
		size += proto.NestedStructureSize(splitHdrChildrenField, h.children[i])
	}

	return size
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

	n, err = proto.UInt64Marshal(hdrEpochField, buf[offset:], h.creatEpoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(hdrPayloadLengthField, buf[offset:], h.payloadLen)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(hdrPayloadHashField, buf[offset:], h.payloadHash)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(hdrObjectTypeField, buf[offset:], int32(h.typ))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(hdrHomomorphicHashField, buf[offset:], h.homoHash)
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
		n, err = proto.NestedStructureMarshal(hdrAttributesField, buf[offset:], h.attr[i])
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
	size += proto.BytesSize(hdrPayloadHashField, h.payloadHash)
	size += proto.EnumSize(hdrObjectTypeField, int32(h.typ))
	size += proto.BytesSize(hdrHomomorphicHashField, h.homoHash)
	size += proto.NestedStructureSize(hdrSessionTokenField, h.sessionToken)
	for i := range h.attr {
		size += proto.NestedStructureSize(hdrAttributesField, h.attr[i])
	}
	size += proto.NestedStructureSize(hdrSplitField, h.split)

	return size
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

	_, err = proto.BytesMarshal(objPayloadField, buf[offset:], o.payload)
	if err != nil {
		return nil, err
	}

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

func (r *PutRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *PutRequestBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *GetResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *GetResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *PutResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *PutResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *DeleteResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *DeleteResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *HeadResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *HeadResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *SearchResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *SearchResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *GetRangeResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *GetRangeResponseBody) StableSize() (size int) {
	panic("not implemented")
}

func (r *GetRangeHashResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *GetRangeHashResponseBody) StableSize() (size int) {
	panic("not implemented")
}
