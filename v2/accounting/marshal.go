package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	decimalValueField     = 1
	decimalPrecisionField = 2

	balanceReqBodyOwnerField = 1

	balanceRespBodyDecimalField = 1
)

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

	n, err = proto.Int64Marshal(decimalValueField, buf[offset:], d.val)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt32Marshal(decimalPrecisionField, buf[offset:], d.prec)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (d *Decimal) StableSize() (size int) {
	if d == nil {
		return 0
	}

	size += proto.Int64Size(decimalValueField, d.val)
	size += proto.UInt32Size(decimalPrecisionField, d.prec)

	return size
}

func (b *BalanceRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if b == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, b.StableSize())
	}

	if b.ownerID != nil {
		_, err := proto.NestedStructureMarshal(balanceReqBodyOwnerField, buf, b.ownerID)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func (b *BalanceRequestBody) StableSize() (size int) {
	if b == nil {
		return 0
	}

	if b.ownerID != nil {
		size = proto.NestedStructureSize(balanceReqBodyOwnerField, b.ownerID)
	}

	return size
}

func (br *BalanceResponseBody) StableMarshal(buf []byte) ([]byte, error) {
	if br == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, br.StableSize())
	}

	if br.bal != nil {
		_, err := proto.NestedStructureMarshal(balanceRespBodyDecimalField, buf, br.bal)
		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func (br *BalanceResponseBody) StableSize() (size int) {
	if br == nil {
		return 0
	}

	if br.bal != nil {
		size = proto.NestedStructureSize(balanceRespBodyDecimalField, br.bal)
	}

	return size
}
