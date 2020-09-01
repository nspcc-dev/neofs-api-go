package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	keyAttributeField   = 1
	valueAttributeField = 2

	keyNodeInfoField        = 1
	addressNodeInfoField    = 2
	attributesNodeInfoField = 3
)

func (p *PlacementPolicy) StableMarshal(buf []byte) ([]byte, error) {
	// todo: implement me
	return nil, nil
}

func (p *PlacementPolicy) StableSize() (size int) {
	// todo: implement me
	return 0
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

	n, err = proto.StringMarshal(keyAttributeField, buf[offset:], a.key)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(valueAttributeField, buf[offset:], a.value)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (a *Attribute) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.StringSize(keyAttributeField, a.key)
	size += proto.StringSize(valueAttributeField, a.value)

	return size
}

func (ni *NodeInfo) StableMarshal(buf []byte) ([]byte, error) {
	if ni == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, ni.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.BytesMarshal(keyNodeInfoField, buf[offset:], ni.publicKey)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(addressNodeInfoField, buf[offset:], ni.address)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range ni.attributes {
		n, err = proto.NestedStructureMarshal(attributesNodeInfoField, buf[offset:], ni.attributes[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (ni *NodeInfo) StableSize() (size int) {
	if ni == nil {
		return 0
	}

	size += proto.BytesSize(keyNodeInfoField, ni.publicKey)
	size += proto.StringSize(addressNodeInfoField, ni.address)
	for i := range ni.attributes {
		size += proto.NestedStructureSize(attributesNodeInfoField, ni.attributes[i])
	}

	return size
}
