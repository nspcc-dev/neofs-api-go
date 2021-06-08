package session_test

import (
	"testing"

	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	sessiontest "github.com/nspcc-dev/neofs-api-go/pkg/session/test"
	sessionv2 "github.com/nspcc-dev/neofs-api-go/v2/session"
	"github.com/stretchr/testify/require"
)

func TestSessionToken_SetID(t *testing.T) {
	token := session.NewToken()

	id := []byte{1, 2, 3}
	token.SetID(id)

	require.Equal(t, id, token.ID())
}

func TestSessionToken_SetOwnerID(t *testing.T) {
	token := session.NewToken()

	ownerID := ownertest.Generate()

	token.SetOwnerID(ownerID)

	require.Equal(t, ownerID, token.OwnerID())
}

func TestSessionToken_SetSessionKey(t *testing.T) {
	token := session.NewToken()

	key := []byte{1, 2, 3}
	token.SetSessionKey(key)

	require.Equal(t, key, token.SessionKey())
}

func TestSessionTokenEncoding(t *testing.T) {
	tok := sessiontest.Generate()

	t.Run("binary", func(t *testing.T) {
		data, err := tok.Marshal()
		require.NoError(t, err)

		tok2 := session.NewToken()
		require.NoError(t, tok2.Unmarshal(data))

		require.Equal(t, tok, tok2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := tok.MarshalJSON()
		require.NoError(t, err)

		tok2 := session.NewToken()
		require.NoError(t, tok2.UnmarshalJSON(data))

		require.Equal(t, tok, tok2)
	})
}

func TestToken_VerifySignature(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var tok *session.Token

		require.False(t, tok.VerifySignature())
	})

	t.Run("unsigned", func(t *testing.T) {
		tok := sessiontest.Generate()

		require.False(t, tok.VerifySignature())
	})

	t.Run("signed", func(t *testing.T) {
		tok := sessiontest.GenerateSigned()

		require.True(t, tok.VerifySignature())
	})
}

var unsupportedContexts = []interface{}{
	123,
	true,
	session.NewToken(),
}

var nonContainerContexts = unsupportedContexts

func TestToken_Context(t *testing.T) {
	tok := session.NewToken()

	for _, item := range []struct {
		ctx      interface{}
		v2assert func(interface{})
	}{
		{
			ctx: sessiontest.ContainerContext(),
			v2assert: func(c interface{}) {
				require.Equal(t, c.(*session.ContainerContext).ToV2(), tok.ToV2().GetBody().GetContext())
			},
		},
	} {
		tok.SetContext(item.ctx)

		require.Equal(t, item.ctx, tok.Context())

		item.v2assert(item.ctx)
	}

	for _, c := range unsupportedContexts {
		tok.SetContext(c)

		require.Nil(t, tok.Context())
	}
}

func TestGetContainerContext(t *testing.T) {
	tok := session.NewToken()

	c := sessiontest.ContainerContext()

	tok.SetContext(c)

	require.Equal(t, c, session.GetContainerContext(tok))

	for _, c := range nonContainerContexts {
		tok.SetContext(c)

		require.Nil(t, session.GetContainerContext(tok))
	}
}

func TestToken_Exp(t *testing.T) {
	tok := session.NewToken()

	const exp = 11

	tok.SetExp(exp)

	require.EqualValues(t, exp, tok.Exp())
}

func TestToken_Nbf(t *testing.T) {
	tok := session.NewToken()

	const nbf = 22

	tok.SetNbf(nbf)

	require.EqualValues(t, nbf, tok.Nbf())
}

func TestToken_Iat(t *testing.T) {
	tok := session.NewToken()

	const iat = 33

	tok.SetIat(iat)

	require.EqualValues(t, iat, tok.Iat())
}

func TestNewTokenFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *sessionv2.SessionToken

		require.Nil(t, session.NewTokenFromV2(x))
	})
}

func TestToken_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *session.Token

		require.Nil(t, x.ToV2())
	})
}

func TestNewToken(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		token := session.NewToken()

		// check initial values
		require.Nil(t, token.Signature())
		require.Nil(t, token.OwnerID())
		require.Nil(t, token.SessionKey())
		require.Nil(t, token.ID())
		require.Zero(t, token.Exp())
		require.Zero(t, token.Iat())
		require.Zero(t, token.Nbf())

		// convert to v2 message
		tokenV2 := token.ToV2()

		require.Nil(t, tokenV2.GetSignature())
		require.Nil(t, tokenV2.GetBody())
	})
}
