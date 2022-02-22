package refs

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

const (
	ownerIDValField = 1

	containerIDValField = 1

	objectIDValField = 1

	addressContainerField = 1
	addressObjectField    = 2

	checksumTypeField  = 1
	checksumValueField = 2

	signatureKeyField    = 1
	signatureValueField  = 2
	signatureSchemeField = 3

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
	return message.Unmarshal(o, data, new(refs.OwnerID))
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
	return message.Unmarshal(c, data, new(refs.ContainerID))
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

// ObjectIDNestedListSize returns byte length of nested
// repeated ObjectID field with fNum number.
func ObjectIDNestedListSize(fNum int64, ids []*ObjectID) (sz int) {
	for i := range ids {
		sz += proto.NestedStructureSize(fNum, ids[i])
	}

	return
}

func (o *ObjectID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(objectIDValField, o.val)
}

// ObjectIDNestedListMarshal writes protobuf repeated ObjectID field
// with fNum number to buf.
func ObjectIDNestedListMarshal(fNum int64, buf []byte, ids []*ObjectID) (off int, err error) {
	for i := range ids {
		var n int

		n, err = proto.NestedStructureMarshal(fNum, buf[off:], ids[i])
		if err != nil {
			return
		}

		off += n
	}

	return
}

func (o *ObjectID) Unmarshal(data []byte) error {
	return message.Unmarshal(o, data, new(refs.ObjectID))
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
	return message.Unmarshal(a, data, new(refs.Address))
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

func (c *Checksum) Unmarshal(data []byte) error {
	return message.Unmarshal(c, data, new(refs.Checksum))
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

	n, err = proto.BytesMarshal(signatureValueField, buf[offset:], s.sign)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.EnumMarshal(signatureSchemeField, buf[offset:], int32(s.scheme))
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
	size += proto.EnumSize(signatureSchemeField, int32(s.scheme))

	return size
}

func (s *Signature) Unmarshal(data []byte) error {
	return message.Unmarshal(s, data, new(refs.Signature))
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
	return message.Unmarshal(v, data, new(refs.Version))
}

// SubnetID message field numbers
const (
	_ = iota
	subnetIDValFNum
)

// StableMarshal marshals SubnetID to NeoFS API V2 binary format (Protocol Buffers with direct field order).
//
// Returns a slice of recorded data. Data is written to the provided buffer if there is enough space.
func (s *SubnetID) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	proto.Fixed32Marshal(subnetIDValFNum, buf, s.value)

	return buf, nil
}

// StableSize returns the number of bytes required to write SubnetID in NeoFS API V2 binary format (see StableMarshal).
func (s *SubnetID) StableSize() (size int) {
	if s != nil {
		size += proto.Fixed32Size(subnetIDValFNum, s.value)
	}

	return
}

// Unmarshal unmarshals SubnetID from NeoFS API V2 binary format (see StableMarshal).
// Must not be called on nil.
//
// Note: empty data corresponds to zero ID value or nil pointer to it.
func (s *SubnetID) Unmarshal(data []byte) error {
	return message.Unmarshal(s, data, new(refs.SubnetID))
}
