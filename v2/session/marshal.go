package session

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	createReqBodyOwnerField    = 1
	createReqBodyLifetimeField = 2

	createRespBodyIDField  = 1
	createRespBodyKeyField = 2
)

func (c *CreateRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.NestedStructureMarshal(createReqBodyOwnerField, buf[offset:], c.ownerID)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(createReqBodyLifetimeField, buf[offset:], c.lifetime)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *CreateRequestBody) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.NestedStructureSize(createReqBodyOwnerField, c.ownerID)
	size += proto.NestedStructureSize(createReqBodyLifetimeField, c.lifetime)

	return size
}

func (c *CreateResponseBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.BytesMarshal(createRespBodyIDField, buf[offset:], c.id)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.BytesMarshal(createRespBodyKeyField, buf[offset:], c.sessionKey)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *CreateResponseBody) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.BytesSize(createRespBodyIDField, c.id)
	size += proto.BytesSize(createRespBodyKeyField, c.sessionKey)

	return size
}
