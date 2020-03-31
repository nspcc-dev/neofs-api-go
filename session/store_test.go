package session

import (
	"crypto/ecdsa"
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/refs"
	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/stretchr/testify/require"
)

type testClient struct {
	*ecdsa.PrivateKey
	OwnerID OwnerID
}

func (c *testClient) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(c.PrivateKey, data)
}

func newTestClient(t *testing.T) *testClient {
	key, err := ecdsa.GenerateKey(defaultCurve(), rand.Reader)
	require.NoError(t, err)

	owner, err := refs.NewOwnerID(&key.PublicKey)
	require.NoError(t, err)

	return &testClient{PrivateKey: key, OwnerID: owner}
}

func signToken(t *testing.T, token *PToken, c *testClient) {
	require.NotNil(t, token)
	token.SetPublicKeys(&c.PublicKey)

	signH, err := c.Sign(token.Header.PublicKey)
	require.NoError(t, err)
	require.NotNil(t, signH)

	// data is not yet signed
	keys := UnmarshalPublicKeys(&token.Token)
	require.False(t, token.Verify(keys...))

	signT, err := c.Sign(token.verificationData())
	require.NoError(t, err)
	require.NotNil(t, signT)

	token.AddSignatures(signH, signT)
	require.True(t, token.Verify(keys...))
}

func TestTokenStore(t *testing.T) {
	s := NewSimpleStore()

	oid, err := refs.NewObjectID()
	require.NoError(t, err)

	c := newTestClient(t)
	require.NotNil(t, c)
	pk := [][]byte{crypto.MarshalPublicKey(&c.PublicKey)}

	// create new token
	token := s.New(TokenParams{
		ObjectID:   []ObjectID{oid},
		OwnerID:    c.OwnerID,
		PublicKeys: pk,
	})
	signToken(t, token, c)

	// check that it can be fetched
	t1 := s.Fetch(token.ID)
	require.NotNil(t, t1)
	require.Equal(t, token, t1)

	// create and sign another token by the same client
	t1 = s.New(TokenParams{
		ObjectID:   []ObjectID{oid},
		OwnerID:    c.OwnerID,
		PublicKeys: pk,
	})

	signToken(t, t1, c)

	data := []byte{1, 2, 3}
	sign, err := t1.SignData(data)
	require.NoError(t, err)
	require.Error(t, token.Header.VerifyData(data, sign))

	sign, err = token.SignData(data)
	require.NoError(t, err)
	require.NoError(t, token.Header.VerifyData(data, sign))

	s.Remove(token.ID)
	require.Nil(t, s.Fetch(token.ID))
	require.NotNil(t, s.Fetch(t1.ID))
}
