package status

import (
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	status "github.com/nspcc-dev/neofs-api-go/v2/status/grpc"
	protoutil "github.com/nspcc-dev/neofs-api-go/v2/util/proto"
)

const (
	_ = iota
	detailIDFNum
	detailValueFNum
)

func (x *Detail) StableMarshal(buf []byte) ([]byte, error) {
	if x == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, x.StableSize())
	}

	var offset int

	offset += protoutil.UInt32Marshal(detailIDFNum, buf[offset:], x.id)
	offset += protoutil.BytesMarshal(detailValueFNum, buf[offset:], x.val)

	return buf, nil
}

func (x *Detail) StableSize() (size int) {
	size += protoutil.UInt32Size(detailIDFNum, x.id)
	size += protoutil.BytesSize(detailValueFNum, x.val)

	return size
}

func (x *Detail) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(status.Status_Detail))
}

const (
	_ = iota
	statusCodeFNum
	statusMsgFNum
	statusDetailsFNum
)

func (x *Status) StableMarshal(buf []byte) ([]byte, error) {
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

	offset += protoutil.UInt32Marshal(statusCodeFNum, buf[offset:], CodeToGRPC(x.code))
	if err != nil {
		return nil, err
	}

	offset += protoutil.StringMarshal(statusMsgFNum, buf[offset:], x.msg)
	if err != nil {
		return nil, err
	}

	for i := range x.details {
		n, err = protoutil.NestedStructureMarshal(statusDetailsFNum, buf[offset:], &x.details[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

func (x *Status) StableSize() (size int) {
	size += protoutil.UInt32Size(statusCodeFNum, CodeToGRPC(x.code))
	size += protoutil.StringSize(statusMsgFNum, x.msg)

	for i := range x.details {
		size += protoutil.NestedStructureSize(statusDetailsFNum, &x.details[i])
	}

	return size
}

func (x *Status) Unmarshal(data []byte) error {
	return message.Unmarshal(x, data, new(status.Status))
}
