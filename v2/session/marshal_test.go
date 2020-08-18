package session_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/service"
	"github.com/nspcc-dev/neofs-api-go/v2/session"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	"github.com/stretchr/testify/require"
)

func TestCreateRequestBody_StableMarshal(t *testing.T) {
	requestFrom := generateCreateSessionRequestBody("Owner ID")
	transport := new(grpc.CreateRequest_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := requestFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		requestTo := session.CreateRequestBodyFromGRPCMessage(transport)
		require.Equal(t, requestFrom, requestTo)
	})
}

func TestCreateResponseBody_StableMarshal(t *testing.T) {
	responseFrom := generateCreateSessionResponseBody("ID", "Session Public Key")
	transport := new(grpc.CreateResponse_Body)

	t.Run("non empty", func(t *testing.T) {
		wire, err := responseFrom.StableMarshal(nil)
		require.NoError(t, err)

		err = transport.Unmarshal(wire)
		require.NoError(t, err)

		responseTo := session.CreateResponseBodyFromGRPCMessage(transport)
		require.Equal(t, responseFrom, responseTo)
	})
}

func generateCreateSessionRequestBody(id string) *session.CreateRequestBody {
	lifetime := new(service.TokenLifetime)
	lifetime.SetIat(1)
	lifetime.SetNbf(2)
	lifetime.SetExp(3)

	owner := new(refs.OwnerID)
	owner.SetValue([]byte(id))

	s := new(session.CreateRequestBody)
	s.SetOwnerID(owner)
	s.SetLifetime(lifetime)

	return s
}

func generateCreateSessionResponseBody(id, key string) *session.CreateResponseBody {
	s := new(session.CreateResponseBody)
	s.SetID([]byte(id))
	s.SetSessionKey([]byte(key))

	return s
}
