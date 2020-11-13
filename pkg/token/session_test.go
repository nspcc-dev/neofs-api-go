package token

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/stretchr/testify/require"
)

func TestSessionToken_SetID(t *testing.T) {
	token := NewSessionToken()

	id := []byte{1, 2, 3}
	token.SetID(id)

	require.Equal(t, id, token.ID())
}

func TestSessionToken_SetOwnerID(t *testing.T) {
	token := NewSessionToken()

	w := new(owner.NEO3Wallet)
	_, err := rand.Read(w.Bytes())
	require.NoError(t, err)

	ownerID := owner.NewID()
	ownerID.SetNeo3Wallet(w)

	token.SetOwnerID(ownerID)

	require.Equal(t, ownerID, token.OwnerID())
}

func TestSessionToken_SetSessionKey(t *testing.T) {
	token := NewSessionToken()

	key := []byte{1, 2, 3}
	token.SetSessionKey(key)

	require.Equal(t, key, token.SessionKey())
}

func TestSessionTokenEncoding(t *testing.T) {
	tok := NewSessionToken()
	tok.SetID([]byte("id"))

	t.Run("binary", func(t *testing.T) {
		data, err := tok.Marshal()
		require.NoError(t, err)

		tok2 := NewSessionToken()
		require.NoError(t, tok2.Unmarshal(data))

		require.Equal(t, tok, tok2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := tok.MarshalJSON()
		require.NoError(t, err)

		tok2 := NewSessionToken()
		require.NoError(t, tok2.UnmarshalJSON(data))

		require.Equal(t, tok, tok2)
	})
}
