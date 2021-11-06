package status_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	statustest "github.com/nspcc-dev/neofs-api-go/v2/status/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return statustest.Detail(empty) },
		func(empty bool) message.Message { return statustest.Status(empty) },
	)
}
