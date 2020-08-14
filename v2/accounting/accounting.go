package accounting

import (
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

type BalanceRequestBody struct {
	ownerID *refs.OwnerID
}

type BalanceRequest struct {
	body *BalanceRequestBody

	metaHeader *service.RequestMetaHeader

	verifyHeader *service.RequestVerificationHeader
}

func (b *BalanceRequestBody) GetOwnerID() *refs.OwnerID {
	if b != nil {
		return b.ownerID
	}

	return nil
}

func (b *BalanceRequestBody) SetOwnerID(v *refs.OwnerID) {
	if b != nil {
		b.ownerID = v
	}
}

func (r *BalanceRequestBody) StableMarshal(buf []byte) ([]byte, error) {
	if r == nil {
		return nil, nil
	}

	// TODO: do not use hack
	_, err := BalanceRequestBodyToGRPCMessage(r).MarshalTo(buf)

	return buf, err
}

func (r *BalanceRequestBody) StableSize() int {
	// TODO: do not use hack
	return BalanceRequestBodyToGRPCMessage(r).Size()
}

func (b *BalanceRequest) GetBody() *BalanceRequestBody {
	if b != nil {
		return b.body
	}

	return nil
}

func (b *BalanceRequest) SetBody(v *BalanceRequestBody) {
	if b != nil {
		b.body = v
	}
}

func (b *BalanceRequest) GetRequestMetaHeader() *service.RequestMetaHeader {
	if b != nil {
		return b.metaHeader
	}

	return nil
}

func (b *BalanceRequest) SetRequestMetaHeader(v *service.RequestMetaHeader) {
	if b != nil {
		b.metaHeader = v
	}
}

func (b *BalanceRequest) GetRequestVerificationHeader() *service.RequestVerificationHeader {
	if b != nil {
		return b.verifyHeader
	}

	return nil
}

func (b *BalanceRequest) SetRequestVerificationHeader(v *service.RequestVerificationHeader) {
	if b != nil {
		b.verifyHeader = v
	}
}
