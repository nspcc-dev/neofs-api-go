package accounting

import (
	accounting "github.com/nspcc-dev/neofs-api-go/v2/accounting/grpc"
	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
)

func BalanceRequestBodyToGRPCMessage(b *BalanceRequestBody) *accounting.BalanceRequest_Body {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest_Body)

	m.SetOwnerId(
		refs.OwnerIDToGRPCMessage(b.GetOwnerID()),
	)

	return m
}

func BalanceRequestBodyFromGRPCMessage(m *accounting.BalanceRequest_Body) *BalanceRequestBody {
	if m == nil {
		return nil
	}

	b := new(BalanceRequestBody)

	b.SetOwnerID(
		refs.OwnerIDFromGRPCMessage(m.GetOwnerId()),
	)

	return b
}

func BalanceRequestToGRPCMessage(b *BalanceRequest) *accounting.BalanceRequest {
	if b == nil {
		return nil
	}

	m := new(accounting.BalanceRequest)

	m.SetBody(
		BalanceRequestBodyToGRPCMessage(b.GetBody()),
	)

	service.RequestHeadersToGRPC(b, m)

	return m
}

func BalanceRequestFromGRPCMessage(m *accounting.BalanceRequest) *BalanceRequest {
	if m == nil {
		return nil
	}

	b := new(BalanceRequest)

	b.SetBody(
		BalanceRequestBodyFromGRPCMessage(m.GetBody()),
	)

	service.RequestHeadersFromGRPC(m, b)

	return b
}

func DecimalToGRPCMessage(d *Decimal) *accounting.Decimal {
	if d == nil {
		return nil
	}

	m := new(accounting.Decimal)

	m.SetValue(d.GetValue())
	m.SetPrecision(d.GetPrecision())

	return m
}

func DecimalFromGRPCMessage(m *accounting.Decimal) *Decimal {
	if m == nil {
		return nil
	}

	d := new(Decimal)

	d.SetValue(m.GetValue())
	d.SetPrecision(m.GetPrecision())

	return d
}
