package v2

import (
	"encoding/binary"
	"math/bits"
)

func (m *OwnerID) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.StableSize())
	}

	var (
		i, n, offset int
	)

	// Write key field.

	buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Value)))
	n = copy(buf[i+1+offset:], m.Value)
	i += 1 + offset + n

	return buf, nil
}

func (m *OwnerID) StableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	ln = len(m.Value)                        // size of key field
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string

	return size
}

// uvarIntSize returns length of varint byte sequence for uint64 value 'x'.
func uvarIntSize(x uint64) int {
	return (bits.Len64(x|1) + 6) / 7
}
