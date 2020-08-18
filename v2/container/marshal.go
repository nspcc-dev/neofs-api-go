package container

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	attributeKeyField   = 1
	attributeValueField = 2

	containerVersionField    = 1
	containerOwnerField      = 2
	containerNonceField      = 3
	containerBasicACLField   = 4
	containerAttributesField = 5
	containerPlacementField  = 6

	putReqBodyContainerField = 1
	putReqBodySignatureField = 2

	putRespBodyIDField = 1

	deleteReqBodyIDField        = 1
	deleteReqBodySignatureField = 2

	getReqBodyIDField = 1

	getRespBodyContainerField = 1

	listReqBodyOwnerField = 1

	listRespBodyIDsField = 1

	setEACLReqBodyTableField     = 1
	setEACLReqBodySignatureField = 2

	getEACLReqBodyIDField = 1

	getEACLRespBodyTableField     = 1
	getEACLRespBodySignatureField = 2
)

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

	size += proto.StringSize(attributeKeyField, a.key)
	size += proto.StringSize(attributeValueField, a.val)

	return size
}

func (c *Container) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(containerVersionField, buf[offset:], c.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(containerOwnerField, buf[offset:], c.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(containerNonceField, buf[offset:], c.nonce)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(containerBasicACLField, buf[offset:], c.basicACL)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range c.attr {
		n, err = proto.NestedStructureMarshal(containerAttributesField, buf[offset:], c.attr[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	_, err = proto.NestedStructureMarshal(containerPlacementField, buf[offset:], c.policy)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Container) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.NestedStructureSize(containerVersionField, c.version)
	size += proto.NestedStructureSize(containerOwnerField, c.ownerID)
	size += proto.BytesSize(containerNonceField, c.nonce)
	size += proto.UInt32Size(containerBasicACLField, c.basicACL)

	for i := range c.attr {
		size += proto.NestedStructureSize(containerAttributesField, c.attr[i])
	}

	size += proto.NestedStructureSize(containerPlacementField, c.policy)

	return size
}

func (r *PutRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.NestedStructureMarshal(putReqBodyContainerField, buf[offset:], r.cnr)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(putReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *PutRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(putReqBodyContainerField, r.cnr)
	size += proto.NestedStructureSize(putReqBodySignatureField, r.sig)

	return size
}

func (r *PutResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		err error
	)

	_, err = proto.NestedStructureMarshal(putRespBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *PutResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(putRespBodyIDField, r.cid)

	return size
}

func (r *DeleteRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.NestedStructureMarshal(deleteReqBodyIDField, buf[offset:], r.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(deleteReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *DeleteRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(deleteReqBodyIDField, r.cid)
	size += proto.NestedStructureSize(deleteReqBodySignatureField, r.sig)

	return size
}

func (r *DeleteResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *DeleteResponseBody) StableSize() (size int) {
	return 0
}

func (r *GetRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(getReqBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getReqBodyIDField, r.cid)

	return size
}

func (r *GetResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(getRespBodyContainerField, buf, r.cnr)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getRespBodyContainerField, r.cnr)

	return size
}

func (r *ListRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(listReqBodyOwnerField, buf, r.ownerID)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *ListRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(listReqBodyOwnerField, r.ownerID)

	return size
}

func (r *ListResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	var (
		n, offset int
		err       error
	)

	for i := range r.cidList {
		n, err = proto.NestedStructureMarshal(listRespBodyIDsField, buf[offset:], r.cidList[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (r *ListResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	for i := range r.cidList {
		size += proto.NestedStructureSize(listRespBodyIDsField, r.cidList[i])
	}

	return size
}

func (r *SetExtendedACLRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.NestedStructureMarshal(setEACLReqBodyTableField, buf[offset:], r.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(setEACLReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *SetExtendedACLRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(setEACLReqBodyTableField, r.eacl)
	size += proto.NestedStructureSize(setEACLReqBodySignatureField, r.sig)

	return size
}

func (r *SetExtendedACLResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *SetExtendedACLResponseBody) StableSize() (size int) {
	return 0
}

func (r *GetExtendedACLRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := proto.NestedStructureMarshal(getEACLReqBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetExtendedACLRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getEACLReqBodyIDField, r.cid)

	return size
}

func (r *GetExtendedACLResponseBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.NestedStructureMarshal(getEACLRespBodyTableField, buf[offset:], r.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(getEACLRespBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetExtendedACLResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.NestedStructureSize(getEACLRespBodyTableField, r.eacl)
	size += proto.NestedStructureSize(getEACLRespBodySignatureField, r.sig)

	return size
}
