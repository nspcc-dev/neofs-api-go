package acl

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	FilterHeaderTypeField = 1
	FilterMatchTypeField  = 2
	FilterNameField       = 3
	FilterValueField      = 4
)

func (t *Table) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (t *Table) StableSize() int {
	panic("not implemented")
}

func (r *Record) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (r *Record) StableSize() int {
	panic("not implemented")
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

func (t *HeaderType) StableMarshal(buf []byte) ([]byte, error) {
	panic("not implemented")
}

func (t *HeaderType) StableSize() int {
	panic("not implemented")
}
