package token_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/acl/eacl"
	"github.com/nspcc-dev/neofs-api-go/pkg/owner"
	"github.com/nspcc-dev/neofs-api-go/pkg/token"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestBearerToken_Issuer(t *testing.T) {
	bearerToken := token.NewBearerToken()

	t.Run("non signed token", func(t *testing.T) {
		require.Nil(t, bearerToken.Issuer())
	})

	t.Run("signed token", func(t *testing.T) {
		key := test.DecodeKey(1)

		wallet, err := owner.NEO3WalletFromPublicKey(&key.PublicKey)
		require.NoError(t, err)

		ownerID := owner.NewIDFromNeo3Wallet(wallet)

		bearerToken.SetEACLTable(eacl.NewTable())
		require.NoError(t, bearerToken.SignToken(key))
		require.True(t, ownerID.Equal(bearerToken.Issuer()))
	})
}

func TestFilterEncoding(t *testing.T) {
	f := token.NewBearerToken()
	f.SetLifetime(1, 2, 3)

	t.Run("binary", func(t *testing.T) {
		data, err := f.Marshal()
		require.NoError(t, err)

		f2 := token.NewBearerToken()
		require.NoError(t, f2.Unmarshal(data))

		require.Equal(t, f, f2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := f.MarshalJSON()
		require.NoError(t, err)

		d2 := token.NewBearerToken()
		require.NoError(t, d2.UnmarshalJSON(data))

		require.Equal(t, f, d2)
	})
}
