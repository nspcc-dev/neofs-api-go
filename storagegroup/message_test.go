package storagegroup_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/rpc/message"
	messagetest "github.com/nspcc-dev/neofs-api-go/v2/rpc/message/test"
	storagegrouptest "github.com/nspcc-dev/neofs-api-go/v2/storagegroup/test"
)

func TestMessageConvert(t *testing.T) {
	messagetest.TestRPCMessage(t,
		func(empty bool) message.Message { return storagegrouptest.GenerateStorageGroup(empty) },
	)
}
