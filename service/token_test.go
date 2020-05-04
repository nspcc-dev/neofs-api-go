package service

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/nspcc-dev/neofs-crypto/test"
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

func TestSignToken(t *testing.T) {
	// nil token
	require.EqualError(t,
		SignToken(nil, nil),
		ErrNilToken.Error(),
	)

	require.EqualError(t,
		VerifyTokenSignature(nil, nil),
		ErrNilToken.Error(),
	)

	var token SessionToken = new(Token)

	// nil key
	require.EqualError(t,
		SignToken(token, nil),
		crypto.ErrEmptyPrivateKey.Error(),
	)

	require.EqualError(t,
		VerifyTokenSignature(token, nil),
		crypto.ErrEmptyPublicKey.Error(),
	)

	// create private key for signing
	sk := test.DecodeKey(0)
	pk := &sk.PublicKey

	id := TokenID{}
	_, err := rand.Read(id[:])
	require.NoError(t, err)
	token.SetID(id)

	ownerID := OwnerID{}
	_, err = rand.Read(ownerID[:])
	require.NoError(t, err)
	token.SetOwnerID(ownerID)

	verb := Token_Info_Verb(1)
	token.SetVerb(verb)

	addr := Address{}
	_, err = rand.Read(addr.ObjectID[:])
	require.NoError(t, err)
	_, err = rand.Read(addr.CID[:])
	require.NoError(t, err)
	token.SetAddress(addr)

	cEpoch := uint64(1)
	token.SetCreationEpoch(cEpoch)

	fEpoch := uint64(2)
	token.SetExpirationEpoch(fEpoch)

	sessionKey := make([]byte, 10)
	_, err = rand.Read(sessionKey[:])
	require.NoError(t, err)
	token.SetSessionKey(sessionKey)

	// sign and verify token
	require.NoError(t, SignToken(token, sk))
	require.NoError(t, VerifyTokenSignature(token, pk))

	items := []struct {
		corrupt func()
		restore func()
	}{
		{ // ID
			corrupt: func() {
				id[0]++
				token.SetID(id)
			},
			restore: func() {
				id[0]--
				token.SetID(id)
			},
		},
		{ // Owner ID
			corrupt: func() {
				ownerID[0]++
				token.SetOwnerID(ownerID)
			},
			restore: func() {
				ownerID[0]--
				token.SetOwnerID(ownerID)
			},
		},
		{ // Verb
			corrupt: func() {
				token.SetVerb(verb + 1)
			},
			restore: func() {
				token.SetVerb(verb)
			},
		},
		{ // ObjectID
			corrupt: func() {
				addr.ObjectID[0]++
				token.SetAddress(addr)
			},
			restore: func() {
				addr.ObjectID[0]--
				token.SetAddress(addr)
			},
		},
		{ // CID
			corrupt: func() {
				addr.CID[0]++
				token.SetAddress(addr)
			},
			restore: func() {
				addr.CID[0]--
				token.SetAddress(addr)
			},
		},
		{ // Creation epoch
			corrupt: func() {
				token.SetCreationEpoch(cEpoch + 1)
			},
			restore: func() {
				token.SetCreationEpoch(cEpoch)
			},
		},
		{ // Expiration epoch
			corrupt: func() {
				token.SetExpirationEpoch(fEpoch + 1)
			},
			restore: func() {
				token.SetExpirationEpoch(fEpoch)
			},
		},
		{ // Session key
			corrupt: func() {
				sessionKey[0]++
				token.SetSessionKey(sessionKey)
			},
			restore: func() {
				sessionKey[0]--
				token.SetSessionKey(sessionKey)
			},
		},
	}

	for _, v := range items {
		v.corrupt()
		require.Error(t, VerifyTokenSignature(token, pk))
		v.restore()
		require.NoError(t, VerifyTokenSignature(token, pk))
	}
}
