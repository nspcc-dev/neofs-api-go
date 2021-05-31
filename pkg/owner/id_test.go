package owner_test

import (
	"strconv"
	"testing"

	"github.com/mr-tron/base58"
	. "github.com/nspcc-dev/neofs-api-go/pkg/owner"
	ownertest "github.com/nspcc-dev/neofs-api-go/pkg/owner/test"
	"github.com/nspcc-dev/neofs-crypto/test"
	"github.com/stretchr/testify/require"
)

func TestIDV2(t *testing.T) {
	id := ownertest.Generate()

	idV2 := id.ToV2()

	require.Equal(t, id, NewIDFromV2(idV2))
}

func TestNewIDFromNeo3Wallet(t *testing.T) {
	wallet, err := NEO3WalletFromPublicKey(&test.DecodeKey(1).PublicKey)
	require.NoError(t, err)

	id := NewIDFromNeo3Wallet(wallet)
	require.Equal(t, id.ToV2().GetValue(), wallet.Bytes())
}

func TestID_Parse(t *testing.T) {
	t.Run("should parse successful", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			j := i
			t.Run(strconv.Itoa(j), func(t *testing.T) {
				wallet, err := NEO3WalletFromPublicKey(&test.DecodeKey(j).PublicKey)
				require.NoError(t, err)

				eid := NewIDFromNeo3Wallet(wallet)
				aid := NewID()

				require.NoError(t, aid.Parse(eid.String()))
				require.Equal(t, eid, aid)
			})
		}
	})

	t.Run("should failure on parse", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			j := i
			t.Run(strconv.Itoa(j), func(t *testing.T) {
				cs := []byte{1, 2, 3, 4, 5, byte(j)}
				str := base58.Encode(cs)
				cid := NewID()

				require.Error(t, cid.Parse(str))
			})
		}
	})
}

func TestIDEncoding(t *testing.T) {
	id := ownertest.Generate()

	t.Run("binary", func(t *testing.T) {
		data, err := id.Marshal()
		require.NoError(t, err)

		id2 := NewID()
		require.NoError(t, id2.Unmarshal(data))

		require.Equal(t, id, id2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := id.MarshalJSON()
		require.NoError(t, err)

		a2 := NewID()
		require.NoError(t, a2.UnmarshalJSON(data))

		require.Equal(t, id, a2)
	})
}

func TestID_Equal(t *testing.T) {
	var (
		data1 = []byte{1, 2, 3}
		data2 = data1
		data3 = append(data1, 255)
	)

	id1 := ownertest.GenerateFromBytes(data1)

	require.True(t, id1.Equal(
		ownertest.GenerateFromBytes(data2),
	))

	require.False(t, id1.Equal(
		ownertest.GenerateFromBytes(data3),
	))
}
