package container_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/container"
	statustest "github.com/nspcc-dev/neofs-api-go/v2/status/test"
)

func TestStatusCodes(t *testing.T) {
	statustest.TestCodes(t, container.LocalizeFailStatus, container.GlobalizeFail,
		container.StatusNotFound, 3072,
		container.StatusEACLNotFound, 3073,
	)
}
