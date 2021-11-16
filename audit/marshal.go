package audit

import (
	audit "github.com/nspcc-dev/neofs-api-go/v2/audit/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	"github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

const (
	_ = iota
	versionFNum
	auditEpochFNum
	cidFNum
	pubKeyFNum
	completeFNum
	requestsFNum
	retriesFNum
	passSGFNum
	failSGFNum
	hitFNum
	missFNum
	failFNum
	passNodesFNum
	failNodesFNum
)

// StableMarshal marshals unified DataAuditResult structure into a protobuf
// binary format without field order shuffle.
func (a *DataAuditResult) StableMarshal(buf []byte) ([]byte, error) {
	if a == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, a.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(versionFNum, buf[offset:], a.version)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.Fixed64Marshal(auditEpochFNum, buf[offset:], a.auditEpoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.NestedStructureMarshal(cidFNum, buf[offset:], a.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(pubKeyFNum, buf[offset:], a.pubKey)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BoolMarshal(completeFNum, buf[offset:], a.complete)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(requestsFNum, buf[offset:], a.requests)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(retriesFNum, buf[offset:], a.retries)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = refs.ObjectIDNestedListMarshal(passSGFNum, buf[offset:], a.passSG)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = refs.ObjectIDNestedListMarshal(failSGFNum, buf[offset:], a.failSG)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(hitFNum, buf[offset:], a.hit)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(missFNum, buf[offset:], a.miss)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(failFNum, buf[offset:], a.fail)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.RepeatedBytesMarshal(passNodesFNum, buf[offset:], a.passNodes)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.RepeatedBytesMarshal(failNodesFNum, buf[offset:], a.failNodes)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// StableSize returns byte length of DataAuditResult structure
// marshaled by StableMarshal function.
func (a *DataAuditResult) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.NestedStructureSize(versionFNum, a.version)
	size += proto.Fixed64Size(auditEpochFNum, a.auditEpoch)
	size += proto.NestedStructureSize(cidFNum, a.cid)
	size += proto.BytesSize(pubKeyFNum, a.pubKey)
	size += proto.BoolSize(completeFNum, a.complete)
	size += proto.UInt32Size(requestsFNum, a.requests)
	size += proto.UInt32Size(retriesFNum, a.retries)
	size += refs.ObjectIDNestedListSize(passSGFNum, a.passSG)
	size += refs.ObjectIDNestedListSize(failSGFNum, a.failSG)
	size += proto.UInt32Size(hitFNum, a.hit)
	size += proto.UInt32Size(missFNum, a.miss)
	size += proto.UInt32Size(failFNum, a.fail)
	size += proto.RepeatedBytesSize(passNodesFNum, a.passNodes)
	size += proto.RepeatedBytesSize(failNodesFNum, a.failNodes)

	return size
}

// Unmarshal unmarshals DataAuditResult structure from its protobuf
// binary representation.
func (a *DataAuditResult) Unmarshal(data []byte) error {
	return message.Unmarshal(a, data, new(audit.DataAuditResult))
}
