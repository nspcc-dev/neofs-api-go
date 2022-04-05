/*
This package contains help functions for stable marshaller. Their usage is
totally optional. One can implement fast stable marshaller without these
runtime function calls.
*/

package proto

import (
	"encoding/binary"
	"math"
	"math/bits"
	"reflect"
)

type (
	stableMarshaller interface {
		StableMarshal([]byte) []byte
		StableSize() int
	}
)

func BytesMarshal(field int, buf, v []byte) int {
	if len(v) == 0 {
		return 0
	}

	prefix := field<<3 | 0x2

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	i += binary.PutUvarint(buf[i:], uint64(len(v)))
	i += copy(buf[i:], v)

	return i
}

func BytesSize(field int, v []byte) int {
	ln := len(v)
	if ln == 0 {
		return 0
	}

	prefix := field<<3 | 0x2

	return VarUIntSize(uint64(prefix)) + VarUIntSize(uint64(ln)) + ln
}

func StringMarshal(field int, buf []byte, v string) int {
	return BytesMarshal(field, buf, []byte(v))
}

func StringSize(field int, v string) int {
	return BytesSize(field, []byte(v))
}

func BoolMarshal(field int, buf []byte, v bool) int {
	if !v {
		return 0
	}

	prefix := field << 3

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	buf[i] = 0x1

	return i + 1
}

func BoolSize(field int, v bool) int {
	if !v {
		return 0
	}

	prefix := field << 3

	return VarUIntSize(uint64(prefix)) + 1 // bool is always 1 byte long
}

func UInt64Marshal(field int, buf []byte, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := field << 3

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	i += binary.PutUvarint(buf[i:], v)

	return i
}

func UInt64Size(field int, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := field << 3

	return VarUIntSize(uint64(prefix)) + VarUIntSize(v)
}

func Int64Marshal(field int, buf []byte, v int64) int {
	return UInt64Marshal(field, buf, uint64(v))
}

func Int64Size(field int, v int64) int {
	return UInt64Size(field, uint64(v))
}

func UInt32Marshal(field int, buf []byte, v uint32) int {
	return UInt64Marshal(field, buf, uint64(v))
}

func UInt32Size(field int, v uint32) int {
	return UInt64Size(field, uint64(v))
}

func Int32Marshal(field int, buf []byte, v int32) int {
	return UInt64Marshal(field, buf, uint64(v))
}

func Int32Size(field int, v int32) int {
	return UInt64Size(field, uint64(v))
}

func EnumMarshal(field int, buf []byte, v int32) int {
	return UInt64Marshal(field, buf, uint64(v))
}

func EnumSize(field int, v int32) int {
	return UInt64Size(field, uint64(v))
}

func RepeatedBytesMarshal(field int, buf []byte, v [][]byte) int {
	var offset int

	for i := range v {
		offset += BytesMarshal(field, buf[offset:], v[i])
	}

	return offset
}

func RepeatedBytesSize(field int, v [][]byte) (size int) {
	for i := range v {
		size += BytesSize(field, v[i])
	}

	return size
}

func RepeatedStringMarshal(field int, buf []byte, v []string) int {
	var offset int

	for i := range v {
		offset += StringMarshal(field, buf[offset:], v[i])
	}

	return offset
}

func RepeatedStringSize(field int, v []string) (size int) {
	for i := range v {
		size += StringSize(field, v[i])
	}

	return size
}

func RepeatedUInt64Marshal(field int, buf []byte, v []uint64) int {
	if len(v) == 0 {
		return 0
	}

	prefix := field<<3 | 0x02
	offset := binary.PutUvarint(buf, uint64(prefix))

	_, arrSize := RepeatedUInt64Size(field, v)
	offset += binary.PutUvarint(buf[offset:], uint64(arrSize))
	for i := range v {
		offset += binary.PutUvarint(buf[offset:], v[i])
	}

	return offset
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

func RepeatedInt64Marshal(field int, buf []byte, v []int64) int {
	if len(v) == 0 {
		return 0
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

func RepeatedUInt32Marshal(field int, buf []byte, v []uint32) int {
	if len(v) == 0 {
		return 0
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

func RepeatedInt32Marshal(field int, buf []byte, v []int32) int {
	if len(v) == 0 {
		return 0
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

func NestedStructureMarshal(field int64, buf []byte, v stableMarshaller) int {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return 0
	}

	prefix, _ := NestedStructurePrefix(field)
	offset := binary.PutUvarint(buf, prefix)

	n := v.StableSize()
	offset += binary.PutUvarint(buf[offset:], uint64(n))
	v.StableMarshal(buf[offset:])

	return offset + n
}

func NestedStructureSize(field int64, v stableMarshaller) (size int) {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return 0
	}

	_, ln := NestedStructurePrefix(field)
	n := v.StableSize()
	size = ln + VarUIntSize(uint64(n)) + n

	return size
}

func Fixed64Marshal(field int, buf []byte, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := field<<3 | 1

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	binary.LittleEndian.PutUint64(buf[i:], v)

	return i + 8
}

func Fixed64Size(fNum int, v uint64) int {
	if v == 0 {
		return 0
	}

	prefix := fNum<<3 | 1

	return VarUIntSize(uint64(prefix)) + 8
}

func Float64Marshal(field int, buf []byte, v float64) int {
	if v == 0 {
		return 0
	}

	prefix := field<<3 | 1

	i := binary.PutUvarint(buf, uint64(prefix))
	binary.LittleEndian.PutUint64(buf[i:], math.Float64bits(v))

	return i + 8
}

func Float64Size(fNum int, v float64) int {
	if v == 0 {
		return 0
	}

	prefix := fNum<<3 | 1

	return VarUIntSize(uint64(prefix)) + 8
}

// Fixed32Marshal encodes uint32 value to Protocol Buffers fixed32 field with specified number,
// and writes it to specified buffer. Returns number of bytes written.
//
// Panics if the buffer is undersized.
func Fixed32Marshal(field int, buf []byte, v uint32) int {
	if v == 0 {
		return 0
	}

	prefix := field<<3 | 5

	// buf length check can prevent panic at PutUvarint, but it will make
	// marshaller a bit slower.
	i := binary.PutUvarint(buf, uint64(prefix))
	binary.LittleEndian.PutUint32(buf[i:], v)

	return i + 4
}

// Fixed32Size returns number of bytes required to encode uint32 value to Protocol Buffers fixed32 field
// with specified number.
func Fixed32Size(fNum int, v uint32) int {
	if v == 0 {
		return 0
	}

	prefix := fNum<<3 | 5

	return VarUIntSize(uint64(prefix)) + 4
}
