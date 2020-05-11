package bootstrap

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestRequestSign(t *testing.T) {
	sk := test.DecodeKey(0)

	type sigType interface {
		service.SignedDataWithToken
		service.SignKeyPairAccumulator
		service.SignKeyPairSource
		SetToken(*service.Token)
	}

	items := []struct {
		constructor    func() sigType
		payloadCorrupt []func(sigType)
	}{
		{ // Request
			constructor: func() sigType {
				return new(Request)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*Request)

					req.SetType(req.GetType() + 1)
				},
				func(s sigType) {
					req := s.(*Request)

					req.SetState(req.GetState() + 1)
				},
				func(s sigType) {
					req := s.(*Request)

					info := req.GetInfo()
					info.Address += "1"

					req.SetInfo(info)
				},
			},
		},
	}

	for _, item := range items {
		{ // token corruptions
			v := item.constructor()

			token := new(service.Token)
			v.SetToken(token)

			require.NoError(t, service.SignDataWithSessionToken(sk, v))

			require.NoError(t, service.VerifyAccumulatedSignaturesWithToken(v))

			token.SetSessionKey(append(token.GetSessionKey(), 1))

			require.Error(t, service.VerifyAccumulatedSignaturesWithToken(v))
		}

		{ // payload corruptions
			for _, corruption := range item.payloadCorrupt {
				v := item.constructor()

				require.NoError(t, service.SignDataWithSessionToken(sk, v))

				require.NoError(t, service.VerifyAccumulatedSignaturesWithToken(v))

				corruption(v)

				require.Error(t, service.VerifyAccumulatedSignaturesWithToken(v))
			}
		}
	}
}
