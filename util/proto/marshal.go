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

// varUIntSize returns length of varint byte sequence for uint64 value 'x'.
func VarUIntSize(x uint64) int {
	return (bits.Len64(x|1) + 6) / 7
}
