package container

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
		SetBearer(*service.BearerTokenMsg)
	}

	items := []struct {
		constructor    func() sigType
		payloadCorrupt []func(sigType)
	}{
		{ // PutRequest
			constructor: func() sigType {
				return new(PutRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*PutRequest)

					id := req.GetMessageID()
					id[0]++

					req.SetMessageID(id)
				},
				func(s sigType) {
					req := s.(*PutRequest)

					req.SetCapacity(req.GetCapacity() + 1)
				},
				func(s sigType) {
					req := s.(*PutRequest)

					owner := req.GetOwnerID()
					owner[0]++

					req.SetOwnerID(owner)
				},
				func(s sigType) {
					req := s.(*PutRequest)

					rules := req.GetRules()
					rules.ReplFactor++

					req.SetRules(rules)
				},
				func(s sigType) {
					req := s.(*PutRequest)

					req.SetBasicACL(req.GetBasicACL() + 1)
				},
			},
		},
		{ // DeleteRequest
			constructor: func() sigType {
				return new(DeleteRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*DeleteRequest)

					cid := req.GetCID()
					cid[0]++

					req.SetCID(cid)
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

					cid := req.GetCID()
					cid[0]++

					req.SetCID(cid)
				},
			},
		},
		{ // ListRequest
			constructor: func() sigType {
				return new(ListRequest)
			},
			payloadCorrupt: []func(sigType){
				func(s sigType) {
					req := s.(*ListRequest)

					owner := req.GetOwnerID()
					owner[0]++

					req.SetOwnerID(owner)
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

		{ // Bearer token corruptions
			v := item.constructor()

			token := new(service.BearerTokenMsg)
			v.SetBearer(token)

			require.NoError(t, service.SignRequestData(sk, v))

			require.NoError(t, service.VerifyRequestData(v))

			token.SetACLRules(append(token.GetACLRules(), 1))

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
