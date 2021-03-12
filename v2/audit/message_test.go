package audit_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	audittest "github.com/nspcc-dev/neofs-api-go/v2/audit/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return audittest.GenerateDataAuditResult(empty) },
	)
}
