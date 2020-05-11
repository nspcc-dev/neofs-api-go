package accounting

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestSignBalanceRequest(t *testing.T) {
	// create test OwnerID
	ownerID := OwnerID{1, 2, 3}

	// create test BalanceRequest
	req := new(BalanceRequest)
	req.SetOwnerID(ownerID)

	// create test private key
	sk := test.DecodeKey(0)

	items := []struct {
		corrupt func()
		restore func()
	}{
		{
			corrupt: func() {
				ownerID[0]++
				req.SetOwnerID(ownerID)
			},
			restore: func() {
				ownerID[0]--
				req.SetOwnerID(ownerID)
			},
		},
	}

	for _, item := range items {
		// sign with private key
		require.NoError(t, service.AddSignatureWithKey(sk, req))

		// ascertain that verification is passed
		require.NoError(t, service.VerifyAccumulatedSignatures(req))

		// corrupt the request
		item.corrupt()

		// ascertain that verification is failed
		require.Error(t, service.VerifyAccumulatedSignatures(req))

		// ascertain that request
		item.restore()
	}
}
