package container

import (
	"encoding/binary"
	"math/bits"

	"github.com/pkg/errors"
)

// StableMarshal marshals auto-generated container structure into
// protobuf-compatible stable byte sequence.
func (m *Container) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.StableSize())
	}

	var (
		i, n, offset int
	)

	// Write owner id field.

	if m.OwnerId != nil {
		buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
		n = m.OwnerId.StableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))

		_, err := m.OwnerId.StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal owner id")
		}

		i += 1 + offset + n
	}

	// Write salt field.

	buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Nonce)))
	n = copy(buf[i+1+offset:], m.Nonce)
	i += 1 + offset + n

	// Write basic acl field.

	buf[i] = 0x18 // id:0x3 << 3 | wiretype:0x0
	offset = binary.PutUvarint(buf[i+1:], uint64(m.BasicAcl))
	i += 1 + offset

	// Write attributes field.

	for j := range m.Attributes {
		buf[i] = 0x22 // id:0x4 << 3 | wiretype:0x2
		n = m.Attributes[j].StableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))

		_, err := m.Attributes[j].StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal attribute %v",
				m.Attributes[i].Key)
		}

		i += 1 + offset + n
	}

	// Write placement rule field.

	if m.PlacementPolicy != nil {
		buf[i] = 0x2A // id:0x5 << 3 | wiretype:0x2
		n = m.PlacementPolicy.StableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))

		_, err := m.PlacementPolicy.StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal attribute %v",
				m.Attributes[i].Key)
		}
	}

	return buf, nil
}

func (m *Container) StableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	if m.OwnerId != nil {
		ln = m.OwnerId.StableSize()
	}
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + bytes length + bytes

	ln = len(m.Nonce)                        // size of salt field
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + bytes length + bytes

	// size of basic acl field
	size += 1 + uvarIntSize(uint64(m.BasicAcl)) // wiretype + varint

	// size of attributes
	for i := range m.Attributes {
		ln = m.Attributes[i].StableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of struct + struct
	}

	// size of placement rule
	if m.PlacementPolicy != nil {
		ln = m.PlacementPolicy.StableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of struct + struct
	}

	return size
}

func (m *Container_Attribute) StableMarshal(buf []byte) ([]byte, error) {
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
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Key)))
	n = copy(buf[i+1+offset:], m.Key)
	i += 1 + offset + n

	// Write value field.

	buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Value)))
	copy(buf[i+1+offset:], m.Value)

	return buf, nil
}

func (m *Container_Attribute) StableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	ln = len(m.Key)                          // size of key field
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string

	ln = len(m.Value)                        // size of value field
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string

	return size
}

func (m *PutRequest_Body) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.StableSize())
	}

	var (
		i, n, offset int
	)

	// Write container field.

	if m.Container != nil {
		buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
		n = m.Container.StableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))

		_, err := m.Container.StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal container")
		}

		i += 1 + offset + n
	}

	// Write public key field.

	buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.PublicKey)))
	n = copy(buf[i+1+offset:], m.PublicKey)
	i += 1 + offset + n

	// Write signature field.

	buf[i] = 0x1A // id:0x3 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Signature)))
	copy(buf[i+1+offset:], m.Signature)

	return buf, nil
}

func (m *PutRequest_Body) StableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	if m.Container != nil {
		ln = m.Container.StableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string
	}

	ln = len(m.PublicKey)
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string

	ln = len(m.Signature)
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of string + string

	return size
}

// uvarIntSize returns length of varint byte sequence for uint64 value 'x'.
func uvarIntSize(x uint64) int {
	return (bits.Len64(x|1) + 6) / 7
}
