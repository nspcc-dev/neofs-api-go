package audit_test

import (
	"testing"

	audittest "github.com/nspcc-dev/neofs-api-go/v2/audit/test"
	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/v2/rpc/message/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return audittest.GenerateDataAuditResult(empty) },
	)
}
