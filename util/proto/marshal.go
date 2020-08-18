/*
This package contains help functions for stable marshaller. Their usage is
totally optional. One can implement fast stable marshaller without these
runtime function calls.
*/

package proto

import (
	"encoding/binary"
	"math/bits"
)

type (
	stableMarshaller interface {
		StableMarshal([]byte) ([]byte, error)
		StableSize() int
	}
)

func BytesMarshal(field int, buf, v []byte) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x2

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	i += binary.PutUvarint(buf[i:], uint64(len(v)))
	i += copy(buf[i:], v)

	return i, nil
}

func BytesSize(field int, v []byte) int {
	ln := len(v)
	if ln == 0 {
		return 0
	}

	prefix := field<<3 | 0x2

	return VarUIntSize(uint64(prefix)) + VarUIntSize(uint64(ln)) + ln
}

func StringMarshal(field int, buf []byte, v string) (int, error) {
	return BytesMarshal(field, buf, []byte(v))
}

func StringSize(field int, v string) int {
	return BytesSize(field, []byte(v))
}

func BoolMarshal(field int, buf []byte, v bool) (int, error) {
	if !v {
		return 0, nil
	}

	prefix := field << 3

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	buf[i] = 0x1

	return i + 1, nil
}

func BoolSize(field int, v bool) int {
	if !v {
		return 0
	}

	prefix := field << 3

	return VarUIntSize(uint64(prefix)) + 1 // bool is always 1 byte long
}

func UInt64Marshal(field int, buf []byte, v uint64) (int, error) {
	if v == 0 {
		return 0, nil
	}

	prefix := field << 3

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	i += binary.PutUvarint(buf[i:], v)

	return i, nil
}

func UInt64Size(field int, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := field << 3

	return VarUIntSize(uint64(prefix)) + VarUIntSize(v)
}

func Int64Marshal(field int, buf []byte, v int64) (int, error) {
	return UInt64Marshal(field, buf, uint64(v))
}

func Int64Size(field int, v int64) int {
	return UInt64Size(field, uint64(v))
}

func UInt32Marshal(field int, buf []byte, v uint32) (int, error) {
	return UInt64Marshal(field, buf, uint64(v))
}

func UInt32Size(field int, v uint32) int {
	return UInt64Size(field, uint64(v))
}

func Int32Marshal(field int, buf []byte, v int32) (int, error) {
	return UInt64Marshal(field, buf, uint64(v))
}

func Int32Size(field int, v int32) int {
	return UInt64Size(field, uint64(v))
}

func EnumMarshal(field int, buf []byte, v int32) (int, error) {
	return UInt64Marshal(field, buf, uint64(v))
}

func EnumSize(field int, v int32) int {
	return UInt64Size(field, uint64(v))
}

func RepeatedBytesMarshal(field int, buf []byte, v [][]byte) (int, error) {
	var offset int

	for i := range v {
		off, err := BytesMarshal(field, buf[offset:], v[i])
		if err != nil {
			return 0, err
		}

		offset += off
	}

	return offset, nil
}

func RepeatedBytesSize(field int, v [][]byte) (size int) {
	for i := range v {
		size += BytesSize(field, v[i])
	}

	return size
}

func RepeatedStringMarshal(field int, buf []byte, v []string) (int, error) {
	var offset int

	for i := range v {
		off, err := StringMarshal(field, buf[offset:], v[i])
		if err != nil {
			return 0, err
		}

		offset += off
	}

	return offset, nil
}

func RepeatedStringSize(field int, v []string) (size int) {
	for i := range v {
		size += StringSize(field, v[i])
	}

	return size
}

func RepeatedUInt64Marshal(field int, buf []byte, v []uint64) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	prefix := field<<3 | 0x02
	offset := binary.PutUvarint(buf, uint64(prefix))

	_, arrSize := RepeatedUInt64Size(field, v)
	offset += binary.PutUvarint(buf[offset:], uint64(arrSize))
	for i := range v {
		offset += binary.PutUvarint(buf[offset:], v[i])
	}

	return offset, nil
}

func RepeatedUInt64Size(field int, v []uint64) (size, arraySize int) {
	if len(v) == 0 {
		return 0, 0
	}

	for i := range v {
		size += VarUIntSize(v[i])
	}
	arraySize = size

	size += VarUIntSize(uint64(size))

	prefix := field<<3 | 0x2
	size += VarUIntSize(uint64(prefix))

	return size, arraySize
}

func RepeatedInt64Marshal(field int, buf []byte, v []int64) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Marshal(field, buf, convert)
}

func RepeatedInt64Size(field int, v []int64) (size, arraySize int) {
	if len(v) == 0 {
		return 0, 0
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Size(field, convert)
}

func RepeatedUInt32Marshal(field int, buf []byte, v []uint32) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Marshal(field, buf, convert)
}

func RepeatedUInt32Size(field int, v []uint32) (size, arraySize int) {
	if len(v) == 0 {
		return 0, 0
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Size(field, convert)
}

func RepeatedInt32Marshal(field int, buf []byte, v []int32) (int, error) {
	if len(v) == 0 {
		return 0, nil
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Marshal(field, buf, convert)
}

func RepeatedInt32Size(field int, v []int32) (size, arraySize int) {
	if len(v) == 0 {
		return 0, 0
	}

	convert := make([]uint64, len(v))
	for i := range v {
		convert[i] = uint64(v[i])
	}

	return RepeatedUInt64Size(field, convert)
}

// varUIntSize returns length of varint byte sequence for uint64 value 'x'.
func VarUIntSize(x uint64) int {
	return (bits.Len64(x|1) + 6) / 7
}

func NestedStructurePrefix(field int64) (prefix uint64, ln int) {
	prefix = uint64(field<<3 | 0x02)
	return prefix, VarUIntSize(prefix)
}

func NestedStructureMarshal(field int64, buf []byte, v stableMarshaller) (int, error) {
	if v == nil {
		return 0, nil
	}

	prefix, _ := NestedStructurePrefix(field)
	offset := binary.PutUvarint(buf, prefix)

	n := v.StableSize()
	offset += binary.PutUvarint(buf[offset:], uint64(n))

	_, err := v.StableMarshal(buf[offset:])
	if err != nil {
		return 0, err
	}

	return offset + n, nil
}

func NestedStructureSize(field int64, v stableMarshaller) (size int) {
	if v == nil {
		return 0
	}

	_, ln := NestedStructurePrefix(field)
	n := v.StableSize()
	size = ln + VarUIntSize(uint64(n)) + n

	return size
}
