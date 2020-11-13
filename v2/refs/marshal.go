package refs

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	goproto "google.golang.org/protobuf/proto"
)

const (
	ownerIDValField = 1

	containerIDValField = 1

	objectIDValField = 1

	addressContainerField = 1
	addressObjectField    = 2

	checksumTypeField  = 1
	checksumValueField = 2

	signatureKeyField   = 1
	signatureValueField = 2

	versionMajorField = 1
	versionMinorField = 2
)

func (o *OwnerID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(ownerIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *OwnerID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(ownerIDValField, o.val)
}

func (o *OwnerID) Unmarshal(data []byte) error {
	m := new(refs.OwnerID)
	if err := goproto.Unmarshal(data, m); err != nil {
		return err
	}

	*o = *OwnerIDFromGRPCMessage(m)

	return nil
}

func (c *ContainerID) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	_, err := proto.BytesMarshal(containerIDValField, buf, c.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *ContainerID) StableSize() int {
	if c == nil {
		return 0
	}

	return proto.BytesSize(containerIDValField, c.val)
}

func (c *ContainerID) Unmarshal(data []byte) error {
	m := new(refs.ContainerID)
	if err := goproto.Unmarshal(data, m); err != nil {
		return err
	}

	*c = *ContainerIDFromGRPCMessage(m)

	return nil
}

func (o *ObjectID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(objectIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *ObjectID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(objectIDValField, o.val)
}

func (o *ObjectID) Unmarshal(data []byte) error {
	m := new(refs.ObjectID)
	if err := goproto.Unmarshal(data, m); err != nil {
		return err
	}

	*o = *ObjectIDFromGRPCMessage(m)

	return nil
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
		err       error
	)

	n, err = proto.NestedStructureMarshal(addressContainerField, buf[offset:], a.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(addressObjectField, buf[offset:], a.oid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (a *Address) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.NestedStructureSize(addressContainerField, a.cid)

	size += proto.NestedStructureSize(addressObjectField, a.oid)

	return size
}

func (a *Address) Unmarshal(data []byte) error {
	addrGRPC := new(refs.Address)
	if err := goproto.Unmarshal(data, addrGRPC); err != nil {
		return err
	}

	*a = *AddressFromGRPCMessage(addrGRPC)

	return nil
}

func (c *Checksum) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.EnumMarshal(checksumTypeField, buf[offset:], int32(c.typ))
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.BytesMarshal(checksumValueField, buf[offset:], c.sum)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Checksum) StableSize() (size int) {
	if c == nil {
		return 0
	}

	size += proto.EnumSize(checksumTypeField, int32(c.typ))
	size += proto.BytesSize(checksumValueField, c.sum)

	return size
}

func (s *Signature) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.BytesMarshal(signatureKeyField, buf[offset:], s.key)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.BytesMarshal(signatureValueField, buf[offset:], s.sign)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Signature) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.BytesSize(signatureKeyField, s.key)
	size += proto.BytesSize(signatureValueField, s.sign)

	return size
}

func (v *Version) StableMarshal(buf []byte) ([]byte, error) {
	if v == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, v.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt32Marshal(versionMajorField, buf[offset:], v.major)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.UInt32Marshal(versionMinorField, buf[offset:], v.minor)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (v *Version) StableSize() (size int) {
	if v == nil {
		return 0
	}

	size += proto.UInt32Size(versionMajorField, v.major)
	size += proto.UInt32Size(versionMinorField, v.minor)

	return size
}

func (v *Version) Unmarshal(data []byte) error {
	m := new(refs.Version)
	if err := goproto.Unmarshal(data, m); err != nil {
		return err
	}

	*v = *VersionFromGRPCMessage(m)

	return nil
}
