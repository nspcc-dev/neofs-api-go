package accounting

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/service"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestSignBalanceRequest(t *testing.T) {
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
		{ // BalanceRequest
			constructor: func() sigType {
				return new(BalanceRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*BalanceRequest)

					owner := req.GetOwnerID()
					owner[0]++

					req.SetOwnerID(owner)
				},
			},
		},
		{ // GetRequest
			constructor: func() sigType {
				return new(GetRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*GetRequest)

					id, err := NewChequeID()
					require.NoError(t, err)

					req.SetID(id)
				},
				func(s sigType) {
					req := s.(*GetRequest)

					id := req.GetOwnerID()
					id[0]++

					req.SetOwnerID(id)
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
