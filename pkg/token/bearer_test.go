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
		require.Equal(t, bearerToken.Issuer().String(), ownerID.String())
	})
}
