package netmap

import (
	"encoding/binary"
	"math/bits"

	"github.com/pkg/errors"
)

func (m *PlacementRule) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.StableSize())
	}

	var (
		i, n, offset int
	)

	// Write replication factor field.

	buf[i] = 0x08 // id:0x1 << 3 | wiretype:0x0
	offset = binary.PutUvarint(buf[i+1:], uint64(m.ReplFactor))
	i += 1 + offset

	// write select/filter groups field
	for j := range m.SfGroups {
		buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2

		n, _ = m.SfGroups[j].stableSizeWithExclude()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))

		_, err := m.SfGroups[j].StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal SFGroup id:%d", j)
		}

		i += 1 + offset + n
	}

	return buf, nil
}

func (m *PlacementRule) StableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)
	_ = ln

	// size of replication factor field
	size += 1 + uvarIntSize(uint64(m.ReplFactor)) // wiretype + varint

	for i := range m.SfGroups {
		ln, _ = m.SfGroups[i].stableSizeWithExclude()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of struct + struct
	}

	return size
}

func (m *PlacementRule_SFGroup) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	size, excludeSize := m.stableSizeWithExclude()
	if buf == nil {
		buf = make([]byte, size)
	}

	var (
		i, n, offset int
	)

	// write filters field
	for j := range m.Filters {
		buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
		n = m.Filters[j].stableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))
		_, err := m.Filters[j].StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal Filter id:%d", j)
		}
		i += 1 + offset + n
	}

	// write selectors field
	for j := range m.Selectors {
		buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
		n = m.Selectors[j].stableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))
		_, err := m.Selectors[j].StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal Selector id:%d", j)
		}
		i += 1 + offset + n
	}

	// write excluded field in packed format
	buf[i] = 0x1A // id:0x3 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(excludeSize))
	i += 1 + offset
	for j := range m.Exclude {
		offset = binary.PutUvarint(buf[i:], uint64(m.Exclude[j]))
		i += offset
	}

	return buf, nil
}

func (m *PlacementRule_SFGroup) stableSizeWithExclude() (int, int) {
	if m == nil {
		return 0, 0
	}

	var (
		ln, size int
	)

	// size of filters field
	for i := range m.Filters {
		ln = m.Filters[i].stableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of struct + struct
	}

	// size of selectors field
	for i := range m.Selectors {
		ln = m.Selectors[i].stableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + size of struct + struct
	}

	// size of exclude field
	ln = 0
	for i := range m.Exclude {
		ln += uvarIntSize(uint64(m.Exclude[i]))
	}
	size += 1 + uvarIntSize(uint64(ln)) + ln // wiretype + packed varints size + packed varints

	return size, ln
}

func (m *PlacementRule_SFGroup_Selector) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.stableSize())
	}

	var (
		i, offset int
	)

	// write count field
	buf[i] = 0x8 // id:0x1 << 3 | wiretype:0x0
	offset = binary.PutUvarint(buf[i+1:], uint64(m.Count))
	i += 1 + offset

	// write key field
	buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Key)))
	copy(buf[i+1+offset:], m.Key)

	return buf, nil
}

func (m *PlacementRule_SFGroup_Selector) stableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	// size of count field
	size += 1 + uvarIntSize(uint64(m.Count))

	// size of key field
	ln = len(m.Key)
	size += 1 + uvarIntSize(uint64(ln)) + ln

	return size
}

func (m *PlacementRule_SFGroup_Filter) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.stableSize())
	}

	var (
		i, n, offset int
	)

	// write key field
	buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
	offset = binary.PutUvarint(buf[i+1:], uint64(len(m.Key)))
	n = copy(buf[i+1+offset:], m.Key)
	i += 1 + offset + n

	// write simple filter field
	if m.F != nil {
		buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
		n = m.F.stableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))
		_, err := m.F.StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrap(err, "can't marshal netmap filter")
		}
	}

	return buf, nil
}

func (m *PlacementRule_SFGroup_Filter) stableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	// size of key field
	ln = len(m.Key)
	size += 1 + uvarIntSize(uint64(ln)) + ln

	// size of simple filter
	if m.F != nil {
		ln = m.F.stableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln
	}

	return size
}

func (m *PlacementRule_SFGroup_Filter_SimpleFilter) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.stableSize())
	}

	var (
		i, n, offset int
	)

	// write key field
	buf[i] = 0x08 // id:0x1 << 3 | wiretype:0x0
	offset = binary.PutUvarint(buf[i+1:], uint64(m.Op))
	i += 1 + offset

	// write value if present
	if val, ok := m.Args.(*PlacementRule_SFGroup_Filter_SimpleFilter_Value); ok {
		buf[i] = 0x12 // id:0x2 << 3 | wiretype:0x2
		offset = binary.PutUvarint(buf[i+1:], uint64(len(val.Value)))
		copy(buf[i+1+offset:], val.Value)
	} else if filters, ok := m.Args.(*PlacementRule_SFGroup_Filter_SimpleFilter_FArgs); ok {
		if filters.FArgs != nil {
			buf[i] = 0x1A // id:0x3 << 3 | wiretype:0x2
			n = filters.FArgs.stableSize()
			offset = binary.PutUvarint(buf[i+1:], uint64(n))
			_, err := filters.FArgs.StableMarshal(buf[i+1+offset:])
			if err != nil {
				return nil, errors.Wrap(err, "can't marshal simple filters")
			}
		}
	}

	return buf, nil
}

func (m *PlacementRule_SFGroup_Filter_SimpleFilter) stableSize() int {
	if m == nil {
		return 0
	}

	var (
		ln, size int
	)

	// size of key field
	size += 1 + uvarIntSize(uint64(m.Op))

	if val, ok := m.Args.(*PlacementRule_SFGroup_Filter_SimpleFilter_Value); ok {
		// size of value if present
		ln = len(val.Value)
		size += 1 + uvarIntSize(uint64(ln)) + ln
	} else if filters, ok := m.Args.(*PlacementRule_SFGroup_Filter_SimpleFilter_FArgs); ok {
		// size of simple filters if present
		if filters.FArgs != nil {
			ln = filters.FArgs.stableSize()
			size += 1 + uvarIntSize(uint64(ln)) + ln
		}
	}

	return size
}

func (m *PlacementRule_SFGroup_Filter_SimpleFilters) StableMarshal(buf []byte) ([]byte, error) {
	if m == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, m.stableSize())
	}

	var (
		i, n, offset int
	)

	// write filters field
	for j := range m.Filters {
		buf[i] = 0x0A // id:0x1 << 3 | wiretype:0x2
		n = m.Filters[j].stableSize()
		offset = binary.PutUvarint(buf[i+1:], uint64(n))
		_, err := m.Filters[j].StableMarshal(buf[i+1+offset:])
		if err != nil {
			return nil, errors.Wrapf(err, "can't marshal simple filter id:%d", j)
		}
		i += 1 + offset + n
	}

	return buf, nil
}

func (m *PlacementRule_SFGroup_Filter_SimpleFilters) stableSize() int {
	if m == nil {
		return 0
	}
	var (
		ln, size int
	)

	// size of key field
	for i := range m.Filters {
		ln = m.Filters[i].stableSize()
		size += 1 + uvarIntSize(uint64(ln)) + ln
	}

	return size
}

// uvarIntSize returns length of varint byte sequence for uint64 value 'x'.
func uvarIntSize(x uint64) int {
	return (bits.Len64(x|1) + 6) / 7
}
