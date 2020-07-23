package eacl

import (
	"encoding/binary"

	"github.com/pkg/errors"
)

const (
	sliceLenSize  = 2 // uint16 for len()
	actionSize    = 4 // uint32
	opTypeSize    = 4 // uint32
	hdrTypeSize   = 4 // uint32
	matchTypeSize = 4 // uint32
	targetSize    = 4 // uint32
)

// MarshalTable encodes Table into a
// binary form and returns the result.
//
// If table is nil, empty slice is returned.
func MarshalTable(table Table) []byte {
	if table == nil {
		return make([]byte, 0)
	}

	// allocate buffer
	buf := make([]byte, tableBinSize(table))

	records := table.Records()

	// write record number
	binary.BigEndian.PutUint16(buf, uint16(len(records)))
	off := sliceLenSize

	// write all records
	for _, record := range records {
		// write action
		binary.BigEndian.PutUint32(buf[off:], uint32(record.Action()))
		off += actionSize

		// write operation type
		binary.BigEndian.PutUint32(buf[off:], uint32(record.OperationType()))
		off += actionSize

		filters := record.HeaderFilters()

		// write filter number
		binary.BigEndian.PutUint16(buf[off:], uint16(len(filters)))
		off += sliceLenSize

		// write all filters
		for _, filter := range filters {
			// write header type
			binary.BigEndian.PutUint32(buf[off:], uint32(filter.HeaderType()))
			off += hdrTypeSize

			// write match type
			binary.BigEndian.PutUint32(buf[off:], uint32(filter.MatchType()))
			off += matchTypeSize

			// write header name size
			name := []byte(filter.Name())
			binary.BigEndian.PutUint16(buf[off:], uint16(len(name)))
			off += sliceLenSize

			// write header name bytes
			off += copy(buf[off:], name)

			// write header value size
			val := []byte(filter.Value())
			binary.BigEndian.PutUint16(buf[off:], uint16(len(val)))
			off += sliceLenSize

			// write header value bytes
			off += copy(buf[off:], val)
		}

		targets := record.TargetList()

		// write target number
		binary.BigEndian.PutUint16(buf[off:], uint16(len(targets)))
		off += sliceLenSize

		// write all targets
		for _, target := range targets {
			// write target group
			binary.BigEndian.PutUint32(buf[off:], uint32(target.Group()))
			off += targetSize

			keys := target.KeyList()

			// write key number
			binary.BigEndian.PutUint16(buf[off:], uint16(len(keys)))
			off += sliceLenSize

			// write keys
			for i := range keys {
				// write key size
				binary.BigEndian.PutUint16(buf[off:], uint16(len(keys[i])))
				off += sliceLenSize

				// write key bytes
				off += copy(buf[off:], keys[i])
			}
		}
	}

	return buf
}

// returns the size of Table in a binary format.
func tableBinSize(table Table) (sz int) {
	sz = sliceLenSize // number of records

	records := table.Records()
	ln := len(records)

	sz += ln * actionSize // action type of each record
	sz += ln * opTypeSize // operation type of each record

	for _, record := range records {
		sz += sliceLenSize // number of filters

		filters := record.HeaderFilters()
		ln := len(filters)

		sz += ln * hdrTypeSize   // header type of each filter
		sz += ln * matchTypeSize // match type of each filter

		for _, filter := range filters {
			sz += sliceLenSize       // header name size
			sz += len(filter.Name()) // header name bytes

			sz += sliceLenSize        // header value size
			sz += len(filter.Value()) // header value bytes
		}

		sz += sliceLenSize // number of targets

		targets := record.TargetList()
		ln = len(targets)

		sz += ln * targetSize // target group of each target

		for _, target := range targets {
			sz += sliceLenSize // number of keys

			for _, key := range target.KeyList() {
				sz += sliceLenSize // key size
				sz += len(key)     // key bytes
			}
		}
	}

	return
}

// UnmarshalTable unmarshals Table from
// a binary representation.
//
// If data is empty, table w/o records is returned.
func UnmarshalTable(data []byte) (Table, error) {
	table := WrapTable(nil)

	if len(data) == 0 {
		return table, nil
	}

	// decode record number
	if len(data) < sliceLenSize {
		return nil, errors.New("could not decode record number")
	}

	recordNum := binary.BigEndian.Uint16(data)
	records := make([]Record, 0, recordNum)

	off := sliceLenSize

	// decode all records one by one
	for i := uint16(0); i < recordNum; i++ {
		record := WrapRecord(nil)

		// decode action
		if len(data[off:]) < actionSize {
			return nil, errors.Errorf("could not decode action of record #%d", i)
		}

		record.SetAction(Action(binary.BigEndian.Uint32(data[off:])))
		off += actionSize

		// decode operation type
		if len(data[off:]) < opTypeSize {
			return nil, errors.Errorf("could not decode operation type of record #%d", i)
		}

		record.SetOperationType(OperationType(binary.BigEndian.Uint32(data[off:])))
		off += opTypeSize

		// decode filter number
		if len(data[off:]) < sliceLenSize {
			return nil, errors.Errorf("could not decode filter number of record #%d", i)
		}

		filterNum := binary.BigEndian.Uint16(data[off:])
		off += sliceLenSize
		filters := make([]HeaderFilter, 0, filterNum)

		// decode filters one by one
		for j := uint16(0); j < filterNum; j++ {
			filter := WrapFilterInfo(nil)

			// decode header type
			if len(data[off:]) < hdrTypeSize {
				return nil, errors.Errorf("could not decode header type of filter #%d of record #%d", j, i)
			}

			filter.SetHeaderType(HeaderType(binary.BigEndian.Uint32(data[off:])) )
			off += hdrTypeSize

			// decode match type
			if len(data[off:]) < matchTypeSize {
				return nil, errors.Errorf("could not decode match type of filter #%d of record #%d", j, i)
			}

			filter.SetMatchType(MatchType(binary.BigEndian.Uint32(data[off:])) )
			off += matchTypeSize

			// decode header name size
			if len(data[off:]) < sliceLenSize {
				return nil, errors.Errorf("could not decode header name size of filter #%d of record #%d", j, i)
			}

			hdrNameSize := int(binary.BigEndian.Uint16(data[off:]))
			off += sliceLenSize

			// decode header name
			if len(data[off:]) < hdrNameSize {
				return nil, errors.Errorf("could not decode header name of filter #%d of record #%d", j, i)
			}

			filter.SetName(string(data[off : off+hdrNameSize]))

			off += hdrNameSize

			// decode header value size
			if len(data[off:]) < sliceLenSize {
				return nil, errors.Errorf("could not decode header value size of filter #%d of record #%d", j, i)
			}

			hdrValSize := int(binary.BigEndian.Uint16(data[off:]))
			off += sliceLenSize

			// decode header value
			if len(data[off:]) < hdrValSize {
				return nil, errors.Errorf("could not decode header value of filter #%d of record #%d", j, i)
			}

			filter.SetValue(string(data[off : off+hdrValSize]))

			off += hdrValSize

			filters = append(filters, filter)
		}

		record.SetHeaderFilters(filters)

		// decode target number
		if len(data[off:]) < sliceLenSize {
			return nil, errors.Errorf("could not decode target number of record #%d", i)
		}

		targetNum := int(binary.BigEndian.Uint16(data[off:]))
		off += sliceLenSize

		targets := make([]Target, 0, targetNum)

		// decode targets one by one
		for j := 0; j < targetNum; j++ {
			target := WrapTarget(nil)

			// decode target group
			if len(data[off:]) < targetSize {
				return nil, errors.Errorf("could not decode target group of target #%d of record #%d", j, i)
			}

			target.SetGroup(				Group(binary.BigEndian.Uint32(data[off:])),				)
			off += targetSize

			// decode key number
			if len(data[off:]) < sliceLenSize {
				return nil, errors.Errorf("could not decode key number of target #%d of record #%d", j, i)
			}

			keyNum := int(binary.BigEndian.Uint16(data[off:]))
			off += sliceLenSize
			keys := make([][]byte, 0, keyNum)

			for k := 0; k < keyNum; k++ {
				// decode key size
				if len(data[off:]) < sliceLenSize {
					return nil, errors.Errorf("could not decode size of key #%d target #%d of record #%d", k, j, i)
				}

				keySz := int(binary.BigEndian.Uint16(data[off:]))
				off += sliceLenSize

				// decode key
				if len(data[off:]) < keySz {
					return nil, errors.Errorf("could not decode key #%d target #%d of record #%d", k, j, i)
				}

				key := make([]byte, keySz)

				off += copy(key, data[off:off+keySz])

				keys = append(keys, key)
			}

			target.SetKeyList(keys)

			targets = append(targets, target)
		}

		record.SetTargetList(targets)

		records = append(records, record)
	}

	table.SetRecords(records)

	return table, nil
}
