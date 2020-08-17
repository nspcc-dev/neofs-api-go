package refs

import (
	"encoding/binary"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	OwnerIDValField = 1

	ContainerIDValField = 1

	ObjectIDValField = 1

	AddressContainerField = 1
	AddressObjectField    = 2
)

func (o *OwnerID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(OwnerIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *OwnerID) StableSize() int {
	if o == nil {
		return 0
	}
	return proto.BytesSize(OwnerIDValField, o.val)
}

func (c *ContainerID) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	_, err := proto.BytesMarshal(ContainerIDValField, buf, c.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *ContainerID) StableSize() int {
	if c == nil {
		return 0
	}
	return proto.BytesSize(ContainerIDValField, c.val)
}

func (o *ObjectID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(ObjectIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *ObjectID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(ObjectIDValField, o.val)
}

func (a *Address) StableMarshal(buf []byte) ([]byte, error) {
	if a == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, a.StableSize())
	}

	var (
		offset, n int
		prefix    uint64
		err       error
	)

	if a.cid != nil {
		prefix, _ = proto.NestedStructurePrefix(AddressContainerField)
		offset = binary.PutUvarint(buf, prefix)

		n = a.cid.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = a.cid.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	if a.oid != nil {
		prefix, _ = proto.NestedStructurePrefix(AddressObjectField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = a.oid.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = a.oid.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func (a *Address) StableSize() (size int) {
	if a == nil {
		return 0
	}

	if a.cid != nil {
		_, ln := proto.NestedStructurePrefix(AddressContainerField)
		n := a.cid.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	if a.oid != nil {
		_, ln := proto.NestedStructurePrefix(AddressObjectField)
		n := a.oid.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}
