package owner

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := NewID()

	wallet := new(NEO3Wallet)

	_, err := rand.Read(wallet.Bytes())
	require.NoError(t, err)

	id.SetNeo3Wallet(wallet)

	idV2 := id.ToV2()

	require.Equal(t, wallet.Bytes(), idV2.GetValue())
}
