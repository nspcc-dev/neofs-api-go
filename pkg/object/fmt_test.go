package object

import (
	"crypto/rand"
	"testing"

	neofsecdsatest "github.com/nspcc-dev/neofs-api-go/crypto/ecdsa/test"
	"github.com/stretchr/testify/require"
)

func TestVerificationFields(t *testing.T) {
	obj := NewRaw()

	payload := make([]byte, 10)
	_, _ = rand.Read(payload)

	obj.SetPayload(payload)
	obj.SetPayloadSize(uint64(len(payload)))

	require.NoError(t, SetVerificationFieldsECDSA(neofsecdsatest.Key(), obj))

	require.NoError(t, CheckVerificationFields(obj.Object()))

	items := []struct {
		corrupt func()
		restore func()
	}{
		{
			corrupt: func() {
				payload[0]++
			},
			restore: func() {
				payload[0]--
			},
		},
		{
			corrupt: func() {
				obj.SetPayloadSize(obj.PayloadSize() + 1)
			},
			restore: func() {
				obj.SetPayloadSize(obj.PayloadSize() - 1)
			},
		},
		{
			corrupt: func() {
				obj.ID().ToV2().GetValue()[0]++
			},
			restore: func() {
				obj.ID().ToV2().GetValue()[0]--
			},
		},
		{
			corrupt: func() {
				obj.Signature().Key()[0]++
			},
			restore: func() {
				obj.Signature().Key()[0]--
			},
		},
		{
			corrupt: func() {
				obj.Signature().Sign()[0]++
			},
			restore: func() {
				obj.Signature().Sign()[0]--
			},
		},
	}

	for _, item := range items {
		item.corrupt()

		require.Error(t, CheckVerificationFields(obj.Object()))

		item.restore()

		require.NoError(t, CheckVerificationFields(obj.Object()))
	}
}
