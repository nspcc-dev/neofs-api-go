package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

const (
	_ = iota
	peerIDValFNum
)

func (x *PeerID) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	_, err := protoutil.BytesMarshal(peerIDValFNum, buf, x.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *PeerID) StableSize() (size int) {
	size += protoutil.BytesSize(peerIDValFNum, x.val)

	return
}

func (x *PeerID) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.PeerID))
}

const (
	_ = iota
	trustPeerFNum
	trustValueFNum
)

func (x *Trust) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.NestedStructureMarshal(trustPeerFNum, buf[offset:], x.peer)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.Float64Marshal(trustValueFNum, buf[offset:], x.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *Trust) StableSize() (size int) {
	size += protoutil.NestedStructureSize(trustPeerFNum, x.peer)
	size += protoutil.Float64Size(trustValueFNum, x.val)

	return
}

func (x *Trust) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.Trust))
}

const (
	_ = iota
	globalTrustBodyManagerFNum
	globalTrustBodyValueFNum
)

func (x *GlobalTrustBody) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	offset, err := protoutil.NestedStructureMarshal(globalTrustBodyManagerFNum, buf, x.manager)
	if err != nil {
		return nil, err
	}

	_, err = protoutil.NestedStructureMarshal(globalTrustBodyValueFNum, buf[offset:], x.trust)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *GlobalTrustBody) StableSize() (size int) {
	size += protoutil.NestedStructureSize(globalTrustBodyManagerFNum, x.manager)
	size += protoutil.NestedStructureSize(globalTrustBodyValueFNum, x.trust)

	return
}

func (x *GlobalTrustBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.GlobalTrust_Body))
}

const (
	_ = iota
	globalTrustVersionFNum
	globalTrustBodyFNum
	globalTrustSigFNum
)

func (x *GlobalTrust) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	offset, err := protoutil.NestedStructureMarshal(globalTrustVersionFNum, buf, x.version)
	if err != nil {
		return nil, err
	}

	n, err := protoutil.NestedStructureMarshal(globalTrustBodyFNum, buf[offset:], x.body)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(globalTrustSigFNum, buf[offset:], x.sig)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *GlobalTrust) StableSize() (size int) {
	size += protoutil.NestedStructureSize(globalTrustVersionFNum, x.version)
	size += protoutil.NestedStructureSize(globalTrustBodyFNum, x.body)
	size += protoutil.NestedStructureSize(globalTrustSigFNum, x.sig)

	return
}

func (x *GlobalTrust) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.GlobalTrust))
}

const (
	_ = iota
	sendLocalTrustBodyEpochFNum
	sendLocalTrustBodyTrustsFNum
)

func (x *SendLocalTrustRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.UInt64Marshal(sendLocalTrustBodyEpochFNum, buf[offset:], x.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range x.trusts {
		n, err = protoutil.NestedStructureMarshal(sendLocalTrustBodyTrustsFNum, buf[offset:], x.trusts[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (x *SendLocalTrustRequestBody) StableSize() (size int) {
	size += protoutil.UInt64Size(sendLocalTrustBodyEpochFNum, x.epoch)

	for i := range x.trusts {
		size += protoutil.NestedStructureSize(sendLocalTrustBodyTrustsFNum, x.trusts[i])
	}

	return
}

func (x *SendLocalTrustRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.SendLocalTrustRequest_Body))
}

func (x *SendLocalTrustResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return buf, nil
}

func (x *SendLocalTrustResponseBody) StableSize() int {
	return 0
}

func (x *SendLocalTrustResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.SendLocalTrustResponse_Body))
}

const (
	_ = iota
	sendInterResBodyIterFNum
	sendInterResBodyTrustFNum
)

func (x *SendIntermediateResultRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	offset, err := protoutil.UInt32Marshal(sendInterResBodyIterFNum, buf, x.iter)
	if err != nil {
		return nil, err
	}

	_, err = protoutil.NestedStructureMarshal(sendInterResBodyTrustFNum, buf[offset:], x.trust)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *SendIntermediateResultRequestBody) StableSize() (size int) {
	size += protoutil.UInt32Size(sendInterResBodyIterFNum, x.iter)
	size += protoutil.NestedStructureSize(sendInterResBodyTrustFNum, x.trust)

	return
}

func (x *SendIntermediateResultRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.SendIntermediateResultRequest_Body))
}

func (x *SendIntermediateResultResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return buf, nil
}

func (x *SendIntermediateResultResponseBody) StableSize() int {
	return 0
}

func (x *SendIntermediateResultResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.SendIntermediateResultResponse_Body))
}
