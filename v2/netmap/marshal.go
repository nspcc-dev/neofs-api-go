package netmap

import (
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	netmap "github.com/nspcc-dev/neofs-api-go/v2/netmap/grpc"
	"google.golang.org/protobuf/proto"
)

const (
	nameFilterField    = 1
	keyFilterField     = 2
	opFilterField      = 3
	valueFilterField   = 4
	filtersFilterField = 5

	nameSelectorField      = 1
	countSelectorField     = 2
	clauseSelectorField    = 3
	attributeSelectorField = 4
	filterSelectorField    = 5

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

	versionInfoResponseBodyField = 1
	nodeInfoResponseBodyField    = 2
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

	n, err = protoutil.StringMarshal(nameFilterField, buf[offset:], f.name)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(keyFilterField, buf[offset:], f.key)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.EnumMarshal(opFilterField, buf[offset:], int32(f.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(valueFilterField, buf[offset:], f.value)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range f.filters {
		n, err = protoutil.NestedStructureMarshal(filtersFilterField, buf[offset:], f.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (f *Filter) StableSize() (size int) {
	size += protoutil.StringSize(nameFilterField, f.name)
	size += protoutil.StringSize(keyFilterField, f.key)
	size += protoutil.EnumSize(opFilterField, int32(f.op))
	size += protoutil.StringSize(valueFilterField, f.value)
	for i := range f.filters {
		size += protoutil.NestedStructureSize(filtersFilterField, f.filters[i])
	}

	return size
}

func (f *Filter) Unmarshal(data []byte) error {
	m := new(netmap.Filter)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*f = *FilterFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.StringMarshal(nameSelectorField, buf[offset:], s.name)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.UInt32Marshal(countSelectorField, buf[offset:], s.count)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.EnumMarshal(clauseSelectorField, buf[offset:], int32(s.clause))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(attributeSelectorField, buf[offset:], s.attribute)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.StringMarshal(filterSelectorField, buf[offset:], s.filter)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *Selector) StableSize() (size int) {
	size += protoutil.StringSize(nameSelectorField, s.name)
	size += protoutil.UInt32Size(countSelectorField, s.count)
	size += protoutil.EnumSize(countSelectorField, int32(s.clause))
	size += protoutil.StringSize(attributeSelectorField, s.attribute)
	size += protoutil.StringSize(filterSelectorField, s.filter)

	return size
}

func (s *Selector) Unmarshal(data []byte) error {
	m := new(netmap.Selector)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*s = *SelectorFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.UInt32Marshal(countReplicaField, buf[offset:], r.count)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.StringMarshal(selectorReplicaField, buf[offset:], r.selector)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *Replica) StableSize() (size int) {
	size += protoutil.UInt32Size(countReplicaField, r.count)
	size += protoutil.StringSize(selectorReplicaField, r.selector)

	return size
}

func (r *Replica) Unmarshal(data []byte) error {
	m := new(netmap.Replica)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*r = *ReplicaFromGRPCMessage(m)

	return nil
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
		n, err = protoutil.NestedStructureMarshal(replicasPolicyField, buf[offset:], p.replicas[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	n, err = protoutil.UInt32Marshal(backupPolicyField, buf[offset:], p.backupFactor)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range p.selectors {
		n, err = protoutil.NestedStructureMarshal(selectorsPolicyField, buf[offset:], p.selectors[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	for i := range p.filters {
		n, err = protoutil.NestedStructureMarshal(filtersPolicyField, buf[offset:], p.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (p *PlacementPolicy) StableSize() (size int) {
	for i := range p.replicas {
		size += protoutil.NestedStructureSize(replicasPolicyField, p.replicas[i])
	}

	size += protoutil.UInt32Size(backupPolicyField, p.backupFactor)

	for i := range p.selectors {
		size += protoutil.NestedStructureSize(selectorsPolicyField, p.selectors[i])
	}

	for i := range p.filters {
		size += protoutil.NestedStructureSize(filtersPolicyField, p.filters[i])
	}

	return size
}

func (p *PlacementPolicy) Unmarshal(data []byte) error {
	m := new(netmap.PlacementPolicy)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*p = *PlacementPolicyFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.StringMarshal(keyAttributeField, buf[offset:], a.key)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(valueAttributeField, buf[offset:], a.value)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range a.parents {
		n, err = protoutil.StringMarshal(parentsAttributeField, buf[offset:], a.parents[i])
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

	size += protoutil.StringSize(keyAttributeField, a.key)
	size += protoutil.StringSize(valueAttributeField, a.value)

	for i := range a.parents {
		size += protoutil.StringSize(parentsAttributeField, a.parents[i])
	}

	return size
}

func (a *Attribute) Unmarshal(data []byte) error {
	m := new(netmap.NodeInfo_Attribute)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*a = *AttributeFromGRPCMessage(m)

	return nil
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

	n, err = protoutil.BytesMarshal(keyNodeInfoField, buf[offset:], ni.publicKey)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.StringMarshal(addressNodeInfoField, buf[offset:], ni.address)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range ni.attributes {
		n, err = protoutil.NestedStructureMarshal(attributesNodeInfoField, buf[offset:], ni.attributes[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	_, err = protoutil.EnumMarshal(stateNodeInfoField, buf[offset:], int32(ni.state))
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (ni *NodeInfo) StableSize() (size int) {
	if ni == nil {
		return 0
	}

	size += protoutil.BytesSize(keyNodeInfoField, ni.publicKey)
	size += protoutil.StringSize(addressNodeInfoField, ni.address)
	for i := range ni.attributes {
		size += protoutil.NestedStructureSize(attributesNodeInfoField, ni.attributes[i])
	}

	size += protoutil.EnumSize(stateNodeInfoField, int32(ni.state))

	return size
}

func (ni *NodeInfo) Unmarshal(data []byte) error {
	m := new(netmap.NodeInfo)
	if err := proto.Unmarshal(data, m); err != nil {
		return err
	}

	*ni = *NodeInfoFromGRPCMessage(m)

	return nil
}

func (l *LocalNodeInfoRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	return nil, nil
}

func (l *LocalNodeInfoRequestBody) StableSize() (size int) {
	return 0
}

func (l *LocalNodeInfoResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if l == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, l.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.NestedStructureMarshal(versionInfoResponseBodyField, buf[offset:], l.version)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(nodeInfoResponseBodyField, buf[offset:], l.nodeInfo)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (l *LocalNodeInfoResponseBody) StableSize() (size int) {
	if l == nil {
		return 0
	}

	size += protoutil.NestedStructureSize(versionInfoResponseBodyField, l.version)
	size += protoutil.NestedStructureSize(nodeInfoResponseBodyField, l.nodeInfo)

	return size
}
