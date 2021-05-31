package session_test

import (
	"testing"

	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	"github.com/nspcc-dev/neofs-api-go/pkg/session"
	sessiontest "github.com/nspcc-dev/neofs-api-go/pkg/session/test"
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
