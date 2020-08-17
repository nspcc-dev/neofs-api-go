package acl

import (
	"encoding/binary"

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
		prefix    uint64
		err       error
	)

	if t.cid != nil {
		prefix, _ = proto.NestedStructurePrefix(tableContainerIDField)
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.cid.StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.cid.StableMarshal(buf[offset:])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	prefix, _ = proto.NestedStructurePrefix(tableRecordsField)

	for i := range t.records {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = t.records[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = t.records[i].StableMarshal(buf[offset:])
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

	if t.cid != nil {
		_, ln := proto.NestedStructurePrefix(tableContainerIDField)
		n := t.cid.StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	_, ln := proto.NestedStructurePrefix(tableRecordsField)

	for i := range t.records {
		n := t.records[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
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
		prefix    uint64
		err       error
	)

	n, err = proto.EnumMarshal(recordOperationField, buf, int32(r.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(recordActionField, buf[offset:], int32(r.action))
	if err != nil {
		return nil, err
	}

	offset += n

	prefix, _ = proto.NestedStructurePrefix(recordFiltersField)

	for i := range r.filters {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.filters[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.filters[i].StableMarshal(buf[offset:])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	prefix, _ = proto.NestedStructurePrefix(recordTargetsField)

	for i := range r.targets {
		offset += binary.PutUvarint(buf[offset:], prefix)

		n = r.targets[i].StableSize()
		offset += binary.PutUvarint(buf[offset:], uint64(n))

		_, err = r.targets[i].StableMarshal(buf[offset:])
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

	_, ln := proto.NestedStructurePrefix(recordFiltersField)

	for i := range r.filters {
		n := r.filters[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	_, ln = proto.NestedStructurePrefix(recordTargetsField)

	for i := range r.targets {
		n := r.targets[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
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

	n, err = proto.EnumMarshal(filterHeaderTypeField, buf, int32(f.hdrType))
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

	n, err = proto.EnumMarshal(targetTypeField, buf, int32(t.target))
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
