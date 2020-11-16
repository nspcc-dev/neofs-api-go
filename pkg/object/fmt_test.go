package object

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestVerificationFields(t *testing.T) {
	obj := NewRaw()

	payload := make([]byte, 10)
	_, _ = rand.Read(payload)

	obj.SetPayload(payload)
	obj.SetPayloadSize(uint64(len(payload)))

	require.NoError(t, SetVerificationFields(test.DecodeKey(-1), obj))

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
				obj.SetPayloadSize(obj.GetPayloadSize() + 1)
			},
			restore: func() {
				obj.SetPayloadSize(obj.GetPayloadSize() - 1)
			},
		},
		{
			corrupt: func() {
				obj.GetID().ToV2().GetValue()[0]++
			},
			restore: func() {
				obj.GetID().ToV2().GetValue()[0]--
			},
		},
		{
			corrupt: func() {
				obj.GetSignature().Key()[0]++
			},
			restore: func() {
				obj.GetSignature().Key()[0]--
			},
		},
		{
			corrupt: func() {
				obj.GetSignature().Sign()[0]++
			},
			restore: func() {
				obj.GetSignature().Sign()[0]--
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
