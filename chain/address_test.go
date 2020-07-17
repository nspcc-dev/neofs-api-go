package chain

import (
	"encoding/hex"
	"testing"

	crypto "github.com/nspcc-dev/neofs-crypto"
	"github.com/stretchr/testify/require"
)

type addressTestCase struct {
	name      string
	publicKey string
	wallet    string
}

func TestKeyToAddress(t *testing.T) {
	tests := []addressTestCase{
		{
			"nil key",
			"",
			"",
		},
		{
			"correct key",
			"031a6c6fbbdf02ca351745fa86b9ba5a9452d785ac4f7fc2b7548ca2a46c4fcf4a",
			"NgzuJWWGVEwFGsRrgzj8knswEYRJrTe7sm",
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			data, err := hex.DecodeString(tests[i].publicKey)
			require.NoError(t, err)

			key := crypto.UnmarshalPublicKey(data)

			require.Equal(t, tests[i].wallet, KeyToAddress(key))
		})
	}
}
