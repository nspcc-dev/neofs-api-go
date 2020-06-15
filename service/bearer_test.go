package service

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

type testBearerToken struct {
	aclRules []byte
	expEpoch uint64
	owner    OwnerID
	key      []byte
	sig      []byte
}

func (s testBearerToken) GetACLRules() []byte {
	return s.aclRules
}

func (s *testBearerToken) SetACLRules(v []byte) {
	s.aclRules = v
}

func (s testBearerToken) ExpirationEpoch() uint64 {
	return s.expEpoch
}

func (s *testBearerToken) SetExpirationEpoch(v uint64) {
	s.expEpoch = v
}

func (s testBearerToken) GetOwnerID() OwnerID {
	return s.owner
}

func (s *testBearerToken) SetOwnerID(v OwnerID) {
	s.owner = v
}

func (s testBearerToken) GetOwnerKey() []byte {
	return s.key
}

func (s *testBearerToken) SetOwnerKey(v []byte) {
	s.key = v
}

func (s testBearerToken) GetSignature() []byte {
	return s.sig
}

func (s *testBearerToken) SetSignature(v []byte) {
	s.sig = v
}

func TestBearerTokenMsgGettersSetters(t *testing.T) {
	var tok BearerToken = new(BearerTokenMsg)

	{ // ACLRules
		rules := []byte{1, 2, 3}

		tok.SetACLRules(rules)

		require.Equal(t, rules, tok.GetACLRules())
	}

	{ // OwnerID
		ownerID := OwnerID{}
		_, err := rand.Read(ownerID[:])
		require.NoError(t, err)

		tok.SetOwnerID(ownerID)

		require.Equal(t, ownerID, tok.GetOwnerID())
	}

	{ // ValidUntil
		e := uint64(5)

		tok.SetExpirationEpoch(e)

		require.Equal(t, e, tok.ExpirationEpoch())
	}

	{ // OwnerKey
		key := make([]byte, 10)
		_, err := rand.Read(key)
		require.NoError(t, err)

		tok.SetOwnerKey(key)

		require.Equal(t, key, tok.GetOwnerKey())
	}

	{ // Signature
		sig := make([]byte, 10)
		_, err := rand.Read(sig)
		require.NoError(t, err)

		tok.SetSignature(sig)

		require.Equal(t, sig, tok.GetSignature())
	}
}

func TestSignVerifyBearerToken(t *testing.T) {
	var token BearerToken = new(testBearerToken)

	// create private key for signing
	sk := test.DecodeKey(0)
	pk := &sk.PublicKey

	rules := []byte{1, 2, 3}
	token.SetACLRules(rules)

	ownerID := OwnerID{}
	_, err := rand.Read(ownerID[:])
	require.NoError(t, err)
	token.SetOwnerID(ownerID)

	fEpoch := uint64(2)
	token.SetExpirationEpoch(fEpoch)

	signedToken := NewSignedBearerToken(token)
	verifiedToken := NewVerifiedBearerToken(token)

	// sign and verify token
	require.NoError(t, AddSignatureWithKey(sk, signedToken))
	require.NoError(t, VerifySignatureWithKey(pk, verifiedToken))

	items := []struct {
		corrupt func()
		restore func()
	}{
		{ // ACLRules
			corrupt: func() {
				token.SetACLRules(append(rules, 1))
			},
			restore: func() {
				token.SetACLRules(rules)
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
		{ // Expiration epoch
			corrupt: func() {
				token.SetExpirationEpoch(fEpoch + 1)
			},
			restore: func() {
				token.SetExpirationEpoch(fEpoch)
			},
		},
	}

	for _, v := range items {
		v.corrupt()
		require.Error(t, VerifySignatureWithKey(pk, verifiedToken))
		v.restore()
		require.NoError(t, VerifySignatureWithKey(pk, verifiedToken))
	}
}
