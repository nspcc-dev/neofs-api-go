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
