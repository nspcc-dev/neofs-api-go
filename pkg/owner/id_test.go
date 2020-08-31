package owner

import (
	"crypto/rand"
	"testing"

	"github.com/nspcc-dev/neofs-api-go/pkg/refs"
	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := new(ID)

	wallet := new(refs.NEO3Wallet)

	_, err := rand.Read(wallet.Bytes())
	require.NoError(t, err)

	id.SetNeo3Wallet(wallet)

	idV2 := id.ToV2()

	id2, err := IDFromV2(idV2)
	require.NoError(t, err)

	require.Equal(t, id, id2)
}
