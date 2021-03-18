/*
This package contains help functions for stable marshaller. Their usage is
totally optional. One can implement fast stable marshaller without these
runtime function calls.
*/

package proto

import (
	"encoding/binary"
	"io"
	"reflect"
)

type (
	stream struct {
		buf [10]byte
		w   io.Writer
	}

	// StreamMarshaler represents streaming marshaling interface.
	StreamMarshaler interface {
		MarshalStream(Stream) (int, error)
		StableSize() int
	}
)

var _ Stream = (*stream)(nil)

func NewStream(w io.Writer) Stream {
	return &stream{w: w}
}

// WriteUvarint implements Stream interface.
func (s *stream) WriteUvarint(v uint64) (int, error) {
	n := binary.PutUvarint(s.buf[:], v)
	return s.w.Write(s.buf[:n])
}

// BytesMarshal implements Stream interface.
func (s *stream) BytesMarshal(field int, v []byte) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	prefix := field<<3 | 0x2

	n, err = s.WriteUvarint(uint64(prefix))
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.WriteUvarint(uint64(len(v)))
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.w.Write(v)
	return offset + n, err
}

// StringMarshal implements Stream interface.
func (s *stream) StringMarshal(field int, v string) (int, error) {
	return s.BytesMarshal(field, []byte(v))
}

// BoolMarshal implements Stream interface.
func (s *stream) BoolMarshal(field int, v bool) (int, error) {
	if !v {
		return 0, nil
	}

	prefix := field << 3

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	n, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return n, err
	}

	offset, err := s.w.Write([]byte{0x1})
	return offset + n, err
}

// UInt64Marshal implements Stream interface.
func (s *stream) UInt64Marshal(field int, v uint64) (int, error) {
	if v == 0 {
		return 0, nil
	}

	prefix := field << 3

	n, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return n, err
	}

	offset, err := s.WriteUvarint(v)
	return offset + n, err
}

// UInt64Size implements Stream interface.
func (s *stream) UInt64Size(field int, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := field << 3

	return VarUIntSize(uint64(prefix)) + VarUIntSize(v)
}

// Int64Marshal implements Stream interface.
func (s *stream) Int64Marshal(field int, v int64) (int, error) {
	return s.UInt64Marshal(field, uint64(v))
}

// UInt32Marshal implements Stream interface.
func (s *stream) UInt32Marshal(field int, v uint32) (int, error) {
	return s.UInt64Marshal(field, uint64(v))
}

// Int32Marshal implements Stream interface.
func (s *stream) Int32Marshal(field int, v int32) (int, error) {
	return s.UInt64Marshal(field, uint64(v))
}

// EnumMarshal implements Stream interface.
func (s *stream) EnumMarshal(field int, v int32) (int, error) {
	return s.UInt64Marshal(field, uint64(v))
}

// RepeatedBytesMarshal implements Stream interface.
func (s *stream) RepeatedBytesMarshal(field int, v [][]byte) (int, error) {
	var offset int

	for i := range v {
		off, err := s.BytesMarshal(field, v[i])
		offset += off
		if err != nil {
			return offset, err
		}
	}
	return offset, nil
}

// RepeatedStringMarshal implements Stream interface.
func (s *stream) RepeatedStringMarshal(field int, v []string) (int, error) {
	var offset int

	for i := range v {
		off, err := s.StringMarshal(field, v[i])
		offset += off
		if err != nil {
			return offset, err
		}

	}

	return offset, nil
}

// RepeatedStringSize implements Stream interface.
func (s *stream) RepeatedStringSize(field int, v []string) (size int) {
	for i := range v {
		size += StringSize(field, v[i])
	}

	return size
}

// RepeatedUInt64Marshal implements Stream interface.
func (s *stream) RepeatedUInt64Marshal(field int, v []uint64) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x02
	offset, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return offset, err
	}

	_, arrSize := RepeatedUInt64Size(field, v)
	off, err := s.WriteUvarint(uint64(arrSize))
	offset += off
	if err != nil {
		return offset, err
	}

	for i := range v {
		off, err := s.WriteUvarint(v[i])
		offset += off
		if err != nil {
			return offset, err
		}
	}

	return offset, nil
}

// RepeatedInt64Marshal implements Stream interface.
func (s *stream) RepeatedInt64Marshal(field int, v []int64) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x02
	offset, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return offset, err
	}

	_, arrSize := RepeatedInt64Size(field, v)
	off, err := s.WriteUvarint(uint64(arrSize))
	offset += off
	if err != nil {
		return offset, err
	}

	for i := range v {
		off, err := s.WriteUvarint(uint64(v[i]))
		offset += off
		if err != nil {
			return offset, err
		}
	}

	return offset, nil
}

// RepeatedUInt32Marshal implements Stream interface.
func (s *stream) RepeatedUInt32Marshal(field int, v []uint32) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x02
	offset, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return offset, err
	}

	_, arrSize := RepeatedUInt32Size(field, v)
	off, err := s.WriteUvarint(uint64(arrSize))
	offset += off
	if err != nil {
		return offset, err
	}

	for i := range v {
		off, err := s.WriteUvarint(uint64(v[i]))
		offset += off
		if err != nil {
			return offset, err
		}
	}

	return offset, nil
}

// RepeatedInt32Marshal implements Stream interface.
func (s *stream) RepeatedInt32Marshal(field int, v []int32) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x02
	offset, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return offset, err
	}

	_, arrSize := RepeatedInt32Size(field, v)
	off, err := s.WriteUvarint(uint64(arrSize))
	offset += off
	if err != nil {
		return offset, err
	}

	for i := range v {
		off, err := s.WriteUvarint(uint64(v[i]))
		offset += off
		if err != nil {
			return offset, err
		}
	}

	return offset, nil
}

// NestedStructureMarshal implements Stream interface.
func (s *stream) NestedStructureMarshal(field int64, v StreamMarshaler) (int, error) {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	prefix, _ := NestedStructurePrefix(field)

	n, err = s.WriteUvarint(prefix)
	if err != nil {
		return n, err
	}

	offset += n

	n, err = s.WriteUvarint(uint64(v.StableSize()))
	if err != nil {
		return offset + n, err
	}

	offset += n
	n, err = v.MarshalStream(s)
	return offset + n, nil
}

// Fixed64Marshal implements Stream interface.
func (s *stream) Fixed64Marshal(field int, v uint64) (int, error) {
	if v == 0 {
		return 0, nil
	}

	prefix := field<<3 | 1

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	n, err := s.WriteUvarint(uint64(prefix))
	if err != nil {
		return n, err
	}

	binary.LittleEndian.PutUint64(s.buf[:], v)
	offset, err := s.w.Write(s.buf[:8])
	return offset + n, err
}
