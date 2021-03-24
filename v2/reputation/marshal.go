package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

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

	n, err = protoutil.BytesMarshal(trustPeerFNum, buf[offset:], x.peer)
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
	size += protoutil.BytesSize(trustPeerFNum, x.peer)
	size += protoutil.Float64Size(trustValueFNum, x.val)

	return
}

func (x *Trust) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.Trust))
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
