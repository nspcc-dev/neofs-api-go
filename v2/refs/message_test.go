package refs_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/rpc/message/test"
	refstest "github.com/nspcc-dev/neofs-api-go/v2/refs/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return refstest.GenerateOwnerID(empty) },
		func(empty bool) message.Message { return refstest.GenerateObjectID(empty) },
		func(empty bool) message.Message { return refstest.GenerateContainerID(empty) },
		func(empty bool) message.Message { return refstest.GenerateAddress(empty) },
		func(empty bool) message.Message { return refstest.GenerateChecksum(empty) },
		func(empty bool) message.Message { return refstest.GenerateSignature(empty) },
		func(empty bool) message.Message { return refstest.GenerateVersion(empty) },
	)
}
