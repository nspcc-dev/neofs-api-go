package service

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	"github.com/stretchr/testify/require"
)

func TestTokenGettersSetters(t *testing.T) {
	var tok SessionToken = new(Token)

	{ // ID
		id, err := refs.NewUUID()
		require.NoError(t, err)

		tok.SetID(id)

		require.Equal(t, id, tok.GetID())
	}

	{ // OwnerID
		ownerID := OwnerID{}
		_, err := rand.Read(ownerID[:])
		require.NoError(t, err)

		tok.SetOwnerID(ownerID)

		require.Equal(t, ownerID, tok.GetOwnerID())
	}

	{ // Verb
		verb := Token_Info_Verb(3)

		tok.SetVerb(verb)

		require.Equal(t, verb, tok.GetVerb())
	}

	{ // Address
		addr := Address{}
		_, err := rand.Read(addr.CID[:])
		require.NoError(t, err)
		_, err = rand.Read(addr.ObjectID[:])
		require.NoError(t, err)

		tok.SetAddress(addr)

		require.Equal(t, addr, tok.GetAddress())
	}

	{ // Created
		e := uint64(5)

		tok.SetCreationEpoch(e)

		require.Equal(t, e, tok.CreationEpoch())
	}

	{ // ValidUntil
		e := uint64(5)

		tok.SetExpirationEpoch(e)

		require.Equal(t, e, tok.ExpirationEpoch())
	}

	{ // SessionKey
		key := make([]byte, 10)
		_, err := rand.Read(key)
		require.NoError(t, err)

		tok.SetSessionKey(key)

		require.Equal(t, key, tok.GetSessionKey())
	}

	{ // Signature
		sig := make([]byte, 10)
		_, err := rand.Read(sig)
		require.NoError(t, err)

		tok.SetSignature(sig)

		require.Equal(t, sig, tok.GetSignature())
	}
}
