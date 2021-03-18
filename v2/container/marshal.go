package container

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/util/proto"
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	container "github.com/nspcc-dev/neofs-api-go/v2/container/grpc"
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

	usedSpaceAnnounceEpochField     = 1
	usedSpaceAnnounceCIDField       = 2
	usedSpaceAnnounceUsedSpaceField = 3

	usedSpaceReqBodyAnnouncementsField = 1
)

func (a *Attribute) MarshalStream(s proto.Stream) (int, error) {
	if a == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.StringMarshal(attributeKeyField, a.key)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.StringMarshal(attributeValueField, a.val)
	return offset + n, err
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
	return message.Unmarshal(a, data, new(container.Container_Attribute))
}

func (c *Container) MarshalStream(s proto.Stream) (int, error) {
	if c == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.NestedStructureMarshal(containerVersionField, c.version)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.NestedStructureMarshal(containerOwnerField, c.ownerID)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.BytesMarshal(containerNonceField, c.nonce)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.UInt32Marshal(containerBasicACLField, c.basicACL)
	if err != nil {
		return offset + n, err
	}

	offset += n

	for i := range c.attr {
		n, err = s.NestedStructureMarshal(containerAttributesField, c.attr[i])
		if err != nil {
			return offset + n, err
		}

		offset += n
	}

	n, err = s.NestedStructureMarshal(containerPlacementField, c.policy)
	return offset + n, err
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
	return message.Unmarshal(c, data, new(container.Container))
}

func (r *PutRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.NestedStructureMarshal(putReqBodyContainerField, r.cnr)
	if err != nil {
		return offset + n, err
	}

	offset += n

	_, err = s.NestedStructureMarshal(putReqBodySignatureField, r.sig)
	return offset + n, err
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

func (r *PutRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.PutRequest_Body))
}

func (r *PutResponseBody) MarshalStream(s proto.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(putRespBodyIDField, r.cid)
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

func (r *PutResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.PutResponse_Body))
}

func (r *DeleteRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.NestedStructureMarshal(deleteReqBodyIDField, r.cid)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.NestedStructureMarshal(deleteReqBodySignatureField, r.sig)
	return offset + n, err
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

func (r *DeleteRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.DeleteRequest_Body))
}

func (r *DeleteResponseBody) MarshalStream(protoutil.Stream) (int, error) {
	return 0, nil
}

func (r *DeleteResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *DeleteResponseBody) StableSize() (size int) {
	return 0
}

func (r *DeleteResponseBody) Unmarshal([]byte) error {
	return nil
}

func (r *GetRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(getReqBodyIDField, r.cid)
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

func (r *GetRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.GetRequest_Body))
}

func (r *GetResponseBody) MarshalStream(s proto.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(getRespBodyContainerField, r.cnr)
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

func (r *GetResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.GetResponse_Body))
}

func (r *ListRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(listReqBodyOwnerField, r.ownerID)
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

func (r *ListRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.ListRequest_Body))
}

func (r *ListResponseBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		n, offset int
		err       error
	)

	for i := range r.cidList {
		n, err = s.NestedStructureMarshal(listRespBodyIDsField, r.cidList[i])
		if err != nil {
			return offset + n, err
		}

		offset += n
	}

	return offset, nil
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

func (r *ListResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.ListResponse_Body))
}

func (r *SetExtendedACLRequestBody) MarshalStream(s proto.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.NestedStructureMarshal(setEACLReqBodyTableField, r.eacl)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.NestedStructureMarshal(setEACLReqBodySignatureField, r.sig)
	return offset + n, nil
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

func (r *SetExtendedACLRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.SetExtendedACLRequest_Body))
}

func (r *SetExtendedACLResponseBody) MarshalStream(_ proto.Stream) (int, error) {
	return 0, nil
}

func (r *SetExtendedACLResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *SetExtendedACLResponseBody) StableSize() (size int) {
	return 0
}

func (r *SetExtendedACLResponseBody) Unmarshal([]byte) error {
	return nil
}

func (r *GetExtendedACLRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(getEACLReqBodyIDField, r.cid)
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

func (r *GetExtendedACLRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.GetExtendedACLRequest_Body))
}

func (r *GetExtendedACLResponseBody) MarshalStream(s proto.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.NestedStructureMarshal(getEACLRespBodyTableField, r.eacl)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.NestedStructureMarshal(getEACLRespBodySignatureField, r.sig)
	return offset + n, err
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

func (r *GetExtendedACLResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.GetExtendedACLResponse_Body))
}

func (a *UsedSpaceAnnouncement) MarshalStream(s protoutil.Stream) (int, error) {
	if a == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.UInt64Marshal(usedSpaceAnnounceEpochField, a.epoch)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.NestedStructureMarshal(usedSpaceAnnounceCIDField, a.cid)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.UInt64Marshal(usedSpaceAnnounceUsedSpaceField, a.usedSpace)
	return offset + n, err
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

	n, err = protoutil.UInt64Marshal(usedSpaceAnnounceEpochField, buf[offset:], a.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

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

	size += protoutil.UInt64Size(usedSpaceAnnounceEpochField, a.epoch)
	size += protoutil.NestedStructureSize(usedSpaceAnnounceCIDField, a.cid)
	size += protoutil.UInt64Size(usedSpaceAnnounceUsedSpaceField, a.usedSpace)

	return size
}

func (a *UsedSpaceAnnouncement) Unmarshal(data []byte) error {
	return message.Unmarshal(a, data, new(container.AnnounceUsedSpaceRequest_Body_Announcement))
}

func (r *AnnounceUsedSpaceRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if r == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	for i := range r.announcements {
		n, err = s.NestedStructureMarshal(usedSpaceReqBodyAnnouncementsField, r.announcements[i])
		if err != nil {
			return offset + n, err
		}

		offset += n
	}

	return offset, nil
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

func (r *AnnounceUsedSpaceRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(r, data, new(container.AnnounceUsedSpaceRequest_Body))
}

func (r *AnnounceUsedSpaceResponseBody) MarshalStream(protoutil.Stream) (int, error) {
	return 0, nil
}

func (r *AnnounceUsedSpaceResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (r *AnnounceUsedSpaceResponseBody) StableSize() (size int) {
	return 0
}

func (r *AnnounceUsedSpaceResponseBody) Unmarshal([]byte) error {
	return nil
}
