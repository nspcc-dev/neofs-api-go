package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	protoutil "github.com/nspcc-dev/neofs-api-go/util/proto"
	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
)

const (
	decimalValueField     = 1
	decimalPrecisionField = 2

	balanceReqBodyOwnerField = 1

	balanceRespBodyDecimalField = 1
)

func (d *Decimal) MarshalStream(s protoutil.Stream) (int, error) {
	if d == nil {
		return 0, nil
	}

	var (
		offset, n int
		err       error
	)

	n, err = s.Int64Marshal(decimalValueField, d.val)
	if err != nil {
		return offset + n, err
	}

	offset += n

	n, err = s.UInt32Marshal(decimalPrecisionField, d.prec)
	return offset + n, err
}

func (d *Decimal) StableMarshal(buf []byte) ([]byte, error) {
	if d == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, d.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = protoutil.Int64Marshal(decimalValueField, buf[offset:], d.val)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = protoutil.UInt32Marshal(decimalPrecisionField, buf[offset:], d.prec)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (d *Decimal) StableSize() (size int) {
	if d == nil {
		return 0
	}

	size += protoutil.Int64Size(decimalValueField, d.val)
	size += protoutil.UInt32Size(decimalPrecisionField, d.prec)

	return size
}

func (d *Decimal) Unmarshal(data []byte) error {
	return message.Unmarshal(d, data, new(accounting.Decimal))
}

func (b *BalanceRequestBody) MarshalStream(s protoutil.Stream) (int, error) {
	if b == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(balanceReqBodyOwnerField, b.ownerID)
}

func (b *BalanceRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if b == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, b.StableSize())
	}

	_, err := protoutil.NestedStructureMarshal(balanceReqBodyOwnerField, buf, b.ownerID)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (b *BalanceRequestBody) StableSize() (size int) {
	if b == nil {
		return 0
	}

	size = protoutil.NestedStructureSize(balanceReqBodyOwnerField, b.ownerID)

	return size
}

func (b *BalanceRequestBody) Unmarshal(data []byte) error {
	return message.Unmarshal(b, data, new(accounting.BalanceRequest_Body))
}

func (br *BalanceResponseBody) MarshalStream(s protoutil.Stream) (int, error) {
	if br == nil {
		return 0, nil
	}

	return s.NestedStructureMarshal(balanceRespBodyDecimalField, br.bal)
}

func (br *BalanceResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if br == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, br.StableSize())
	}

	_, err := protoutil.NestedStructureMarshal(balanceRespBodyDecimalField, buf, br.bal)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (br *BalanceResponseBody) StableSize() (size int) {
	if br == nil {
		return 0
	}

	size = protoutil.NestedStructureSize(balanceRespBodyDecimalField, br.bal)

	return size
}

func (br *BalanceResponseBody) Unmarshal(data []byte) error {
	return message.Unmarshal(br, data, new(accounting.BalanceResponse_Body))
}
