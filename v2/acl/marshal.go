package acl

import (
	"encoding/binary"

	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	FilterHeaderTypeField = 1
	FilterMatchTypeField  = 2
	FilterNameField       = 3
	FilterValueField      = 4

	TargetTypeField = 1
	TargetKeysField = 2

	RecordOperationField = 1
	RecordActionField    = 2
	RecordFiltersField   = 3
	RecordTargetsField   = 4
)

func (t *Table) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (t *Table) StableSize() int {
	panic("not implemented")
}

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

	n, err = proto.EnumMarshal(RecordOperationField, buf, int32(r.op))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(RecordActionField, buf[offset:], int32(r.action))
	if err != nil {
		return nil, err
	}

	offset += n

	prefix, _ = proto.NestedStructurePrefix(RecordFiltersField)

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

	prefix, _ = proto.NestedStructurePrefix(RecordTargetsField)

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

func (r *Record) StableSize() (size int) {
	if r == nil {
		return 0
	}

	size += proto.EnumSize(RecordOperationField, int32(r.op))
	size += proto.EnumSize(RecordActionField, int32(r.op))

	_, ln := proto.NestedStructurePrefix(RecordFiltersField)

	for i := range r.filters {
		n := r.filters[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	_, ln = proto.NestedStructurePrefix(RecordTargetsField)

	for i := range r.targets {
		n := r.targets[i].StableSize()
		size += ln + proto.VarUIntSize(uint64(n)) + n
	}

	return size
}

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

	n, err = proto.EnumMarshal(FilterHeaderTypeField, buf, int32(f.hdrType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.EnumMarshal(FilterMatchTypeField, buf[offset:], int32(f.matchType))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(FilterNameField, buf[offset:], f.name)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.StringMarshal(FilterValueField, buf[offset:], f.value)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (f *HeaderFilter) StableSize() (size int) {
	if f == nil {
		return 0
	}

	size += proto.EnumSize(FilterHeaderTypeField, int32(f.hdrType))
	size += proto.EnumSize(FilterMatchTypeField, int32(f.matchType))
	size += proto.StringSize(FilterNameField, f.name)
	size += proto.StringSize(FilterValueField, f.value)

	return size
}

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

	n, err = proto.EnumMarshal(TargetTypeField, buf, int32(t.target))
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.RepeatedBytesMarshal(TargetKeysField, buf[offset:], t.keys)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (t *TargetInfo) StableSize() (size int) {
	if t == nil {
		return 0
	}

	size += proto.EnumSize(TargetTypeField, int32(t.target))
	size += proto.RepeatedBytesSize(TargetKeysField, t.keys)

	return size
}
