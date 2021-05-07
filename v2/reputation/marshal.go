package reputation

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	reputation "github.com/nspcc-dev/neofs-api-go/v2/reputation/grpc"
)

const (
	_ = iota
	peerIDPubKeyFNum
)

func (x *PeerID) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	_, err := protoutil.BytesMarshal(peerIDPubKeyFNum, buf, x.publicKey)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *PeerID) StableSize() (size int) {
	size += protoutil.BytesSize(peerIDPubKeyFNum, x.publicKey)

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
	p2pTrustTrustingFNum
	p2pTrustValueFNum
)

func (x *PeerToPeerTrust) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	offset, err := protoutil.NestedStructureMarshal(p2pTrustTrustingFNum, buf, x.trusting)
	if err != nil {
		return nil, err
	}

	_, err = protoutil.NestedStructureMarshal(p2pTrustValueFNum, buf[offset:], x.trust)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *PeerToPeerTrust) StableSize() (size int) {
	size += protoutil.NestedStructureSize(p2pTrustTrustingFNum, x.trusting)
	size += protoutil.NestedStructureSize(p2pTrustValueFNum, x.trust)

	return
}

func (x *PeerToPeerTrust) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.PeerToPeerTrust))
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
	announceLocalTrustBodyEpochFNum
	announceLocalTrustBodyTrustsFNum
)

func (x *AnnounceLocalTrustRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = protoutil.UInt64Marshal(announceLocalTrustBodyEpochFNum, buf[offset:], x.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range x.trusts {
		n, err = protoutil.NestedStructureMarshal(announceLocalTrustBodyTrustsFNum, buf[offset:], x.trusts[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (x *AnnounceLocalTrustRequestBody) StableSize() (size int) {
	size += protoutil.UInt64Size(announceLocalTrustBodyEpochFNum, x.epoch)

	for i := range x.trusts {
		size += protoutil.NestedStructureSize(announceLocalTrustBodyTrustsFNum, x.trusts[i])
	}

	return
}

func (x *AnnounceLocalTrustRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.AnnounceLocalTrustRequest_Body))
}

func (x *AnnounceLocalTrustResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return buf, nil
}

func (x *AnnounceLocalTrustResponseBody) StableSize() int {
	return 0
}

func (x *AnnounceLocalTrustResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.AnnounceLocalTrustResponse_Body))
}

const (
	_ = iota
	announceInterResBodyEpochFNum
	announceInterResBodyIterFNum
	announceInterResBodyTrustFNum
)

func (x *AnnounceIntermediateResultRequestBody) StableMarshal(buf []byte) ([]byte, error) {
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

	n, err = protoutil.UInt64Marshal(announceInterResBodyEpochFNum, buf, x.epoch)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = protoutil.UInt32Marshal(announceInterResBodyIterFNum, buf[offset:], x.iter)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.NestedStructureMarshal(announceInterResBodyTrustFNum, buf[offset:], x.trust)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (x *AnnounceIntermediateResultRequestBody) StableSize() (size int) {
	size += protoutil.UInt64Size(announceInterResBodyEpochFNum, x.epoch)
	size += protoutil.UInt32Size(announceInterResBodyIterFNum, x.iter)
	size += protoutil.NestedStructureSize(announceInterResBodyTrustFNum, x.trust)

	return
}

func (x *AnnounceIntermediateResultRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.AnnounceIntermediateResultRequest_Body))
}

func (x *AnnounceIntermediateResultResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	return buf, nil
}

func (x *AnnounceIntermediateResultResponseBody) StableSize() int {
	return 0
}

func (x *AnnounceIntermediateResultResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(reputation.AnnounceIntermediateResultResponse_Body))
}
