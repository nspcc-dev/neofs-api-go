package accounting_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	accountingtest "github.com/nspcc-dev/neofs-api-go/v2/accounting/test"
)

func TestMessage(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return accountingtest.GenerateDecimal(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceRequestBody(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceRequest(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceResponseBody(empty) },
		func(empty bool) message.Message { return accountingtest.GenerateBalanceResponse(empty) },
	)
}
