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

func BalanceResponseBodyToGRPCMessage(br *BalanceResponseBody) *accounting.BalanceResponse_Body {
	if br == nil {
		return nil
	}

	m := new(accounting.BalanceResponse_Body)

	m.SetBalance(
		DecimalToGRPCMessage(br.GetBalance()),
	)

	return m
}

func BalanceResponseBodyFromGRPCMessage(m *accounting.BalanceResponse_Body) *BalanceResponseBody {
	if m == nil {
		return nil
	}

	br := new(BalanceResponseBody)

	br.SetBalance(
		DecimalFromGRPCMessage(m.GetBalance()),
	)

	return br
}

func BalanceResponseToGRPCMessage(br *BalanceResponse) *accounting.BalanceResponse {
	if br == nil {
		return nil
	}

	m := new(accounting.BalanceResponse)

	m.SetBody(
		BalanceResponseBodyToGRPCMessage(br.GetBody()),
	)

	service.ResponseHeadersToGRPC(br, m)

	return m
}

func BalanceResponseFromGRPCMessage(m *accounting.BalanceResponse) *BalanceResponse {
	if m == nil {
		return nil
	}

	br := new(BalanceResponse)

	br.SetBody(
		BalanceResponseBodyFromGRPCMessage(m.GetBody()),
	)

	service.ResponseHeadersFromGRPC(m, br)

	return br
}
