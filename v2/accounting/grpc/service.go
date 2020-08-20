package accounting

import (
	refs "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	service "github.com/nspcc-dev/neofs-api-go/v2/service/grpc"
)

// SetOwnerId sets identifier of the account owner.
func (m *BalanceRequest_Body) SetOwnerId(v *refs.OwnerID) {
	if m != nil {
		m.OwnerId = v
	}
}

// SetBody sets body of the request.
func (m *BalanceRequest) SetBody(v *BalanceRequest_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the request.
func (m *BalanceRequest) SetMetaHeader(v *service.RequestMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the request.
func (m *BalanceRequest) SetVerifyHeader(v *service.RequestVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}

// SetBalance sets balance value of the response.
func (m *BalanceResponse_Body) SetBalance(v *Decimal) {
	if m != nil {
		m.Balance = v
	}
}

// SetBody sets body of the response.
func (m *BalanceResponse) SetBody(v *BalanceResponse_Body) {
	if m != nil {
		m.Body = v
	}
}

// SetMetaHeader sets meta header of the response.
func (m *BalanceResponse) SetMetaHeader(v *service.ResponseMetaHeader) {
	if m != nil {
		m.MetaHeader = v
	}
}

// SetVerifyHeader sets verification header of the response.
func (m *BalanceResponse) SetVerifyHeader(v *service.ResponseVerificationHeader) {
	if m != nil {
		m.VerifyHeader = v
	}
}
