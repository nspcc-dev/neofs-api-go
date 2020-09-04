package netmap

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	nameFilterField    = 1
	keyFilterField     = 2
	opFilterField      = 3
	valueFilterField   = 4
	filtersFilterField = 5

	nameSelectorField      = 1
	countSelectorField     = 2
	attributeSelectorField = 3
	filterSelectorField    = 4

	countReplicaField    = 1
	selectorReplicaField = 2

	replicasPolicyField  = 1
	backupPolicyField    = 2
	selectorsPolicyField = 3
	filtersPolicyField   = 4

	keyAttributeField     = 1
	valueAttributeField   = 2
	parentsAttributeField = 3

	keyNodeInfoField        = 1
	addressNodeInfoField    = 2
	attributesNodeInfoField = 3
	stateNodeInfoField      = 4
)

func (f *Filter) StableMarshal(buf []byte) ([]byte, error) {
	if f == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, f.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.StringMarshal(nameFilterField, buf[offset:], f.name)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(keyFilterField, buf[offset:], f.key)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(opFilterField, buf[offset:], int32(f.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(valueFilterField, buf[offset:], f.value)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range f.filters {
		n, err = proto.NestedStructureMarshal(filtersFilterField, buf[offset:], f.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (f *Filter) StableSize() (size int) {
	size += proto.StringSize(nameFilterField, f.name)
	size += proto.StringSize(keyFilterField, f.key)
	size += proto.EnumSize(opFilterField, int32(f.op))
	size += proto.StringSize(valueFilterField, f.value)
	for i := range f.filters {
		size += proto.NestedStructureSize(filtersFilterField, f.filters[i])
	}

	return size
}

func (s *Selector) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.StringMarshal(nameSelectorField, buf[offset:], s.name)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(countSelectorField, buf[offset:], s.count)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(attributeSelectorField, buf[offset:], s.attribute)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(filterSelectorField, buf[offset:], s.filter)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Selector) StableSize() (size int) {
	size += proto.StringSize(nameSelectorField, s.name)
	size += proto.UInt32Size(countSelectorField, s.count)
	size += proto.StringSize(attributeSelectorField, s.attribute)
	size += proto.StringSize(filterSelectorField, s.filter)

	return size
}

func (r *Replica) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.UInt32Marshal(countReplicaField, buf[offset:], r.count)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(selectorReplicaField, buf[offset:], r.selector)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *Replica) StableSize() (size int) {
	size += proto.UInt32Size(countReplicaField, r.count)
	size += proto.StringSize(selectorReplicaField, r.selector)

	return size
}

func (p *PlacementPolicy) StableMarshal(buf []byte) ([]byte, error) {
	if p == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, p.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	for i := range p.replicas {
		n, err = proto.NestedStructureMarshal(replicasPolicyField, buf[offset:], p.replicas[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	n, err = proto.UInt32Marshal(backupPolicyField, buf[offset:], p.backupFactor)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range p.selectors {
		n, err = proto.NestedStructureMarshal(selectorsPolicyField, buf[offset:], p.selectors[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	for i := range p.filters {
		n, err = proto.NestedStructureMarshal(filtersPolicyField, buf[offset:], p.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (p *PlacementPolicy) StableSize() (size int) {
	for i := range p.replicas {
		size += proto.NestedStructureSize(replicasPolicyField, p.replicas[i])
	}

	size += proto.UInt32Size(backupPolicyField, p.backupFactor)

	for i := range p.selectors {
		size += proto.NestedStructureSize(selectorsPolicyField, p.selectors[i])
	}

	for i := range p.filters {
		size += proto.NestedStructureSize(filtersPolicyField, p.filters[i])
	}

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

	n, err = proto.StringMarshal(keyAttributeField, buf[offset:], a.key)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(valueAttributeField, buf[offset:], a.value)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range a.parents {
		n, err = proto.StringMarshal(parentsAttributeField, buf[offset:], a.parents[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (a *Attribute) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.StringSize(keyAttributeField, a.key)
	size += proto.StringSize(valueAttributeField, a.value)

	for i := range a.parents {
		size += proto.StringSize(parentsAttributeField, a.parents[i])
	}

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

	n, err = proto.EnumMarshal(stateNodeInfoField, buf[offset:], int32(ni.state))
	if err != nil {
		return nil, err
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

	size += proto.EnumSize(stateNodeInfoField, int32(ni.state))

	return size
}
