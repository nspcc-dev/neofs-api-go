package container

import (
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
	"google.golang.org/protobuf/proto"
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

	usedSpaceAnnounceCIDField       = 1
	usedSpaceAnnounceUsedSpaceField = 2

	usedSpaceReqBodyAnnouncementsField = 1
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

	n, err = protoutil.StringMarshal(attributeKeyField, buf[offset:], a.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.StringMarshal(attributeValueField, buf[offset:], a.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (a *Attribute) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += protoutil.StringSize(attributeKeyField, a.key)
	size += protoutil.StringSize(attributeValueField, a.val)

	return size
}

func (a *Attribute) Unmarshal(data []byte) error {
	m := new(container.Container_Attribute)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*a = *AttributeFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.NestedStructureMarshal(containerVersionField, buf[offset:], c.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.NestedStructureMarshal(containerOwnerField, buf[offset:], c.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.BytesMarshal(containerNonceField, buf[offset:], c.nonce)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.UInt32Marshal(containerBasicACLField, buf[offset:], c.basicACL)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range c.attr {
		n, err = protoutil.NestedStructureMarshal(containerAttributesField, buf[offset:], c.attr[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	_, err = protoutil.NestedStructureMarshal(containerPlacementField, buf[offset:], c.policy)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Container) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(containerVersionField, c.version)
	size += protoutil.NestedStructureSize(containerOwnerField, c.ownerID)
	size += protoutil.BytesSize(containerNonceField, c.nonce)
	size += protoutil.UInt32Size(containerBasicACLField, c.basicACL)

	for i := range c.attr {
		size += protoutil.NestedStructureSize(containerAttributesField, c.attr[i])
	}

	size += protoutil.NestedStructureSize(containerPlacementField, c.policy)

	return size
}

func (c *Container) Unmarshal(data []byte) error {
	m := new(container.Container)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*c = *ContainerFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.NestedStructureMarshal(putReqBodyContainerField, buf[offset:], r.cnr)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(putReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *PutRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(putReqBodyContainerField, r.cnr)
	size += protoutil.NestedStructureSize(putReqBodySignatureField, r.sig)

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

	_, err = protoutil.NestedStructureMarshal(putRespBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *PutResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(putRespBodyIDField, r.cid)

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

	n, err = protoutil.NestedStructureMarshal(deleteReqBodyIDField, buf[offset:], r.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(deleteReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *DeleteRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(deleteReqBodyIDField, r.cid)
	size += protoutil.NestedStructureSize(deleteReqBodySignatureField, r.sig)

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

	_, err := protoutil.NestedStructureMarshal(getReqBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(getReqBodyIDField, r.cid)

	return size
}

func (r *GetResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := protoutil.NestedStructureMarshal(getRespBodyContainerField, buf, r.cnr)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(getRespBodyContainerField, r.cnr)

	return size
}

func (r *ListRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, r.StableSize())
	}

	_, err := protoutil.NestedStructureMarshal(listReqBodyOwnerField, buf, r.ownerID)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *ListRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(listReqBodyOwnerField, r.ownerID)

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
		n, err = protoutil.NestedStructureMarshal(listRespBodyIDsField, buf[offset:], r.cidList[i])
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
		size += protoutil.NestedStructureSize(listRespBodyIDsField, r.cidList[i])
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

	n, err = protoutil.NestedStructureMarshal(setEACLReqBodyTableField, buf[offset:], r.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(setEACLReqBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *SetExtendedACLRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(setEACLReqBodyTableField, r.eacl)
	size += protoutil.NestedStructureSize(setEACLReqBodySignatureField, r.sig)

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

	_, err := protoutil.NestedStructureMarshal(getEACLReqBodyIDField, buf, r.cid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetExtendedACLRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(getEACLReqBodyIDField, r.cid)

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

	n, err = protoutil.NestedStructureMarshal(getEACLRespBodyTableField, buf[offset:], r.eacl)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(getEACLRespBodySignatureField, buf[offset:], r.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *GetExtendedACLResponseBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(getEACLRespBodyTableField, r.eacl)
	size += protoutil.NestedStructureSize(getEACLRespBodySignatureField, r.sig)

	return size
}

func (a *UsedSpaceAnnouncement) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = protoutil.NestedStructureMarshal(usedSpaceAnnounceCIDField, buf[offset:], a.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.UInt64Marshal(usedSpaceAnnounceUsedSpaceField, buf[offset:], a.usedSpace)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (a *UsedSpaceAnnouncement) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(usedSpaceAnnounceCIDField, a.cid)
	size += protoutil.UInt64Size(usedSpaceAnnounceUsedSpaceField, a.usedSpace)

	return size
}

func (r *AnnounceUsedSpaceRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	for i := range r.announcements {
		n, err = protoutil.NestedStructureMarshal(usedSpaceReqBodyAnnouncementsField, buf[offset:], r.announcements[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (r *AnnounceUsedSpaceRequestBody) StableSize() (size int) {
	if r == nil {
		return 0
	}

	for i := range r.announcements {
		size += protoutil.NestedStructureSize(usedSpaceReqBodyAnnouncementsField, r.announcements[i])
	}

	return size
}

func (r *AnnounceUsedSpaceResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *AnnounceUsedSpaceResponseBody) StableSize() (size int) {
	return 0
}
