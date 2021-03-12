package accountingtest

import (
	"github.com/nspcc-dev/neofs-api-go/v2/accounting"
	accountingtest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
	sessiontest "github.com/nspcc-dev/neofs-api-go/v2/session/test"
)

func GenerateBalanceRequest(empty bool) *accounting.BalanceRequest {
	m := new(accounting.BalanceRequest)

	m.SetBody(GenerateBalanceRequestBody(empty))
	m.SetMetaHeader(sessiontest.GenerateRequestMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateRequestVerificationHeader(empty))

	return m
}

func GenerateBalanceRequestBody(empty bool) *accounting.BalanceRequestBody {
	m := new(accounting.BalanceRequestBody)

	m.SetOwnerID(accountingtest.GenerateOwnerID(empty))

	return m
}

func GenerateBalanceResponse(empty bool) *accounting.BalanceResponse {
	m := new(accounting.BalanceResponse)

	m.SetBody(GenerateBalanceResponseBody(empty))
	m.SetMetaHeader(sessiontest.GenerateResponseMetaHeader(empty))
	m.SetVerificationHeader(sessiontest.GenerateResponseVerificationHeader(empty))

	return m
}

func GenerateBalanceResponseBody(empty bool) *accounting.BalanceResponseBody {
	m := new(accounting.BalanceResponseBody)

	m.SetBalance(GenerateDecimal(empty))

	return m
}

func GenerateDecimal(empty bool) *accounting.Decimal {
	m := new(accounting.Decimal)

	if !empty {
		m.SetValue(1)
		m.SetPrecision(2)
	}

	return m
}
