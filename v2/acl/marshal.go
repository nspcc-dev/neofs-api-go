package acl

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	filterHeaderTypeField = 1
	filterMatchTypeField  = 2
	filterNameField       = 3
	filterValueField      = 4

	targetTypeField = 1
	targetKeysField = 2

	recordOperationField = 1
	recordActionField    = 2
	recordFiltersField   = 3
	recordTargetsField   = 4

	tableContainerIDField = 1
	tableRecordsField     = 2
)

// StableMarshal marshals unified acl table structure in a protobuf
// compatible way without field order shuffle.
func (t *Table) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(tableContainerIDField, buf[offset:], t.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range t.records {
		n, err = proto.NestedStructureMarshal(tableRecordsField, buf[offset:], t.records[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of acl table structure marshalled by StableMarshal function.
func (t *Table) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.NestedStructureSize(tableContainerIDField, t.cid)

	for i := range t.records {
		size += proto.NestedStructureSize(tableRecordsField, t.records[i])
	}

	return size
}

// StableMarshal marshals unified acl record structure in a protobuf
// compatible way without field order shuffle.
func (r *Record) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.EnumMarshal(recordOperationField, buf[offset:], int32(r.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(recordActionField, buf[offset:], int32(r.action))
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range r.filters {
		n, err = proto.NestedStructureMarshal(recordFiltersField, buf[offset:], r.filters[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	for i := range r.targets {
		n, err = proto.NestedStructureMarshal(recordTargetsField, buf[offset:], r.targets[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of acl record structure marshalled by StableMarshal function.
func (r *Record) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.EnumSize(recordOperationField, int32(r.op))
	size += proto.EnumSize(recordActionField, int32(r.op))

	for i := range r.filters {
		size += proto.NestedStructureSize(recordFiltersField, r.filters[i])
	}

	for i := range r.targets {
		size += proto.NestedStructureSize(recordTargetsField, r.targets[i])
	}

	return size
}

// StableMarshal marshals unified header filter structure in a protobuf
// compatible way without field order shuffle.
func (f *HeaderFilter) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = proto.EnumMarshal(filterHeaderTypeField, buf[offset:], int32(f.hdrType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(filterMatchTypeField, buf[offset:], int32(f.matchType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(filterNameField, buf[offset:], f.name)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.StringMarshal(filterValueField, buf[offset:], f.value)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize of header filter structure marshalled by StableMarshal function.
func (f *HeaderFilter) StableSize() (size int) {
	if f == nil {
		return 0
	}

	size += proto.EnumSize(filterHeaderTypeField, int32(f.hdrType))
	size += proto.EnumSize(filterMatchTypeField, int32(f.matchType))
	size += proto.StringSize(filterNameField, f.name)
	size += proto.StringSize(filterValueField, f.value)

	return size
}

// StableMarshal marshals unified target info structure in a protobuf
// compatible way without field order shuffle.
func (t *TargetInfo) StableMarshal(buf []byte) ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, t.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.EnumMarshal(targetTypeField, buf[offset:], int32(t.target))
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.RepeatedBytesMarshal(targetKeysField, buf[offset:], t.keys)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize of target info structure marshalled by StableMarshal function.
func (t *TargetInfo) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.EnumSize(targetTypeField, int32(t.target))
	size += proto.RepeatedBytesSize(targetKeysField, t.keys)

	return size
}
