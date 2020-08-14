package v2

type BalanceRequestBody struct {
	ownerID *OwnerID
}

type BalanceRequest struct {
	body *BalanceRequestBody

	metaHeader *RequestMetaHeader

	verifyHeader *RequestVerificationHeader
}

func (b *BalanceRequestBody) GetOwnerID() *OwnerID {
	if b != nil {
		return b.ownerID
	}

	return nil
}

func (b *BalanceRequestBody) SetOwnerID(v *OwnerID) {
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

func (b *BalanceRequest) GetRequestMetaHeader() *RequestMetaHeader {
	if b != nil {
		return b.metaHeader
	}

	return nil
}

func (b *BalanceRequest) SetRequestMetaHeader(v *RequestMetaHeader) {
	if b != nil {
		b.metaHeader = v
	}
}

func (b *BalanceRequest) GetRequestVerificationHeader() *RequestVerificationHeader {
	if b != nil {
		return b.verifyHeader
	}

	return nil
}

func (b *BalanceRequest) SetRequestVerificationHeader(v *RequestVerificationHeader) {
	if b != nil {
		b.verifyHeader = v
	}
}
