package state

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestRequestSign(t *testing.T) {
	sk := test.DecodeKey(0)

	type sigType interface {
		service.RequestData
		service.SignKeyPairAccumulator
		service.SignKeyPairSource
		SetToken(*service.Token)
	}

	items := []struct {
		constructor    func() sigType
		payloadCorrupt []func(sigType)
	}{
		{ // NetmapRequest
			constructor: func() sigType {
				return new(NetmapRequest)
			},
		},
		{ // MetricsRequest
			constructor: func() sigType {
				return new(MetricsRequest)
			},
		},
		{ // HealthRequest
			constructor: func() sigType {
				return new(HealthRequest)
			},
		},
		{ // DumpRequest
			constructor: func() sigType {
				return new(DumpRequest)
			},
		},
		{ // DumpVarsRequest
			constructor: func() sigType {
				return new(DumpVarsRequest)
			},
		},
		{
			constructor: func() sigType {
				return new(ChangeStateRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*ChangeStateRequest)

					req.SetState(req.GetState() + 1)
				},
			},
		},
	}

	for _, item := range items {
		{ // token corruptions
			v := item.constructor()

			token := new(service.Token)
			v.SetToken(token)

			require.NoError(t, service.SignRequestData(sk, v))

			require.NoError(t, service.VerifyRequestData(v))

			token.SetSessionKey(append(token.GetSessionKey(), 1))

			require.Error(t, service.VerifyRequestData(v))
		}

		{ // payload corruptions
			for _, corruption := range item.payloadCorrupt {
				v := item.constructor()

				require.NoError(t, service.SignRequestData(sk, v))

				require.NoError(t, service.VerifyRequestData(v))

				corruption(v)

				require.Error(t, service.VerifyRequestData(v))
			}
		}
	}
}
